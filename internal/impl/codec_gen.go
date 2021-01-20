// NOT FUCINKING EGEN
package impl

import (
	"math"
	"unicode/utf8"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// sizeBool returns the size of wire encoding a bool pointer as a Bool.
func sizeBool(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := *p.Bool()
	return f.tagsize + protowire.SizeVarint(protowire.EncodeBool(v))
}

// appendBool wire encodes a bool pointer as a Bool.
func appendBool(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Bool()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendVarint(b, protowire.EncodeBool(v))
	return b, nil
}

// consumeBool wire decodes a bool pointer as a Bool.
func consumeBool(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.VarintType {
		return out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return out, errDecode
	}
	*p.Bool() = protowire.DecodeBool(v)
	out.n = n
	return out, nil
}

var coderBool = pointerCoderFuncs{
	size:      sizeBool,
	marshal:   appendBool,
	unmarshal: consumeBool,
	merge:     mergeBool,
}

// sizeBoolNoZero returns the size of wire encoding a bool pointer as a Bool.
// The zero value is not encoded.
func sizeBoolNoZero(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := *p.Bool()
	if v == false {
		return 0
	}
	return f.tagsize + protowire.SizeVarint(protowire.EncodeBool(v))
}

// appendBoolNoZero wire encodes a bool pointer as a Bool.
// The zero value is not encoded.
func appendBoolNoZero(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Bool()
	if v == false {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendVarint(b, protowire.EncodeBool(v))
	return b, nil
}

var coderBoolNoZero = pointerCoderFuncs{
	size:      sizeBoolNoZero,
	marshal:   appendBoolNoZero,
	unmarshal: consumeBool,
	merge:     mergeBoolNoZero,
}

// sizeBoolPtr returns the size of wire encoding a *bool pointer as a Bool.
// It panics if the pointer is nil.
func sizeBoolPtr(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := **p.BoolPtr()
	return f.tagsize + protowire.SizeVarint(protowire.EncodeBool(v))
}

// appendBoolPtr wire encodes a *bool pointer as a Bool.
// It panics if the pointer is nil.
func appendBoolPtr(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := **p.BoolPtr()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendVarint(b, protowire.EncodeBool(v))
	return b, nil
}

// consumeBoolPtr wire decodes a *bool pointer as a Bool.
func consumeBoolPtr(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.VarintType {
		return out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return out, errDecode
	}
	vp := p.BoolPtr()
	if *vp == nil {
		*vp = new(bool)
	}
	**vp = protowire.DecodeBool(v)
	out.n = n
	return out, nil
}

var coderBoolPtr = pointerCoderFuncs{
	size:      sizeBoolPtr,
	marshal:   appendBoolPtr,
	unmarshal: consumeBoolPtr,
	merge:     mergeBoolPtr,
}

// sizeBoolSlice returns the size of wire encoding a []bool pointer as a repeated Bool.
func sizeBoolSlice(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	s := *p.BoolSlice()
	for _, v := range s {
		size += f.tagsize + protowire.SizeVarint(protowire.EncodeBool(v))
	}
	return size
}

// appendBoolSlice encodes a []bool pointer as a repeated Bool.
func appendBoolSlice(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.BoolSlice()
	for _, v := range s {
		b = protowire.AppendVarint(b, f.wiretag)
		b = protowire.AppendVarint(b, protowire.EncodeBool(v))
	}
	return b, nil
}

// consumeBoolSlice wire decodes a []bool pointer as a repeated Bool.
func consumeBoolSlice(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	sp := p.BoolSlice()
	if wtyp == protowire.BytesType {
		s := *sp
		b, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return out, errDecode
		}
		for len(b) > 0 {
			var v uint64
			var n int
			if len(b) >= 1 && b[0] < 0x80 {
				v = uint64(b[0])
				n = 1
			} else if len(b) >= 2 && b[1] < 128 {
				v = uint64(b[0]&0x7f) + uint64(b[1])<<7
				n = 2
			} else {
				v, n = protowire.ConsumeVarint(b)
			}
			if n < 0 {
				return out, errDecode
			}
			s = append(s, protowire.DecodeBool(v))
			b = b[n:]
		}
		*sp = s
		out.n = n
		return out, nil
	}
	if wtyp != protowire.VarintType {
		return out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return out, errDecode
	}
	*sp = append(*sp, protowire.DecodeBool(v))
	out.n = n
	return out, nil
}

var coderBoolSlice = pointerCoderFuncs{
	size:      sizeBoolSlice,
	marshal:   appendBoolSlice,
	unmarshal: consumeBoolSlice,
	merge:     mergeBoolSlice,
}

// sizeBoolPackedSlice returns the size of wire encoding a []bool pointer as a packed repeated Bool.
func sizeBoolPackedSlice(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	s := *p.BoolSlice()
	if len(s) == 0 {
		return 0
	}
	n := 0
	for _, v := range s {
		n += protowire.SizeVarint(protowire.EncodeBool(v))
	}
	return f.tagsize + protowire.SizeBytes(n)
}

// appendBoolPackedSlice encodes a []bool pointer as a packed repeated Bool.
func appendBoolPackedSlice(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.BoolSlice()
	if len(s) == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	n := 0
	for _, v := range s {
		n += protowire.SizeVarint(protowire.EncodeBool(v))
	}
	b = protowire.AppendVarint(b, uint64(n))
	for _, v := range s {
		b = protowire.AppendVarint(b, protowire.EncodeBool(v))
	}
	return b, nil
}

var coderBoolPackedSlice = pointerCoderFuncs{
	size:      sizeBoolPackedSlice,
	marshal:   appendBoolPackedSlice,
	unmarshal: consumeBoolSlice,
	merge:     mergeBoolSlice,
}

// sizeBoolValue returns the size of wire encoding a bool value as a Bool.
func sizeBoolValue(v protoreflect.Value, tagsize int, opts marshalOptions) int {
	return tagsize + protowire.SizeVarint(protowire.EncodeBool(v.Bool()))
}

// appendBoolValue encodes a bool value as a Bool.
func appendBoolValue(b []byte, v protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	b = protowire.AppendVarint(b, wiretag)
	b = protowire.AppendVarint(b, protowire.EncodeBool(v.Bool()))
	return b, nil
}

// consumeBoolValue decodes a bool value as a Bool.
func consumeBoolValue(b []byte, _ protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	if wtyp != protowire.VarintType {
		return protoreflect.Value{}, out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	out.n = n
	return protoreflect.ValueOfBool(protowire.DecodeBool(v)), out, nil
}

var coderBoolValue = valueCoderFuncs{
	size:      sizeBoolValue,
	marshal:   appendBoolValue,
	unmarshal: consumeBoolValue,
	merge:     mergeScalarValue,
}

// sizeBoolSliceValue returns the size of wire encoding a []bool value as a repeated Bool.
func sizeBoolSliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	for i, llen := 0, list.Len(); i < llen; i++ {
		v := list.Get(i)
		size += tagsize + protowire.SizeVarint(protowire.EncodeBool(v.Bool()))
	}
	return size
}

// appendBoolSliceValue encodes a []bool value as a repeated Bool.
func appendBoolSliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	for i, llen := 0, list.Len(); i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendVarint(b, wiretag)
		b = protowire.AppendVarint(b, protowire.EncodeBool(v.Bool()))
	}
	return b, nil
}

// consumeBoolSliceValue wire decodes a []bool value as a repeated Bool.
func consumeBoolSliceValue(b []byte, listv protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	list := listv.List()
	if wtyp == protowire.BytesType {
		b, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return protoreflect.Value{}, out, errDecode
		}
		for len(b) > 0 {
			var v uint64
			var n int
			if len(b) >= 1 && b[0] < 0x80 {
				v = uint64(b[0])
				n = 1
			} else if len(b) >= 2 && b[1] < 128 {
				v = uint64(b[0]&0x7f) + uint64(b[1])<<7
				n = 2
			} else {
				v, n = protowire.ConsumeVarint(b)
			}
			if n < 0 {
				return protoreflect.Value{}, out, errDecode
			}
			list.Append(protoreflect.ValueOfBool(protowire.DecodeBool(v)))
			b = b[n:]
		}
		out.n = n
		return listv, out, nil
	}
	if wtyp != protowire.VarintType {
		return protoreflect.Value{}, out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	list.Append(protoreflect.ValueOfBool(protowire.DecodeBool(v)))
	out.n = n
	return listv, out, nil
}

var coderBoolSliceValue = valueCoderFuncs{
	size:      sizeBoolSliceValue,
	marshal:   appendBoolSliceValue,
	unmarshal: consumeBoolSliceValue,
	merge:     mergeListValue,
}

// sizeBoolPackedSliceValue returns the size of wire encoding a []bool value as a packed repeated Bool.
func sizeBoolPackedSliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	llen := list.Len()
	if llen == 0 {
		return 0
	}
	n := 0
	for i, llen := 0, llen; i < llen; i++ {
		v := list.Get(i)
		n += protowire.SizeVarint(protowire.EncodeBool(v.Bool()))
	}
	return tagsize + protowire.SizeBytes(n)
}

// appendBoolPackedSliceValue encodes a []bool value as a packed repeated Bool.
func appendBoolPackedSliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	llen := list.Len()
	if llen == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, wiretag)
	n := 0
	for i := 0; i < llen; i++ {
		v := list.Get(i)
		n += protowire.SizeVarint(protowire.EncodeBool(v.Bool()))
	}
	b = protowire.AppendVarint(b, uint64(n))
	for i := 0; i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendVarint(b, protowire.EncodeBool(v.Bool()))
	}
	return b, nil
}

var coderBoolPackedSliceValue = valueCoderFuncs{
	size:      sizeBoolPackedSliceValue,
	marshal:   appendBoolPackedSliceValue,
	unmarshal: consumeBoolSliceValue,
	merge:     mergeListValue,
}

// sizeEnumValue returns the size of wire encoding a  value as a Enum.
func sizeEnumValue(v protoreflect.Value, tagsize int, opts marshalOptions) int {
	return tagsize + protowire.SizeVarint(uint64(v.Enum()))
}

// appendEnumValue encodes a  value as a Enum.
func appendEnumValue(b []byte, v protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	b = protowire.AppendVarint(b, wiretag)
	b = protowire.AppendVarint(b, uint64(v.Enum()))
	return b, nil
}

// consumeEnumValue decodes a  value as a Enum.
func consumeEnumValue(b []byte, _ protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	if wtyp != protowire.VarintType {
		return protoreflect.Value{}, out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	out.n = n
	return protoreflect.ValueOfEnum(protoreflect.EnumNumber(v)), out, nil
}

var coderEnumValue = valueCoderFuncs{
	size:      sizeEnumValue,
	marshal:   appendEnumValue,
	unmarshal: consumeEnumValue,
	merge:     mergeScalarValue,
}

// sizeEnumSliceValue returns the size of wire encoding a [] value as a repeated Enum.
func sizeEnumSliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	for i, llen := 0, list.Len(); i < llen; i++ {
		v := list.Get(i)
		size += tagsize + protowire.SizeVarint(uint64(v.Enum()))
	}
	return size
}

// appendEnumSliceValue encodes a [] value as a repeated Enum.
func appendEnumSliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	for i, llen := 0, list.Len(); i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendVarint(b, wiretag)
		b = protowire.AppendVarint(b, uint64(v.Enum()))
	}
	return b, nil
}

// consumeEnumSliceValue wire decodes a [] value as a repeated Enum.
func consumeEnumSliceValue(b []byte, listv protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	list := listv.List()
	if wtyp == protowire.BytesType {
		b, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return protoreflect.Value{}, out, errDecode
		}
		for len(b) > 0 {
			var v uint64
			var n int
			if len(b) >= 1 && b[0] < 0x80 {
				v = uint64(b[0])
				n = 1
			} else if len(b) >= 2 && b[1] < 128 {
				v = uint64(b[0]&0x7f) + uint64(b[1])<<7
				n = 2
			} else {
				v, n = protowire.ConsumeVarint(b)
			}
			if n < 0 {
				return protoreflect.Value{}, out, errDecode
			}
			list.Append(protoreflect.ValueOfEnum(protoreflect.EnumNumber(v)))
			b = b[n:]
		}
		out.n = n
		return listv, out, nil
	}
	if wtyp != protowire.VarintType {
		return protoreflect.Value{}, out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	list.Append(protoreflect.ValueOfEnum(protoreflect.EnumNumber(v)))
	out.n = n
	return listv, out, nil
}

var coderEnumSliceValue = valueCoderFuncs{
	size:      sizeEnumSliceValue,
	marshal:   appendEnumSliceValue,
	unmarshal: consumeEnumSliceValue,
	merge:     mergeListValue,
}

// sizeEnumPackedSliceValue returns the size of wire encoding a [] value as a packed repeated Enum.
func sizeEnumPackedSliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	llen := list.Len()
	if llen == 0 {
		return 0
	}
	n := 0
	for i, llen := 0, llen; i < llen; i++ {
		v := list.Get(i)
		n += protowire.SizeVarint(uint64(v.Enum()))
	}
	return tagsize + protowire.SizeBytes(n)
}

// appendEnumPackedSliceValue encodes a [] value as a packed repeated Enum.
func appendEnumPackedSliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	llen := list.Len()
	if llen == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, wiretag)
	n := 0
	for i := 0; i < llen; i++ {
		v := list.Get(i)
		n += protowire.SizeVarint(uint64(v.Enum()))
	}
	b = protowire.AppendVarint(b, uint64(n))
	for i := 0; i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendVarint(b, uint64(v.Enum()))
	}
	return b, nil
}

var coderEnumPackedSliceValue = valueCoderFuncs{
	size:      sizeEnumPackedSliceValue,
	marshal:   appendEnumPackedSliceValue,
	unmarshal: consumeEnumSliceValue,
	merge:     mergeListValue,
}

// sizeInt32 returns the size of wire encoding a int32 pointer as a Int32.
func sizeInt32(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := *p.Int32()
	return f.tagsize + protowire.SizeVarint(uint64(v))
}

// appendInt32 wire encodes a int32 pointer as a Int32.
func appendInt32(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Int32()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendVarint(b, uint64(v))
	return b, nil
}

// consumeInt32 wire decodes a int32 pointer as a Int32.
func consumeInt32(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.VarintType {
		return out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return out, errDecode
	}
	*p.Int32() = int32(v)
	out.n = n
	return out, nil
}

var coderInt32 = pointerCoderFuncs{
	size:      sizeInt32,
	marshal:   appendInt32,
	unmarshal: consumeInt32,
	merge:     mergeInt32,
}

// sizeInt32NoZero returns the size of wire encoding a int32 pointer as a Int32.
// The zero value is not encoded.
func sizeInt32NoZero(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := *p.Int32()
	if v == 0 {
		return 0
	}
	return f.tagsize + protowire.SizeVarint(uint64(v))
}

// appendInt32NoZero wire encodes a int32 pointer as a Int32.
// The zero value is not encoded.
func appendInt32NoZero(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Int32()
	if v == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendVarint(b, uint64(v))
	return b, nil
}

var coderInt32NoZero = pointerCoderFuncs{
	size:      sizeInt32NoZero,
	marshal:   appendInt32NoZero,
	unmarshal: consumeInt32,
	merge:     mergeInt32NoZero,
}

// sizeInt32Ptr returns the size of wire encoding a *int32 pointer as a Int32.
// It panics if the pointer is nil.
func sizeInt32Ptr(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := **p.Int32Ptr()
	return f.tagsize + protowire.SizeVarint(uint64(v))
}

// appendInt32Ptr wire encodes a *int32 pointer as a Int32.
// It panics if the pointer is nil.
func appendInt32Ptr(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := **p.Int32Ptr()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendVarint(b, uint64(v))
	return b, nil
}

// consumeInt32Ptr wire decodes a *int32 pointer as a Int32.
func consumeInt32Ptr(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.VarintType {
		return out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return out, errDecode
	}
	vp := p.Int32Ptr()
	if *vp == nil {
		*vp = new(int32)
	}
	**vp = int32(v)
	out.n = n
	return out, nil
}

var coderInt32Ptr = pointerCoderFuncs{
	size:      sizeInt32Ptr,
	marshal:   appendInt32Ptr,
	unmarshal: consumeInt32Ptr,
	merge:     mergeInt32Ptr,
}

// sizeInt32Slice returns the size of wire encoding a []int32 pointer as a repeated Int32.
func sizeInt32Slice(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	s := *p.Int32Slice()
	for _, v := range s {
		size += f.tagsize + protowire.SizeVarint(uint64(v))
	}
	return size
}

// appendInt32Slice encodes a []int32 pointer as a repeated Int32.
func appendInt32Slice(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.Int32Slice()
	for _, v := range s {
		b = protowire.AppendVarint(b, f.wiretag)
		b = protowire.AppendVarint(b, uint64(v))
	}
	return b, nil
}

// consumeInt32Slice wire decodes a []int32 pointer as a repeated Int32.
func consumeInt32Slice(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	sp := p.Int32Slice()
	if wtyp == protowire.BytesType {
		s := *sp
		b, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return out, errDecode
		}
		for len(b) > 0 {
			var v uint64
			var n int
			if len(b) >= 1 && b[0] < 0x80 {
				v = uint64(b[0])
				n = 1
			} else if len(b) >= 2 && b[1] < 128 {
				v = uint64(b[0]&0x7f) + uint64(b[1])<<7
				n = 2
			} else {
				v, n = protowire.ConsumeVarint(b)
			}
			if n < 0 {
				return out, errDecode
			}
			s = append(s, int32(v))
			b = b[n:]
		}
		*sp = s
		out.n = n
		return out, nil
	}
	if wtyp != protowire.VarintType {
		return out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return out, errDecode
	}
	*sp = append(*sp, int32(v))
	out.n = n
	return out, nil
}

var coderInt32Slice = pointerCoderFuncs{
	size:      sizeInt32Slice,
	marshal:   appendInt32Slice,
	unmarshal: consumeInt32Slice,
	merge:     mergeInt32Slice,
}

// sizeInt32PackedSlice returns the size of wire encoding a []int32 pointer as a packed repeated Int32.
func sizeInt32PackedSlice(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	s := *p.Int32Slice()
	if len(s) == 0 {
		return 0
	}
	n := 0
	for _, v := range s {
		n += protowire.SizeVarint(uint64(v))
	}
	return f.tagsize + protowire.SizeBytes(n)
}

