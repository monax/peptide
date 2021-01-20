package internal_gengo

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

func customType(qualifiedName string) protogen.GoIdent {
	dot := strings.LastIndex(qualifiedName, ".")
	if dot == -1 {
		return protogen.GoIdent{GoName: qualifiedName}
	}
	return protogen.GoIdent{
		GoName:       qualifiedName[dot+1:],
		GoImportPath: protogen.GoImportPath(qualifiedName[:dot]),
	}
}

func enumValueName(g *protogen.GeneratedFile, value *protogen.EnumValue) string {
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
