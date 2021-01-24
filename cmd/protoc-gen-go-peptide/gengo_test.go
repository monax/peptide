package main

import (
	"fmt"
	"google.golang.org/protobuf/types/descriptorpb"
	"io/ioutil"
	"path"
	"runtime"
	"strings"
	"testing"

	"github.com/monax/peptide/cmd/protoc-gen-go-peptide/gengo"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

//go:generate protoc -I ../../types -I ../../ cmd/protoc-gen-go-peptide/testdata/gogo/customtype.proto --include_imports -o customtype.pb.bin

const customtypeFile = "customtype.pb.bin"

func TestGengo(t *testing.T) {
	bs, err := ioutil.ReadFile(path.Join(sourceFileDir(), customtypeFile))
	require.NoError(t, err)

	set := new(descriptorpb.FileDescriptorSet)
	err = proto.Unmarshal(bs, set)
	require.NoError(t, err)

	filenames := make([]string, len(set.File))
	for i, f := range set.File {
		filenames[i] = *f.Name
	}
	gen, err := protogen.Options{}.New(&pluginpb.CodeGeneratorRequest{
		FileToGenerate: filenames,
		ProtoFile:      set.File,
	})
	require.NoError(t, err)
	var file *protogen.File
	for _, f := range gen.Files {
		if strings.HasSuffix(*f.Proto.Name, "customtype.proto") {
			file = f
			break
		}
	}
	genFile := gengo.GenerateFile(gen, file)
	bs, err = genFile.Content()
	require.NoError(t, err)
	fmt.Println(string(bs))
}

type foopb struct {
	Address []byte
}

type foo struct {
	foopb
	Address [10]byte
}

// Get the directory of source file of the caller
func sourceFileDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Dir(filename)
}