// appendInt32PackedSlice encodes a []int32 pointer as a packed repeated Int32.
func appendInt32PackedSlice(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.Int32Slice()
	if len(s) == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	n := 0
	for _, v := range s {
		n += protowire.SizeVarint(uint64(v))
	}
	b = protowire.AppendVarint(b, uint64(n))
	for _, v := range s {
		b = protowire.AppendVarint(b, uint64(v))
	}
	return b, nil
}

var coderInt32PackedSlice = pointerCoderFuncs{
	size:      sizeInt32PackedSlice,
	marshal:   appendInt32PackedSlice,
	unmarshal: consumeInt32Slice,
	merge:     mergeInt32Slice,
}

// sizeInt32Value returns the size of wire encoding a int32 value as a Int32.
func sizeInt32Value(v protoreflect.Value, tagsize int, opts marshalOptions) int {
	return tagsize + protowire.SizeVarint(uint64(int32(v.Int())))
}

// appendInt32Value encodes a int32 value as a Int32.
func appendInt32Value(b []byte, v protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	b = protowire.AppendVarint(b, wiretag)
	b = protowire.AppendVarint(b, uint64(int32(v.Int())))
	return b, nil
}

// consumeInt32Value decodes a int32 value as a Int32.
func consumeInt32Value(b []byte, _ protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	if wtyp != protowire.VarintType {
		return protoreflect.Value{}, out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	out.n = n
	return protoreflect.ValueOfInt32(int32(v)), out, nil
}

var coderInt32Value = valueCoderFuncs{
	size:      sizeInt32Value,
	marshal:   appendInt32Value,
	unmarshal: consumeInt32Value,
	merge:     mergeScalarValue,
}

// sizeInt32SliceValue returns the size of wire encoding a []int32 value as a repeated Int32.
func sizeInt32SliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	for i, llen := 0, list.Len(); i < llen; i++ {
		v := list.Get(i)
		size += tagsize + protowire.SizeVarint(uint64(int32(v.Int())))
	}
	return size
}

// appendInt32SliceValue encodes a []int32 value as a repeated Int32.
func appendInt32SliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	for i, llen := 0, list.Len(); i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendVarint(b, wiretag)
		b = protowire.AppendVarint(b, uint64(int32(v.Int())))
	}
	return b, nil
}

// consumeInt32SliceValue wire decodes a []int32 value as a repeated Int32.
func consumeInt32SliceValue(b []byte, listv protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	list := listv.List()
	if wtyp == protowire.BytesType {
		b, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return protoreflect.Value{}, out, errDecode
		}
		for len(b) > 0 {
			var v uint64
			var n int
			if len(b) >= 1 && b[0] < 0x80 {
				v = uint64(b[0])
				n = 1
			} else if len(b) >= 2 && b[1] < 128 {
				v = uint64(b[0]&0x7f) + uint64(b[1])<<7
				n = 2
			} else {
				v, n = protowire.ConsumeVarint(b)
			}
			if n < 0 {
				return protoreflect.Value{}, out, errDecode
			}
			list.Append(protoreflect.ValueOfInt32(int32(v)))
			b = b[n:]
		}
		out.n = n
		return listv, out, nil
	}
	if wtyp != protowire.VarintType {
		return protoreflect.Value{}, out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	list.Append(protoreflect.ValueOfInt32(int32(v)))
	out.n = n
	return listv, out, nil
}

var coderInt32SliceValue = valueCoderFuncs{
	size:      sizeInt32SliceValue,
	marshal:   appendInt32SliceValue,
	unmarshal: consumeInt32SliceValue,
	merge:     mergeListValue,
}

// sizeInt32PackedSliceValue returns the size of wire encoding a []int32 value as a packed repeated Int32.
func sizeInt32PackedSliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	llen := list.Len()
	if llen == 0 {
		return 0
	}
	n := 0
	for i, llen := 0, llen; i < llen; i++ {
		v := list.Get(i)
		n += protowire.SizeVarint(uint64(int32(v.Int())))
	}
	return tagsize + protowire.SizeBytes(n)
}

// appendInt32PackedSliceValue encodes a []int32 value as a packed repeated Int32.
func appendInt32PackedSliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	llen := list.Len()
	if llen == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, wiretag)
	n := 0
	for i := 0; i < llen; i++ {
		v := list.Get(i)
		n += protowire.SizeVarint(uint64(int32(v.Int())))
	}
	b = protowire.AppendVarint(b, uint64(n))
	for i := 0; i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendVarint(b, uint64(int32(v.Int())))
	}
	return b, nil
}

var coderInt32PackedSliceValue = valueCoderFuncs{
	size:      sizeInt32PackedSliceValue,
	marshal:   appendInt32PackedSliceValue,
	unmarshal: consumeInt32SliceValue,
	merge:     mergeListValue,
}

// sizeSint32 returns the size of wire encoding a int32 pointer as a Sint32.
func sizeSint32(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := *p.Int32()
	return f.tagsize + protowire.SizeVarint(protowire.EncodeZigZag(int64(v)))
}

// appendSint32 wire encodes a int32 pointer as a Sint32.
func appendSint32(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Int32()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendVarint(b, protowire.EncodeZigZag(int64(v)))
	return b, nil
}

// consumeSint32 wire decodes a int32 pointer as a Sint32.
func consumeSint32(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.VarintType {
		return out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return out, errDecode
	}
	*p.Int32() = int32(protowire.DecodeZigZag(v & math.MaxUint32))
	out.n = n
	return out, nil
}

var coderSint32 = pointerCoderFuncs{
	size:      sizeSint32,
	marshal:   appendSint32,
	unmarshal: consumeSint32,
	merge:     mergeInt32,
}

// sizeSint32NoZero returns the size of wire encoding a int32 pointer as a Sint32.
// The zero value is not encoded.
func sizeSint32NoZero(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := *p.Int32()
	if v == 0 {
		return 0
	}
	return f.tagsize + protowire.SizeVarint(protowire.EncodeZigZag(int64(v)))
}

// appendSint32NoZero wire encodes a int32 pointer as a Sint32.
// The zero value is not encoded.
func appendSint32NoZero(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Int32()
	if v == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendVarint(b, protowire.EncodeZigZag(int64(v)))
	return b, nil
}

var coderSint32NoZero = pointerCoderFuncs{
	size:      sizeSint32NoZero,
	marshal:   appendSint32NoZero,
	unmarshal: consumeSint32,
	merge:     mergeInt32NoZero,
}

// sizeSint32Ptr returns the size of wire encoding a *int32 pointer as a Sint32.
// It panics if the pointer is nil.
func sizeSint32Ptr(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := **p.Int32Ptr()
	return f.tagsize + protowire.SizeVarint(protowire.EncodeZigZag(int64(v)))
}

// appendSint32Ptr wire encodes a *int32 pointer as a Sint32.
// It panics if the pointer is nil.
func appendSint32Ptr(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := **p.Int32Ptr()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendVarint(b, protowire.EncodeZigZag(int64(v)))
	return b, nil
}

// consumeSint32Ptr wire decodes a *int32 pointer as a Sint32.
func consumeSint32Ptr(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.VarintType {
		return out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return out, errDecode
	}
	vp := p.Int32Ptr()
	if *vp == nil {
		*vp = new(int32)
	}
	**vp = int32(protowire.DecodeZigZag(v & math.MaxUint32))
	out.n = n
	return out, nil
}

var coderSint32Ptr = pointerCoderFuncs{
	size:      sizeSint32Ptr,
	marshal:   appendSint32Ptr,
	unmarshal: consumeSint32Ptr,
	merge:     mergeInt32Ptr,
}

// sizeSint32Slice returns the size of wire encoding a []int32 pointer as a repeated Sint32.
func sizeSint32Slice(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	s := *p.Int32Slice()
	for _, v := range s {
		size += f.tagsize + protowire.SizeVarint(protowire.EncodeZigZag(int64(v)))
	}
	return size
}

// appendSint32Slice encodes a []int32 pointer as a repeated Sint32.
func appendSint32Slice(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.Int32Slice()
	for _, v := range s {
		b = protowire.AppendVarint(b, f.wiretag)
		b = protowire.AppendVarint(b, protowire.EncodeZigZag(int64(v)))
	}
	return b, nil
}

// consumeSint32Slice wire decodes a []int32 pointer as a repeated Sint32.
func consumeSint32Slice(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	sp := p.Int32Slice()
	if wtyp == protowire.BytesType {
		s := *sp
		b, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return out, errDecode
		}
		for len(b) > 0 {
			var v uint64
			var n int
			if len(b) >= 1 && b[0] < 0x80 {
				v = uint64(b[0])
				n = 1
			} else if len(b) >= 2 && b[1] < 128 {
				v = uint64(b[0]&0x7f) + uint64(b[1])<<7
				n = 2
			} else {
				v, n = protowire.ConsumeVarint(b)
			}
			if n < 0 {
				return out, errDecode
			}
			s = append(s, int32(protowire.DecodeZigZag(v&math.MaxUint32)))
			b = b[n:]
		}
		*sp = s
		out.n = n
		return out, nil
	}
	if wtyp != protowire.VarintType {
		return out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return out, errDecode
	}
	*sp = append(*sp, int32(protowire.DecodeZigZag(v&math.MaxUint32)))
	out.n = n
	return out, nil
}

var coderSint32Slice = pointerCoderFuncs{
	size:      sizeSint32Slice,
	marshal:   appendSint32Slice,
	unmarshal: consumeSint32Slice,
	merge:     mergeInt32Slice,
}

// sizeSint32PackedSlice returns the size of wire encoding a []int32 pointer as a packed repeated Sint32.
func sizeSint32PackedSlice(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	s := *p.Int32Slice()
	if len(s) == 0 {
		return 0
	}
	n := 0
	for _, v := range s {
		n += protowire.SizeVarint(protowire.EncodeZigZag(int64(v)))
	}
	return f.tagsize + protowire.SizeBytes(n)
}

// appendSint32PackedSlice encodes a []int32 pointer as a packed repeated Sint32.
func appendSint32PackedSlice(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.Int32Slice()
	if len(s) == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	n := 0
	for _, v := range s {
		n += protowire.SizeVarint(protowire.EncodeZigZag(int64(v)))
	}
	b = protowire.AppendVarint(b, uint64(n))
	for _, v := range s {
		b = protowire.AppendVarint(b, protowire.EncodeZigZag(int64(v)))
	}
	return b, nil
}

var coderSint32PackedSlice = pointerCoderFuncs{
	size:      sizeSint32PackedSlice,
	marshal:   appendSint32PackedSlice,
	unmarshal: consumeSint32Slice,
	merge:     mergeInt32Slice,
}

// sizeSint32Value returns the size of wire encoding a int32 value as a Sint32.
func sizeSint32Value(v protoreflect.Value, tagsize int, opts marshalOptions) int {
	return tagsize + protowire.SizeVarint(protowire.EncodeZigZag(int64(int32(v.Int()))))
}

// appendSint32Value encodes a int32 value as a Sint32.
func appendSint32Value(b []byte, v protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	b = protowire.AppendVarint(b, wiretag)
	b = protowire.AppendVarint(b, protowire.EncodeZigZag(int64(int32(v.Int()))))
	return b, nil
}

// consumeSint32Value decodes a int32 value as a Sint32.
func consumeSint32Value(b []byte, _ protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	if wtyp != protowire.VarintType {
		return protoreflect.Value{}, out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	out.n = n
	return protoreflect.ValueOfInt32(int32(protowire.DecodeZigZag(v & math.MaxUint32))), out, nil
}

var coderSint32Value = valueCoderFuncs{
	size:      sizeSint32Value,
	marshal:   appendSint32Value,
	unmarshal: consumeSint32Value,
	merge:     mergeScalarValue,
}

// sizeSint32SliceValue returns the size of wire encoding a []int32 value as a repeated Sint32.
func sizeSint32SliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	for i, llen := 0, list.Len(); i < llen; i++ {
		v := list.Get(i)
		size += tagsize + protowire.SizeVarint(protowire.EncodeZigZag(int64(int32(v.Int()))))
	}
	return size
}

// appendSint32SliceValue encodes a []int32 value as a repeated Sint32.
func appendSint32SliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	for i, llen := 0, list.Len(); i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendVarint(b, wiretag)
		b = protowire.AppendVarint(b, protowire.EncodeZigZag(int64(int32(v.Int()))))
	}
	return b, nil
}

// consumeSint32SliceValue wire decodes a []int32 value as a repeated Sint32.
func consumeSint32SliceValue(b []byte, listv protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	list := listv.List()
	if wtyp == protowire.BytesType {
		b, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return protoreflect.Value{}, out, errDecode
		}
		for len(b) > 0 {
			var v uint64
			var n int
			if len(b) >= 1 && b[0] < 0x80 {
				v = uint64(b[0])
				n = 1
			} else if len(b) >= 2 && b[1] < 128 {
				v = uint64(b[0]&0x7f) + uint64(b[1])<<7
				n = 2
			} else {
				v, n = protowire.ConsumeVarint(b)
			}
			if n < 0 {
				return protoreflect.Value{}, out, errDecode
			}
			list.Append(protoreflect.ValueOfInt32(int32(protowire.DecodeZigZag(v & math.MaxUint32))))
			b = b[n:]
		}
		out.n = n
		return listv, out, nil
	}
	if wtyp != protowire.VarintType {
		return protoreflect.Value{}, out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	list.Append(protoreflect.ValueOfInt32(int32(protowire.DecodeZigZag(v & math.MaxUint32))))
	out.n = n
	return listv, out, nil
}

var coderSint32SliceValue = valueCoderFuncs{
	size:      sizeSint32SliceValue,
	marshal:   appendSint32SliceValue,
	unmarshal: consumeSint32SliceValue,
	merge:     mergeListValue,
}

// sizeSint32PackedSliceValue returns the size of wire encoding a []int32 value as a packed repeated Sint32.
func sizeSint32PackedSliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	llen := list.Len()
	if llen == 0 {
		return 0
	}
	n := 0
	for i, llen := 0, llen; i < llen; i++ {
		v := list.Get(i)
		n += protowire.SizeVarint(protowire.EncodeZigZag(int64(int32(v.Int()))))
	}
	return tagsize + protowire.SizeBytes(n)
}

// appendSint32PackedSliceValue encodes a []int32 value as a packed repeated Sint32.
func appendSint32PackedSliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	llen := list.Len()
	if llen == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, wiretag)
	n := 0
	for i := 0; i < llen; i++ {
		v := list.Get(i)
		n += protowire.SizeVarint(protowire.EncodeZigZag(int64(int32(v.Int()))))
	}
	b = protowire.AppendVarint(b, uint64(n))
	for i := 0; i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendVarint(b, protowire.EncodeZigZag(int64(int32(v.Int()))))
	}
	return b, nil
}

var coderSint32PackedSliceValue = valueCoderFuncs{
	size:      sizeSint32PackedSliceValue,
	marshal:   appendSint32PackedSliceValue,
	unmarshal: consumeSint32SliceValue,
	merge:     mergeListValue,
}

// sizeUint32 returns the size of wire encoding a uint32 pointer as a Uint32.
func sizeUint32(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := *p.Uint32()
	return f.tagsize + protowire.SizeVarint(uint64(v))
}

// appendUint32 wire encodes a uint32 pointer as a Uint32.
func appendUint32(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Uint32()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendVarint(b, uint64(v))
	return b, nil
}

// consumeUint32 wire decodes a uint32 pointer as a Uint32.
func consumeUint32(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.VarintType {
		return out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return out, errDecode
	}
	*p.Uint32() = uint32(v)
	out.n = n
	return out, nil
}

var coderUint32 = pointerCoderFuncs{
	size:      sizeUint32,
	marshal:   appendUint32,
	unmarshal: consumeUint32,
	merge:     mergeUint32,
}

// sizeUint32NoZero returns the size of wire encoding a uint32 pointer as a Uint32.
// The zero value is not encoded.
func sizeUint32NoZero(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := *p.Uint32()
	if v == 0 {
		return 0
	}
	return f.tagsize + protowire.SizeVarint(uint64(v))
}

// appendUint32NoZero wire encodes a uint32 pointer as a Uint32.
// The zero value is not encoded.
func appendUint32NoZero(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Uint32()
	if v == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendVarint(b, uint64(v))
	return b, nil
}

var coderUint32NoZero = pointerCoderFuncs{
	size:      sizeUint32NoZero,
	marshal:   appendUint32NoZero,
	unmarshal: consumeUint32,
	merge:     mergeUint32NoZero,
}

// sizeUint32Ptr returns the size of wire encoding a *uint32 pointer as a Uint32.
// It panics if the pointer is nil.
func sizeUint32Ptr(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := **p.Uint32Ptr()
	return f.tagsize + protowire.SizeVarint(uint64(v))
}

// appendUint32Ptr wire encodes a *uint32 pointer as a Uint32.
// It panics if the pointer is nil.
func appendUint32Ptr(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := **p.Uint32Ptr()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendVarint(b, uint64(v))
	return b, nil
}

// consumeUint32Ptr wire decodes a *uint32 pointer as a Uint32.
func consumeUint32Ptr(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.VarintType {
		return out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return out, errDecode
	}
	vp := p.Uint32Ptr()
	if *vp == nil {
		*vp = new(uint32)
	}
	**vp = uint32(v)
	out.n = n
	return out, nil
}

var coderUint32Ptr = pointerCoderFuncs{
	size:      sizeUint32Ptr,
	marshal:   appendUint32Ptr,
	unmarshal: consumeUint32Ptr,
	merge:     mergeUint32Ptr,
}

// sizeUint32Slice returns the size of wire encoding a []uint32 pointer as a repeated Uint32.
func sizeUint32Slice(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	s := *p.Uint32Slice()
	for _, v := range s {
		size += f.tagsize + protowire.SizeVarint(uint64(v))
	}
	return size
}

// appendUint32Slice encodes a []uint32 pointer as a repeated Uint32.
func appendUint32Slice(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.Uint32Slice()
	for _, v := range s {
		b = protowire.AppendVarint(b, f.wiretag)
		b = protowire.AppendVarint(b, uint64(v))
	}
	return b, nil
}

