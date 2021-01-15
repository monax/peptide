// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func TestRegistry(t *testing.T) {
	var hasFiles bool
	protoregistry.GlobalFiles.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		hasFiles = true
		return true
	})
	if !hasFiles {
		t.Errorf("protoregistry.GlobalFiles is empty")
	}
}
