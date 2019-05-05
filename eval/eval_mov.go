package eval

import (
	"fmt"
	"github.com/hangingman/gosk/ast"
	"github.com/hangingman/gosk/object"
	"github.com/hangingman/gosk/token"
	"log"
	"strings"
)

// 複雑なオペコードはファイルを分割する方針で
func evalMOVStatement(stmt *ast.MnemonicStatement) object.Object {
	toks := stmt.Name.Tokens
	bin := &object.Binary{Value: []byte{}}

	switch {
	case IsR8(toks[1]) && IsImm8(toks[2]):
		// MOV r8 , imm8
		log.Println(fmt.Sprintf("info: MOV r8 (%s), imm8 (%s)", toks[1], toks[2]))
		bin.Value = append(bin.Value, plusRb(0xb0, toks[1].Literal))
		bin.Value = append(bin.Value, imm8ToByte(toks[2])...)
	case IsR16(toks[1]) && IsImm16(toks[2]):
		// MOV r16, imm16
		log.Println(fmt.Sprintf("info: MOV r16 (%s), imm16 (%s)", toks[1], toks[2]))
		// 0xB8+rw
		bin.Value = append(bin.Value, plusRw(0xb8, toks[1].Literal))
		bin.Value = append(bin.Value, imm16ToWord(toks[2])...)
	case IsR32(toks[1]) && IsImm32(toks[2]):
		// MOV r32, imm32
		log.Println(fmt.Sprintf("info: MOV r32 (%s), imm32 (%s)", toks[1], toks[2]))
		// 0xB8+rd
		bin.Value = append(bin.Value, plusRd(0xb8, toks[1].Literal))
		bin.Value = append(bin.Value, imm32ToDword(toks[2])...)
	case IsR8(toks[1]) && toks[2].Type == token.LBRACKET && toks[4].Type == token.RBRACKET:
		// MOV r8 , imm8 で immが参照（ex: [SI]）
		log.Println(fmt.Sprintf("info: MOV r8 (%s), disp8 (%s)", toks[1], toks[3]))
		disp := "[" + toks[3].Literal + "]"
		// 0x8a
		bin.Value = append(bin.Value, 0x8a)
		bin.Value = append(bin.Value, generateModRMSlashN(0x8a, RegReg, disp, "/0"))
	case IsR16(toks[1]) && toks[2].Type == token.LBRACKET && toks[4].Type == token.RBRACKET:
		// MOV r16 , imm16 で immが参照
		log.Println(fmt.Sprintf("info: MOV r16 (%s), disp16 (%s)", toks[1], toks[3]))
		disp := "[" + toks[3].Literal + "]"
		// 0x8b
		bin.Value = append(bin.Value, 0x8b)
		bin.Value = append(bin.Value, generateModRMSlashN(0x8b, RegReg, disp, "/0"))
	case IsR32(toks[1]) && toks[2].Type == token.LBRACKET && toks[4].Type == token.RBRACKET:
		// MOV r32 , imm32 で immが参照
		log.Println(fmt.Sprintf("info: MOV r32 (%s), disp32 (%s)", toks[1], toks[3]))
		disp := "[" + toks[3].Literal + "]"
		// 0x8b
		bin.Value = append(bin.Value, 0x8b)
		bin.Value = append(bin.Value, generateModRMSlashN(0x8b, RegReg, disp, "/0"))
	case IsR8(toks[1]) && toks[2].Type == token.IDENT:
		// MOV r8 , imm8 で immがラベル
		// callbackを配置し今のバイト数を設定する
		log.Println(fmt.Sprintf("info: MOV r8 (%s), imm8(label) (%s)", toks[1], toks[2]))
		// 0xB0+rb
		opcode := plusRb(0xb0, toks[1].Literal)
		bin.Value = append(bin.Value, opcode)
		bin.Value = append(bin.Value, 0x00)

		labelManage.AddLabelCallback(
			[]byte{opcode},
			toks[2].Literal,
			bin,
			-dollarPosition,
			int2Byte,
		)
	case IsR16(toks[1]) && toks[2].Type == token.IDENT:
		// MOV r16 , imm16 で immがラベルの場合
		// callbackを配置し今のバイト数を設定する
		log.Println(fmt.Sprintf("info: MOV r16 (%s), imm16(label) (%s)", toks[1], toks[2]))
		// 0xB8+rw
		opcode := plusRw(0xb8, toks[1].Literal)
		bin.Value = append(bin.Value, opcode)
		bin.Value = append(bin.Value, 0x00)
		bin.Value = append(bin.Value, 0x00)

		labelManage.AddLabelCallback(
			[]byte{opcode},
			toks[2].Literal,
			bin,
			-dollarPosition,
			int2Word,
		)
	case IsR32(toks[1]) && toks[2].Type == token.IDENT:
		// MOV r32 , imm32 で immがラベルの場合
		// callbackを配置し今のバイト数を設定する
		log.Println(fmt.Sprintf("info: MOV r32 (%s), imm32(label) (%s)", toks[1], toks[2]))
		// 0xB8+rd
		opcode := plusRd(0xb8, toks[1].Literal)
		bin.Value = append(bin.Value, opcode)
		bin.Value = append(bin.Value, 0x00)
		bin.Value = append(bin.Value, 0x00)
		bin.Value = append(bin.Value, 0x00)
		bin.Value = append(bin.Value, 0x00)

		labelManage.AddLabelCallback(
			[]byte{opcode},
			toks[2].Literal,
			bin,
			-dollarPosition,
			int2Dword,
		)
	case IsSreg(toks[1]) && IsR16(toks[2]):
		// MOV Sreg, r/m16
		log.Println(fmt.Sprintf("info: MOV Sreg (%s), r/m16 (%s)", toks[1], toks[2]))
		// 0x8E /r
		bin.Value = append(bin.Value, 0x8e)
		bin.Value = append(bin.Value, generateModRMSlashR(0x8e, Reg, toks[1].Literal))
	}

	tokStrArray := []string{}
	for _, tok := range toks {
		tokStrArray = append(tokStrArray, fmt.Sprintf("%s: %s", tok.Type, tok.Literal))
	}

	log.Println(fmt.Sprintf("info: [%s]", strings.Join(tokStrArray, ", ")))

	stmt.Bin = bin
	return stmt.Bin
}
