package eval

import (
	"github.com/hangingman/gosk/ast"
	"github.com/hangingman/gosk/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.BinaryLiteral:
		return &object.Binary{Value: node.Value}
	}

	return nil
}
