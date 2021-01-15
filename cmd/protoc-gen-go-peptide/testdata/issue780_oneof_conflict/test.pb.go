// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go-peptide. DO NOT EDIT. BUMssss
// source: cmd/protoc-gen-go-peptide/testdata/issue780_oneof_conflict/test.proto

package issue780_oneof_conflict

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type Foo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Bar:
	//	*Foo_GetBar
	Bar isFoo_Bar `protobuf_oneof:"bar"`
}

func (x *Foo) Reset() {
	*x = Foo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_protoc_gen_go_peptide_testdata_issue780_oneof_conflict_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Foo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Foo) ProtoMessage() {}

func (x *Foo) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_peptide_testdata_issue780_oneof_conflict_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Foo.ProtoReflect.Descriptor instead.
func (*Foo) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_peptide_testdata_issue780_oneof_conflict_test_proto_rawDescGZIP(), []int{0}
}

func (m *Foo) GetBar() isFoo_Bar {
	if m != nil {
		return m.Bar
	}
	return nil
}

func (x *Foo) GetGetBar() string {
	if x, ok := x.GetBar().(*Foo_GetBar); ok {
		return x.GetBar
	}
	return ""
}

type isFoo_Bar interface {
	isFoo_Bar()
}

type Foo_GetBar struct {
	GetBar string `protobuf:"bytes,1,opt,name=get_bar,json=getBar,oneof"`
}

func (*Foo_GetBar) isFoo_Bar() {}

var File_cmd_protoc_gen_go_peptide_testdata_issue780_oneof_conflict_test_proto protoreflect.FileDescriptor

var file_cmd_protoc_gen_go_peptide_testdata_issue780_oneof_conflict_test_proto_rawDesc = []byte{
	0x0a, 0x45, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2d, 0x70, 0x65, 0x70, 0x74, 0x69, 0x64, 0x65, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x37, 0x38, 0x30, 0x5f, 0x6f, 0x6e,
	0x65, 0x6f, 0x66, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x74, 0x65,
	0x73, 0x74, 0x22, 0x27, 0x0a, 0x03, 0x46, 0x6f, 0x6f, 0x12, 0x19, 0x0a, 0x07, 0x67, 0x65, 0x74,
	0x5f, 0x62, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x67, 0x65,
	0x74, 0x42, 0x61, 0x72, 0x42, 0x05, 0x0a, 0x03, 0x62, 0x61, 0x72, 0x42, 0x55, 0x5a, 0x53, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x6e, 0x61, 0x78, 0x2f,
	0x70, 0x65, 0x70, 0x74, 0x69, 0x64, 0x65, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x70, 0x65, 0x70, 0x74, 0x69, 0x64,
	0x65, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x37, 0x38, 0x30, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x69,
	0x63, 0x74,
}

var (
	file_cmd_protoc_gen_go_peptide_testdata_issue780_oneof_conflict_test_proto_rawDescOnce sync.Once
	file_cmd_protoc_gen_go_peptide_testdata_issue780_oneof_conflict_test_proto_rawDescData = file_cmd_protoc_gen_go_peptide_testdata_issue780_oneof_conflict_test_proto_rawDesc
)

func file_cmd_protoc_gen_go_peptide_testdata_issue780_oneof_conflict_test_proto_rawDescGZIP() []byte {
	file_cmd_protoc_gen_go_peptide_testdata_issue780_oneof_conflict_test_proto_rawDescOnce.Do(func() {
		file_cmd_protoc_gen_go_peptide_testdata_issue780_oneof_conflict_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_protoc_gen_go_peptide_testdata_issue780_oneof_conflict_test_proto_rawDescData)
	})
	return file_cmd_protoc_gen_go_peptide_testdata_issue780_oneof_conflict_test_proto_rawDescData
}

var file_cmd_protoc_gen_go_peptide_testdata_issue780_oneof_conflict_test_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cmd_protoc_gen_go_peptide_testdata_issue780_oneof_conflict_test_proto_goTypes = []interface{}{
	(*Foo)(nil), // 0: oneoftest.Foo
}
var file_cmd_protoc_gen_go_peptide_testdata_issue780_oneof_conflict_test_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmd_protoc_gen_go_peptide_testdata_issue780_oneof_conflict_test_proto_init() }
func file_cmd_protoc_gen_go_peptide_testdata_issue780_oneof_conflict_test_proto_init() {
	if File_cmd_protoc_gen_go_peptide_testdata_issue780_oneof_conflict_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_protoc_gen_go_peptide_testdata_issue780_oneof_conflict_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Foo); i {
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
	file_cmd_protoc_gen_go_peptide_testdata_issue780_oneof_conflict_test_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Foo_GetBar)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_protoc_gen_go_peptide_testdata_issue780_oneof_conflict_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_protoc_gen_go_peptide_testdata_issue780_oneof_conflict_test_proto_goTypes,
		DependencyIndexes: file_cmd_protoc_gen_go_peptide_testdata_issue780_oneof_conflict_test_proto_depIdxs,
		MessageInfos:      file_cmd_protoc_gen_go_peptide_testdata_issue780_oneof_conflict_test_proto_msgTypes,
	}.Build()
	File_cmd_protoc_gen_go_peptide_testdata_issue780_oneof_conflict_test_proto = out.File
	file_cmd_protoc_gen_go_peptide_testdata_issue780_oneof_conflict_test_proto_rawDesc = nil
	file_cmd_protoc_gen_go_peptide_testdata_issue780_oneof_conflict_test_proto_goTypes = nil
	file_cmd_protoc_gen_go_peptide_testdata_issue780_oneof_conflict_test_proto_depIdxs = nil
}
