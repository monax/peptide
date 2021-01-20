package gogo

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoiface"
)

type Hash []byte

//func (h Hash) PBValueOf(value reflect.Value) protoreflect.Value {
//	panic("implement me")
//}
//
//func (h Hash) GoValueOf(value protoreflect.Value) reflect.Value {
//	panic("implement me")
//}
//
//func (h Hash) IsValidPB(value protoreflect.Value) bool {
//	panic("implement me")
//}
//
//func (h Hash) IsValidGo(value reflect.Value) bool {
//	panic("implement me")
//}
//
//func (h Hash) New() protoreflect.Value {
//	panic("implement me")
//}
//
//func (h Hash) Zero() protoreflect.Value {
//	panic("implement me")
//}

// Extender

func (x *CustomTypeMessage) ProtoExtend(message protoreflect.Message) protoreflect.Message {
	return &customTypeMessage{message}
}

type customTypeMessage struct {
	protoreflect.Message
}

func (c customTypeMessage) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	c.Message.Range(f)
}

func (c customTypeMessage) ProtoMethods() *protoiface.Methods {
	return nil
}

// Converter
