package eval

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

func generateModRM(opcode byte, m Mod, srcReg string, dstReg string) byte {
	return 0x00
}
