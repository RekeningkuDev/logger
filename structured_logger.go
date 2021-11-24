package logger

import (
	"go.uber.org/zap"
)

type StructuredLog struct {
	*zap.Logger
	level Level
}


//type Field interface {
//	String(key string, val string) Field
//	Error(err error) Field
//	Any(key string, val interface{}) Field
//}
//
//var LogField = StructuredLogField{}
//
//type StructuredLogField struct {
//	zap.Field
//}

type Field struct {
	zap.Field
}

func (l *StructuredLog) DebugString(key string, val string) zap.Field {
	if l.level >= DebugLevel {
		return zap.String(key, val)
	}
	return zap.Skip()
}


func (l *StructuredLog) String(key string, val string) zap.Field {
	return zap.String(key, val)
}


func (l *StructuredLog) DebugAny(key string, val interface{}) zap.Field {
	if l.level >= DebugLevel {
		return zap.Any(key, val)
	}

	return zap.Skip()
}

func (l *StructuredLog) Any(key string, val interface{}) zap.Field {
	return zap.Any(key, val)
}


//func (s StructuredLogField) Error(err error) Field {
//	panic("implement me")
//}
//
//func (s StructuredLogField) Any(key string, val interface{}) Field {
//	panic("implement me")
//}

func NewStructuredLog(level Level) StructuredLog {
	if level >= DebugLevel {
		l, _ := zap.NewDevelopment()
		return StructuredLog{l, level}
	}
	l, _ := zap.NewProduction()
	return StructuredLog{l, level}
}
