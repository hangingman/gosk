package eval

import (
	"fmt"
	"github.com/hangingman/gosk/token"
	"log"
	"strconv"
)

type Mod int

const (
	RegReg    Mod = iota // mod=00: [レジスター+レジスター]
	RegDisp8             // mod=01: [レジスター+disp8]
	RegDisp16            // mod=10: [レジスター+disp16/32]
	Reg                  // mod=11: レジスター
)

func (c Mod) String() string {
	switch c {
	case RegReg:
		return "RegReg"
	case RegDisp8:
		return "RegDisp8"
	case RegDisp16:
		return "RegDisp16"
	case Reg:
		return "Reg"
	default:
		return "Unknown"
	}
}

// ModR/Mの条件と2bitのデータの対応
var mod2byteMap = map[Mod]string{
	RegReg: "00", RegDisp8: "01", RegDisp16: "10", Reg: "11",
}

// /0,/1.../7の条件と3bitのデータの対応
var slash2bitMap = map[string]string{
	"/0": "000",
	"/1": "001",
	"/2": "010",
	"/3": "011",
	"/4": "100",
	"/5": "101",
	"/6": "110",
	"/7": "111",
}

func getRMFromReg(srcReg string) string {
	log.Println(fmt.Sprintf("info: Get rm reg=%s", srcReg))

	var regBits int

	switch {
	case IsR8(token.Token{Type: token.REGISTER, Literal: srcReg}):
		regBits = r8CodeMap[srcReg]
	case IsR16(token.Token{Type: token.REGISTER, Literal: srcReg}):
		regBits = r16CodeMap[srcReg]
	case IsR32(token.Token{Type: token.REGISTER, Literal: srcReg}):
		regBits = r32CodeMap[srcReg]
	case IsSreg(token.Token{Type: token.SEG_REGISTER, Literal: srcReg}):
		regBits = sregCodeMap[srcReg]
	default:
	}

	// [<SIB>], [<SIB>+disp8], [<SIB>+disp32]
	//return "100";

	ans := fmt.Sprintf("%03b", regBits)
	log.Println(fmt.Sprintf("info: Get rm = %s", ans))
	return ans
}

// generateModRM オペコードと２つのレジスタについてModR/Mを作成する
// 仕様書に '/r' の形式でModR/Mを求められる場合に使用する
func generateModRMSlashR(opcode byte, m Mod, dstReg string) byte {
	log.Println(fmt.Sprintf("info: ModR/M /r opcode=%x type=%s dst=%s", opcode, m, dstReg))
	//
	// Generate ModR/M byte with arguments
	// [mod] 2bit
	// [reg] 3bit
	// [r/m] 3bit
	//
	modrm := mod2byteMap[m]       // [mod]
	modrm += getRMFromReg(dstReg) // [reg]
	modrm += "000"                // [r/m]

	i, _ := strconv.ParseUint(modrm, 2, 0)
	log.Println(fmt.Sprintf("info: ModR/M => %s(%x)", modrm, i))
	return byte(i)
}

// generateModRM オペコードと２つのレジスタについてModR/Mを作成する
// 仕様書に '/r' の形式でModR/Mを求められる場合に使用する
// @param opcode   オペコード
// @param m        どの形式の命令か
// @param dstReg   宛先のレジスタ
// @param regField '/0', '/1' ... '/7' までの文字列
func generateModRMSlashN(opcode byte, m Mod, dstReg string, regField string) byte {
	log.Println(fmt.Sprintf("info: ModR/M %s opcode=%x type=%s dst=%s", regField, opcode, m, dstReg))
	//
	// Generate ModR/M byte with arguments
	// [mod] 2bit
	// [reg] 3bit
	// [r/m] 3bit
	//
	modrm := mod2byteMap[m]         // [mod]
	modrm += slash2bitMap[regField] // [reg]
	modrm += getRMFromReg(dstReg)   // [r/m]

	i, _ := strconv.ParseUint(modrm, 2, 0)
	log.Println(fmt.Sprintf("info: ModR/M => %s(%x)", modrm, i))
	return byte(i)
}