// consumeUint32Slice wire decodes a []uint32 pointer as a repeated Uint32.
func consumeUint32Slice(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	sp := p.Uint32Slice()
	if wtyp == protowire.BytesType {
		s := *sp
		b, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return out, errDecode
		}
		for len(b) > 0 {
			var v uint64
			var n int
			if len(b) >= 1 && b[0] < 0x80 {
				v = uint64(b[0])
				n = 1
			} else if len(b) >= 2 && b[1] < 128 {
				v = uint64(b[0]&0x7f) + uint64(b[1])<<7
				n = 2
			} else {
				v, n = protowire.ConsumeVarint(b)
			}
			if n < 0 {
				return out, errDecode
			}
			s = append(s, uint32(v))
			b = b[n:]
		}
		*sp = s
		out.n = n
		return out, nil
	}
	if wtyp != protowire.VarintType {
		return out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return out, errDecode
	}
	*sp = append(*sp, uint32(v))
	out.n = n
	return out, nil
}

var coderUint32Slice = pointerCoderFuncs{
	size:      sizeUint32Slice,
	marshal:   appendUint32Slice,
	unmarshal: consumeUint32Slice,
	merge:     mergeUint32Slice,
}

// sizeUint32PackedSlice returns the size of wire encoding a []uint32 pointer as a packed repeated Uint32.
func sizeUint32PackedSlice(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	s := *p.Uint32Slice()
	if len(s) == 0 {
		return 0
	}
	n := 0
	for _, v := range s {
		n += protowire.SizeVarint(uint64(v))
	}
	return f.tagsize + protowire.SizeBytes(n)
}

// appendUint32PackedSlice encodes a []uint32 pointer as a packed repeated Uint32.
func appendUint32PackedSlice(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.Uint32Slice()
	if len(s) == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	n := 0
	for _, v := range s {
		n += protowire.SizeVarint(uint64(v))
	}
	b = protowire.AppendVarint(b, uint64(n))
	for _, v := range s {
		b = protowire.AppendVarint(b, uint64(v))
	}
	return b, nil
}

var coderUint32PackedSlice = pointerCoderFuncs{
	size:      sizeUint32PackedSlice,
	marshal:   appendUint32PackedSlice,
	unmarshal: consumeUint32Slice,
	merge:     mergeUint32Slice,
}

// sizeUint32Value returns the size of wire encoding a uint32 value as a Uint32.
func sizeUint32Value(v protoreflect.Value, tagsize int, opts marshalOptions) int {
	return tagsize + protowire.SizeVarint(uint64(uint32(v.Uint())))
}

// appendUint32Value encodes a uint32 value as a Uint32.
func appendUint32Value(b []byte, v protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	b = protowire.AppendVarint(b, wiretag)
	b = protowire.AppendVarint(b, uint64(uint32(v.Uint())))
	return b, nil
}

// consumeUint32Value decodes a uint32 value as a Uint32.
func consumeUint32Value(b []byte, _ protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	if wtyp != protowire.VarintType {
		return protoreflect.Value{}, out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	out.n = n
	return protoreflect.ValueOfUint32(uint32(v)), out, nil
}

var coderUint32Value = valueCoderFuncs{
	size:      sizeUint32Value,
	marshal:   appendUint32Value,
	unmarshal: consumeUint32Value,
	merge:     mergeScalarValue,
}

// sizeUint32SliceValue returns the size of wire encoding a []uint32 value as a repeated Uint32.
func sizeUint32SliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	for i, llen := 0, list.Len(); i < llen; i++ {
		v := list.Get(i)
		size += tagsize + protowire.SizeVarint(uint64(uint32(v.Uint())))
	}
	return size
}

// appendUint32SliceValue encodes a []uint32 value as a repeated Uint32.
func appendUint32SliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	for i, llen := 0, list.Len(); i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendVarint(b, wiretag)
		b = protowire.AppendVarint(b, uint64(uint32(v.Uint())))
	}
	return b, nil
}

// consumeUint32SliceValue wire decodes a []uint32 value as a repeated Uint32.
func consumeUint32SliceValue(b []byte, listv protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	list := listv.List()
	if wtyp == protowire.BytesType {
		b, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return protoreflect.Value{}, out, errDecode
		}
		for len(b) > 0 {
			var v uint64
			var n int
			if len(b) >= 1 && b[0] < 0x80 {
				v = uint64(b[0])
				n = 1
			} else if len(b) >= 2 && b[1] < 128 {
				v = uint64(b[0]&0x7f) + uint64(b[1])<<7
				n = 2
			} else {
				v, n = protowire.ConsumeVarint(b)
			}
			if n < 0 {
				return protoreflect.Value{}, out, errDecode
			}
			list.Append(protoreflect.ValueOfUint32(uint32(v)))
			b = b[n:]
		}
		out.n = n
		return listv, out, nil
	}
	if wtyp != protowire.VarintType {
		return protoreflect.Value{}, out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	list.Append(protoreflect.ValueOfUint32(uint32(v)))
	out.n = n
	return listv, out, nil
}

var coderUint32SliceValue = valueCoderFuncs{
	size:      sizeUint32SliceValue,
	marshal:   appendUint32SliceValue,
	unmarshal: consumeUint32SliceValue,
	merge:     mergeListValue,
}

// sizeUint32PackedSliceValue returns the size of wire encoding a []uint32 value as a packed repeated Uint32.
func sizeUint32PackedSliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	llen := list.Len()
	if llen == 0 {
		return 0
	}
	n := 0
	for i, llen := 0, llen; i < llen; i++ {
		v := list.Get(i)
		n += protowire.SizeVarint(uint64(uint32(v.Uint())))
	}
	return tagsize + protowire.SizeBytes(n)
}

// appendUint32PackedSliceValue encodes a []uint32 value as a packed repeated Uint32.
func appendUint32PackedSliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	llen := list.Len()
	if llen == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, wiretag)
	n := 0
	for i := 0; i < llen; i++ {
		v := list.Get(i)
		n += protowire.SizeVarint(uint64(uint32(v.Uint())))
	}
	b = protowire.AppendVarint(b, uint64(n))
	for i := 0; i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendVarint(b, uint64(uint32(v.Uint())))
	}
	return b, nil
}

var coderUint32PackedSliceValue = valueCoderFuncs{
	size:      sizeUint32PackedSliceValue,
	marshal:   appendUint32PackedSliceValue,
	unmarshal: consumeUint32SliceValue,
	merge:     mergeListValue,
}

// sizeInt64 returns the size of wire encoding a int64 pointer as a Int64.
func sizeInt64(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := *p.Int64()
	return f.tagsize + protowire.SizeVarint(uint64(v))
}

// appendInt64 wire encodes a int64 pointer as a Int64.
func appendInt64(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Int64()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendVarint(b, uint64(v))
	return b, nil
}

// consumeInt64 wire decodes a int64 pointer as a Int64.
func consumeInt64(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.VarintType {
		return out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return out, errDecode
	}
	*p.Int64() = int64(v)
	out.n = n
	return out, nil
}

var coderInt64 = pointerCoderFuncs{
	size:      sizeInt64,
	marshal:   appendInt64,
	unmarshal: consumeInt64,
	merge:     mergeInt64,
}

// sizeInt64NoZero returns the size of wire encoding a int64 pointer as a Int64.
// The zero value is not encoded.
func sizeInt64NoZero(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := *p.Int64()
	if v == 0 {
		return 0
	}
	return f.tagsize + protowire.SizeVarint(uint64(v))
}

// appendInt64NoZero wire encodes a int64 pointer as a Int64.
// The zero value is not encoded.
func appendInt64NoZero(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Int64()
	if v == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendVarint(b, uint64(v))
	return b, nil
}

var coderInt64NoZero = pointerCoderFuncs{
	size:      sizeInt64NoZero,
	marshal:   appendInt64NoZero,
	unmarshal: consumeInt64,
	merge:     mergeInt64NoZero,
}

// sizeInt64Ptr returns the size of wire encoding a *int64 pointer as a Int64.
// It panics if the pointer is nil.
func sizeInt64Ptr(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := **p.Int64Ptr()
	return f.tagsize + protowire.SizeVarint(uint64(v))
}

// appendInt64Ptr wire encodes a *int64 pointer as a Int64.
// It panics if the pointer is nil.
func appendInt64Ptr(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := **p.Int64Ptr()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendVarint(b, uint64(v))
	return b, nil
}

// consumeInt64Ptr wire decodes a *int64 pointer as a Int64.
func consumeInt64Ptr(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.VarintType {
		return out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return out, errDecode
	}
	vp := p.Int64Ptr()
	if *vp == nil {
		*vp = new(int64)
	}
	**vp = int64(v)
	out.n = n
	return out, nil
}

var coderInt64Ptr = pointerCoderFuncs{
	size:      sizeInt64Ptr,
	marshal:   appendInt64Ptr,
	unmarshal: consumeInt64Ptr,
	merge:     mergeInt64Ptr,
}

// sizeInt64Slice returns the size of wire encoding a []int64 pointer as a repeated Int64.
func sizeInt64Slice(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	s := *p.Int64Slice()
	for _, v := range s {
		size += f.tagsize + protowire.SizeVarint(uint64(v))
	}
	return size
}

// appendInt64Slice encodes a []int64 pointer as a repeated Int64.
func appendInt64Slice(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.Int64Slice()
	for _, v := range s {
		b = protowire.AppendVarint(b, f.wiretag)
		b = protowire.AppendVarint(b, uint64(v))
	}
	return b, nil
}

// consumeInt64Slice wire decodes a []int64 pointer as a repeated Int64.
func consumeInt64Slice(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	sp := p.Int64Slice()
	if wtyp == protowire.BytesType {
		s := *sp
		b, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return out, errDecode
		}
		for len(b) > 0 {
			var v uint64
			var n int
			if len(b) >= 1 && b[0] < 0x80 {
				v = uint64(b[0])
				n = 1
			} else if len(b) >= 2 && b[1] < 128 {
				v = uint64(b[0]&0x7f) + uint64(b[1])<<7
				n = 2
			} else {
				v, n = protowire.ConsumeVarint(b)
			}
			if n < 0 {
				return out, errDecode
			}
			s = append(s, int64(v))
			b = b[n:]
		}
		*sp = s
		out.n = n
		return out, nil
	}
	if wtyp != protowire.VarintType {
		return out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return out, errDecode
	}
	*sp = append(*sp, int64(v))
	out.n = n
	return out, nil
}

var coderInt64Slice = pointerCoderFuncs{
	size:      sizeInt64Slice,
	marshal:   appendInt64Slice,
	unmarshal: consumeInt64Slice,
	merge:     mergeInt64Slice,
}

// sizeInt64PackedSlice returns the size of wire encoding a []int64 pointer as a packed repeated Int64.
func sizeInt64PackedSlice(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	s := *p.Int64Slice()
	if len(s) == 0 {
		return 0
	}
	n := 0
	for _, v := range s {
		n += protowire.SizeVarint(uint64(v))
	}
	return f.tagsize + protowire.SizeBytes(n)
}

// appendInt64PackedSlice encodes a []int64 pointer as a packed repeated Int64.
func appendInt64PackedSlice(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.Int64Slice()
	if len(s) == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	n := 0
	for _, v := range s {
		n += protowire.SizeVarint(uint64(v))
	}
	b = protowire.AppendVarint(b, uint64(n))
	for _, v := range s {
		b = protowire.AppendVarint(b, uint64(v))
	}
	return b, nil
}

var coderInt64PackedSlice = pointerCoderFuncs{
	size:      sizeInt64PackedSlice,
	marshal:   appendInt64PackedSlice,
	unmarshal: consumeInt64Slice,
	merge:     mergeInt64Slice,
}

// sizeInt64Value returns the size of wire encoding a int64 value as a Int64.
func sizeInt64Value(v protoreflect.Value, tagsize int, opts marshalOptions) int {
	return tagsize + protowire.SizeVarint(uint64(v.Int()))
}

// appendInt64Value encodes a int64 value as a Int64.
func appendInt64Value(b []byte, v protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	b = protowire.AppendVarint(b, wiretag)
	b = protowire.AppendVarint(b, uint64(v.Int()))
	return b, nil
}

// consumeInt64Value decodes a int64 value as a Int64.
func consumeInt64Value(b []byte, _ protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	if wtyp != protowire.VarintType {
		return protoreflect.Value{}, out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	out.n = n
	return protoreflect.ValueOfInt64(int64(v)), out, nil
}

var coderInt64Value = valueCoderFuncs{
	size:      sizeInt64Value,
	marshal:   appendInt64Value,
	unmarshal: consumeInt64Value,
	merge:     mergeScalarValue,
}

// sizeInt64SliceValue returns the size of wire encoding a []int64 value as a repeated Int64.
func sizeInt64SliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	for i, llen := 0, list.Len(); i < llen; i++ {
		v := list.Get(i)
		size += tagsize + protowire.SizeVarint(uint64(v.Int()))
	}
	return size
}

// appendInt64SliceValue encodes a []int64 value as a repeated Int64.
func appendInt64SliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	for i, llen := 0, list.Len(); i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendVarint(b, wiretag)
		b = protowire.AppendVarint(b, uint64(v.Int()))
	}
	return b, nil
}

// consumeInt64SliceValue wire decodes a []int64 value as a repeated Int64.
func consumeInt64SliceValue(b []byte, listv protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	list := listv.List()
	if wtyp == protowire.BytesType {
		b, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return protoreflect.Value{}, out, errDecode
		}
		for len(b) > 0 {
			var v uint64
			var n int
			if len(b) >= 1 && b[0] < 0x80 {
				v = uint64(b[0])
				n = 1
			} else if len(b) >= 2 && b[1] < 128 {
				v = uint64(b[0]&0x7f) + uint64(b[1])<<7
				n = 2
			} else {
				v, n = protowire.ConsumeVarint(b)
			}
			if n < 0 {
				return protoreflect.Value{}, out, errDecode
			}
			list.Append(protoreflect.ValueOfInt64(int64(v)))
			b = b[n:]
		}
		out.n = n
		return listv, out, nil
	}
	if wtyp != protowire.VarintType {
		return protoreflect.Value{}, out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	list.Append(protoreflect.ValueOfInt64(int64(v)))
	out.n = n
	return listv, out, nil
}

var coderInt64SliceValue = valueCoderFuncs{
	size:      sizeInt64SliceValue,
	marshal:   appendInt64SliceValue,
	unmarshal: consumeInt64SliceValue,
	merge:     mergeListValue,
}

// sizeInt64PackedSliceValue returns the size of wire encoding a []int64 value as a packed repeated Int64.
func sizeInt64PackedSliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	llen := list.Len()
	if llen == 0 {
		return 0
	}
	n := 0
	for i, llen := 0, llen; i < llen; i++ {
		v := list.Get(i)
		n += protowire.SizeVarint(uint64(v.Int()))
	}
	return tagsize + protowire.SizeBytes(n)
}

// appendInt64PackedSliceValue encodes a []int64 value as a packed repeated Int64.
func appendInt64PackedSliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	llen := list.Len()
	if llen == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, wiretag)
	n := 0
	for i := 0; i < llen; i++ {
		v := list.Get(i)
		n += protowire.SizeVarint(uint64(v.Int()))
	}
	b = protowire.AppendVarint(b, uint64(n))
	for i := 0; i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendVarint(b, uint64(v.Int()))
	}
	return b, nil
}

var coderInt64PackedSliceValue = valueCoderFuncs{
	size:      sizeInt64PackedSliceValue,
	marshal:   appendInt64PackedSliceValue,
	unmarshal: consumeInt64SliceValue,
	merge:     mergeListValue,
}

// sizeSint64 returns the size of wire encoding a int64 pointer as a Sint64.
func sizeSint64(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := *p.Int64()
	return f.tagsize + protowire.SizeVarint(protowire.EncodeZigZag(v))
}

// appendSint64 wire encodes a int64 pointer as a Sint64.
func appendSint64(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Int64()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendVarint(b, protowire.EncodeZigZag(v))
	return b, nil
}

// consumeSint64 wire decodes a int64 pointer as a Sint64.
func consumeSint64(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.VarintType {
		return out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return out, errDecode
	}
	*p.Int64() = protowire.DecodeZigZag(v)
	out.n = n
	return out, nil
}

var coderSint64 = pointerCoderFuncs{
	size:      sizeSint64,
	marshal:   appendSint64,
	unmarshal: consumeSint64,
	merge:     mergeInt64,
}

// sizeSint64NoZero returns the size of wire encoding a int64 pointer as a Sint64.
// The zero value is not encoded.
func sizeSint64NoZero(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := *p.Int64()
	if v == 0 {
		return 0
	}
	return f.tagsize + protowire.SizeVarint(protowire.EncodeZigZag(v))
}

// appendSint64NoZero wire encodes a int64 pointer as a Sint64.
// The zero value is not encoded.
func appendSint64NoZero(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Int64()
	if v == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendVarint(b, protowire.EncodeZigZag(v))
	return b, nil
}

var coderSint64NoZero = pointerCoderFuncs{
	size:      sizeSint64NoZero,
	marshal:   appendSint64NoZero,
	unmarshal: consumeSint64,
	merge:     mergeInt64NoZero,
}

// sizeSint64Ptr returns the size of wire encoding a *int64 pointer as a Sint64.
// It panics if the pointer is nil.
func sizeSint64Ptr(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := **p.Int64Ptr()
	return f.tagsize + protowire.SizeVarint(protowire.EncodeZigZag(v))
}

// appendSint64Ptr wire encodes a *int64 pointer as a Sint64.
// It panics if the pointer is nil.
func appendSint64Ptr(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := **p.Int64Ptr()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendVarint(b, protowire.EncodeZigZag(v))
	return b, nil
}

// consumeSint64Ptr wire decodes a *int64 pointer as a Sint64.
func consumeSint64Ptr(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.VarintType {
		return out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return out, errDecode
	}
	vp := p.Int64Ptr()
	if *vp == nil {
		*vp = new(int64)
	}
	**vp = protowire.DecodeZigZag(v)
	out.n = n
	return out, nil
}

var coderSint64Ptr = pointerCoderFuncs{
	size:      sizeSint64Ptr,
	marshal:   appendSint64Ptr,
	unmarshal: consumeSint64Ptr,
	merge:     mergeInt64Ptr,
}

