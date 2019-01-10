package object

import (
	"encoding/hex"
)

type ObjectType string

const (
	BINARY_OBJ = "BINARY"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Binary struct {
	Value []byte
}

func (b *Binary) Type() ObjectType {
	return BINARY_OBJ
}

func (b *Binary) Inspect() string {
	return hex.EncodeToString(b.Value)
}
