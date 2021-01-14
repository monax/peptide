// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: cmd/protoc-gen-go-peptide/testdata/proto3/fields.proto

package proto3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FieldTestMessage_Enum int32

const (
	FieldTestMessage_ZERO FieldTestMessage_Enum = 0
)

// Enum value maps for FieldTestMessage_Enum.
var (
	FieldTestMessage_Enum_name = map[int32]string{
		0: "ZERO",
	}
	FieldTestMessage_Enum_value = map[string]int32{
		"ZERO": 0,
	}
)

func (x FieldTestMessage_Enum) Enum() *FieldTestMessage_Enum {
	p := new(FieldTestMessage_Enum)
	*p = x
	return p
}

func (x FieldTestMessage_Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FieldTestMessage_Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_enumTypes[0].Descriptor()
}

func (FieldTestMessage_Enum) Type() protoreflect.EnumType {
	return &file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_enumTypes[0]
}

func (x FieldTestMessage_Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FieldTestMessage_Enum.Descriptor instead.
func (FieldTestMessage_Enum) EnumDescriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_rawDescGZIP(), []int{0, 0}
}

type FieldTestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OptionalBool     string                               `protobuf:"bytes,1,opt,name=optional_bool,json=optionalBool,proto3" json:"optional_bool,omitempty"`
	OptionalEnum     FieldTestMessage_Enum                `protobuf:"varint,2,opt,name=optional_enum,json=optionalEnum,proto3,enum=goproto.protoc.proto3.FieldTestMessage_Enum" json:"optional_enum,omitempty"`
	OptionalInt32    int32                                `protobuf:"varint,3,opt,name=optional_int32,json=optionalInt32,proto3" json:"optional_int32,omitempty"`
	OptionalSint32   int32                                `protobuf:"zigzag32,4,opt,name=optional_sint32,json=optionalSint32,proto3" json:"optional_sint32,omitempty"`
	OptionalUint32   uint32                               `protobuf:"varint,5,opt,name=optional_uint32,json=optionalUint32,proto3" json:"optional_uint32,omitempty"`
	OptionalInt64    int64                                `protobuf:"varint,6,opt,name=optional_int64,json=optionalInt64,proto3" json:"optional_int64,omitempty"`
	OptionalSint64   int64                                `protobuf:"zigzag64,7,opt,name=optional_sint64,json=optionalSint64,proto3" json:"optional_sint64,omitempty"`
	OptionalUint64   uint64                               `protobuf:"varint,8,opt,name=optional_uint64,json=optionalUint64,proto3" json:"optional_uint64,omitempty"`
	OptionalSfixed32 int32                                `protobuf:"fixed32,9,opt,name=optional_sfixed32,json=optionalSfixed32,proto3" json:"optional_sfixed32,omitempty"`
	OptionalFixed32  uint32                               `protobuf:"fixed32,10,opt,name=optional_fixed32,json=optionalFixed32,proto3" json:"optional_fixed32,omitempty"`
	OptionalFloat    float32                              `protobuf:"fixed32,11,opt,name=optional_float,json=optionalFloat,proto3" json:"optional_float,omitempty"`
	OptionalSfixed64 int64                                `protobuf:"fixed64,12,opt,name=optional_sfixed64,json=optionalSfixed64,proto3" json:"optional_sfixed64,omitempty"`
	OptionalFixed64  uint64                               `protobuf:"fixed64,13,opt,name=optional_fixed64,json=optionalFixed64,proto3" json:"optional_fixed64,omitempty"`
	OptionalDouble   float64                              `protobuf:"fixed64,14,opt,name=optional_double,json=optionalDouble,proto3" json:"optional_double,omitempty"`
	OptionalString   string                               `protobuf:"bytes,15,opt,name=optional_string,json=optionalString,proto3" json:"optional_string,omitempty"`
	OptionalBytes    []byte                               `protobuf:"bytes,16,opt,name=optional_bytes,json=optionalBytes,proto3" json:"optional_bytes,omitempty"`
	Optional_Message *FieldTestMessage_Message            `protobuf:"bytes,17,opt,name=optional_Message,json=optionalMessage,proto3" json:"optional_Message,omitempty"`
	RepeatedBool     []bool                               `protobuf:"varint,201,rep,packed,name=repeated_bool,json=repeatedBool,proto3" json:"repeated_bool,omitempty"`
	RepeatedEnum     []FieldTestMessage_Enum              `protobuf:"varint,202,rep,packed,name=repeated_enum,json=repeatedEnum,proto3,enum=goproto.protoc.proto3.FieldTestMessage_Enum" json:"repeated_enum,omitempty"`
	RepeatedInt32    []int32                              `protobuf:"varint,203,rep,packed,name=repeated_int32,json=repeatedInt32,proto3" json:"repeated_int32,omitempty"`
	RepeatedSint32   []int32                              `protobuf:"zigzag32,204,rep,packed,name=repeated_sint32,json=repeatedSint32,proto3" json:"repeated_sint32,omitempty"`
	RepeatedUint32   []uint32                             `protobuf:"varint,205,rep,packed,name=repeated_uint32,json=repeatedUint32,proto3" json:"repeated_uint32,omitempty"`
	RepeatedInt64    []int64                              `protobuf:"varint,206,rep,packed,name=repeated_int64,json=repeatedInt64,proto3" json:"repeated_int64,omitempty"`
	RepeatedSint64   []int64                              `protobuf:"zigzag64,207,rep,packed,name=repeated_sint64,json=repeatedSint64,proto3" json:"repeated_sint64,omitempty"`
	RepeatedUint64   []uint64                             `protobuf:"varint,208,rep,packed,name=repeated_uint64,json=repeatedUint64,proto3" json:"repeated_uint64,omitempty"`
	RepeatedSfixed32 []int32                              `protobuf:"fixed32,209,rep,packed,name=repeated_sfixed32,json=repeatedSfixed32,proto3" json:"repeated_sfixed32,omitempty"`
	RepeatedFixed32  []uint32                             `protobuf:"fixed32,210,rep,packed,name=repeated_fixed32,json=repeatedFixed32,proto3" json:"repeated_fixed32,omitempty"`
	RepeatedFloat    []float32                            `protobuf:"fixed32,211,rep,packed,name=repeated_float,json=repeatedFloat,proto3" json:"repeated_float,omitempty"`
	RepeatedSfixed64 []int64                              `protobuf:"fixed64,212,rep,packed,name=repeated_sfixed64,json=repeatedSfixed64,proto3" json:"repeated_sfixed64,omitempty"`
	RepeatedFixed64  []uint64                             `protobuf:"fixed64,213,rep,packed,name=repeated_fixed64,json=repeatedFixed64,proto3" json:"repeated_fixed64,omitempty"`
	RepeatedDouble   []float64                            `protobuf:"fixed64,214,rep,packed,name=repeated_double,json=repeatedDouble,proto3" json:"repeated_double,omitempty"`
	RepeatedString   []string                             `protobuf:"bytes,215,rep,name=repeated_string,json=repeatedString,proto3" json:"repeated_string,omitempty"`
	RepeatedBytes    [][]byte                             `protobuf:"bytes,216,rep,name=repeated_bytes,json=repeatedBytes,proto3" json:"repeated_bytes,omitempty"`
	Repeated_Message []*FieldTestMessage_Message          `protobuf:"bytes,217,rep,name=repeated_Message,json=repeatedMessage,proto3" json:"repeated_Message,omitempty"`
	MapInt32Int64    map[int32]int64                      `protobuf:"bytes,500,rep,name=map_int32_int64,json=mapInt32Int64,proto3" json:"map_int32_int64,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	MapStringMessage map[string]*FieldTestMessage_Message `protobuf:"bytes,501,rep,name=map_string_message,json=mapStringMessage,proto3" json:"map_string_message,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MapFixed64Enum   map[uint64]FieldTestMessage_Enum     `protobuf:"bytes,502,rep,name=map_fixed64_enum,json=mapFixed64Enum,proto3" json:"map_fixed64_enum,omitempty" protobuf_key:"fixed64,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=goproto.protoc.proto3.FieldTestMessage_Enum"`
}