// sizeSint64Slice returns the size of wire encoding a []int64 pointer as a repeated Sint64.
func sizeSint64Slice(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	s := *p.Int64Slice()
	for _, v := range s {
		size += f.tagsize + protowire.SizeVarint(protowire.EncodeZigZag(v))
	}
	return size
}

// appendSint64Slice encodes a []int64 pointer as a repeated Sint64.
func appendSint64Slice(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.Int64Slice()
	for _, v := range s {
		b = protowire.AppendVarint(b, f.wiretag)
		b = protowire.AppendVarint(b, protowire.EncodeZigZag(v))
	}
	return b, nil
}

// consumeSint64Slice wire decodes a []int64 pointer as a repeated Sint64.
func consumeSint64Slice(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	sp := p.Int64Slice()
	if wtyp == protowire.BytesType {
		s := *sp
		b, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return out, errDecode
		}
		for len(b) > 0 {
			var v uint64
			var n int
			if len(b) >= 1 && b[0] < 0x80 {
				v = uint64(b[0])
				n = 1
			} else if len(b) >= 2 && b[1] < 128 {
				v = uint64(b[0]&0x7f) + uint64(b[1])<<7
				n = 2
			} else {
				v, n = protowire.ConsumeVarint(b)
			}
			if n < 0 {
				return out, errDecode
			}
			s = append(s, protowire.DecodeZigZag(v))
			b = b[n:]
		}
		*sp = s
		out.n = n
		return out, nil
	}
	if wtyp != protowire.VarintType {
		return out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return out, errDecode
	}
	*sp = append(*sp, protowire.DecodeZigZag(v))
	out.n = n
	return out, nil
}

var coderSint64Slice = pointerCoderFuncs{
	size:      sizeSint64Slice,
	marshal:   appendSint64Slice,
	unmarshal: consumeSint64Slice,
	merge:     mergeInt64Slice,
}

// sizeSint64PackedSlice returns the size of wire encoding a []int64 pointer as a packed repeated Sint64.
func sizeSint64PackedSlice(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	s := *p.Int64Slice()
	if len(s) == 0 {
		return 0
	}
	n := 0
	for _, v := range s {
		n += protowire.SizeVarint(protowire.EncodeZigZag(v))
	}
	return f.tagsize + protowire.SizeBytes(n)
}

// appendSint64PackedSlice encodes a []int64 pointer as a packed repeated Sint64.
func appendSint64PackedSlice(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.Int64Slice()
	if len(s) == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	n := 0
	for _, v := range s {
		n += protowire.SizeVarint(protowire.EncodeZigZag(v))
	}
	b = protowire.AppendVarint(b, uint64(n))
	for _, v := range s {
		b = protowire.AppendVarint(b, protowire.EncodeZigZag(v))
	}
	return b, nil
}

var coderSint64PackedSlice = pointerCoderFuncs{
	size:      sizeSint64PackedSlice,
	marshal:   appendSint64PackedSlice,
	unmarshal: consumeSint64Slice,
	merge:     mergeInt64Slice,
}

// sizeSint64Value returns the size of wire encoding a int64 value as a Sint64.
func sizeSint64Value(v protoreflect.Value, tagsize int, opts marshalOptions) int {
	return tagsize + protowire.SizeVarint(protowire.EncodeZigZag(v.Int()))
}

// appendSint64Value encodes a int64 value as a Sint64.
func appendSint64Value(b []byte, v protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	b = protowire.AppendVarint(b, wiretag)
	b = protowire.AppendVarint(b, protowire.EncodeZigZag(v.Int()))
	return b, nil
}

// consumeSint64Value decodes a int64 value as a Sint64.
func consumeSint64Value(b []byte, _ protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	if wtyp != protowire.VarintType {
		return protoreflect.Value{}, out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	out.n = n
	return protoreflect.ValueOfInt64(protowire.DecodeZigZag(v)), out, nil
}

var coderSint64Value = valueCoderFuncs{
	size:      sizeSint64Value,
	marshal:   appendSint64Value,
	unmarshal: consumeSint64Value,
	merge:     mergeScalarValue,
}

// sizeSint64SliceValue returns the size of wire encoding a []int64 value as a repeated Sint64.
func sizeSint64SliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	for i, llen := 0, list.Len(); i < llen; i++ {
		v := list.Get(i)
		size += tagsize + protowire.SizeVarint(protowire.EncodeZigZag(v.Int()))
	}
	return size
}

// appendSint64SliceValue encodes a []int64 value as a repeated Sint64.
func appendSint64SliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	for i, llen := 0, list.Len(); i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendVarint(b, wiretag)
		b = protowire.AppendVarint(b, protowire.EncodeZigZag(v.Int()))
	}
	return b, nil
}

// consumeSint64SliceValue wire decodes a []int64 value as a repeated Sint64.
func consumeSint64SliceValue(b []byte, listv protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	list := listv.List()
	if wtyp == protowire.BytesType {
		b, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return protoreflect.Value{}, out, errDecode
		}
		for len(b) > 0 {
			var v uint64
			var n int
			if len(b) >= 1 && b[0] < 0x80 {
				v = uint64(b[0])
				n = 1
			} else if len(b) >= 2 && b[1] < 128 {
				v = uint64(b[0]&0x7f) + uint64(b[1])<<7
				n = 2
			} else {
				v, n = protowire.ConsumeVarint(b)
			}
			if n < 0 {
				return protoreflect.Value{}, out, errDecode
			}
			list.Append(protoreflect.ValueOfInt64(protowire.DecodeZigZag(v)))
			b = b[n:]
		}
		out.n = n
		return listv, out, nil
	}
	if wtyp != protowire.VarintType {
		return protoreflect.Value{}, out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	list.Append(protoreflect.ValueOfInt64(protowire.DecodeZigZag(v)))
	out.n = n
	return listv, out, nil
}

var coderSint64SliceValue = valueCoderFuncs{
	size:      sizeSint64SliceValue,
	marshal:   appendSint64SliceValue,
	unmarshal: consumeSint64SliceValue,
	merge:     mergeListValue,
}

// sizeSint64PackedSliceValue returns the size of wire encoding a []int64 value as a packed repeated Sint64.
func sizeSint64PackedSliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	llen := list.Len()
	if llen == 0 {
		return 0
	}
	n := 0
	for i, llen := 0, llen; i < llen; i++ {
		v := list.Get(i)
		n += protowire.SizeVarint(protowire.EncodeZigZag(v.Int()))
	}
	return tagsize + protowire.SizeBytes(n)
}

// appendSint64PackedSliceValue encodes a []int64 value as a packed repeated Sint64.
func appendSint64PackedSliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	llen := list.Len()
	if llen == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, wiretag)
	n := 0
	for i := 0; i < llen; i++ {
		v := list.Get(i)
		n += protowire.SizeVarint(protowire.EncodeZigZag(v.Int()))
	}
	b = protowire.AppendVarint(b, uint64(n))
	for i := 0; i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendVarint(b, protowire.EncodeZigZag(v.Int()))
	}
	return b, nil
}

var coderSint64PackedSliceValue = valueCoderFuncs{
	size:      sizeSint64PackedSliceValue,
	marshal:   appendSint64PackedSliceValue,
	unmarshal: consumeSint64SliceValue,
	merge:     mergeListValue,
}

// sizeUint64 returns the size of wire encoding a uint64 pointer as a Uint64.
func sizeUint64(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := *p.Uint64()
	return f.tagsize + protowire.SizeVarint(v)
}

// appendUint64 wire encodes a uint64 pointer as a Uint64.
func appendUint64(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Uint64()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendVarint(b, v)
	return b, nil
}

// consumeUint64 wire decodes a uint64 pointer as a Uint64.
func consumeUint64(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.VarintType {
		return out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return out, errDecode
	}
	*p.Uint64() = v
	out.n = n
	return out, nil
}

var coderUint64 = pointerCoderFuncs{
	size:      sizeUint64,
	marshal:   appendUint64,
	unmarshal: consumeUint64,
	merge:     mergeUint64,
}

// sizeUint64NoZero returns the size of wire encoding a uint64 pointer as a Uint64.
// The zero value is not encoded.
func sizeUint64NoZero(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := *p.Uint64()
	if v == 0 {
		return 0
	}
	return f.tagsize + protowire.SizeVarint(v)
}

// appendUint64NoZero wire encodes a uint64 pointer as a Uint64.
// The zero value is not encoded.
func appendUint64NoZero(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Uint64()
	if v == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendVarint(b, v)
	return b, nil
}

var coderUint64NoZero = pointerCoderFuncs{
	size:      sizeUint64NoZero,
	marshal:   appendUint64NoZero,
	unmarshal: consumeUint64,
	merge:     mergeUint64NoZero,
}

// sizeUint64Ptr returns the size of wire encoding a *uint64 pointer as a Uint64.
// It panics if the pointer is nil.
func sizeUint64Ptr(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := **p.Uint64Ptr()
	return f.tagsize + protowire.SizeVarint(v)
}

// appendUint64Ptr wire encodes a *uint64 pointer as a Uint64.
// It panics if the pointer is nil.
func appendUint64Ptr(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := **p.Uint64Ptr()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendVarint(b, v)
	return b, nil
}

// consumeUint64Ptr wire decodes a *uint64 pointer as a Uint64.
func consumeUint64Ptr(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.VarintType {
		return out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return out, errDecode
	}
	vp := p.Uint64Ptr()
	if *vp == nil {
		*vp = new(uint64)
	}
	**vp = v
	out.n = n
	return out, nil
}

var coderUint64Ptr = pointerCoderFuncs{
	size:      sizeUint64Ptr,
	marshal:   appendUint64Ptr,
	unmarshal: consumeUint64Ptr,
	merge:     mergeUint64Ptr,
}

// sizeUint64Slice returns the size of wire encoding a []uint64 pointer as a repeated Uint64.
func sizeUint64Slice(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	s := *p.Uint64Slice()
	for _, v := range s {
		size += f.tagsize + protowire.SizeVarint(v)
	}
	return size
}

// appendUint64Slice encodes a []uint64 pointer as a repeated Uint64.
func appendUint64Slice(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.Uint64Slice()
	for _, v := range s {
		b = protowire.AppendVarint(b, f.wiretag)
		b = protowire.AppendVarint(b, v)
	}
	return b, nil
}

// consumeUint64Slice wire decodes a []uint64 pointer as a repeated Uint64.
func consumeUint64Slice(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	sp := p.Uint64Slice()
	if wtyp == protowire.BytesType {
		s := *sp
		b, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return out, errDecode
		}
		for len(b) > 0 {
			var v uint64
			var n int
			if len(b) >= 1 && b[0] < 0x80 {
				v = uint64(b[0])
				n = 1
			} else if len(b) >= 2 && b[1] < 128 {
				v = uint64(b[0]&0x7f) + uint64(b[1])<<7
				n = 2
			} else {
				v, n = protowire.ConsumeVarint(b)
			}
			if n < 0 {
				return out, errDecode
			}
			s = append(s, v)
			b = b[n:]
		}
		*sp = s
		out.n = n
		return out, nil
	}
	if wtyp != protowire.VarintType {
		return out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return out, errDecode
	}
	*sp = append(*sp, v)
	out.n = n
	return out, nil
}

var coderUint64Slice = pointerCoderFuncs{
	size:      sizeUint64Slice,
	marshal:   appendUint64Slice,
	unmarshal: consumeUint64Slice,
	merge:     mergeUint64Slice,
}

// sizeUint64PackedSlice returns the size of wire encoding a []uint64 pointer as a packed repeated Uint64.
func sizeUint64PackedSlice(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	s := *p.Uint64Slice()
	if len(s) == 0 {
		return 0
	}
	n := 0
	for _, v := range s {
		n += protowire.SizeVarint(v)
	}
	return f.tagsize + protowire.SizeBytes(n)
}

// appendUint64PackedSlice encodes a []uint64 pointer as a packed repeated Uint64.
func appendUint64PackedSlice(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.Uint64Slice()
	if len(s) == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	n := 0
	for _, v := range s {
		n += protowire.SizeVarint(v)
	}
	b = protowire.AppendVarint(b, uint64(n))
	for _, v := range s {
		b = protowire.AppendVarint(b, v)
	}
	return b, nil
}

var coderUint64PackedSlice = pointerCoderFuncs{
	size:      sizeUint64PackedSlice,
	marshal:   appendUint64PackedSlice,
	unmarshal: consumeUint64Slice,
	merge:     mergeUint64Slice,
}

// sizeUint64Value returns the size of wire encoding a uint64 value as a Uint64.
func sizeUint64Value(v protoreflect.Value, tagsize int, opts marshalOptions) int {
	return tagsize + protowire.SizeVarint(v.Uint())
}

// appendUint64Value encodes a uint64 value as a Uint64.
func appendUint64Value(b []byte, v protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	b = protowire.AppendVarint(b, wiretag)
	b = protowire.AppendVarint(b, v.Uint())
	return b, nil
}

// consumeUint64Value decodes a uint64 value as a Uint64.
func consumeUint64Value(b []byte, _ protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	if wtyp != protowire.VarintType {
		return protoreflect.Value{}, out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	out.n = n
	return protoreflect.ValueOfUint64(v), out, nil
}

var coderUint64Value = valueCoderFuncs{
	size:      sizeUint64Value,
	marshal:   appendUint64Value,
	unmarshal: consumeUint64Value,
	merge:     mergeScalarValue,
}

// sizeUint64SliceValue returns the size of wire encoding a []uint64 value as a repeated Uint64.
func sizeUint64SliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	for i, llen := 0, list.Len(); i < llen; i++ {
		v := list.Get(i)
		size += tagsize + protowire.SizeVarint(v.Uint())
	}
	return size
}

// appendUint64SliceValue encodes a []uint64 value as a repeated Uint64.
func appendUint64SliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	for i, llen := 0, list.Len(); i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendVarint(b, wiretag)
		b = protowire.AppendVarint(b, v.Uint())
	}
	return b, nil
}

// consumeUint64SliceValue wire decodes a []uint64 value as a repeated Uint64.
func consumeUint64SliceValue(b []byte, listv protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	list := listv.List()
	if wtyp == protowire.BytesType {
		b, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return protoreflect.Value{}, out, errDecode
		}
		for len(b) > 0 {
			var v uint64
			var n int
			if len(b) >= 1 && b[0] < 0x80 {
				v = uint64(b[0])
				n = 1
			} else if len(b) >= 2 && b[1] < 128 {
				v = uint64(b[0]&0x7f) + uint64(b[1])<<7
				n = 2
			} else {
				v, n = protowire.ConsumeVarint(b)
			}
			if n < 0 {
				return protoreflect.Value{}, out, errDecode
			}
			list.Append(protoreflect.ValueOfUint64(v))
			b = b[n:]
		}
		out.n = n
		return listv, out, nil
	}
	if wtyp != protowire.VarintType {
		return protoreflect.Value{}, out, errUnknown
	}
	var v uint64
	var n int
	if len(b) >= 1 && b[0] < 0x80 {
		v = uint64(b[0])
		n = 1
	} else if len(b) >= 2 && b[1] < 128 {
		v = uint64(b[0]&0x7f) + uint64(b[1])<<7
		n = 2
	} else {
		v, n = protowire.ConsumeVarint(b)
	}
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	list.Append(protoreflect.ValueOfUint64(v))
	out.n = n
	return listv, out, nil
}

var coderUint64SliceValue = valueCoderFuncs{
	size:      sizeUint64SliceValue,
	marshal:   appendUint64SliceValue,
	unmarshal: consumeUint64SliceValue,
	merge:     mergeListValue,
}

// sizeUint64PackedSliceValue returns the size of wire encoding a []uint64 value as a packed repeated Uint64.
func sizeUint64PackedSliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	llen := list.Len()
	if llen == 0 {
		return 0
	}
	n := 0
	for i, llen := 0, llen; i < llen; i++ {
		v := list.Get(i)
		n += protowire.SizeVarint(v.Uint())
	}
	return tagsize + protowire.SizeBytes(n)
}

// appendUint64PackedSliceValue encodes a []uint64 value as a packed repeated Uint64.
func appendUint64PackedSliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	llen := list.Len()
	if llen == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, wiretag)
	n := 0
	for i := 0; i < llen; i++ {
		v := list.Get(i)
		n += protowire.SizeVarint(v.Uint())
	}
	b = protowire.AppendVarint(b, uint64(n))
	for i := 0; i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendVarint(b, v.Uint())
	}
	return b, nil
}

var coderUint64PackedSliceValue = valueCoderFuncs{
	size:      sizeUint64PackedSliceValue,
	marshal:   appendUint64PackedSliceValue,
	unmarshal: consumeUint64SliceValue,
	merge:     mergeListValue,
}

