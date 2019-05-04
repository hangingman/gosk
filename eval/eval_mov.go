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
	bin := []byte{}

	switch {
	case IsR8(toks[1]) && IsImm8(toks[2]):
		// MOV r8 , imm8
		log.Println(fmt.Sprintf("info: MOV r8 (%s), imm8 (%s)", toks[1], toks[2]))
		bin = []byte{} // 0xB0+rb
		bin = append(bin, plusRb(0xb0, toks[1].Literal))
		bin = append(bin, imm8ToByte(toks[2])...)
	case IsR16(toks[1]) && IsImm16(toks[2]):
		// MOV r16, imm16
		log.Println(fmt.Sprintf("info: MOV r16 (%s), imm16 (%s)", toks[1], toks[2]))
		bin = []byte{} // 0xB8+rw
		bin = append(bin, plusRw(0xb8, toks[1].Literal))
		bin = append(bin, imm16ToWord(toks[2])...)
	case IsR32(toks[1]) && IsImm32(toks[2]):
		// MOV r32, imm32
		log.Println(fmt.Sprintf("info: MOV r32 (%s), imm32 (%s)", toks[1], toks[2]))
		bin = []byte{} // 0xB8+rd
		bin = append(bin, plusRd(0xb8, toks[1].Literal))
		bin = append(bin, imm32ToDword(toks[2])...)

	case IsR8(toks[1]) && toks[2].Type == token.LBRACKET && toks[4].Type == token.RBRACKET:
		// MOV r8 , imm8 で immが参照（ex: [SI]）
		log.Println(fmt.Sprintf("info: MOV r8 (%s), disp8 (%s)", toks[1], toks[3]))
		disp := "[" + toks[3].Literal + "]"
		bin = []byte{} // 0x8a
		bin = append(bin, 0x8a)
		bin = append(bin, generateModRMSlashN(0x8a, RegReg, disp, "/0"))

	case IsR16(toks[1]) && toks[2].Type == token.LBRACKET && toks[4].Type == token.RBRACKET:
		// MOV r16 , imm16 で immが参照
		log.Println(fmt.Sprintf("info: MOV r16 (%s), disp16 (%s)", toks[1], toks[3]))
		disp := "[" + toks[3].Literal + "]"
		bin = []byte{} // 0x8b
		bin = append(bin, 0x8b)
		bin = append(bin, generateModRMSlashN(0x8b, RegReg, disp, "/0"))

	case IsR32(toks[1]) && toks[2].Type == token.LBRACKET && toks[4].Type == token.RBRACKET:
		// MOV r32 , imm32 で immが参照
		log.Println(fmt.Sprintf("info: MOV r32 (%s), disp32 (%s)", toks[1], toks[3]))
		disp := "[" + toks[3].Literal + "]"
		bin = []byte{} // 0x8b
		bin = append(bin, 0x8b)
		bin = append(bin, generateModRMSlashN(0x8b, RegReg, disp, "/0"))

	case IsR8(toks[1]) && toks[2].Type == token.IDENT:
		// MOV r8 , imm8 で immがラベル
		// callbackを配置し今のバイト数を設定する
		log.Println(fmt.Sprintf("info: MOV r8 (%s), imm8(label) (%s)", toks[1], toks[2]))
		bin = []byte{} // 0xB0+rb
		bin = append(bin, plusRb(0xb0, toks[1].Literal))
		labelManage.AddLabelCallback(
			[]byte{0x00},
			toks[2].Literal,
			&object.Binary{Value: bin},
			-dollarPosition,
			int2Byte,
		)
	case IsR16(toks[1]) && toks[2].Type == token.IDENT:
		// MOV r16 , imm16 で immがラベルの場合
		// callbackを配置し今のバイト数を設定する
		log.Println(fmt.Sprintf("info: MOV r16 (%s), imm16(label) (%s)", toks[1], toks[2]))
		bin = []byte{} // 0xB8+rw
		bin = append(bin, plusRw(0xb8, toks[1].Literal))
		labelManage.AddLabelCallback(
			[]byte{0x00, 0x00},
			toks[2].Literal,
			&object.Binary{Value: bin},
			-dollarPosition,
			int2Word,
		)
	case IsR32(toks[1]) && toks[2].Type == token.IDENT:
		// MOV r32 , imm32 で immがラベルの場合
		// callbackを配置し今のバイト数を設定する
		log.Println(fmt.Sprintf("info: MOV r32 (%s), imm32(label) (%s)", toks[1], toks[2]))
		bin = []byte{} // 0xB8+rd
		bin = append(bin, plusRd(0xb8, toks[1].Literal))
		labelManage.AddLabelCallback(
			[]byte{0x00, 0x00, 0x00, 0x00},
			toks[2].Literal,
			&object.Binary{Value: bin},
			-dollarPosition,
			int2Dword,
		)

	case IsSreg(toks[1]) && IsR16(toks[2]):
		// MOV Sreg, r/m16
		log.Println(fmt.Sprintf("info: MOV Sreg (%s), r/m16 (%s)", toks[1], toks[2]))
		bin = []byte{} // 0x8E /r
		bin = append(bin, 0x8e)
		bin = append(bin, generateModRMSlashR(0x8e, Reg, toks[1].Literal))
	}

	tokStrArray := []string{}
	for _, tok := range toks {
		tokStrArray = append(tokStrArray, fmt.Sprintf("%s: %s", tok.Type, tok.Literal))
	}

	log.Println(fmt.Sprintf("info: [%s]", strings.Join(tokStrArray, ", ")))
	return &object.Binary{Value: bin}
}
