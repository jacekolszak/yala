# YALA - Yet Another Logging Abstraction

Simple logging abstraction with adapters for most popular logging Go libraries and easy way to roll your own.

## Supported logging implementations

[logrus](adapter/logrusadapter), [zap](adapter/zapadapter), [zerolog](adapter/zerologadapter), [glog](adapter/glogadapter), [log15](adapter/log15adapter) and [standard fmt and log packages](adapter/printer)

## When to use?

* If you are a module/package author
* And you want to participate in a caller logging system (log messages using the logger provided by the consumer)
* And you don't want to add dependency to any specific logging library to your code
* And you don't want to manually inject logger to every possible place where you want to log something
* If you need nice and elegant API with a bunch of useful functions, but at the same time you don't want your clients spend hours on writing their own logging adapter.

## Installation

```shell
# Add recent version of yala to Go module:
go get -d -u github.com/jacekolszak/yala        
```

## How to use

### Set logger implementation (globally)

```go
import (
	"github.com/jacekolszak/yala/adapter/printer"
	"github.com/jacekolszak/yala/logger"
)
...
logger.SetAdapter(printer.StdoutAdapter())
```

### Log message in any function

```go
logger.Debug(ctx, "Debug message")
logger.With(ctx, "field_name", "value").Info("Message with field")
logger.WithError(ctx, err).Error("Message with error")
```

### Why context.Context is a parameter?

`context.Context` can very useful in transiting request-scoped tags or logger. logger.Adapter implementation might use them
making possible to log messages instrumented with tags.

### Why global state?

Logging is a special kind of dependency. It is used all over the place. Adding it as an explicit dependency to every
function, struct etc. can be cumbersome. Still though, you have an option to use local logger by injecting
logger.Adapter into your library:

```go
// your library code:
func NewLibrary(adapter logger.Adapter) YourLib {
    localLogger := logger.Local(adapter)
    return YourLib{localLogger: localLogger}
}

func (l YourLib) Method(ctx context.Context) {
    l.localLogger.Debug(ctx, "message from local logger")
}

// client code
adapter := printer.StdoutAdapter()
lib := NewLibrary(adapter)
```

### Difference between logger.Logger and logger.Adapter

* logger.Logger is a struct for logging messages (optionally with fields and error). It is used by packages in your module.
* logger.Adapter is an abstraction which should be implemented by adapters, such as logrusadapter (consumer of your module should provide it)

### Writing your own adapter

```go
type Adapter struct{}

func (s Adapter) Log(ctx context.Context, entry logger.Entry) {
    // here you can do whatever you want with the log entry 
}
```

### Why just don't create my own abstraction instead of using yala?

Yes, you can also create your own. Very often it just an interface with a single method, like this:

```go
type ImaginaryLogger interface {
    Log(context.Context, Entry)
}
```

But there are limitations for such solution:

* such interface is not very easy to use in your module
* someone who is using your module is supposed to write his own adapter of this interface (or you can provide adapters which
  of course takes your valuable time)
* it is not obvious how logging API should look like

### YALA limitations

* even though your module will be independent of any specific logging implementation, you still have to import 
  `github.com/jacekolszak/yala/logger`. This package is relatively small though, compared to real logging libraries
  (about ~240 lines of production code) and it does not import any external libraries.
* yala is optimized for the ease of use (both for the developer who logs messages and for the developer writing
  adapter). It is not optimized for performance, because this would hurt the user experience and readability of the
  created code.