func (x *FieldTestMessage) Reset() {
	*x = FieldTestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldTestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldTestMessage) ProtoMessage() {}

func (x *FieldTestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldTestMessage.ProtoReflect.Descriptor instead.
func (*FieldTestMessage) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_rawDescGZIP(), []int{0}
}

func (x *FieldTestMessage) GetOptionalBool() string {
	if x != nil {
		return x.OptionalBool
	}
	return ""
}

func (x *FieldTestMessage) GetOptionalEnum() FieldTestMessage_Enum {
	if x != nil {
		return x.OptionalEnum
	}
	return FieldTestMessage_ZERO
}

func (x *FieldTestMessage) GetOptionalInt32() int32 {
	if x != nil {
		return x.OptionalInt32
	}
	return 0
}

func (x *FieldTestMessage) GetOptionalSint32() int32 {
	if x != nil {
		return x.OptionalSint32
	}
	return 0
}

func (x *FieldTestMessage) GetOptionalUint32() uint32 {
	if x != nil {
		return x.OptionalUint32
	}
	return 0
}

func (x *FieldTestMessage) GetOptionalInt64() int64 {
	if x != nil {
		return x.OptionalInt64
	}
	return 0
}

func (x *FieldTestMessage) GetOptionalSint64() int64 {
	if x != nil {
		return x.OptionalSint64
	}
	return 0
}

