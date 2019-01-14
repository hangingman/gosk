package eval

import (
	"fmt"
	"github.com/hangingman/gosk/ast"
	"github.com/hangingman/gosk/object"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestEvalStatements(t *testing.T) {
	d1 := &ast.DummyStatement{}
	d2 := &ast.DummyStatement{}
	d3 := &ast.DummyStatement{}
	dArr := []ast.Statement{d1, d2, d3}

	evaluated := evalStatements(dArr)
	assert.Equal(t, "*object.ObjectArray", reflect.TypeOf(evaluated).String())
	objArr, ok := evaluated.(*object.ObjectArray)

	assert.True(t, ok)
	assert.Equal(t, 3, len(*objArr))
	for _, obj := range *objArr {
		assert.Equal(t, "*object.Binary", reflect.TypeOf(obj).String())
		fmt.Printf("%s: %s\n", reflect.TypeOf(obj), obj.Inspect())
		bin, _ := obj.(*object.Binary)
		fmt.Printf("%s\n", bin.Inspect())
	}
}
