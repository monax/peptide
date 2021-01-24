package main

import (
	"fmt"
	"google.golang.org/protobuf/types/descriptorpb"
	"io/ioutil"
	"path"
	"runtime"
	"testing"

	"github.com/monax/peptide/cmd/protoc-gen-go-peptide/gengo"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

//go:generate protoc -I ../../types -I ../../ cmd/protoc-gen-go-peptide/testdata/gogo/customtype.proto --include_imports -o customtype.pb

const customtypeFile = "customtype.pb"

func TestGengo(t *testing.T) {
	fmt.Println(sourceFileDir())
	bs, err := ioutil.ReadFile(path.Join(sourceFileDir(), customtypeFile))
	require.NoError(t, err)

	set := new(descriptorpb.FileDescriptorSet)
	err = proto.Unmarshal(bs, set)
	require.NoError(t, err)

	files := make([]string, len(set.File))
	for i, f := range set.File {
		files[i] = *f.Name
	}
	gen, err := protogen.Options{}.New(&pluginpb.CodeGeneratorRequest{
		FileToGenerate: files,
		ProtoFile:      set.File,
	})
	require.NoError(t, err)
	genFile := gengo.GenerateFile(gen, gen.Files[0])
	bs, err = genFile.Content()
	require.NoError(t, err)
	fmt.Println(string(bs))

}

// Get the directory of source file of the caller
func sourceFileDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Dir(filename)
}
