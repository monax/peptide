package peptide

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

// The extend interface is intended to be implemented on a generated type from non-generated code in the same package.
//
// Generators should wrap their existing ProtoReflect() function with a call to ProtoExtend(), like:
//
//  func (x *TestMessage) ProtoReflect() protoreflect.Message {
//  	return x.ProtoExtend(x.protoReflect())
//  }
//
// Where `TestMessage` is a generated protobuf message and `protoReflect()` is the 'original' function implementing
// proto.Message.
//
// This allows the default ProtoMessage to be extended without reimplementing the entire interface
type Extender interface {
	ProtoExtend(message protoreflect.Message) protoreflect.Message
}

// This struct implements Extender and is intended to be embedded into a generated message type by the generator.
// If other code implements ProtoExtend() directly on the generator struct then that implementation will be run, otherwise
// the behaviour is unchanged from the generated code
type NoopExtender struct {
}

func (_ NoopExtender) ProtoExtend(message protoreflect.Message) protoreflect.Message {
	//
	return message
}
