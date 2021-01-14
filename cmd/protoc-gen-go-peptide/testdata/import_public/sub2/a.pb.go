// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: cmd/protoc-gen-go-peptide/testdata/import_public/sub2/a.proto

package sub2

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

type Sub2Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Sub2Message) Reset() {
	*x = Sub2Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_protoc_gen_go_peptide_testdata_import_public_sub2_a_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sub2Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sub2Message) ProtoMessage() {}

func (x *Sub2Message) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_peptide_testdata_import_public_sub2_a_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sub2Message.ProtoReflect.Descriptor instead.
func (*Sub2Message) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_peptide_testdata_import_public_sub2_a_proto_rawDescGZIP(), []int{0}
}

var File_cmd_protoc_gen_go_peptide_testdata_import_public_sub2_a_proto protoreflect.FileDescriptor

var file_cmd_protoc_gen_go_peptide_testdata_import_public_sub2_a_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2d, 0x70, 0x65, 0x70, 0x74, 0x69, 0x64, 0x65, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x2f, 0x73, 0x75, 0x62, 0x32, 0x2f, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x21, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e,
	0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x73, 0x75,
	0x62, 0x32, 0x22, 0x0d, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x32, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x42, 0x50, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6d, 0x6f, 0x6e, 0x61, 0x78, 0x2f, 0x70, 0x65, 0x70, 0x74, 0x69, 0x64, 0x65, 0x2f, 0x63, 0x6d,
	0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d,
	0x70, 0x65, 0x70, 0x74, 0x69, 0x64, 0x65, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61,
	0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x73,
	0x75, 0x62, 0x32,
}

var (
	file_cmd_protoc_gen_go_peptide_testdata_import_public_sub2_a_proto_rawDescOnce sync.Once
	file_cmd_protoc_gen_go_peptide_testdata_import_public_sub2_a_proto_rawDescData = file_cmd_protoc_gen_go_peptide_testdata_import_public_sub2_a_proto_rawDesc
)

func file_cmd_protoc_gen_go_peptide_testdata_import_public_sub2_a_proto_rawDescGZIP() []byte {
	file_cmd_protoc_gen_go_peptide_testdata_import_public_sub2_a_proto_rawDescOnce.Do(func() {
		file_cmd_protoc_gen_go_peptide_testdata_import_public_sub2_a_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_protoc_gen_go_peptide_testdata_import_public_sub2_a_proto_rawDescData)
	})
	return file_cmd_protoc_gen_go_peptide_testdata_import_public_sub2_a_proto_rawDescData
}

var file_cmd_protoc_gen_go_peptide_testdata_import_public_sub2_a_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cmd_protoc_gen_go_peptide_testdata_import_public_sub2_a_proto_goTypes = []interface{}{
	(*Sub2Message)(nil), // 0: goproto.protoc.import_public.sub2.Sub2Message
}
var file_cmd_protoc_gen_go_peptide_testdata_import_public_sub2_a_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmd_protoc_gen_go_peptide_testdata_import_public_sub2_a_proto_init() }
func file_cmd_protoc_gen_go_peptide_testdata_import_public_sub2_a_proto_init() {
	if File_cmd_protoc_gen_go_peptide_testdata_import_public_sub2_a_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_protoc_gen_go_peptide_testdata_import_public_sub2_a_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sub2Message); i {
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
			RawDescriptor: file_cmd_protoc_gen_go_peptide_testdata_import_public_sub2_a_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_protoc_gen_go_peptide_testdata_import_public_sub2_a_proto_goTypes,
		DependencyIndexes: file_cmd_protoc_gen_go_peptide_testdata_import_public_sub2_a_proto_depIdxs,
		MessageInfos:      file_cmd_protoc_gen_go_peptide_testdata_import_public_sub2_a_proto_msgTypes,
	}.Build()
	File_cmd_protoc_gen_go_peptide_testdata_import_public_sub2_a_proto = out.File
	file_cmd_protoc_gen_go_peptide_testdata_import_public_sub2_a_proto_rawDesc = nil
	file_cmd_protoc_gen_go_peptide_testdata_import_public_sub2_a_proto_goTypes = nil
	file_cmd_protoc_gen_go_peptide_testdata_import_public_sub2_a_proto_depIdxs = nil
}
