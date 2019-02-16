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

func getRMFromReg(srcReg string) string {
	log.Println(fmt.Sprintf("info: Get rm from %s", srcReg))

	// const size_t imm = get_imm_size_evenif_bracket(src_reg);
	// if (imm == imm16) {
	//  return "110";
	// } else if (imm == imm32) {
	//  return "110";
	// }

	// fmt.Printf("%03b\n", 3)
	var regBits int

	switch {
	case IsR8(token.Token{Type: token.REGISTER, Literal: srcReg}):
		regBits = r8CodeMap[srcReg]
	case IsR16(token.Token{Type: token.REGISTER, Literal: srcReg}):
		regBits = r16CodeMap[srcReg]
	case IsR32(token.Token{Type: token.REGISTER, Literal: srcReg}):
		regBits = r32CodeMap[srcReg]
	case IsSreg(token.Token{Type: token.SEG_REGISTER, Literal: srcReg}):
		log.Println(fmt.Sprintf("info: HiHi!!!"))
		regBits = sregCodeMap[srcReg]
	}

	// } else if (SEGMENT_REGISTERS_SSS_MAP.count(src_reg)) {
	//  return SEGMENT_REGISTERS_SSS_MAP.at(src_reg);
	// } else if (REGISTERS_MMM_MAP.count(src_reg)) {
	//  return REGISTERS_MMM_MAP.at(src_reg);
	// }

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
	modrm += getRMFromReg(dstReg) // [r/m]

	i, _ := strconv.ParseUint(modrm, 2, 0)
	log.Println(fmt.Sprintf("info: ModR/M => %s(%x)", modrm, i))
	return byte(i)
}