func (x *FieldTestMessage) GetOptionalUint64() uint64 {
	if x != nil {
		return x.OptionalUint64
	}
	return 0
}

func (x *FieldTestMessage) GetOptionalSfixed32() int32 {
	if x != nil {
		return x.OptionalSfixed32
	}
	return 0
}

func (x *FieldTestMessage) GetOptionalFixed32() uint32 {
	if x != nil {
		return x.OptionalFixed32
	}
	return 0
}

func (x *FieldTestMessage) GetOptionalFloat() float32 {
	if x != nil {
		return x.OptionalFloat
	}
	return 0
}

func (x *FieldTestMessage) GetOptionalSfixed64() int64 {
	if x != nil {
		return x.OptionalSfixed64
	}
	return 0
}

func (x *FieldTestMessage) GetOptionalFixed64() uint64 {
	if x != nil {
		return x.OptionalFixed64
	}
	return 0
}

func (x *FieldTestMessage) GetOptionalDouble() float64 {
	if x != nil {
		return x.OptionalDouble
	}
	return 0
}

func (x *FieldTestMessage) GetOptionalString() string {
	if x != nil {
		return x.OptionalString
	}
	return ""
}

func (x *FieldTestMessage) GetOptionalBytes() []byte {
	if x != nil {
		return x.OptionalBytes
	}
	return nil
}

func (x *FieldTestMessage) GetOptional_Message() *FieldTestMessage_Message {
	if x != nil {
		return x.Optional_Message
	}
	return nil
}

func (x *FieldTestMessage) GetRepeatedBool() []bool {
	if x != nil {
		return x.RepeatedBool
	}
	return nil
}

func (x *FieldTestMessage) GetRepeatedEnum() []FieldTestMessage_Enum {
	if x != nil {
		return x.RepeatedEnum
	}
	return nil
}

func (x *FieldTestMessage) GetRepeatedInt32() []int32 {
	if x != nil {
		return x.RepeatedInt32
	}
	return nil
}

func (x *FieldTestMessage) GetRepeatedSint32() []int32 {
	if x != nil {
		return x.RepeatedSint32
	}
	return nil
}

func (x *FieldTestMessage) GetRepeatedUint32() []uint32 {
	if x != nil {
		return x.RepeatedUint32
	}
	return nil
}

func (x *FieldTestMessage) GetRepeatedInt64() []int64 {
	if x != nil {
		return x.RepeatedInt64
	}
	return nil
}

func (x *FieldTestMessage) GetRepeatedSint64() []int64 {
	if x != nil {
		return x.RepeatedSint64
	}
	return nil
}

func (x *FieldTestMessage) GetRepeatedUint64() []uint64 {
	if x != nil {
		return x.RepeatedUint64
	}
	return nil
}

func (x *FieldTestMessage) GetRepeatedSfixed32() []int32 {
	if x != nil {
		return x.RepeatedSfixed32
	}
	return nil
}

func (x *FieldTestMessage) GetRepeatedFixed32() []uint32 {
	if x != nil {
		return x.RepeatedFixed32
	}
	return nil
}

func (x *FieldTestMessage) GetRepeatedFloat() []float32 {
	if x != nil {
		return x.RepeatedFloat
	}
	return nil
}

func (x *FieldTestMessage) GetRepeatedSfixed64() []int64 {
	if x != nil {
		return x.RepeatedSfixed64
	}
	return nil
}

func (x *FieldTestMessage) GetRepeatedFixed64() []uint64 {
	if x != nil {
		return x.RepeatedFixed64
	}
	return nil
}

func (x *FieldTestMessage) GetRepeatedDouble() []float64 {
	if x != nil {
		return x.RepeatedDouble
	}
	return nil
}

