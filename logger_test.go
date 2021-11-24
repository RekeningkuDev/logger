package logger

import (
	"testing"
)

func BenchmarkLoggerDebugAllocation(b *testing.B){
	l := New(InfoLevel)
	for i := 0; i < b.N; i ++{
		l.WithFields(map[string]interface{}{
			"test": "test",
		}).Info()
	}
}


func BenchmarkStructuredLogger(b *testing.B){
	l := NewStructuredLog(InfoLevel)
	test := New(DebugLevel)
	for i := 0; i < b.N; i ++ {
		l.OnDebug(func() {
			l.Debug("failed to fetch URL", l.DebugAny("test", test))
		})
	}
	l.Sync()
}