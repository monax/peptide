package main

import (
	"testing"

	"github.com/monax/peptide/cmd/protoc-gen-go-peptide/testdata/gogo"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestGogo(t *testing.T) {
	t.Run("Test gogogo", func(t *testing.T) {
		msg := &gogo.TestMessage{
			WithMoreTags: "tagly",
			WithJsonTag:  false,
			Fruit:        gogo.Apples,
		}
		bs, err := proto.Marshal(msg)
		require.NoError(t, err)
		if err != nil {
			t.Fatal(err)
		}
		msgOut := new(gogo.TestMessage)
		err = proto.Unmarshal(bs, msgOut)
		require.NoError(t, err)
		require.True(t, proto.Equal(msg, msgOut))
	})
}
