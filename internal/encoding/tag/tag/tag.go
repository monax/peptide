// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tag marshals and unmarshals the legacy struct tags as generated
// by historical versions of protoc-gen-go.
package tag

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	pref "google.golang.org/protobuf/reflect/protoreflect"
)

var byteType = reflect.TypeOf(byte(0))

// Format is the serialization format used to represent the default value.
type Format int

const (
	_ Format = iota

	// Descriptor uses the serialization format that protoc uses with the
	// google.protobuf.FieldDescriptorProto.default_value field.
	Descriptor

	// GoTag uses the historical serialization format in Go struct field tags.
	GoTag
)


// Marshal encodes the protoreflect.FieldDescriptor as a tag.
//
// The enumName must be provided if the kind is an enum.
// Historically, the formulation of the enum "name" was the proto package
// dot-concatenated with the generated Go identifier for the enum type.
// Depending on the context on how Marshal is called, there are different ways
// through which that information is determined. As such it is the caller's
// responsibility to provide a function to obtain that information.
func Marshal(fd pref.FieldDescriptor, enumName string) string {
	var tag []string
	switch fd.Kind() {
	case pref.BoolKind, pref.EnumKind, pref.Int32Kind, pref.Uint32Kind, pref.Int64Kind, pref.Uint64Kind:
		tag = append(tag, "varint")
	case pref.Sint32Kind:
		tag = append(tag, "zigzag32")
	case pref.Sint64Kind:
		tag = append(tag, "zigzag64")
	case pref.Sfixed32Kind, pref.Fixed32Kind, pref.FloatKind:
		tag = append(tag, "fixed32")
	case pref.Sfixed64Kind, pref.Fixed64Kind, pref.DoubleKind:
		tag = append(tag, "fixed64")
	case pref.StringKind, pref.BytesKind, pref.MessageKind:
		tag = append(tag, "bytes")
	case pref.GroupKind:
		tag = append(tag, "group")
	}
	tag = append(tag, strconv.Itoa(int(fd.Number())))
	switch fd.Cardinality() {
	case pref.Optional:
		tag = append(tag, "opt")
	case pref.Required:
		tag = append(tag, "req")
	case pref.Repeated:
		tag = append(tag, "rep")
	}
	if fd.IsPacked() {
		tag = append(tag, "packed")
	}
	name := string(fd.Name())
	if fd.Kind() == pref.GroupKind {
		// The name of the FieldDescriptor for a group field is
		// lowercased. To find the original capitalization, we
		// look in the field's MessageType.
		name = string(fd.Message().Name())
	}
	tag = append(tag, "name="+name)
	if jsonName := fd.JSONName(); jsonName != "" && jsonName != name && !fd.IsExtension() {
		// NOTE: The jsonName != name condition is suspect, but it preserve
		// the exact same semantics from the previous generator.
		tag = append(tag, "json="+jsonName)
	}
	if fd.IsWeak() {
		tag = append(tag, "weak="+string(fd.Message().FullName()))
	}
	// The previous implementation does not tag extension fields as proto3,
	// even when the field is defined in a proto3 file. Match that behavior
	// for consistency.
	if fd.Syntax() == pref.Proto3 && !fd.IsExtension() {
		tag = append(tag, "proto3")
	}
	if fd.Kind() == pref.EnumKind && enumName != "" {
		tag = append(tag, "enum="+enumName)
	}
	if fd.ContainingOneof() != nil {
		tag = append(tag, "oneof")
	}
	// This must appear last in the tag, since commas in strings aren't escaped.
	if fd.HasDefault() {
		def, _ := defaultMarshal(fd.Default(), fd.DefaultEnumValue(), fd.Kind(), GoTag)
		tag = append(tag, "def="+def)
	}
	return strings.Join(tag, ",")
}

// Marshal serializes v as the default string according to the given kind k.
// When specifying the Descriptor format for an enum kind, the associated
// enum value descriptor must be provided.
func defaultMarshal(v pref.Value, ev pref.EnumValueDescriptor, k pref.Kind, f Format) (string, error) {
	switch k {
	case pref.BoolKind:
		if f == GoTag {
			if v.Bool() {
				return "1", nil
			} else {
				return "0", nil
			}
		} else {
			if v.Bool() {
				return "true", nil
			} else {
				return "false", nil
			}
		}
	case pref.EnumKind:
		if f == GoTag {
			return strconv.FormatInt(int64(v.Enum()), 10), nil
		} else {
			return string(ev.Name()), nil
		}
	case pref.Int32Kind, pref.Sint32Kind, pref.Sfixed32Kind, pref.Int64Kind, pref.Sint64Kind, pref.Sfixed64Kind:
		return strconv.FormatInt(v.Int(), 10), nil
	case pref.Uint32Kind, pref.Fixed32Kind, pref.Uint64Kind, pref.Fixed64Kind:
		return strconv.FormatUint(v.Uint(), 10), nil
	case pref.FloatKind, pref.DoubleKind:
		f := v.Float()
		switch {
		case math.IsInf(f, -1):
			return "-inf", nil
		case math.IsInf(f, +1):
			return "inf", nil
		case math.IsNaN(f):
			return "nan", nil
		default:
			if k == pref.FloatKind {
				return strconv.FormatFloat(f, 'g', -1, 32), nil
			} else {
				return strconv.FormatFloat(f, 'g', -1, 64), nil
			}
		}
	case pref.StringKind:
		// String values are serialized as is without any escaping.
		return v.String(), nil
	case pref.BytesKind:
		if s, ok := marshalBytes(v.Bytes()); ok {
			return s, nil
		}
	}
	return "", fmt.Errorf("could not format value for %v: %v", k, v)
}

// JSONCamelCase converts a snake_case identifier to a camelCase identifier,
// according to the protobuf JSON specification.
func jsonCamelCase(s string) string {
	var b []byte
	var wasUnderscore bool
	for i := 0; i < len(s); i++ { // proto identifiers are always ASCII
		c := s[i]
		if c != '_' {
			if wasUnderscore && isASCIILower(c) {
				c -= 'a' - 'A' // convert to uppercase
			}
			b = append(b, c)
		}
		wasUnderscore = c == '_'
	}
	return string(b)
}

func isASCIILower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

// marshalBytes serializes bytes by using C escaping.
// To match the exact output of protoc, this is identical to the
// CEscape function in strutil.cc of the protoc source code.
func marshalBytes(b []byte) (string, bool) {
	var s []byte
	for _, c := range b {
		switch c {
		case '\n':
			s = append(s, `\n`...)
		case '\r':
			s = append(s, `\r`...)
		case '\t':
			s = append(s, `\t`...)
		case '"':
			s = append(s, `\"`...)
		case '\'':
			s = append(s, `\'`...)
		case '\\':
			s = append(s, `\\`...)
		default:
			if printableASCII := c >= 0x20 && c <= 0x7e; printableASCII {
				s = append(s, c)
			} else {
				s = append(s, fmt.Sprintf(`\%03o`, c)...)
			}
		}
	}
	return string(s), true
}
