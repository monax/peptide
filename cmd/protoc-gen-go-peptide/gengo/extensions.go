package gengo

import (
	"strings"

	"github.com/monax/peptide/types/gogoproto"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func getFieldExtensionString(field *protogen.Field, extTypes ...protoreflect.ExtensionType) string {
	ext := getFieldExtension(field, extTypes...)
	if ext == nil {
		return ""
	}
	return ext.(string)
}

func getFieldExtension(field *protogen.Field, extTypes ...protoreflect.ExtensionType) interface{} {
	opts := field.Desc.Options().(*descriptorpb.FieldOptions)
	for _, extType := range extTypes {
		if proto.HasExtension(opts, extType) {
			return proto.GetExtension(opts, extType)
		}
	}
	return nil
}

func getCustomType(file *protogen.File, qualifiedName string) protogen.GoIdent {
	dot := strings.LastIndex(qualifiedName, ".")
	if dot == -1 {
		return protogen.GoIdent{
			GoImportPath: file.GoImportPath,
			GoName:       qualifiedName,
		}
	}
	return protogen.GoIdent{
		GoImportPath: protogen.GoImportPath(qualifiedName[:dot]),
		GoName:       qualifiedName[dot+1:],
	}
}

func getEnumValueName(g *protogen.GeneratedFile, value *protogen.EnumValue) string {
	enumValueOpts := value.Desc.Options().(*descriptorpb.EnumValueOptions)
	valueName := proto.GetExtension(enumValueOpts, gogoproto.E_EnumvalueCustomname).(string)
	if valueName != "" {
		return valueName
	}
	// duplicating behavior of g.P, but directly taking the GoName should
	// also work because of the syntactic context (const declaration)
	return g.QualifiedGoIdent(value.GoIdent)
}

func getMessageExtension(message *protogen.Message, extTypes ...protoreflect.ExtensionType) interface{} {
	opts := message.Desc.Options().(*descriptorpb.MessageOptions)
	for _, extType := range extTypes {
		if proto.HasExtension(opts, extType) {
			return proto.GetExtension(opts, extType)
		}
	}
	return nil
}

func getMessageExtensionBool(message *protogen.Message, extType protoreflect.ExtensionType, defaultValue bool) bool {
	flag := getMessageExtension(message, extType)
	if flag == nil {
		return defaultValue
	}
	return flag.(bool)
}

type customField struct {
	*protogen.Field
	originalIdent protogen.GoIdent
	pointer       bool
}

func (f *customField) goType(g *protogen.GeneratedFile) string {
	name := g.QualifiedGoIdent(f.GoIdent)
	if f.pointer {
		return "*" + name
	}
	return name
}

func getCustomFields(file *protogen.File, message *protogen.Message) []*customField {
	var fields []*customField
	for _, f := range message.Fields {
		generate := false
		field := &customField{
			Field: &protogen.Field{
				Desc:     f.Desc,
				GoName:   f.GoName,
				GoIdent:  f.GoIdent,
				Parent:   f.Parent,
				Oneof:    f.Oneof,
				Extendee: f.Extendee,
				Enum:     f.Enum,
				Message:  f.Message,
				Location: f.Location,
				Comments: f.Comments,
			},
			originalIdent: f.GoIdent,
			pointer:       true,
		}
		customTypeName := getFieldExtensionString(f, gogoproto.E_Customtype)
		if customTypeName != "" {
			field.GoIdent = getCustomType(file, customTypeName)
			generate = true
		}
		nullable := getFieldExtension(f, gogoproto.E_Nullable)
		if nullable != nil {
			field.pointer = nullable.(bool)
			generate = true
		}
		if generate {
			fields = append(fields, field)
		}
	}
	return fields
}