func (x *FieldTestMessage) GetRepeatedString() []string {
	if x != nil {
		return x.RepeatedString
	}
	return nil
}

func (x *FieldTestMessage) GetRepeatedBytes() [][]byte {
	if x != nil {
		return x.RepeatedBytes
	}
	return nil
}

func (x *FieldTestMessage) GetRepeated_Message() []*FieldTestMessage_Message {
	if x != nil {
		return x.Repeated_Message
	}
	return nil
}

func (x *FieldTestMessage) GetMapInt32Int64() map[int32]int64 {
	if x != nil {
		return x.MapInt32Int64
	}
	return nil
}

func (x *FieldTestMessage) GetMapStringMessage() map[string]*FieldTestMessage_Message {
	if x != nil {
		return x.MapStringMessage
	}
	return nil
}

func (x *FieldTestMessage) GetMapFixed64Enum() map[uint64]FieldTestMessage_Enum {
	if x != nil {
		return x.MapFixed64Enum
	}
	return nil
}

type FieldTestMessage_Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FieldTestMessage_Message) Reset() {
	*x = FieldTestMessage_Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldTestMessage_Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldTestMessage_Message) ProtoMessage() {}

func (x *FieldTestMessage_Message) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldTestMessage_Message.ProtoReflect.Descriptor instead.
func (*FieldTestMessage_Message) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_rawDescGZIP(), []int{0, 3}
}

var File_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto protoreflect.FileDescriptor

var file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_rawDesc = []byte{
	0x0a, 0x36, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2d, 0x70, 0x65, 0x70, 0x74, 0x69, 0x64, 0x65, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x64, 0x61, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x22,
	0xd0, 0x11, 0x0a, 0x10, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x51, 0x0a, 0x0d, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x65,
	0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0c,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x25, 0x0a, 0x0e,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f,
	0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x11, 0x52, 0x0e, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x27, 0x0a, 0x0f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x55,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x27, 0x0a, 0x0f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x12, 0x52, 0x0e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x2b,
	0x0a, 0x11, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x33, 0x32, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x10, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x29, 0x0a, 0x10, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x07, 0x52, 0x0f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x46,
	0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x2b, 0x0a,
	0x11, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64,
	0x36, 0x34, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x10, 0x52, 0x10, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x29, 0x0a, 0x10, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x06, 0x52, 0x0f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x46, 0x69,
	0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x27,
	0x0a, 0x0f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x5a,
	0x0a, 0x10, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0xc9, 0x01, 0x20, 0x03,
	0x28, 0x08, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x6f, 0x6f, 0x6c,
	0x12, 0x52, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x75,
	0x6d, 0x18, 0xca, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x45, 0x6e, 0x75, 0x6d, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0xcb, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0d, 0x72,
	0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x28, 0x0a, 0x0f,
	0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18,
	0xcc, 0x01, 0x20, 0x03, 0x28, 0x11, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x53, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0xcd, 0x01, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x18, 0xce, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0xcf, 0x01, 0x20, 0x03,
	0x28, 0x12, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x75,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0xd0, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0e, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x2c, 0x0a, 0x11,
	0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33,
	0x32, 0x18, 0xd1, 0x01, 0x20, 0x03, 0x28, 0x0f, 0x52, 0x10, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0xd2,
	0x01, 0x20, 0x03, 0x28, 0x07, 0x52, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x46,
	0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0xd3, 0x01, 0x20, 0x03, 0x28, 0x02, 0x52,
	0x0d, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x2c,
	0x0a, 0x11, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x36, 0x34, 0x18, 0xd4, 0x01, 0x20, 0x03, 0x28, 0x10, 0x52, 0x10, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x2a, 0x0a, 0x10,
	0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34,
	0x18, 0xd5, 0x01, 0x20, 0x03, 0x28, 0x06, 0x52, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0xd6, 0x01, 0x20, 0x03,
	0x28, 0x01, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x44, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0xd7, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x0a, 0x0e,
	0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0xd8,
	0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x12, 0x5b, 0x0a, 0x10, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0xd9, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2f, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x65, 0x73,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x63, 0x0a, 0x0f, 0x6d, 0x61, 0x70, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x18, 0xf4, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x49, 0x6e, 0x74,
	0x36, 0x34, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x6d, 0x61, 0x70, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x6c, 0x0a, 0x12, 0x6d, 0x61, 0x70, 0x5f, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0xf5, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x61, 0x70,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x10, 0x6d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x66, 0x0a, 0x10, 0x6d, 0x61, 0x70, 0x5f, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x36, 0x34, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0xf6, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x3b, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x65, 0x73,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x46, 0x69, 0x78, 0x65,
	0x64, 0x36, 0x34, 0x45, 0x6e, 0x75, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x6d, 0x61,
	0x70, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x45, 0x6e, 0x75, 0x6d, 0x1a, 0x40, 0x0a, 0x12,
	0x4d, 0x61, 0x70, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x74,
	0x0a, 0x15, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x45, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x6f, 0x0a, 0x13, 0x4d, 0x61, 0x70, 0x46, 0x69, 0x78, 0x65, 0x64,
	0x36, 0x34, 0x45, 0x6e, 0x75, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x06, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x42, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x67,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x09, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x10, 0x0a, 0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x08, 0x0a, 0x04, 0x5a, 0x45, 0x52, 0x4f,
	0x10, 0x00, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x6f, 0x6e, 0x61, 0x78, 0x2f, 0x70, 0x65, 0x70, 0x74, 0x69, 0x64, 0x65, 0x2f, 0x63,
	0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f,
	0x2d, 0x70, 0x65, 0x70, 0x74, 0x69, 0x64, 0x65, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74,
	0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_rawDescOnce sync.Once
	file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_rawDescData = file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_rawDesc
)

