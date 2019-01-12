package object

import (
	"encoding/hex"
	"strings"
)

type ObjectType string

const (
	BINARY_OBJ       = "BINARY"
	BINARY_ARRAY_OBJ = "BINARY_ARRAY"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

// Binary
type Binary struct {
	Value []byte
}

func (b *Binary) Type() ObjectType {
	return BINARY_OBJ
}

func (b *Binary) Inspect() string {
	return hex.EncodeToString(b.Value)
}

// BinaryArray
type BinaryArray struct {
	Value []Binary
}

func (b *BinaryArray) Type() ObjectType {
	return BINARY_ARRAY_OBJ
}

func (b *BinaryArray) Inspect() string {
	var barray []string
	for _, binary := range b.Value {
		barray = append(barray, binary.Inspect())
	}
	return strings.Join(barray, ",")
}
