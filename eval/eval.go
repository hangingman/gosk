package eval

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/hangingman/gosk/ast"
	"github.com/hangingman/gosk/object"
	"github.com/hangingman/gosk/token"
	"reflect"
	"strconv"
	"strings"
)

func isNil(x interface{}) bool {
	return x == nil || reflect.ValueOf(x).IsNil()
}

func isNotNil(x interface{}) bool {
	return !isNil(x)
}

func int2Byte(i int) []byte {
	return []byte{uint8(i)}
}

func int2Word(i int) []byte {
	bs := make([]byte, 2)
	binary.LittleEndian.PutUint16(bs, uint16(i))
	return bs
}

func int2Dword(i int) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(i))
	return bs
}

func evalDStatements(stmt *ast.MnemonicStatement, f func(int) []byte) object.Object {
	toks := []string{}
	bytes := []byte{}
	for _, tok := range stmt.Name.Tokens {
		if tok.Type == token.HEX_LIT {
			// 0xを取り除いて処理
			bs, _ := hex.DecodeString(string([]rune(tok.Literal)[2:]))
			bytes = append(bytes, bs...)
		} else if tok.Type == token.STR_LIT {
			// "を取り除いて処理
			strLength := len(tok.Literal)
			bs := []byte(tok.Literal[1 : strLength-1])
			bytes = append(bytes, bs...)
		} else if tok.Type == token.INT {
			// Go言語のintは常にint64 -> uint8
			int64Val, _ := strconv.Atoi(tok.Literal)
			bs := f(int64Val)
			bytes = append(bytes, bs...)
		}
		toks = append(toks, fmt.Sprintf("%s: %s", tok.Type, tok.Literal))
	}
	fmt.Printf("[%s]\n", strings.Join(toks, ", "))
	return &object.Binary{Value: bytes}
}

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
	return evalDStatements(stmt, int2Byte)
}

func evalDWStatement(stmt *ast.MnemonicStatement) object.Object {
	return evalDStatements(stmt, int2Word)
}

func evalDDStatement(stmt *ast.MnemonicStatement) object.Object {
	return evalDStatements(stmt, int2Dword)
}