func file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_rawDescGZIP() []byte {
	file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_rawDescOnce.Do(func() {
		file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_rawDescData)
	})
	return file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_rawDescData
}

var file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_goTypes = []interface{}{
	(FieldTestMessage_Enum)(0),       // 0: goproto.protoc.proto3.FieldTestMessage.Enum
	(*FieldTestMessage)(nil),         // 1: goproto.protoc.proto3.FieldTestMessage
	nil,                              // 2: goproto.protoc.proto3.FieldTestMessage.MapInt32Int64Entry
	nil,                              // 3: goproto.protoc.proto3.FieldTestMessage.MapStringMessageEntry
	nil,                              // 4: goproto.protoc.proto3.FieldTestMessage.MapFixed64EnumEntry
	(*FieldTestMessage_Message)(nil), // 5: goproto.protoc.proto3.FieldTestMessage.Message
}
var file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_depIdxs = []int32{
	0, // 0: goproto.protoc.proto3.FieldTestMessage.optional_enum:type_name -> goproto.protoc.proto3.FieldTestMessage.Enum
	5, // 1: goproto.protoc.proto3.FieldTestMessage.optional_Message:type_name -> goproto.protoc.proto3.FieldTestMessage.Message
	0, // 2: goproto.protoc.proto3.FieldTestMessage.repeated_enum:type_name -> goproto.protoc.proto3.FieldTestMessage.Enum
	5, // 3: goproto.protoc.proto3.FieldTestMessage.repeated_Message:type_name -> goproto.protoc.proto3.FieldTestMessage.Message
	2, // 4: goproto.protoc.proto3.FieldTestMessage.map_int32_int64:type_name -> goproto.protoc.proto3.FieldTestMessage.MapInt32Int64Entry
	3, // 5: goproto.protoc.proto3.FieldTestMessage.map_string_message:type_name -> goproto.protoc.proto3.FieldTestMessage.MapStringMessageEntry
	4, // 6: goproto.protoc.proto3.FieldTestMessage.map_fixed64_enum:type_name -> goproto.protoc.proto3.FieldTestMessage.MapFixed64EnumEntry
	5, // 7: goproto.protoc.proto3.FieldTestMessage.MapStringMessageEntry.value:type_name -> goproto.protoc.proto3.FieldTestMessage.Message
	0, // 8: goproto.protoc.proto3.FieldTestMessage.MapFixed64EnumEntry.value:type_name -> goproto.protoc.proto3.FieldTestMessage.Enum
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_init() }
func file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_init() {
	if File_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldTestMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldTestMessage_Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_goTypes,
		DependencyIndexes: file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_depIdxs,
		EnumInfos:         file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_enumTypes,
		MessageInfos:      file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_msgTypes,
	}.Build()
	File_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto = out.File
	file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_rawDesc = nil
	file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_goTypes = nil
	file_cmd_protoc_gen_go_peptide_testdata_proto3_fields_proto_depIdxs = nil
}
