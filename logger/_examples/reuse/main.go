package main

import (
	"context"

	"github.com/elgopher/yala/adapter/console"
	"github.com/elgopher/yala/logger"
)

// This example shows how to reuse loggers
func main() {
	ctx := context.Background()

	log := logger.Local{Adapter: console.StdoutAdapter()}

	// requestLogger will log all messages with at least two fields: request_id and user
	requestLogger := log.With(ctx, "request_id", "123").With("user", "elgopher")

	requestLogger.Debug("request started")
	requestLogger.With("rows_updated", 3).With("table", "gophers").Debug("sql update executed")
	requestLogger.Debug("request finished")
}