// sizeSfixed32 returns the size of wire encoding a int32 pointer as a Sfixed32.
func sizeSfixed32(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {

	return f.tagsize + protowire.SizeFixed32()
}

// appendSfixed32 wire encodes a int32 pointer as a Sfixed32.
func appendSfixed32(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Int32()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendFixed32(b, uint32(v))
	return b, nil
}

// consumeSfixed32 wire decodes a int32 pointer as a Sfixed32.
func consumeSfixed32(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.Fixed32Type {
		return out, errUnknown
	}
	v, n := protowire.ConsumeFixed32(b)
	if n < 0 {
		return out, errDecode
	}
	*p.Int32() = int32(v)
	out.n = n
	return out, nil
}

var coderSfixed32 = pointerCoderFuncs{
	size:      sizeSfixed32,
	marshal:   appendSfixed32,
	unmarshal: consumeSfixed32,
	merge:     mergeInt32,
}

// sizeSfixed32NoZero returns the size of wire encoding a int32 pointer as a Sfixed32.
// The zero value is not encoded.
func sizeSfixed32NoZero(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := *p.Int32()
	if v == 0 {
		return 0
	}
	return f.tagsize + protowire.SizeFixed32()
}

// appendSfixed32NoZero wire encodes a int32 pointer as a Sfixed32.
// The zero value is not encoded.
func appendSfixed32NoZero(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Int32()
	if v == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendFixed32(b, uint32(v))
	return b, nil
}

var coderSfixed32NoZero = pointerCoderFuncs{
	size:      sizeSfixed32NoZero,
	marshal:   appendSfixed32NoZero,
	unmarshal: consumeSfixed32,
	merge:     mergeInt32NoZero,
}

// sizeSfixed32Ptr returns the size of wire encoding a *int32 pointer as a Sfixed32.
// It panics if the pointer is nil.
func sizeSfixed32Ptr(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	return f.tagsize + protowire.SizeFixed32()
}

// appendSfixed32Ptr wire encodes a *int32 pointer as a Sfixed32.
// It panics if the pointer is nil.
func appendSfixed32Ptr(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := **p.Int32Ptr()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendFixed32(b, uint32(v))
	return b, nil
}

// consumeSfixed32Ptr wire decodes a *int32 pointer as a Sfixed32.
func consumeSfixed32Ptr(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.Fixed32Type {
		return out, errUnknown
	}
	v, n := protowire.ConsumeFixed32(b)
	if n < 0 {
		return out, errDecode
	}
	vp := p.Int32Ptr()
	if *vp == nil {
		*vp = new(int32)
	}
	**vp = int32(v)
	out.n = n
	return out, nil
}

var coderSfixed32Ptr = pointerCoderFuncs{
	size:      sizeSfixed32Ptr,
	marshal:   appendSfixed32Ptr,
	unmarshal: consumeSfixed32Ptr,
	merge:     mergeInt32Ptr,
}

// sizeSfixed32Slice returns the size of wire encoding a []int32 pointer as a repeated Sfixed32.
func sizeSfixed32Slice(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	s := *p.Int32Slice()
	size = len(s) * (f.tagsize + protowire.SizeFixed32())
	return size
}

// appendSfixed32Slice encodes a []int32 pointer as a repeated Sfixed32.
func appendSfixed32Slice(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.Int32Slice()
	for _, v := range s {
		b = protowire.AppendVarint(b, f.wiretag)
		b = protowire.AppendFixed32(b, uint32(v))
	}
	return b, nil
}

// consumeSfixed32Slice wire decodes a []int32 pointer as a repeated Sfixed32.
func consumeSfixed32Slice(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	sp := p.Int32Slice()
	if wtyp == protowire.BytesType {
		s := *sp
		b, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return out, errDecode
		}
		for len(b) > 0 {
			v, n := protowire.ConsumeFixed32(b)
			if n < 0 {
				return out, errDecode
			}
			s = append(s, int32(v))
			b = b[n:]
		}
		*sp = s
		out.n = n
		return out, nil
	}
	if wtyp != protowire.Fixed32Type {
		return out, errUnknown
	}
	v, n := protowire.ConsumeFixed32(b)
	if n < 0 {
		return out, errDecode
	}
	*sp = append(*sp, int32(v))
	out.n = n
	return out, nil
}

var coderSfixed32Slice = pointerCoderFuncs{
	size:      sizeSfixed32Slice,
	marshal:   appendSfixed32Slice,
	unmarshal: consumeSfixed32Slice,
	merge:     mergeInt32Slice,
}

// sizeSfixed32PackedSlice returns the size of wire encoding a []int32 pointer as a packed repeated Sfixed32.
func sizeSfixed32PackedSlice(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	s := *p.Int32Slice()
	if len(s) == 0 {
		return 0
	}
	n := len(s) * protowire.SizeFixed32()
	return f.tagsize + protowire.SizeBytes(n)
}

// appendSfixed32PackedSlice encodes a []int32 pointer as a packed repeated Sfixed32.
func appendSfixed32PackedSlice(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.Int32Slice()
	if len(s) == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	n := len(s) * protowire.SizeFixed32()
	b = protowire.AppendVarint(b, uint64(n))
	for _, v := range s {
		b = protowire.AppendFixed32(b, uint32(v))
	}
	return b, nil
}

var coderSfixed32PackedSlice = pointerCoderFuncs{
	size:      sizeSfixed32PackedSlice,
	marshal:   appendSfixed32PackedSlice,
	unmarshal: consumeSfixed32Slice,
	merge:     mergeInt32Slice,
}

// sizeSfixed32Value returns the size of wire encoding a int32 value as a Sfixed32.
func sizeSfixed32Value(v protoreflect.Value, tagsize int, opts marshalOptions) int {
	return tagsize + protowire.SizeFixed32()
}

// appendSfixed32Value encodes a int32 value as a Sfixed32.
func appendSfixed32Value(b []byte, v protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	b = protowire.AppendVarint(b, wiretag)
	b = protowire.AppendFixed32(b, uint32(v.Int()))
	return b, nil
}

// consumeSfixed32Value decodes a int32 value as a Sfixed32.
func consumeSfixed32Value(b []byte, _ protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	if wtyp != protowire.Fixed32Type {
		return protoreflect.Value{}, out, errUnknown
	}
	v, n := protowire.ConsumeFixed32(b)
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	out.n = n
	return protoreflect.ValueOfInt32(int32(v)), out, nil
}

var coderSfixed32Value = valueCoderFuncs{
	size:      sizeSfixed32Value,
	marshal:   appendSfixed32Value,
	unmarshal: consumeSfixed32Value,
	merge:     mergeScalarValue,
}

// sizeSfixed32SliceValue returns the size of wire encoding a []int32 value as a repeated Sfixed32.
func sizeSfixed32SliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	size = list.Len() * (tagsize + protowire.SizeFixed32())
	return size
}

// appendSfixed32SliceValue encodes a []int32 value as a repeated Sfixed32.
func appendSfixed32SliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	for i, llen := 0, list.Len(); i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendVarint(b, wiretag)
		b = protowire.AppendFixed32(b, uint32(v.Int()))
	}
	return b, nil
}

// consumeSfixed32SliceValue wire decodes a []int32 value as a repeated Sfixed32.
func consumeSfixed32SliceValue(b []byte, listv protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	list := listv.List()
	if wtyp == protowire.BytesType {
		b, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return protoreflect.Value{}, out, errDecode
		}
		for len(b) > 0 {
			v, n := protowire.ConsumeFixed32(b)
			if n < 0 {
				return protoreflect.Value{}, out, errDecode
			}
			list.Append(protoreflect.ValueOfInt32(int32(v)))
			b = b[n:]
		}
		out.n = n
		return listv, out, nil
	}
	if wtyp != protowire.Fixed32Type {
		return protoreflect.Value{}, out, errUnknown
	}
	v, n := protowire.ConsumeFixed32(b)
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	list.Append(protoreflect.ValueOfInt32(int32(v)))
	out.n = n
	return listv, out, nil
}

var coderSfixed32SliceValue = valueCoderFuncs{
	size:      sizeSfixed32SliceValue,
	marshal:   appendSfixed32SliceValue,
	unmarshal: consumeSfixed32SliceValue,
	merge:     mergeListValue,
}

// sizeSfixed32PackedSliceValue returns the size of wire encoding a []int32 value as a packed repeated Sfixed32.
func sizeSfixed32PackedSliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	llen := list.Len()
	if llen == 0 {
		return 0
	}
	n := llen * protowire.SizeFixed32()
	return tagsize + protowire.SizeBytes(n)
}

// appendSfixed32PackedSliceValue encodes a []int32 value as a packed repeated Sfixed32.
func appendSfixed32PackedSliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	llen := list.Len()
	if llen == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, wiretag)
	n := llen * protowire.SizeFixed32()
	b = protowire.AppendVarint(b, uint64(n))
	for i := 0; i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendFixed32(b, uint32(v.Int()))
	}
	return b, nil
}

var coderSfixed32PackedSliceValue = valueCoderFuncs{
	size:      sizeSfixed32PackedSliceValue,
	marshal:   appendSfixed32PackedSliceValue,
	unmarshal: consumeSfixed32SliceValue,
	merge:     mergeListValue,
}

// sizeFixed32 returns the size of wire encoding a uint32 pointer as a Fixed32.
func sizeFixed32(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {

	return f.tagsize + protowire.SizeFixed32()
}

// appendFixed32 wire encodes a uint32 pointer as a Fixed32.
func appendFixed32(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Uint32()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendFixed32(b, v)
	return b, nil
}

// consumeFixed32 wire decodes a uint32 pointer as a Fixed32.
func consumeFixed32(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.Fixed32Type {
		return out, errUnknown
	}
	v, n := protowire.ConsumeFixed32(b)
	if n < 0 {
		return out, errDecode
	}
	*p.Uint32() = v
	out.n = n
	return out, nil
}

var coderFixed32 = pointerCoderFuncs{
	size:      sizeFixed32,
	marshal:   appendFixed32,
	unmarshal: consumeFixed32,
	merge:     mergeUint32,
}

// sizeFixed32NoZero returns the size of wire encoding a uint32 pointer as a Fixed32.
// The zero value is not encoded.
func sizeFixed32NoZero(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := *p.Uint32()
	if v == 0 {
		return 0
	}
	return f.tagsize + protowire.SizeFixed32()
}

// appendFixed32NoZero wire encodes a uint32 pointer as a Fixed32.
// The zero value is not encoded.
func appendFixed32NoZero(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Uint32()
	if v == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendFixed32(b, v)
	return b, nil
}

var coderFixed32NoZero = pointerCoderFuncs{
	size:      sizeFixed32NoZero,
	marshal:   appendFixed32NoZero,
	unmarshal: consumeFixed32,
	merge:     mergeUint32NoZero,
}

// sizeFixed32Ptr returns the size of wire encoding a *uint32 pointer as a Fixed32.
// It panics if the pointer is nil.
func sizeFixed32Ptr(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	return f.tagsize + protowire.SizeFixed32()
}

// appendFixed32Ptr wire encodes a *uint32 pointer as a Fixed32.
// It panics if the pointer is nil.
func appendFixed32Ptr(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := **p.Uint32Ptr()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendFixed32(b, v)
	return b, nil
}

// consumeFixed32Ptr wire decodes a *uint32 pointer as a Fixed32.
func consumeFixed32Ptr(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.Fixed32Type {
		return out, errUnknown
	}
	v, n := protowire.ConsumeFixed32(b)
	if n < 0 {
		return out, errDecode
	}
	vp := p.Uint32Ptr()
	if *vp == nil {
		*vp = new(uint32)
	}
	**vp = v
	out.n = n
	return out, nil
}

var coderFixed32Ptr = pointerCoderFuncs{
	size:      sizeFixed32Ptr,
	marshal:   appendFixed32Ptr,
	unmarshal: consumeFixed32Ptr,
	merge:     mergeUint32Ptr,
}

// sizeFixed32Slice returns the size of wire encoding a []uint32 pointer as a repeated Fixed32.
func sizeFixed32Slice(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	s := *p.Uint32Slice()
	size = len(s) * (f.tagsize + protowire.SizeFixed32())
	return size
}

// appendFixed32Slice encodes a []uint32 pointer as a repeated Fixed32.
func appendFixed32Slice(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.Uint32Slice()
	for _, v := range s {
		b = protowire.AppendVarint(b, f.wiretag)
		b = protowire.AppendFixed32(b, v)
	}
	return b, nil
}

// consumeFixed32Slice wire decodes a []uint32 pointer as a repeated Fixed32.
func consumeFixed32Slice(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	sp := p.Uint32Slice()
	if wtyp == protowire.BytesType {
		s := *sp
		b, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return out, errDecode
		}
		for len(b) > 0 {
			v, n := protowire.ConsumeFixed32(b)
			if n < 0 {
				return out, errDecode
			}
			s = append(s, v)
			b = b[n:]
		}
		*sp = s
		out.n = n
		return out, nil
	}
	if wtyp != protowire.Fixed32Type {
		return out, errUnknown
	}
	v, n := protowire.ConsumeFixed32(b)
	if n < 0 {
		return out, errDecode
	}
	*sp = append(*sp, v)
	out.n = n
	return out, nil
}

var coderFixed32Slice = pointerCoderFuncs{
	size:      sizeFixed32Slice,
	marshal:   appendFixed32Slice,
	unmarshal: consumeFixed32Slice,
	merge:     mergeUint32Slice,
}

// sizeFixed32PackedSlice returns the size of wire encoding a []uint32 pointer as a packed repeated Fixed32.
func sizeFixed32PackedSlice(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	s := *p.Uint32Slice()
	if len(s) == 0 {
		return 0
	}
	n := len(s) * protowire.SizeFixed32()
	return f.tagsize + protowire.SizeBytes(n)
}

// appendFixed32PackedSlice encodes a []uint32 pointer as a packed repeated Fixed32.
func appendFixed32PackedSlice(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.Uint32Slice()
	if len(s) == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	n := len(s) * protowire.SizeFixed32()
	b = protowire.AppendVarint(b, uint64(n))
	for _, v := range s {
		b = protowire.AppendFixed32(b, v)
	}
	return b, nil
}

var coderFixed32PackedSlice = pointerCoderFuncs{
	size:      sizeFixed32PackedSlice,
	marshal:   appendFixed32PackedSlice,
	unmarshal: consumeFixed32Slice,
	merge:     mergeUint32Slice,
}

// sizeFixed32Value returns the size of wire encoding a uint32 value as a Fixed32.
func sizeFixed32Value(v protoreflect.Value, tagsize int, opts marshalOptions) int {
	return tagsize + protowire.SizeFixed32()
}

// appendFixed32Value encodes a uint32 value as a Fixed32.
func appendFixed32Value(b []byte, v protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	b = protowire.AppendVarint(b, wiretag)
	b = protowire.AppendFixed32(b, uint32(v.Uint()))
	return b, nil
}

// consumeFixed32Value decodes a uint32 value as a Fixed32.
func consumeFixed32Value(b []byte, _ protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	if wtyp != protowire.Fixed32Type {
		return protoreflect.Value{}, out, errUnknown
	}
	v, n := protowire.ConsumeFixed32(b)
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	out.n = n
	return protoreflect.ValueOfUint32(uint32(v)), out, nil
}

var coderFixed32Value = valueCoderFuncs{
	size:      sizeFixed32Value,
	marshal:   appendFixed32Value,
	unmarshal: consumeFixed32Value,
	merge:     mergeScalarValue,
}

// sizeFixed32SliceValue returns the size of wire encoding a []uint32 value as a repeated Fixed32.
func sizeFixed32SliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	size = list.Len() * (tagsize + protowire.SizeFixed32())
	return size
}

// appendFixed32SliceValue encodes a []uint32 value as a repeated Fixed32.
func appendFixed32SliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	for i, llen := 0, list.Len(); i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendVarint(b, wiretag)
		b = protowire.AppendFixed32(b, uint32(v.Uint()))
	}
	return b, nil
}

// consumeFixed32SliceValue wire decodes a []uint32 value as a repeated Fixed32.
func consumeFixed32SliceValue(b []byte, listv protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	list := listv.List()
	if wtyp == protowire.BytesType {
		b, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return protoreflect.Value{}, out, errDecode
		}
		for len(b) > 0 {
			v, n := protowire.ConsumeFixed32(b)
			if n < 0 {
				return protoreflect.Value{}, out, errDecode
			}
			list.Append(protoreflect.ValueOfUint32(uint32(v)))
			b = b[n:]
		}
		out.n = n
		return listv, out, nil
	}
	if wtyp != protowire.Fixed32Type {
		return protoreflect.Value{}, out, errUnknown
	}
	v, n := protowire.ConsumeFixed32(b)
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	list.Append(protoreflect.ValueOfUint32(uint32(v)))
	out.n = n
	return listv, out, nil
}

var coderFixed32SliceValue = valueCoderFuncs{
	size:      sizeFixed32SliceValue,
	marshal:   appendFixed32SliceValue,
	unmarshal: consumeFixed32SliceValue,
	merge:     mergeListValue,
}

// sizeFixed32PackedSliceValue returns the size of wire encoding a []uint32 value as a packed repeated Fixed32.
func sizeFixed32PackedSliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	llen := list.Len()
	if llen == 0 {
		return 0
	}
	n := llen * protowire.SizeFixed32()
	return tagsize + protowire.SizeBytes(n)
}

// appendFixed32PackedSliceValue encodes a []uint32 value as a packed repeated Fixed32.
func appendFixed32PackedSliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	llen := list.Len()
	if llen == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, wiretag)
	n := llen * protowire.SizeFixed32()
	b = protowire.AppendVarint(b, uint64(n))
	for i := 0; i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendFixed32(b, uint32(v.Uint()))
	}
	return b, nil
}

var coderFixed32PackedSliceValue = valueCoderFuncs{
	size:      sizeFixed32PackedSliceValue,
	marshal:   appendFixed32PackedSliceValue,
	unmarshal: consumeFixed32SliceValue,
	merge:     mergeListValue,
}

