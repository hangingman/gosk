package eval

import (
	"fmt"
	"github.com/hangingman/gosk/ast"
	"github.com/hangingman/gosk/object"
	"reflect"
	"strings"
)

func isNil(x interface{}) bool {
	return x == nil || reflect.ValueOf(x).IsNil()
}

func isNotNil(x interface{}) bool {
	return !isNil(x)
}

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.DummyStatement:
		return &object.Binary{Value: []byte{1, 2, 3}}
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
	results := object.ObjectArray{}

	// 文を評価して、結果としてobject.ObjectArrayを返す
	for _, stmt := range stmts {
		if isNotNil(stmt) {
			results = append(results, Eval(stmt))
		}
	}

	return &results
}

func evalMnemonicStatement(stmt *ast.MnemonicStatement) object.Object {
	switch stmt.Name.Tokens[0].Literal {
	case "DB":
		return evalDBStatement(stmt)
	case "DW":
		return evalDWStatement(stmt)
	case "DD":
		return evalDDStatement(stmt)
	}

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

func evalDBStatement(stmt *ast.MnemonicStatement) object.Object {
	toks := []string{}
	for _, tok := range stmt.Name.Tokens {
		toks = append(toks, fmt.Sprintf("%s: %s", tok.Type, tok.Literal))
	}
	fmt.Printf("[%s]\n", strings.Join(toks, ", "))
	return &object.Binary{Value: []byte{0}}
}

func evalDWStatement(stmt *ast.MnemonicStatement) object.Object {
	toks := []string{}
	for _, tok := range stmt.Name.Tokens {
		toks = append(toks, fmt.Sprintf("%s: %s", tok.Type, tok.Literal))
	}
	fmt.Printf("[%s]\n", strings.Join(toks, ", "))
	return &object.Binary{Value: []byte{0}}
}

func evalDDStatement(stmt *ast.MnemonicStatement) object.Object {
	toks := []string{}
	for _, tok := range stmt.Name.Tokens {
		toks = append(toks, fmt.Sprintf("%s: %s", tok.Type, tok.Literal))
	}
	fmt.Printf("[%s]\n", strings.Join(toks, ", "))
	return &object.Binary{Value: []byte{0}}
}
