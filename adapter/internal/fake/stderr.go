package fake

import (
	"os"
	"testing"
)

func UseFakeStderr(t *testing.T) Std {
	t.Helper()

	return useStd(t,
		func() *os.File {
			return os.Stderr
		},
		func(f *os.File) {
			os.Stderr = f
		},
	)
}
