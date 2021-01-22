package peptide

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"reflect"
)

var converterType = reflect.TypeOf((*Converter)(nil)).Elem()

func AsConverter(t reflect.Type) Converter {
	if t.Implements(converterType) {
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		i := reflect.New(t).Interface()
		return i.(Converter)
	}
	return nil
}

// A Converter coverts to/from Go reflect.Value types and protobuf protoreflect.Value types.
type Converter interface {
	// PBValueOf converts a reflect.Value to a protoreflect.Value.
	PBValueOf(reflect.Value) protoreflect.Value

	// GoValueOf converts a protoreflect.Value to a reflect.Value.
	GoValueOf(protoreflect.Value) reflect.Value

	// IsValidPB returns whether a protoreflect.Value is compatible with this type.
	IsValidPB(protoreflect.Value) bool

	// IsValidGo returns whether a reflect.Value is compatible with this type.
	IsValidGo(reflect.Value) bool

	// New returns a new field value.
	// For scalars, it returns the default value of the field.
	// For composite types, it returns a new mutable value.
	New() protoreflect.Value

	// Zero returns a new field value.
	// For scalars, it returns the default value of the field.
	// For composite types, it returns an immutable, empty value.
	Zero() protoreflect.Value
}
