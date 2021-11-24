package logger

import (
	"testing"
)

func BenchmarkLoggerDebugAllocation(b *testing.B){
	l := New(DebugLevel)
	for i := 0; i < b.N; i ++{
		l.WithFields(map[string]interface{}{
			"test": "test",
		}).Debug()
	}
}

func BenchmarkStructuredLogger(b *testing.B){
	l := NewStructuredLog(InfoLevel)
	for i := 0; i < b.N; i ++ {
		l.Info("failed to fetch URL", String("test", "test"))
	}
	l.Sync()
}