// sizeFloat returns the size of wire encoding a float32 pointer as a Float.
func sizeFloat(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {

	return f.tagsize + protowire.SizeFixed32()
}

// appendFloat wire encodes a float32 pointer as a Float.
func appendFloat(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Float32()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendFixed32(b, math.Float32bits(v))
	return b, nil
}

// consumeFloat wire decodes a float32 pointer as a Float.
func consumeFloat(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.Fixed32Type {
		return out, errUnknown
	}
	v, n := protowire.ConsumeFixed32(b)
	if n < 0 {
		return out, errDecode
	}
	*p.Float32() = math.Float32frombits(v)
	out.n = n
	return out, nil
}

var coderFloat = pointerCoderFuncs{
	size:      sizeFloat,
	marshal:   appendFloat,
	unmarshal: consumeFloat,
	merge:     mergeFloat32,
}

// sizeFloatNoZero returns the size of wire encoding a float32 pointer as a Float.
// The zero value is not encoded.
func sizeFloatNoZero(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := *p.Float32()
	if v == 0 && !math.Signbit(float64(v)) {
		return 0
	}
	return f.tagsize + protowire.SizeFixed32()
}

// appendFloatNoZero wire encodes a float32 pointer as a Float.
// The zero value is not encoded.
func appendFloatNoZero(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Float32()
	if v == 0 && !math.Signbit(float64(v)) {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendFixed32(b, math.Float32bits(v))
	return b, nil
}

var coderFloatNoZero = pointerCoderFuncs{
	size:      sizeFloatNoZero,
	marshal:   appendFloatNoZero,
	unmarshal: consumeFloat,
	merge:     mergeFloat32NoZero,
}

// sizeFloatPtr returns the size of wire encoding a *float32 pointer as a Float.
// It panics if the pointer is nil.
func sizeFloatPtr(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	return f.tagsize + protowire.SizeFixed32()
}

// appendFloatPtr wire encodes a *float32 pointer as a Float.
// It panics if the pointer is nil.
func appendFloatPtr(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := **p.Float32Ptr()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendFixed32(b, math.Float32bits(v))
	return b, nil
}

// consumeFloatPtr wire decodes a *float32 pointer as a Float.
func consumeFloatPtr(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.Fixed32Type {
		return out, errUnknown
	}
	v, n := protowire.ConsumeFixed32(b)
	if n < 0 {
		return out, errDecode
	}
	vp := p.Float32Ptr()
	if *vp == nil {
		*vp = new(float32)
	}
	**vp = math.Float32frombits(v)
	out.n = n
	return out, nil
}

var coderFloatPtr = pointerCoderFuncs{
	size:      sizeFloatPtr,
	marshal:   appendFloatPtr,
	unmarshal: consumeFloatPtr,
	merge:     mergeFloat32Ptr,
}

// sizeFloatSlice returns the size of wire encoding a []float32 pointer as a repeated Float.
func sizeFloatSlice(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	s := *p.Float32Slice()
	size = len(s) * (f.tagsize + protowire.SizeFixed32())
	return size
}

// appendFloatSlice encodes a []float32 pointer as a repeated Float.
func appendFloatSlice(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.Float32Slice()
	for _, v := range s {
		b = protowire.AppendVarint(b, f.wiretag)
		b = protowire.AppendFixed32(b, math.Float32bits(v))
	}
	return b, nil
}

// consumeFloatSlice wire decodes a []float32 pointer as a repeated Float.
func consumeFloatSlice(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	sp := p.Float32Slice()
	if wtyp == protowire.BytesType {
		s := *sp
		b, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return out, errDecode
		}
		for len(b) > 0 {
			v, n := protowire.ConsumeFixed32(b)
			if n < 0 {
				return out, errDecode
			}
			s = append(s, math.Float32frombits(v))
			b = b[n:]
		}
		*sp = s
		out.n = n
		return out, nil
	}
	if wtyp != protowire.Fixed32Type {
		return out, errUnknown
	}
	v, n := protowire.ConsumeFixed32(b)
	if n < 0 {
		return out, errDecode
	}
	*sp = append(*sp, math.Float32frombits(v))
	out.n = n
	return out, nil
}

var coderFloatSlice = pointerCoderFuncs{
	size:      sizeFloatSlice,
	marshal:   appendFloatSlice,
	unmarshal: consumeFloatSlice,
	merge:     mergeFloat32Slice,
}

// sizeFloatPackedSlice returns the size of wire encoding a []float32 pointer as a packed repeated Float.
func sizeFloatPackedSlice(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	s := *p.Float32Slice()
	if len(s) == 0 {
		return 0
	}
	n := len(s) * protowire.SizeFixed32()
	return f.tagsize + protowire.SizeBytes(n)
}

// appendFloatPackedSlice encodes a []float32 pointer as a packed repeated Float.
func appendFloatPackedSlice(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.Float32Slice()
	if len(s) == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	n := len(s) * protowire.SizeFixed32()
	b = protowire.AppendVarint(b, uint64(n))
	for _, v := range s {
		b = protowire.AppendFixed32(b, math.Float32bits(v))
	}
	return b, nil
}

var coderFloatPackedSlice = pointerCoderFuncs{
	size:      sizeFloatPackedSlice,
	marshal:   appendFloatPackedSlice,
	unmarshal: consumeFloatSlice,
	merge:     mergeFloat32Slice,
}

// sizeFloatValue returns the size of wire encoding a float32 value as a Float.
func sizeFloatValue(v protoreflect.Value, tagsize int, opts marshalOptions) int {
	return tagsize + protowire.SizeFixed32()
}

// appendFloatValue encodes a float32 value as a Float.
func appendFloatValue(b []byte, v protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	b = protowire.AppendVarint(b, wiretag)
	b = protowire.AppendFixed32(b, math.Float32bits(float32(v.Float())))
	return b, nil
}

// consumeFloatValue decodes a float32 value as a Float.
func consumeFloatValue(b []byte, _ protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	if wtyp != protowire.Fixed32Type {
		return protoreflect.Value{}, out, errUnknown
	}
	v, n := protowire.ConsumeFixed32(b)
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	out.n = n
	return protoreflect.ValueOfFloat32(math.Float32frombits(uint32(v))), out, nil
}

var coderFloatValue = valueCoderFuncs{
	size:      sizeFloatValue,
	marshal:   appendFloatValue,
	unmarshal: consumeFloatValue,
	merge:     mergeScalarValue,
}

// sizeFloatSliceValue returns the size of wire encoding a []float32 value as a repeated Float.
func sizeFloatSliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	size = list.Len() * (tagsize + protowire.SizeFixed32())
	return size
}

// appendFloatSliceValue encodes a []float32 value as a repeated Float.
func appendFloatSliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	for i, llen := 0, list.Len(); i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendVarint(b, wiretag)
		b = protowire.AppendFixed32(b, math.Float32bits(float32(v.Float())))
	}
	return b, nil
}

// consumeFloatSliceValue wire decodes a []float32 value as a repeated Float.
func consumeFloatSliceValue(b []byte, listv protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	list := listv.List()
	if wtyp == protowire.BytesType {
		b, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return protoreflect.Value{}, out, errDecode
		}
		for len(b) > 0 {
			v, n := protowire.ConsumeFixed32(b)
			if n < 0 {
				return protoreflect.Value{}, out, errDecode
			}
			list.Append(protoreflect.ValueOfFloat32(math.Float32frombits(uint32(v))))
			b = b[n:]
		}
		out.n = n
		return listv, out, nil
	}
	if wtyp != protowire.Fixed32Type {
		return protoreflect.Value{}, out, errUnknown
	}
	v, n := protowire.ConsumeFixed32(b)
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	list.Append(protoreflect.ValueOfFloat32(math.Float32frombits(uint32(v))))
	out.n = n
	return listv, out, nil
}

var coderFloatSliceValue = valueCoderFuncs{
	size:      sizeFloatSliceValue,
	marshal:   appendFloatSliceValue,
	unmarshal: consumeFloatSliceValue,
	merge:     mergeListValue,
}

// sizeFloatPackedSliceValue returns the size of wire encoding a []float32 value as a packed repeated Float.
func sizeFloatPackedSliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	llen := list.Len()
	if llen == 0 {
		return 0
	}
	n := llen * protowire.SizeFixed32()
	return tagsize + protowire.SizeBytes(n)
}

// appendFloatPackedSliceValue encodes a []float32 value as a packed repeated Float.
func appendFloatPackedSliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	llen := list.Len()
	if llen == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, wiretag)
	n := llen * protowire.SizeFixed32()
	b = protowire.AppendVarint(b, uint64(n))
	for i := 0; i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendFixed32(b, math.Float32bits(float32(v.Float())))
	}
	return b, nil
}

var coderFloatPackedSliceValue = valueCoderFuncs{
	size:      sizeFloatPackedSliceValue,
	marshal:   appendFloatPackedSliceValue,
	unmarshal: consumeFloatSliceValue,
	merge:     mergeListValue,
}

// sizeSfixed64 returns the size of wire encoding a int64 pointer as a Sfixed64.
func sizeSfixed64(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {

	return f.tagsize + protowire.SizeFixed64()
}

// appendSfixed64 wire encodes a int64 pointer as a Sfixed64.
func appendSfixed64(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Int64()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendFixed64(b, uint64(v))
	return b, nil
}

// consumeSfixed64 wire decodes a int64 pointer as a Sfixed64.
func consumeSfixed64(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.Fixed64Type {
		return out, errUnknown
	}
	v, n := protowire.ConsumeFixed64(b)
	if n < 0 {
		return out, errDecode
	}
	*p.Int64() = int64(v)
	out.n = n
	return out, nil
}

var coderSfixed64 = pointerCoderFuncs{
	size:      sizeSfixed64,
	marshal:   appendSfixed64,
	unmarshal: consumeSfixed64,
	merge:     mergeInt64,
}

// sizeSfixed64NoZero returns the size of wire encoding a int64 pointer as a Sfixed64.
// The zero value is not encoded.
func sizeSfixed64NoZero(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := *p.Int64()
	if v == 0 {
		return 0
	}
	return f.tagsize + protowire.SizeFixed64()
}

// appendSfixed64NoZero wire encodes a int64 pointer as a Sfixed64.
// The zero value is not encoded.
func appendSfixed64NoZero(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Int64()
	if v == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendFixed64(b, uint64(v))
	return b, nil
}

var coderSfixed64NoZero = pointerCoderFuncs{
	size:      sizeSfixed64NoZero,
	marshal:   appendSfixed64NoZero,
	unmarshal: consumeSfixed64,
	merge:     mergeInt64NoZero,
}

// sizeSfixed64Ptr returns the size of wire encoding a *int64 pointer as a Sfixed64.
// It panics if the pointer is nil.
func sizeSfixed64Ptr(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	return f.tagsize + protowire.SizeFixed64()
}

// appendSfixed64Ptr wire encodes a *int64 pointer as a Sfixed64.
// It panics if the pointer is nil.
func appendSfixed64Ptr(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := **p.Int64Ptr()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendFixed64(b, uint64(v))
	return b, nil
}

// consumeSfixed64Ptr wire decodes a *int64 pointer as a Sfixed64.
func consumeSfixed64Ptr(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.Fixed64Type {
		return out, errUnknown
	}
	v, n := protowire.ConsumeFixed64(b)
	if n < 0 {
		return out, errDecode
	}
	vp := p.Int64Ptr()
	if *vp == nil {
		*vp = new(int64)
	}
	**vp = int64(v)
	out.n = n
	return out, nil
}

var coderSfixed64Ptr = pointerCoderFuncs{
	size:      sizeSfixed64Ptr,
	marshal:   appendSfixed64Ptr,
	unmarshal: consumeSfixed64Ptr,
	merge:     mergeInt64Ptr,
}

// sizeSfixed64Slice returns the size of wire encoding a []int64 pointer as a repeated Sfixed64.
func sizeSfixed64Slice(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	s := *p.Int64Slice()
	size = len(s) * (f.tagsize + protowire.SizeFixed64())
	return size
}

// appendSfixed64Slice encodes a []int64 pointer as a repeated Sfixed64.
func appendSfixed64Slice(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.Int64Slice()
	for _, v := range s {
		b = protowire.AppendVarint(b, f.wiretag)
		b = protowire.AppendFixed64(b, uint64(v))
	}
	return b, nil
}

// consumeSfixed64Slice wire decodes a []int64 pointer as a repeated Sfixed64.
func consumeSfixed64Slice(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	sp := p.Int64Slice()
	if wtyp == protowire.BytesType {
		s := *sp
		b, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return out, errDecode
		}
		for len(b) > 0 {
			v, n := protowire.ConsumeFixed64(b)
			if n < 0 {
				return out, errDecode
			}
			s = append(s, int64(v))
			b = b[n:]
		}
		*sp = s
		out.n = n
		return out, nil
	}
	if wtyp != protowire.Fixed64Type {
		return out, errUnknown
	}
	v, n := protowire.ConsumeFixed64(b)
	if n < 0 {
		return out, errDecode
	}
	*sp = append(*sp, int64(v))
	out.n = n
	return out, nil
}

var coderSfixed64Slice = pointerCoderFuncs{
	size:      sizeSfixed64Slice,
	marshal:   appendSfixed64Slice,
	unmarshal: consumeSfixed64Slice,
	merge:     mergeInt64Slice,
}

// sizeSfixed64PackedSlice returns the size of wire encoding a []int64 pointer as a packed repeated Sfixed64.
func sizeSfixed64PackedSlice(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	s := *p.Int64Slice()
	if len(s) == 0 {
		return 0
	}
	n := len(s) * protowire.SizeFixed64()
	return f.tagsize + protowire.SizeBytes(n)
}

// appendSfixed64PackedSlice encodes a []int64 pointer as a packed repeated Sfixed64.
func appendSfixed64PackedSlice(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.Int64Slice()
	if len(s) == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	n := len(s) * protowire.SizeFixed64()
	b = protowire.AppendVarint(b, uint64(n))
	for _, v := range s {
		b = protowire.AppendFixed64(b, uint64(v))
	}
	return b, nil
}

var coderSfixed64PackedSlice = pointerCoderFuncs{
	size:      sizeSfixed64PackedSlice,
	marshal:   appendSfixed64PackedSlice,
	unmarshal: consumeSfixed64Slice,
	merge:     mergeInt64Slice,
}

// sizeSfixed64Value returns the size of wire encoding a int64 value as a Sfixed64.
func sizeSfixed64Value(v protoreflect.Value, tagsize int, opts marshalOptions) int {
	return tagsize + protowire.SizeFixed64()
}

// appendSfixed64Value encodes a int64 value as a Sfixed64.
func appendSfixed64Value(b []byte, v protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	b = protowire.AppendVarint(b, wiretag)
	b = protowire.AppendFixed64(b, uint64(v.Int()))
	return b, nil
}

// consumeSfixed64Value decodes a int64 value as a Sfixed64.
func consumeSfixed64Value(b []byte, _ protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	if wtyp != protowire.Fixed64Type {
		return protoreflect.Value{}, out, errUnknown
	}
	v, n := protowire.ConsumeFixed64(b)
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	out.n = n
	return protoreflect.ValueOfInt64(int64(v)), out, nil
}

var coderSfixed64Value = valueCoderFuncs{
	size:      sizeSfixed64Value,
	marshal:   appendSfixed64Value,
	unmarshal: consumeSfixed64Value,
	merge:     mergeScalarValue,
}

// sizeSfixed64SliceValue returns the size of wire encoding a []int64 value as a repeated Sfixed64.
func sizeSfixed64SliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	size = list.Len() * (tagsize + protowire.SizeFixed64())
	return size
}

// appendSfixed64SliceValue encodes a []int64 value as a repeated Sfixed64.
func appendSfixed64SliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	for i, llen := 0, list.Len(); i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendVarint(b, wiretag)
		b = protowire.AppendFixed64(b, uint64(v.Int()))
	}
	return b, nil
}

// consumeSfixed64SliceValue wire decodes a []int64 value as a repeated Sfixed64.
func consumeSfixed64SliceValue(b []byte, listv protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	list := listv.List()
	if wtyp == protowire.BytesType {
		b, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return protoreflect.Value{}, out, errDecode
		}
		for len(b) > 0 {
			v, n := protowire.ConsumeFixed64(b)
			if n < 0 {
				return protoreflect.Value{}, out, errDecode
			}
			list.Append(protoreflect.ValueOfInt64(int64(v)))
			b = b[n:]
		}
		out.n = n
		return listv, out, nil
	}
	if wtyp != protowire.Fixed64Type {
		return protoreflect.Value{}, out, errUnknown
	}
	v, n := protowire.ConsumeFixed64(b)
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	list.Append(protoreflect.ValueOfInt64(int64(v)))
	out.n = n
	return listv, out, nil
}

var coderSfixed64SliceValue = valueCoderFuncs{
	size:      sizeSfixed64SliceValue,
	marshal:   appendSfixed64SliceValue,
	unmarshal: consumeSfixed64SliceValue,
	merge:     mergeListValue,
}

// sizeSfixed64PackedSliceValue returns the size of wire encoding a []int64 value as a packed repeated Sfixed64.
func sizeSfixed64PackedSliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	llen := list.Len()
	if llen == 0 {
		return 0
	}
	n := llen * protowire.SizeFixed64()
	return tagsize + protowire.SizeBytes(n)
}

// appendSfixed64PackedSliceValue encodes a []int64 value as a packed repeated Sfixed64.
func appendSfixed64PackedSliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	llen := list.Len()
	if llen == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, wiretag)
	n := llen * protowire.SizeFixed64()
	b = protowire.AppendVarint(b, uint64(n))
	for i := 0; i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendFixed64(b, uint64(v.Int()))
	}
	return b, nil
}

var coderSfixed64PackedSliceValue = valueCoderFuncs{
	size:      sizeSfixed64PackedSliceValue,
	marshal:   appendSfixed64PackedSliceValue,
	unmarshal: consumeSfixed64SliceValue,
	merge:     mergeListValue,
}

