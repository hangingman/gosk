package object

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryInspect(t *testing.T) {
	bin := Binary{Value: []byte{1, 2, 3}}
	assert.Equal(t, bin.Value, []byte{1, 2, 3})
	assert.Equal(t, bin.Inspect(), "010203")
}

func TestBinaryArrayInspect(t *testing.T) {
	bin1 := Binary{Value: []byte{1, 2, 3}}
	bin2 := Binary{Value: []byte{4, 5, 6}}

	barray := BinaryArray{Value: []Binary{bin1, bin2}}
	assert.Equal(t, barray.Inspect(), "010203,040506")
}
