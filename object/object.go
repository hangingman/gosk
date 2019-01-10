package object

import (
	"encoding/hex"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Binary struct {
	Value []byte
}

func (b *Binary) Inspect() string {
	return hex.EncodeToString(b.Value)
}