// sizeFixed64 returns the size of wire encoding a uint64 pointer as a Fixed64.
func sizeFixed64(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {

	return f.tagsize + protowire.SizeFixed64()
}

// appendFixed64 wire encodes a uint64 pointer as a Fixed64.
func appendFixed64(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Uint64()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendFixed64(b, v)
	return b, nil
}

// consumeFixed64 wire decodes a uint64 pointer as a Fixed64.
func consumeFixed64(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.Fixed64Type {
		return out, errUnknown
	}
	v, n := protowire.ConsumeFixed64(b)
	if n < 0 {
		return out, errDecode
	}
	*p.Uint64() = v
	out.n = n
	return out, nil
}

var coderFixed64 = pointerCoderFuncs{
	size:      sizeFixed64,
	marshal:   appendFixed64,
	unmarshal: consumeFixed64,
	merge:     mergeUint64,
}

// sizeFixed64NoZero returns the size of wire encoding a uint64 pointer as a Fixed64.
// The zero value is not encoded.
func sizeFixed64NoZero(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := *p.Uint64()
	if v == 0 {
		return 0
	}
	return f.tagsize + protowire.SizeFixed64()
}

// appendFixed64NoZero wire encodes a uint64 pointer as a Fixed64.
// The zero value is not encoded.
func appendFixed64NoZero(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Uint64()
	if v == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendFixed64(b, v)
	return b, nil
}

var coderFixed64NoZero = pointerCoderFuncs{
	size:      sizeFixed64NoZero,
	marshal:   appendFixed64NoZero,
	unmarshal: consumeFixed64,
	merge:     mergeUint64NoZero,
}

// sizeFixed64Ptr returns the size of wire encoding a *uint64 pointer as a Fixed64.
// It panics if the pointer is nil.
func sizeFixed64Ptr(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	return f.tagsize + protowire.SizeFixed64()
}

// appendFixed64Ptr wire encodes a *uint64 pointer as a Fixed64.
// It panics if the pointer is nil.
func appendFixed64Ptr(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := **p.Uint64Ptr()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendFixed64(b, v)
	return b, nil
}

// consumeFixed64Ptr wire decodes a *uint64 pointer as a Fixed64.
func consumeFixed64Ptr(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.Fixed64Type {
		return out, errUnknown
	}
	v, n := protowire.ConsumeFixed64(b)
	if n < 0 {
		return out, errDecode
	}
	vp := p.Uint64Ptr()
	if *vp == nil {
		*vp = new(uint64)
	}
	**vp = v
	out.n = n
	return out, nil
}

var coderFixed64Ptr = pointerCoderFuncs{
	size:      sizeFixed64Ptr,
	marshal:   appendFixed64Ptr,
	unmarshal: consumeFixed64Ptr,
	merge:     mergeUint64Ptr,
}

// sizeFixed64Slice returns the size of wire encoding a []uint64 pointer as a repeated Fixed64.
func sizeFixed64Slice(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	s := *p.Uint64Slice()
	size = len(s) * (f.tagsize + protowire.SizeFixed64())
	return size
}

// appendFixed64Slice encodes a []uint64 pointer as a repeated Fixed64.
func appendFixed64Slice(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.Uint64Slice()
	for _, v := range s {
		b = protowire.AppendVarint(b, f.wiretag)
		b = protowire.AppendFixed64(b, v)
	}
	return b, nil
}

// consumeFixed64Slice wire decodes a []uint64 pointer as a repeated Fixed64.
func consumeFixed64Slice(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	sp := p.Uint64Slice()
	if wtyp == protowire.BytesType {
		s := *sp
		b, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return out, errDecode
		}
		for len(b) > 0 {
			v, n := protowire.ConsumeFixed64(b)
			if n < 0 {
				return out, errDecode
			}
			s = append(s, v)
			b = b[n:]
		}
		*sp = s
		out.n = n
		return out, nil
	}
	if wtyp != protowire.Fixed64Type {
		return out, errUnknown
	}
	v, n := protowire.ConsumeFixed64(b)
	if n < 0 {
		return out, errDecode
	}
	*sp = append(*sp, v)
	out.n = n
	return out, nil
}

var coderFixed64Slice = pointerCoderFuncs{
	size:      sizeFixed64Slice,
	marshal:   appendFixed64Slice,
	unmarshal: consumeFixed64Slice,
	merge:     mergeUint64Slice,
}

// sizeFixed64PackedSlice returns the size of wire encoding a []uint64 pointer as a packed repeated Fixed64.
func sizeFixed64PackedSlice(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	s := *p.Uint64Slice()
	if len(s) == 0 {
		return 0
	}
	n := len(s) * protowire.SizeFixed64()
	return f.tagsize + protowire.SizeBytes(n)
}

// appendFixed64PackedSlice encodes a []uint64 pointer as a packed repeated Fixed64.
func appendFixed64PackedSlice(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.Uint64Slice()
	if len(s) == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	n := len(s) * protowire.SizeFixed64()
	b = protowire.AppendVarint(b, uint64(n))
	for _, v := range s {
		b = protowire.AppendFixed64(b, v)
	}
	return b, nil
}

var coderFixed64PackedSlice = pointerCoderFuncs{
	size:      sizeFixed64PackedSlice,
	marshal:   appendFixed64PackedSlice,
	unmarshal: consumeFixed64Slice,
	merge:     mergeUint64Slice,
}

// sizeFixed64Value returns the size of wire encoding a uint64 value as a Fixed64.
func sizeFixed64Value(v protoreflect.Value, tagsize int, opts marshalOptions) int {
	return tagsize + protowire.SizeFixed64()
}

// appendFixed64Value encodes a uint64 value as a Fixed64.
func appendFixed64Value(b []byte, v protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	b = protowire.AppendVarint(b, wiretag)
	b = protowire.AppendFixed64(b, v.Uint())
	return b, nil
}

// consumeFixed64Value decodes a uint64 value as a Fixed64.
func consumeFixed64Value(b []byte, _ protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	if wtyp != protowire.Fixed64Type {
		return protoreflect.Value{}, out, errUnknown
	}
	v, n := protowire.ConsumeFixed64(b)
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	out.n = n
	return protoreflect.ValueOfUint64(v), out, nil
}

var coderFixed64Value = valueCoderFuncs{
	size:      sizeFixed64Value,
	marshal:   appendFixed64Value,
	unmarshal: consumeFixed64Value,
	merge:     mergeScalarValue,
}

// sizeFixed64SliceValue returns the size of wire encoding a []uint64 value as a repeated Fixed64.
func sizeFixed64SliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	size = list.Len() * (tagsize + protowire.SizeFixed64())
	return size
}

// appendFixed64SliceValue encodes a []uint64 value as a repeated Fixed64.
func appendFixed64SliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	for i, llen := 0, list.Len(); i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendVarint(b, wiretag)
		b = protowire.AppendFixed64(b, v.Uint())
	}
	return b, nil
}

// consumeFixed64SliceValue wire decodes a []uint64 value as a repeated Fixed64.
func consumeFixed64SliceValue(b []byte, listv protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	list := listv.List()
	if wtyp == protowire.BytesType {
		b, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return protoreflect.Value{}, out, errDecode
		}
		for len(b) > 0 {
			v, n := protowire.ConsumeFixed64(b)
			if n < 0 {
				return protoreflect.Value{}, out, errDecode
			}
			list.Append(protoreflect.ValueOfUint64(v))
			b = b[n:]
		}
		out.n = n
		return listv, out, nil
	}
	if wtyp != protowire.Fixed64Type {
		return protoreflect.Value{}, out, errUnknown
	}
	v, n := protowire.ConsumeFixed64(b)
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	list.Append(protoreflect.ValueOfUint64(v))
	out.n = n
	return listv, out, nil
}

var coderFixed64SliceValue = valueCoderFuncs{
	size:      sizeFixed64SliceValue,
	marshal:   appendFixed64SliceValue,
	unmarshal: consumeFixed64SliceValue,
	merge:     mergeListValue,
}

// sizeFixed64PackedSliceValue returns the size of wire encoding a []uint64 value as a packed repeated Fixed64.
func sizeFixed64PackedSliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	llen := list.Len()
	if llen == 0 {
		return 0
	}
	n := llen * protowire.SizeFixed64()
	return tagsize + protowire.SizeBytes(n)
}

// appendFixed64PackedSliceValue encodes a []uint64 value as a packed repeated Fixed64.
func appendFixed64PackedSliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	llen := list.Len()
	if llen == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, wiretag)
	n := llen * protowire.SizeFixed64()
	b = protowire.AppendVarint(b, uint64(n))
	for i := 0; i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendFixed64(b, v.Uint())
	}
	return b, nil
}

var coderFixed64PackedSliceValue = valueCoderFuncs{
	size:      sizeFixed64PackedSliceValue,
	marshal:   appendFixed64PackedSliceValue,
	unmarshal: consumeFixed64SliceValue,
	merge:     mergeListValue,
}

// sizeDouble returns the size of wire encoding a float64 pointer as a Double.
func sizeDouble(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {

	return f.tagsize + protowire.SizeFixed64()
}

// appendDouble wire encodes a float64 pointer as a Double.
func appendDouble(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Float64()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendFixed64(b, math.Float64bits(v))
	return b, nil
}

// consumeDouble wire decodes a float64 pointer as a Double.
func consumeDouble(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.Fixed64Type {
		return out, errUnknown
	}
	v, n := protowire.ConsumeFixed64(b)
	if n < 0 {
		return out, errDecode
	}
	*p.Float64() = math.Float64frombits(v)
	out.n = n
	return out, nil
}

var coderDouble = pointerCoderFuncs{
	size:      sizeDouble,
	marshal:   appendDouble,
	unmarshal: consumeDouble,
	merge:     mergeFloat64,
}

// sizeDoubleNoZero returns the size of wire encoding a float64 pointer as a Double.
// The zero value is not encoded.
func sizeDoubleNoZero(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := *p.Float64()
	if v == 0 && !math.Signbit(float64(v)) {
		return 0
	}
	return f.tagsize + protowire.SizeFixed64()
}

// appendDoubleNoZero wire encodes a float64 pointer as a Double.
// The zero value is not encoded.
func appendDoubleNoZero(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Float64()
	if v == 0 && !math.Signbit(float64(v)) {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendFixed64(b, math.Float64bits(v))
	return b, nil
}

var coderDoubleNoZero = pointerCoderFuncs{
	size:      sizeDoubleNoZero,
	marshal:   appendDoubleNoZero,
	unmarshal: consumeDouble,
	merge:     mergeFloat64NoZero,
}

// sizeDoublePtr returns the size of wire encoding a *float64 pointer as a Double.
// It panics if the pointer is nil.
func sizeDoublePtr(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	return f.tagsize + protowire.SizeFixed64()
}

// appendDoublePtr wire encodes a *float64 pointer as a Double.
// It panics if the pointer is nil.
func appendDoublePtr(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := **p.Float64Ptr()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendFixed64(b, math.Float64bits(v))
	return b, nil
}

// consumeDoublePtr wire decodes a *float64 pointer as a Double.
func consumeDoublePtr(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.Fixed64Type {
		return out, errUnknown
	}
	v, n := protowire.ConsumeFixed64(b)
	if n < 0 {
		return out, errDecode
	}
	vp := p.Float64Ptr()
	if *vp == nil {
		*vp = new(float64)
	}
	**vp = math.Float64frombits(v)
	out.n = n
	return out, nil
}

var coderDoublePtr = pointerCoderFuncs{
	size:      sizeDoublePtr,
	marshal:   appendDoublePtr,
	unmarshal: consumeDoublePtr,
	merge:     mergeFloat64Ptr,
}

// sizeDoubleSlice returns the size of wire encoding a []float64 pointer as a repeated Double.
func sizeDoubleSlice(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	s := *p.Float64Slice()
	size = len(s) * (f.tagsize + protowire.SizeFixed64())
	return size
}

// appendDoubleSlice encodes a []float64 pointer as a repeated Double.
func appendDoubleSlice(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.Float64Slice()
	for _, v := range s {
		b = protowire.AppendVarint(b, f.wiretag)
		b = protowire.AppendFixed64(b, math.Float64bits(v))
	}
	return b, nil
}

// consumeDoubleSlice wire decodes a []float64 pointer as a repeated Double.
func consumeDoubleSlice(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	sp := p.Float64Slice()
	if wtyp == protowire.BytesType {
		s := *sp
		b, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return out, errDecode
		}
		for len(b) > 0 {
			v, n := protowire.ConsumeFixed64(b)
			if n < 0 {
				return out, errDecode
			}
			s = append(s, math.Float64frombits(v))
			b = b[n:]
		}
		*sp = s
		out.n = n
		return out, nil
	}
	if wtyp != protowire.Fixed64Type {
		return out, errUnknown
	}
	v, n := protowire.ConsumeFixed64(b)
	if n < 0 {
		return out, errDecode
	}
	*sp = append(*sp, math.Float64frombits(v))
	out.n = n
	return out, nil
}

var coderDoubleSlice = pointerCoderFuncs{
	size:      sizeDoubleSlice,
	marshal:   appendDoubleSlice,
	unmarshal: consumeDoubleSlice,
	merge:     mergeFloat64Slice,
}

// sizeDoublePackedSlice returns the size of wire encoding a []float64 pointer as a packed repeated Double.
func sizeDoublePackedSlice(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	s := *p.Float64Slice()
	if len(s) == 0 {
		return 0
	}
	n := len(s) * protowire.SizeFixed64()
	return f.tagsize + protowire.SizeBytes(n)
}

// appendDoublePackedSlice encodes a []float64 pointer as a packed repeated Double.
func appendDoublePackedSlice(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.Float64Slice()
	if len(s) == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	n := len(s) * protowire.SizeFixed64()
	b = protowire.AppendVarint(b, uint64(n))
	for _, v := range s {
		b = protowire.AppendFixed64(b, math.Float64bits(v))
	}
	return b, nil
}

var coderDoublePackedSlice = pointerCoderFuncs{
	size:      sizeDoublePackedSlice,
	marshal:   appendDoublePackedSlice,
	unmarshal: consumeDoubleSlice,
	merge:     mergeFloat64Slice,
}

// sizeDoubleValue returns the size of wire encoding a float64 value as a Double.
func sizeDoubleValue(v protoreflect.Value, tagsize int, opts marshalOptions) int {
	return tagsize + protowire.SizeFixed64()
}

// appendDoubleValue encodes a float64 value as a Double.
func appendDoubleValue(b []byte, v protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	b = protowire.AppendVarint(b, wiretag)
	b = protowire.AppendFixed64(b, math.Float64bits(v.Float()))
	return b, nil
}

// consumeDoubleValue decodes a float64 value as a Double.
func consumeDoubleValue(b []byte, _ protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	if wtyp != protowire.Fixed64Type {
		return protoreflect.Value{}, out, errUnknown
	}
	v, n := protowire.ConsumeFixed64(b)
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	out.n = n
	return protoreflect.ValueOfFloat64(math.Float64frombits(v)), out, nil
}

var coderDoubleValue = valueCoderFuncs{
	size:      sizeDoubleValue,
	marshal:   appendDoubleValue,
	unmarshal: consumeDoubleValue,
	merge:     mergeScalarValue,
}

// sizeDoubleSliceValue returns the size of wire encoding a []float64 value as a repeated Double.
func sizeDoubleSliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	size = list.Len() * (tagsize + protowire.SizeFixed64())
	return size
}

// appendDoubleSliceValue encodes a []float64 value as a repeated Double.
func appendDoubleSliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	for i, llen := 0, list.Len(); i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendVarint(b, wiretag)
		b = protowire.AppendFixed64(b, math.Float64bits(v.Float()))
	}
	return b, nil
}

// consumeDoubleSliceValue wire decodes a []float64 value as a repeated Double.
func consumeDoubleSliceValue(b []byte, listv protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	list := listv.List()
	if wtyp == protowire.BytesType {
		b, n := protowire.ConsumeBytes(b)
		if n < 0 {
			return protoreflect.Value{}, out, errDecode
		}
		for len(b) > 0 {
			v, n := protowire.ConsumeFixed64(b)
			if n < 0 {
				return protoreflect.Value{}, out, errDecode
			}
			list.Append(protoreflect.ValueOfFloat64(math.Float64frombits(v)))
			b = b[n:]
		}
		out.n = n
		return listv, out, nil
	}
	if wtyp != protowire.Fixed64Type {
		return protoreflect.Value{}, out, errUnknown
	}
	v, n := protowire.ConsumeFixed64(b)
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	list.Append(protoreflect.ValueOfFloat64(math.Float64frombits(v)))
	out.n = n
	return listv, out, nil
}

var coderDoubleSliceValue = valueCoderFuncs{
	size:      sizeDoubleSliceValue,
	marshal:   appendDoubleSliceValue,
	unmarshal: consumeDoubleSliceValue,
	merge:     mergeListValue,
}

// sizeDoublePackedSliceValue returns the size of wire encoding a []float64 value as a packed repeated Double.
func sizeDoublePackedSliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	llen := list.Len()
	if llen == 0 {
		return 0
	}
	n := llen * protowire.SizeFixed64()
	return tagsize + protowire.SizeBytes(n)
}

// appendDoublePackedSliceValue encodes a []float64 value as a packed repeated Double.
func appendDoublePackedSliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	llen := list.Len()
	if llen == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, wiretag)
	n := llen * protowire.SizeFixed64()
	b = protowire.AppendVarint(b, uint64(n))
	for i := 0; i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendFixed64(b, math.Float64bits(v.Float()))
	}
	return b, nil
}

var coderDoublePackedSliceValue = valueCoderFuncs{
	size:      sizeDoublePackedSliceValue,
	marshal:   appendDoublePackedSliceValue,
	unmarshal: consumeDoubleSliceValue,
	merge:     mergeListValue,
}

// sizeString returns the size of wire encoding a string pointer as a String.
func sizeString(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := *p.String()
	return f.tagsize + protowire.SizeBytes(len(v))
}

// appendString wire encodes a string pointer as a String.
func appendString(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.String()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendString(b, v)
	return b, nil
}

// consumeString wire decodes a string pointer as a String.
func consumeString(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.BytesType {
		return out, errUnknown
	}
	v, n := protowire.ConsumeBytes(b)
	if n < 0 {
		return out, errDecode
	}
	*p.String() = string(v)
	out.n = n
	return out, nil
}

var coderString = pointerCoderFuncs{
	size:      sizeString,
	marshal:   appendString,
	unmarshal: consumeString,
	merge:     mergeString,
}

// appendStringValidateUTF8 wire encodes a string pointer as a String.
func appendStringValidateUTF8(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.String()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendString(b, v)
	if !utf8.ValidString(v) {
		return b, errInvalidUTF8{}
	}
	return b, nil
}

// consumeStringValidateUTF8 wire decodes a string pointer as a String.
func consumeStringValidateUTF8(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.BytesType {
		return out, errUnknown
	}
	v, n := protowire.ConsumeBytes(b)
	if n < 0 {
		return out, errDecode
	}
	if !utf8.Valid(v) {
		return out, errInvalidUTF8{}
	}
	*p.String() = string(v)
	out.n = n
	return out, nil
}

var coderStringValidateUTF8 = pointerCoderFuncs{
	size:      sizeString,
	marshal:   appendStringValidateUTF8,
	unmarshal: consumeStringValidateUTF8,
	merge:     mergeString,
}

// sizeStringNoZero returns the size of wire encoding a string pointer as a String.
// The zero value is not encoded.
func sizeStringNoZero(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := *p.String()
	if len(v) == 0 {
		return 0
	}
	return f.tagsize + protowire.SizeBytes(len(v))
}

// appendStringNoZero wire encodes a string pointer as a String.
// The zero value is not encoded.
func appendStringNoZero(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.String()
	if len(v) == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendString(b, v)
	return b, nil
}

var coderStringNoZero = pointerCoderFuncs{
	size:      sizeStringNoZero,
	marshal:   appendStringNoZero,
	unmarshal: consumeString,
	merge:     mergeStringNoZero,
}

