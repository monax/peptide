// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go-peptide. DO NOT EDIT.
// source: cmd/protoc-gen-go-peptide/testdata/gogo/customtype.proto

package gogo

import (
	protoimpl "github.com/monax/peptide/protoimpl"
	_ "github.com/monax/peptide/types/gogoproto"
	peptide "github.com/monax/peptide/types/peptide"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	reflect "reflect"
	sync "sync"
)

type CustomTypeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	peptide.NoopExtender

	Hash  Hash  `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	Count int64 `protobuf:"varint,2,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (x *CustomTypeMessage) Reset() {
	*x = CustomTypeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_protoc_gen_go_peptide_testdata_gogo_customtype_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomTypeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomTypeMessage) ProtoMessage() {}

func (x *CustomTypeMessage) ProtoReflect() protoreflect.Message {
	return x.ProtoExtend(x.protoReflect())
}

func (x *CustomTypeMessage) protoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_peptide_testdata_gogo_customtype_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomTypeMessage.ProtoReflect.Descriptor instead.
func (*CustomTypeMessage) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_peptide_testdata_gogo_customtype_proto_rawDescGZIP(), []int{0}
}

func (x *CustomTypeMessage) GetHash() Hash {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *CustomTypeMessage) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_cmd_protoc_gen_go_peptide_testdata_gogo_customtype_proto protoreflect.FileDescriptor

var file_cmd_protoc_gen_go_peptide_testdata_gogo_customtype_proto_rawDesc = []byte{
	0x0a, 0x38, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2d, 0x70, 0x65, 0x70, 0x74, 0x69, 0x64, 0x65, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x64, 0x61, 0x74, 0x61, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x65, 0x70, 0x74,
	0x69, 0x64, 0x65, 0x2e, 0x67, 0x6f, 0x67, 0x6f, 0x67, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x88, 0x01, 0x0a, 0x11, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x5d, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x42, 0x49, 0xfa, 0xde, 0x1f, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x6e, 0x61, 0x78, 0x2f, 0x70, 0x65, 0x70, 0x74, 0x69,
	0x64, 0x65, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x70, 0x65, 0x70, 0x74, 0x69, 0x64, 0x65, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x52,
	0x04, 0x48, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x42, 0x5a, 0x40, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x6e, 0x61, 0x78, 0x2f,
	0x70, 0x65, 0x70, 0x74, 0x69, 0x64, 0x65, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x70, 0x65, 0x70, 0x74, 0x69, 0x64,
	0x65, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_protoc_gen_go_peptide_testdata_gogo_customtype_proto_rawDescOnce sync.Once
	file_cmd_protoc_gen_go_peptide_testdata_gogo_customtype_proto_rawDescData = file_cmd_protoc_gen_go_peptide_testdata_gogo_customtype_proto_rawDesc
)

func file_cmd_protoc_gen_go_peptide_testdata_gogo_customtype_proto_rawDescGZIP() []byte {
	file_cmd_protoc_gen_go_peptide_testdata_gogo_customtype_proto_rawDescOnce.Do(func() {
		file_cmd_protoc_gen_go_peptide_testdata_gogo_customtype_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_protoc_gen_go_peptide_testdata_gogo_customtype_proto_rawDescData)
	})
	return file_cmd_protoc_gen_go_peptide_testdata_gogo_customtype_proto_rawDescData
}

var file_cmd_protoc_gen_go_peptide_testdata_gogo_customtype_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cmd_protoc_gen_go_peptide_testdata_gogo_customtype_proto_goTypes = []interface{}{
	(*CustomTypeMessage)(nil), // 0: peptide.gogogo.CustomTypeMessage
}
var file_cmd_protoc_gen_go_peptide_testdata_gogo_customtype_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmd_protoc_gen_go_peptide_testdata_gogo_customtype_proto_init() }
func file_cmd_protoc_gen_go_peptide_testdata_gogo_customtype_proto_init() {
	if File_cmd_protoc_gen_go_peptide_testdata_gogo_customtype_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_protoc_gen_go_peptide_testdata_gogo_customtype_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomTypeMessage); i {
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
			RawDescriptor: file_cmd_protoc_gen_go_peptide_testdata_gogo_customtype_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_protoc_gen_go_peptide_testdata_gogo_customtype_proto_goTypes,
		DependencyIndexes: file_cmd_protoc_gen_go_peptide_testdata_gogo_customtype_proto_depIdxs,
		MessageInfos:      file_cmd_protoc_gen_go_peptide_testdata_gogo_customtype_proto_msgTypes,
	}.Build()
	File_cmd_protoc_gen_go_peptide_testdata_gogo_customtype_proto = out.File
	file_cmd_protoc_gen_go_peptide_testdata_gogo_customtype_proto_rawDesc = nil
	file_cmd_protoc_gen_go_peptide_testdata_gogo_customtype_proto_goTypes = nil
	file_cmd_protoc_gen_go_peptide_testdata_gogo_customtype_proto_depIdxs = nil
}