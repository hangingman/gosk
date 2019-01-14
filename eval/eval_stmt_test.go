package eval

import (
	"github.com/hangingman/gosk/ast"
	"github.com/hangingman/gosk/object"
	// "github.com/hangingman/gosk/eval"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestEvalStatements(t *testing.T) {
	// bin1 := Binary{Value: []byte{1, 2, 3}}
	// bin2 := Binary{Value: []byte{4, 5, 6}}
	// barray := BinaryArray{Value: []Binary{bin1, bin2}}
	// assert.Equal(t, barray.Inspect(), "010203,040506")

	d1 := &ast.DummyStatement{}
	d2 := &ast.DummyStatement{}
	d3 := &ast.DummyStatement{}
	dArr := []ast.Statement{d1, d2, d3}

	objArr, ok := evalStatements(dArr).(*object.ObjectArray)
	assert.True(t, ok)
	assert.Equal(t, len(*objArr), 3)
	for _, obj := range *objArr {
		fmt.Printf("%s: %s\n", reflect.TypeOf(obj), obj.Inspect())
		bin, _ := obj.(*object.Binary)
		fmt.Printf("%s\n", bin.Inspect())
	}
}