// appendStringNoZeroValidateUTF8 wire encodes a string pointer as a String.
// The zero value is not encoded.
func appendStringNoZeroValidateUTF8(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.String()
	if len(v) == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendString(b, v)
	if !utf8.ValidString(v) {
		return b, errInvalidUTF8{}
	}
	return b, nil
}

var coderStringNoZeroValidateUTF8 = pointerCoderFuncs{
	size:      sizeStringNoZero,
	marshal:   appendStringNoZeroValidateUTF8,
	unmarshal: consumeStringValidateUTF8,
	merge:     mergeStringNoZero,
}

// sizeStringPtr returns the size of wire encoding a *string pointer as a String.
// It panics if the pointer is nil.
func sizeStringPtr(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := **p.StringPtr()
	return f.tagsize + protowire.SizeBytes(len(v))
}

// appendStringPtr wire encodes a *string pointer as a String.
// It panics if the pointer is nil.
func appendStringPtr(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := **p.StringPtr()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendString(b, v)
	return b, nil
}

// consumeStringPtr wire decodes a *string pointer as a String.
func consumeStringPtr(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.BytesType {
		return out, errUnknown
	}
	v, n := protowire.ConsumeBytes(b)
	if n < 0 {
		return out, errDecode
	}
	vp := p.StringPtr()
	if *vp == nil {
		*vp = new(string)
	}
	**vp = string(v)
	out.n = n
	return out, nil
}

var coderStringPtr = pointerCoderFuncs{
	size:      sizeStringPtr,
	marshal:   appendStringPtr,
	unmarshal: consumeStringPtr,
	merge:     mergeStringPtr,
}

// appendStringPtrValidateUTF8 wire encodes a *string pointer as a String.
// It panics if the pointer is nil.
func appendStringPtrValidateUTF8(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := **p.StringPtr()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendString(b, v)
	if !utf8.ValidString(v) {
		return b, errInvalidUTF8{}
	}
	return b, nil
}

// consumeStringPtrValidateUTF8 wire decodes a *string pointer as a String.
func consumeStringPtrValidateUTF8(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.BytesType {
		return out, errUnknown
	}
	v, n := protowire.ConsumeBytes(b)
	if n < 0 {
		return out, errDecode
	}
	if !utf8.Valid(v) {
		return out, errInvalidUTF8{}
	}
	vp := p.StringPtr()
	if *vp == nil {
		*vp = new(string)
	}
	**vp = string(v)
	out.n = n
	return out, nil
}

var coderStringPtrValidateUTF8 = pointerCoderFuncs{
	size:      sizeStringPtr,
	marshal:   appendStringPtrValidateUTF8,
	unmarshal: consumeStringPtrValidateUTF8,
	merge:     mergeStringPtr,
}

// sizeStringSlice returns the size of wire encoding a []string pointer as a repeated String.
func sizeStringSlice(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	s := *p.StringSlice()
	for _, v := range s {
		size += f.tagsize + protowire.SizeBytes(len(v))
	}
	return size
}

// appendStringSlice encodes a []string pointer as a repeated String.
func appendStringSlice(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.StringSlice()
	for _, v := range s {
		b = protowire.AppendVarint(b, f.wiretag)
		b = protowire.AppendString(b, v)
	}
	return b, nil
}

// consumeStringSlice wire decodes a []string pointer as a repeated String.
func consumeStringSlice(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	sp := p.StringSlice()
	if wtyp != protowire.BytesType {
		return out, errUnknown
	}
	v, n := protowire.ConsumeBytes(b)
	if n < 0 {
		return out, errDecode
	}
	*sp = append(*sp, string(v))
	out.n = n
	return out, nil
}

var coderStringSlice = pointerCoderFuncs{
	size:      sizeStringSlice,
	marshal:   appendStringSlice,
	unmarshal: consumeStringSlice,
	merge:     mergeStringSlice,
}

// appendStringSliceValidateUTF8 encodes a []string pointer as a repeated String.
func appendStringSliceValidateUTF8(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.StringSlice()
	for _, v := range s {
		b = protowire.AppendVarint(b, f.wiretag)
		b = protowire.AppendString(b, v)
		if !utf8.ValidString(v) {
			return b, errInvalidUTF8{}
		}
	}
	return b, nil
}

// consumeStringSliceValidateUTF8 wire decodes a []string pointer as a repeated String.
func consumeStringSliceValidateUTF8(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.BytesType {
		return out, errUnknown
	}
	v, n := protowire.ConsumeBytes(b)
	if n < 0 {
		return out, errDecode
	}
	if !utf8.Valid(v) {
		return out, errInvalidUTF8{}
	}
	sp := p.StringSlice()
	*sp = append(*sp, string(v))
	out.n = n
	return out, nil
}

var coderStringSliceValidateUTF8 = pointerCoderFuncs{
	size:      sizeStringSlice,
	marshal:   appendStringSliceValidateUTF8,
	unmarshal: consumeStringSliceValidateUTF8,
	merge:     mergeStringSlice,
}

// sizeStringValue returns the size of wire encoding a string value as a String.
func sizeStringValue(v protoreflect.Value, tagsize int, opts marshalOptions) int {
	return tagsize + protowire.SizeBytes(len(v.String()))
}

// appendStringValue encodes a string value as a String.
func appendStringValue(b []byte, v protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	b = protowire.AppendVarint(b, wiretag)
	b = protowire.AppendString(b, v.String())
	return b, nil
}

// consumeStringValue decodes a string value as a String.
func consumeStringValue(b []byte, _ protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	if wtyp != protowire.BytesType {
		return protoreflect.Value{}, out, errUnknown
	}
	v, n := protowire.ConsumeBytes(b)
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	out.n = n
	return protoreflect.ValueOfString(string(v)), out, nil
}

var coderStringValue = valueCoderFuncs{
	size:      sizeStringValue,
	marshal:   appendStringValue,
	unmarshal: consumeStringValue,
	merge:     mergeScalarValue,
}

// appendStringValueValidateUTF8 encodes a string value as a String.
func appendStringValueValidateUTF8(b []byte, v protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	b = protowire.AppendVarint(b, wiretag)
	b = protowire.AppendString(b, v.String())
	if !utf8.ValidString(v.String()) {
		return b, errInvalidUTF8{}
	}
	return b, nil
}

// consumeStringValueValidateUTF8 decodes a string value as a String.
func consumeStringValueValidateUTF8(b []byte, _ protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	if wtyp != protowire.BytesType {
		return protoreflect.Value{}, out, errUnknown
	}
	v, n := protowire.ConsumeBytes(b)
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	if !utf8.Valid(v) {
		return protoreflect.Value{}, out, errInvalidUTF8{}
	}
	out.n = n
	return protoreflect.ValueOfString(string(v)), out, nil
}

var coderStringValueValidateUTF8 = valueCoderFuncs{
	size:      sizeStringValue,
	marshal:   appendStringValueValidateUTF8,
	unmarshal: consumeStringValueValidateUTF8,
	merge:     mergeScalarValue,
}

// sizeStringSliceValue returns the size of wire encoding a []string value as a repeated String.
func sizeStringSliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	for i, llen := 0, list.Len(); i < llen; i++ {
		v := list.Get(i)
		size += tagsize + protowire.SizeBytes(len(v.String()))
	}
	return size
}

// appendStringSliceValue encodes a []string value as a repeated String.
func appendStringSliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	for i, llen := 0, list.Len(); i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendVarint(b, wiretag)
		b = protowire.AppendString(b, v.String())
	}
	return b, nil
}

// consumeStringSliceValue wire decodes a []string value as a repeated String.
func consumeStringSliceValue(b []byte, listv protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	list := listv.List()
	if wtyp != protowire.BytesType {
		return protoreflect.Value{}, out, errUnknown
	}
	v, n := protowire.ConsumeBytes(b)
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	list.Append(protoreflect.ValueOfString(string(v)))
	out.n = n
	return listv, out, nil
}

var coderStringSliceValue = valueCoderFuncs{
	size:      sizeStringSliceValue,
	marshal:   appendStringSliceValue,
	unmarshal: consumeStringSliceValue,
	merge:     mergeListValue,
}

// sizeBytes returns the size of wire encoding a []byte pointer as a Bytes.
func sizeBytes(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := *p.Bytes()
	return f.tagsize + protowire.SizeBytes(len(v))
}

// appendBytes wire encodes a []byte pointer as a Bytes.
func appendBytes(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Bytes()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendBytes(b, v)
	return b, nil
}

// consumeBytes wire decodes a []byte pointer as a Bytes.
func consumeBytes(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.BytesType {
		return out, errUnknown
	}
	v, n := protowire.ConsumeBytes(b)
	if n < 0 {
		return out, errDecode
	}
	*p.Bytes() = append(emptyBuf[:], v...)
	out.n = n
	return out, nil
}

var coderBytes = pointerCoderFuncs{
	size:      sizeBytes,
	marshal:   appendBytes,
	unmarshal: consumeBytes,
	merge:     mergeBytes,
}

// appendBytesValidateUTF8 wire encodes a []byte pointer as a Bytes.
func appendBytesValidateUTF8(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Bytes()
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendBytes(b, v)
	if !utf8.Valid(v) {
		return b, errInvalidUTF8{}
	}
	return b, nil
}

// consumeBytesValidateUTF8 wire decodes a []byte pointer as a Bytes.
func consumeBytesValidateUTF8(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.BytesType {
		return out, errUnknown
	}
	v, n := protowire.ConsumeBytes(b)
	if n < 0 {
		return out, errDecode
	}
	if !utf8.Valid(v) {
		return out, errInvalidUTF8{}
	}
	*p.Bytes() = append(emptyBuf[:], v...)
	out.n = n
	return out, nil
}

var coderBytesValidateUTF8 = pointerCoderFuncs{
	size:      sizeBytes,
	marshal:   appendBytesValidateUTF8,
	unmarshal: consumeBytesValidateUTF8,
	merge:     mergeBytes,
}

// sizeBytesNoZero returns the size of wire encoding a []byte pointer as a Bytes.
// The zero value is not encoded.
func sizeBytesNoZero(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	v := *p.Bytes()
	if len(v) == 0 {
		return 0
	}
	return f.tagsize + protowire.SizeBytes(len(v))
}

// appendBytesNoZero wire encodes a []byte pointer as a Bytes.
// The zero value is not encoded.
func appendBytesNoZero(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Bytes()
	if len(v) == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendBytes(b, v)
	return b, nil
}

// consumeBytesNoZero wire decodes a []byte pointer as a Bytes.
// The zero value is not decoded.
func consumeBytesNoZero(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.BytesType {
		return out, errUnknown
	}
	v, n := protowire.ConsumeBytes(b)
	if n < 0 {
		return out, errDecode
	}
	*p.Bytes() = append(([]byte)(nil), v...)
	out.n = n
	return out, nil
}

var coderBytesNoZero = pointerCoderFuncs{
	size:      sizeBytesNoZero,
	marshal:   appendBytesNoZero,
	unmarshal: consumeBytesNoZero,
	merge:     mergeBytesNoZero,
}

// appendBytesNoZeroValidateUTF8 wire encodes a []byte pointer as a Bytes.
// The zero value is not encoded.
func appendBytesNoZeroValidateUTF8(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	v := *p.Bytes()
	if len(v) == 0 {
		return b, nil
	}
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendBytes(b, v)
	if !utf8.Valid(v) {
		return b, errInvalidUTF8{}
	}
	return b, nil
}

// consumeBytesNoZeroValidateUTF8 wire decodes a []byte pointer as a Bytes.
func consumeBytesNoZeroValidateUTF8(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.BytesType {
		return out, errUnknown
	}
	v, n := protowire.ConsumeBytes(b)
	if n < 0 {
		return out, errDecode
	}
	if !utf8.Valid(v) {
		return out, errInvalidUTF8{}
	}
	*p.Bytes() = append(([]byte)(nil), v...)
	out.n = n
	return out, nil
}

var coderBytesNoZeroValidateUTF8 = pointerCoderFuncs{
	size:      sizeBytesNoZero,
	marshal:   appendBytesNoZeroValidateUTF8,
	unmarshal: consumeBytesNoZeroValidateUTF8,
	merge:     mergeBytesNoZero,
}

// sizeBytesSlice returns the size of wire encoding a [][]byte pointer as a repeated Bytes.
func sizeBytesSlice(p pointer, f *coderFieldInfo, opts marshalOptions) (size int) {
	s := *p.BytesSlice()
	for _, v := range s {
		size += f.tagsize + protowire.SizeBytes(len(v))
	}
	return size
}

// appendBytesSlice encodes a [][]byte pointer as a repeated Bytes.
func appendBytesSlice(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.BytesSlice()
	for _, v := range s {
		b = protowire.AppendVarint(b, f.wiretag)
		b = protowire.AppendBytes(b, v)
	}
	return b, nil
}

// consumeBytesSlice wire decodes a [][]byte pointer as a repeated Bytes.
func consumeBytesSlice(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	sp := p.BytesSlice()
	if wtyp != protowire.BytesType {
		return out, errUnknown
	}
	v, n := protowire.ConsumeBytes(b)
	if n < 0 {
		return out, errDecode
	}
	*sp = append(*sp, append(emptyBuf[:], v...))
	out.n = n
	return out, nil
}

var coderBytesSlice = pointerCoderFuncs{
	size:      sizeBytesSlice,
	marshal:   appendBytesSlice,
	unmarshal: consumeBytesSlice,
	merge:     mergeBytesSlice,
}

// appendBytesSliceValidateUTF8 encodes a [][]byte pointer as a repeated Bytes.
func appendBytesSliceValidateUTF8(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := *p.BytesSlice()
	for _, v := range s {
		b = protowire.AppendVarint(b, f.wiretag)
		b = protowire.AppendBytes(b, v)
		if !utf8.Valid(v) {
			return b, errInvalidUTF8{}
		}
	}
	return b, nil
}

// consumeBytesSliceValidateUTF8 wire decodes a [][]byte pointer as a repeated Bytes.
func consumeBytesSliceValidateUTF8(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.BytesType {
		return out, errUnknown
	}
	v, n := protowire.ConsumeBytes(b)
	if n < 0 {
		return out, errDecode
	}
	if !utf8.Valid(v) {
		return out, errInvalidUTF8{}
	}
	sp := p.BytesSlice()
	*sp = append(*sp, append(emptyBuf[:], v...))
	out.n = n
	return out, nil
}

var coderBytesSliceValidateUTF8 = pointerCoderFuncs{
	size:      sizeBytesSlice,
	marshal:   appendBytesSliceValidateUTF8,
	unmarshal: consumeBytesSliceValidateUTF8,
	merge:     mergeBytesSlice,
}

// sizeBytesValue returns the size of wire encoding a []byte value as a Bytes.
func sizeBytesValue(v protoreflect.Value, tagsize int, opts marshalOptions) int {
	return tagsize + protowire.SizeBytes(len(v.Bytes()))
}

// appendBytesValue encodes a []byte value as a Bytes.
func appendBytesValue(b []byte, v protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	b = protowire.AppendVarint(b, wiretag)
	b = protowire.AppendBytes(b, v.Bytes())
	return b, nil
}

// consumeBytesValue decodes a []byte value as a Bytes.
func consumeBytesValue(b []byte, _ protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	if wtyp != protowire.BytesType {
		return protoreflect.Value{}, out, errUnknown
	}
	v, n := protowire.ConsumeBytes(b)
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	out.n = n
	return protoreflect.ValueOfBytes(append(emptyBuf[:], v...)), out, nil
}

var coderBytesValue = valueCoderFuncs{
	size:      sizeBytesValue,
	marshal:   appendBytesValue,
	unmarshal: consumeBytesValue,
	merge:     mergeBytesValue,
}

// sizeBytesSliceValue returns the size of wire encoding a [][]byte value as a repeated Bytes.
func sizeBytesSliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) (size int) {
	list := listv.List()
	for i, llen := 0, list.Len(); i < llen; i++ {
		v := list.Get(i)
		size += tagsize + protowire.SizeBytes(len(v.Bytes()))
	}
	return size
}

// appendBytesSliceValue encodes a [][]byte value as a repeated Bytes.
func appendBytesSliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	for i, llen := 0, list.Len(); i < llen; i++ {
		v := list.Get(i)
		b = protowire.AppendVarint(b, wiretag)
		b = protowire.AppendBytes(b, v.Bytes())
	}
	return b, nil
}

// consumeBytesSliceValue wire decodes a [][]byte value as a repeated Bytes.
func consumeBytesSliceValue(b []byte, listv protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	list := listv.List()
	if wtyp != protowire.BytesType {
		return protoreflect.Value{}, out, errUnknown
	}
	v, n := protowire.ConsumeBytes(b)
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	list.Append(protoreflect.ValueOfBytes(append(emptyBuf[:], v...)))
	out.n = n
	return listv, out, nil
}

var coderBytesSliceValue = valueCoderFuncs{
	size:      sizeBytesSliceValue,
	marshal:   appendBytesSliceValue,
	unmarshal: consumeBytesSliceValue,
	merge:     mergeBytesListValue,
}

// We append to an empty array rather than a nil []byte to get non-nil zero-length byte slices.
var emptyBuf [0]byte

var wireTypes = map[protoreflect.Kind]protowire.Type{
	protoreflect.BoolKind:     protowire.VarintType,
	protoreflect.EnumKind:     protowire.VarintType,
	protoreflect.Int32Kind:    protowire.VarintType,
	protoreflect.Sint32Kind:   protowire.VarintType,
	protoreflect.Uint32Kind:   protowire.VarintType,
	protoreflect.Int64Kind:    protowire.VarintType,
	protoreflect.Sint64Kind:   protowire.VarintType,
	protoreflect.Uint64Kind:   protowire.VarintType,
	protoreflect.Sfixed32Kind: protowire.Fixed32Type,
	protoreflect.Fixed32Kind:  protowire.Fixed32Type,
	protoreflect.FloatKind:    protowire.Fixed32Type,
	protoreflect.Sfixed64Kind: protowire.Fixed64Type,
	protoreflect.Fixed64Kind:  protowire.Fixed64Type,
	protoreflect.DoubleKind:   protowire.Fixed64Type,
	protoreflect.StringKind:   protowire.BytesType,
	protoreflect.BytesKind:    protowire.BytesType,
	protoreflect.MessageKind:  protowire.BytesType,
	protoreflect.GroupKind:    protowire.StartGroupType,
}
