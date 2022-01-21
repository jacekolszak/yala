package benchmark

import (
	"context"
	"testing"

	"github.com/jacekolszak/yala/logger"
)

func Adapter(b *testing.B, adapter logger.Adapter) {
	b.Helper()

	ctx := context.Background()

	b.Run("global logger info", func(b *testing.B) {
		logger.SetAdapter(adapter)

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			logger.Info(ctx, "msg")
		}
	})

	b.Run("local logger info", func(b *testing.B) {
		localLogger := logger.Local(adapter)

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			localLogger.Info(ctx, "msg")
		}
	})
}
