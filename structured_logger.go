package logger

import (
	"go.uber.org/zap"
)

type StructuredLog struct {
	*zap.Logger
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

func String(key string, val string) zap.Field {
	return zap.String(key, val)
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
		return StructuredLog{l}
	}
	l, _ := zap.NewProduction()
	return StructuredLog{l}
}
