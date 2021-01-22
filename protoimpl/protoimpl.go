package protoimpl

import (
	"github.com/monax/peptide/internal/filedesc"
	"github.com/monax/peptide/internal/filetype"
	"github.com/monax/peptide/internal/impl"
)

var X impl.Export

const UnsafeEnabled = impl.UnsafeEnabled

type (
	// Types used by generated code in init functions.
	DescBuilder = filedesc.Builder
	TypeBuilder = filetype.Builder

	// Types used by generated code to implement EnumType, MessageType, and ExtensionType.
	EnumInfo      = impl.EnumInfo
	MessageInfo   = impl.MessageInfo
	ExtensionInfo = impl.ExtensionInfo

	// Types embedded in generated messages.
	MessageState     = impl.MessageState
	SizeCache        = impl.SizeCache
	WeakFields       = impl.WeakFields
	UnknownFields    = impl.UnknownFields
	ExtensionFields  = impl.ExtensionFields
	ExtensionFieldV1 = impl.ExtensionField

	Pointer = impl.Pointer
)
