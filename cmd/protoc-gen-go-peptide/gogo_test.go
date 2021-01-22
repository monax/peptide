package main

import (
	"reflect"
	"testing"

	"github.com/monax/peptide/cmd/protoc-gen-go-peptide/testdata/gogo"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestGogo(t *testing.T) {
	//roundtrip(t, &gogo.TestMessage{
	//	WithMoreTags: "tagly",
	//	WithJsonTag:  false,
	//	Fruit:        gogo.Apples,
	//})
	roundtrip(t, &gogo.CustomTypeMessage{
		Hash: gogo.Hash{1, 2, 3},
	})
}

func roundtrip(t *testing.T, msg proto.Message) {
	bs, err := proto.Marshal(msg)
	require.NoError(t, err)
	if err != nil {
		t.Fatal(err)
	}
	rt := reflect.TypeOf(msg).Elem()
	msgOut := reflect.New(rt).Interface().(proto.Message)
	err = proto.Unmarshal(bs, msgOut)
	require.NoError(t, err)
	require.True(t, proto.Equal(msg, msgOut))
}
