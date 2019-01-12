package eval

import (
	"github.com/hangingman/gosk/ast"
	"github.com/hangingman/gosk/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.MnemonicStatement:
		return evalMnemonicStatement(node)
	case *ast.SettingStatement:
		return evalSettingStatement(node)
	case *ast.LabelStatement:
		return evalLabelStatement(node)
	case *ast.EquStatement:
		return evalEquStatement(node)
	case *ast.BinaryLiteral:
		return &object.Binary{Value: node.Value}
	}

	return nil
}

// evalStatements は文を評価する
func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	// 文を評価するが最終的にすべて文だった場合は何も返さないことになる
	for _, stmt := range stmts {
		result = Eval(stmt)
	}

	return result
}

func evalMnemonicStatement(stmt *ast.MnemonicStatement) object.Object {
	return nil
}

func evalSettingStatement(stmt *ast.SettingStatement) object.Object {
	return nil
}

func evalLabelStatement(stmt *ast.LabelStatement) object.Object {
	return nil
}

func evalEquStatement(stmt *ast.EquStatement) object.Object {
	return nil
}
