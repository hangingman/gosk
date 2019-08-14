package token

// TokenType は stringのエイリアス
type TokenType string

// Token は文字列をトークン化した後の情報を保持する
// TokenType にはtoken内で定義したconst値
// Literal には実際に読み取った文字列
type Token struct {
	Type    TokenType
	Literal string
}

func (tok *Token) IsOperator() bool {
	return tok.Type == PLUS ||
		tok.Type == MINUS ||
		tok.Type == SLASH ||
		tok.Type == ASTERISK
}

const (
	ILLEGAL      = "ILLEGAL"
	EOF          = "EOF"
	IDENT        = "IDENT"
	INT          = "INT"
	ASSIGN       = "="
	PLUS         = "+"
	MINUS        = "-"
	BANG         = "!"
	ASTERISK     = "*"
	SLASH        = "/"
	COMMA        = ","
	COLON        = ":"
	SEMICOLON    = ";"
	SHARP        = "#"
	LT           = "<"
	GT           = ">"
	LPAREN       = "("
	RPAREN       = ")"
	LBRACE       = "{"
	RBRACE       = "}"
	LBRACKET     = "["
	RBRACKET     = "]"
	DOUBLE_QT    = "\""
	DOLLAR       = "$"
	EQU          = "EQU"
	GLOBAL       = "GLOBAL"
	OPCODE       = "OPCODE"
	SETTING      = "SETTING"
	STR_LIT      = "STR_LIT"
	HEX_LIT      = "HEX_LIT"
	LABEL        = "LABEL"
	REGISTER     = "REGISTER"
	SEG_REGISTER = "SEG_REGISTER"
	CTL_REGISTER = "CTL_REGISTER"
	DATA_TYPE    = "DATA_TYPE"
)

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

var keywords = map[string]TokenType{
	// naskの予約語
	"EQU":    EQU,
	"GLOBAL": GLOBAL,
	// naskがサポートするレジスタ
	"AL":  REGISTER,
	"BL":  REGISTER,
	"CL":  REGISTER,
	"DL":  REGISTER,
	"EAX": REGISTER,
	"EBX": REGISTER,
	"ECX": REGISTER,
	"EDX": REGISTER,
	"AX":  REGISTER,
	"BX":  REGISTER,
	"CX":  REGISTER,
	"DX":  REGISTER,
	"AH":  REGISTER,
	"BH":  REGISTER,
	"CH":  REGISTER,
	"DH":  REGISTER,
	"ESP": REGISTER,
	"EDI": REGISTER,
	"EBP": REGISTER,
	"ESI": REGISTER,
	"SP":  REGISTER,
	"DI":  REGISTER,
	"BP":  REGISTER,
	"SI":  REGISTER,
	"CS":  SEG_REGISTER, // コード
	"DS":  SEG_REGISTER, // データ
	"ES":  SEG_REGISTER, // エクストラ
	"SS":  SEG_REGISTER, // スタック
	"FS":  SEG_REGISTER,
	"GS":  SEG_REGISTER,
	"CR0": CTL_REGISTER,
	"CR1": CTL_REGISTER,
	"CR2": CTL_REGISTER,
	"CR3": CTL_REGISTER,
	"CR4": CTL_REGISTER,
	"BYTE": DATA_TYPE,
	"WORD": DATA_TYPE,
	"DWORD": DATA_TYPE,
	// naskで用意されている設定用命令
	"BITS":     SETTING,
	"INSTRSET": SETTING,
	"OPTIMIZE": SETTING,
	"FORMAT":   SETTING,
	"PADDING":  SETTING,
	"PADSET":   SETTING,
	"OPTION":   SETTING,
	"SECTION":  SETTING,
	"ABSOLUTE": SETTING,
	"FILE":     SETTING,
	// naskでサポートされているオペコード
	"AAA":     OPCODE,
	"AAD":     OPCODE,
	"AAS":     OPCODE,
	"AAM":     OPCODE,
	"ADC":     OPCODE,
	"ADD":     OPCODE,
	"AND":     OPCODE,
	"ALIGN":   OPCODE,
	"ALIGNB":  OPCODE,
	"ARPL":    OPCODE,
	"BOUND":   OPCODE,
	"BSF":     OPCODE,
	"BSR":     OPCODE,
	"BSWAP":   OPCODE,
	"BT":      OPCODE,
	"BTC":     OPCODE,
	"BTR":     OPCODE,
	"BTS":     OPCODE,
	"CALL":    OPCODE,
	"CBW":     OPCODE,
	"CDQ":     OPCODE,
	"CLC":     OPCODE,
	"CLD":     OPCODE,
	"CLI":     OPCODE,
	"CLTS":    OPCODE,
	"CMC":     OPCODE,
	"CMP":     OPCODE,
	"CMPSB":   OPCODE,
	"CMPSD":   OPCODE,
	"CMPSW":   OPCODE,
	"CMPXCHG": OPCODE,
	"CPUID":   OPCODE,
	"CWD":     OPCODE,
	"CWDE":    OPCODE,
	"DAA":     OPCODE,
	"DAS":     OPCODE,
	"DB":      OPCODE,
	"DD":      OPCODE,
	"DEC":     OPCODE,
	"DIV":     OPCODE,
	"DQ":      OPCODE,
	"DT":      OPCODE,
	"DW":      OPCODE,
	"END":     OPCODE,
	"ENTER":   OPCODE,
	"EXTERN":  OPCODE,
	"F2XM1":   OPCODE,
	"FABS":    OPCODE,
	"FADD":    OPCODE,
	"FADDP":   OPCODE,
	"FBLD":    OPCODE,
	"FBSTP":   OPCODE,
	"FCHS":    OPCODE,
	"FCLEX":   OPCODE,
	"FCOM":    OPCODE,
	"FCOMP":   OPCODE,
	"FCOMPP":  OPCODE,
	"FCOS":    OPCODE,
	"FDECSTP": OPCODE,
	"FDISI":   OPCODE,
	"FDIV":    OPCODE,
	"FDIVP":   OPCODE,
	"FDIVR":   OPCODE,
	"FDIVRP":  OPCODE,
	"FENI":    OPCODE,
	"FFREE":   OPCODE,
	"FIADD":   OPCODE,
	"FICOM":   OPCODE,
	"FICOMP":  OPCODE,
	"FIDIV":   OPCODE,
	"FIDIVR":  OPCODE,
	"FILD":    OPCODE,
	"FIMUL":   OPCODE,
	"FINCSTP": OPCODE,
	"FINIT":   OPCODE,
	"FIST":    OPCODE,
	"FISTP":   OPCODE,
	"FISUB":   OPCODE,
	"FISUBR":  OPCODE,
	"FLD":     OPCODE,
	"FLD1":    OPCODE,
	"FLDCW":   OPCODE,
	"FLDENV":  OPCODE,
	"FLDL2E":  OPCODE,
	"FLDL2T":  OPCODE,
	"FLDLG2":  OPCODE,
	"FLDLN2":  OPCODE,
	"FLDPI":   OPCODE,
	"FLDZ":    OPCODE,
	"FMUL":    OPCODE,
	"FMULP":   OPCODE,
	"FNCLEX":  OPCODE,
	"FNDISI":  OPCODE,
	"FNENI":   OPCODE,
	"FNINIT":  OPCODE,
	"FNOP":    OPCODE,
	"FNSAVE":  OPCODE,
	"FNSTCW":  OPCODE,
	"FNSTENV": OPCODE,
	"FNSTSW":  OPCODE,
	"FPATAN":  OPCODE,
	"FPTAN":   OPCODE,
	"FPREM":   OPCODE,
	"FPREM1":  OPCODE,
	"FRNDINT": OPCODE,
	"FRSTOR":  OPCODE,
	"FSAVE":   OPCODE,
	"FSCALE":  OPCODE,
	"FSETPM":  OPCODE,
	"FSIN":    OPCODE,
	"FSINCOS": OPCODE,
	"FSQRT":   OPCODE,
	"FST":     OPCODE,
	"FSTCW":   OPCODE,
	"FSTENV":  OPCODE,
	"FSTP":    OPCODE,
	"FSTSW":   OPCODE,
	"FSUB":    OPCODE,
	"FSUBP":   OPCODE,
	"FSUBR":   OPCODE,
	"FSUBRP":  OPCODE,
	"FTST":    OPCODE,
	"FUCOM":   OPCODE,
	"FUCOMP":  OPCODE,
	"FUCOMPP": OPCODE,
	"FXAM":    OPCODE,
	"FXCH":    OPCODE,
	"FXTRACT": OPCODE,
	"FYL2X":   OPCODE,
	"FYL2XP1": OPCODE,
	"HLT":     OPCODE,
	"IDIV":    OPCODE,
	"IMUL":    OPCODE,
	"IN":      OPCODE,
	"INC":     OPCODE,
	"INCO":    OPCODE,
	"INSB":    OPCODE,
	"INSD":    OPCODE,
	"INSW":    OPCODE,
	"INT":     OPCODE,
	"INT3":    OPCODE,
	"INTO":    OPCODE,
	"INVD":    OPCODE,
	"INVLPG":  OPCODE,
	"IRET":    OPCODE,
	"IRETD":   OPCODE,
	"IRETW":   OPCODE,
	"JA":      OPCODE,
	"JAE":     OPCODE,
	"JB":      OPCODE,
	"JBE":     OPCODE,
	"JC":      OPCODE,
	"JCXZ":    OPCODE,
	"JE":      OPCODE,
	"JECXZ":   OPCODE,
	"JG":      OPCODE,
	"JGE":     OPCODE,
	"JL":      OPCODE,
	"JLE":     OPCODE,
	"JMP":     OPCODE,
	"JNA":     OPCODE,
	"JNAE":    OPCODE,
	"JNB":     OPCODE,
	"JNBE":    OPCODE,
	"JNC":     OPCODE,
	"JNE":     OPCODE,
	"JNG":     OPCODE,
	"JNGE":    OPCODE,
	"JNL":     OPCODE,
	"JNLE":    OPCODE,
	"JNO":     OPCODE,
	"JNP":     OPCODE,
	"JNS":     OPCODE,
	"JNZ":     OPCODE,
	"JO":      OPCODE,
	"JP":      OPCODE,
	"JPE":     OPCODE,
	"JPO":     OPCODE,
	"JS":      OPCODE,
	"JZ":      OPCODE,
	"LAHF":    OPCODE,
	"LAR":     OPCODE,
	"LDS":     OPCODE,
	"LEA":     OPCODE,
	"LEAVE":   OPCODE,
	"LES":     OPCODE,
	"LFS":     OPCODE,
	"LGDT":    OPCODE,
	"LGS":     OPCODE,
	"LIDT":    OPCODE,
	"LLDT":    OPCODE,
	"LMSW":    OPCODE,
	"LOCK":    OPCODE,
	"LODSB":   OPCODE,
	"LODSD":   OPCODE,
	"LODSW":   OPCODE,
	"LOOP":    OPCODE,
	"LOOPE":   OPCODE,
	"LOOPNE":  OPCODE,
	"LOOPNZ":  OPCODE,
	"LOOPZ":   OPCODE,
	"LSL":     OPCODE,
	"LSS":     OPCODE,
	"LTR":     OPCODE,
	"MOV":     OPCODE,
	"MOVSB":   OPCODE,
	"MOVSD":   OPCODE,
	"MOVSW":   OPCODE,
	"MOVSX":   OPCODE,
	"MOVZX":   OPCODE,
	"MUL":     OPCODE,
	"NEG":     OPCODE,
	"NOP":     OPCODE,
	"NOT":     OPCODE,
	"OR":      OPCODE,
	"ORG":     OPCODE,
	"OUT":     OPCODE,
	"OUTSB":   OPCODE,
	"OUTSD":   OPCODE,
	"OUTSW":   OPCODE,
	"POP":     OPCODE,
	"POPA":    OPCODE,
	"POPAD":   OPCODE,
	"POPAW":   OPCODE,
	"POPF":    OPCODE,
	"POPFD":   OPCODE,
	"POPFW":   OPCODE,
	"PUSH":    OPCODE,
	"PUSHA":   OPCODE,
	"PUSHD":   OPCODE,
	"PUSHAD":  OPCODE,
	"PUSHAW":  OPCODE,
	"PUSHF":   OPCODE,
	"PUSHFD":  OPCODE,
	"PUSHFW":  OPCODE,
	"RCL":     OPCODE,
	"RCR":     OPCODE,
	"RDMSR":   OPCODE,
	"RDPMC":   OPCODE,
	"REP":     OPCODE,
	"REPE":    OPCODE,
	"REPNE":   OPCODE,
	"REPNZ":   OPCODE,
	"REPZ":    OPCODE,
	"RESB":    OPCODE,
	"RESD":    OPCODE,
	"RESQ":    OPCODE,
	"REST":    OPCODE,
	"RESW":    OPCODE,
	"RET":     OPCODE,
	"RETF":    OPCODE,
	"RETN":    OPCODE,
	"ROL":     OPCODE,
	"ROR":     OPCODE,
	"RSM":     OPCODE,
	"SAHF":    OPCODE,
	"SAL":     OPCODE,
	"SAR":     OPCODE,
	"SBB":     OPCODE,
	"SCASB":   OPCODE,
	"SCASD":   OPCODE,
	"SCASW":   OPCODE,
	"SETA":    OPCODE,
	"SETAE":   OPCODE,
	"SETB":    OPCODE,
	"SETBE":   OPCODE,
	"SETC":    OPCODE,
	"SETE":    OPCODE,
	"SETG":    OPCODE,
	"SETGE":   OPCODE,
	"SETL":    OPCODE,
	"SETLE":   OPCODE,
	"SETNA":   OPCODE,
	"SETNAE":  OPCODE,
	"SETNB":   OPCODE,
	"SETNBE":  OPCODE,
	"SETNC":   OPCODE,
	"SETNE":   OPCODE,
	"SETNG":   OPCODE,
	"SETNGE":  OPCODE,
	"SETNL":   OPCODE,
	"SETNLE":  OPCODE,
	"SETNO":   OPCODE,
	"SETNP":   OPCODE,
	"SETNS":   OPCODE,
	"SETNZ":   OPCODE,
	"SETO":    OPCODE,
	"SETP":    OPCODE,
	"SETPE":   OPCODE,
	"SETPO":   OPCODE,
	"SETS":    OPCODE,
	"SETZ":    OPCODE,
	"SGDT":    OPCODE,
	"SHL":     OPCODE,
	"SHLD":    OPCODE,
	"SHR":     OPCODE,
	"SHRD":    OPCODE,
	"SIDT":    OPCODE,
	"SLDT":    OPCODE,
	"SMSW":    OPCODE,
	"STC":     OPCODE,
	"STD":     OPCODE,
	"STI":     OPCODE,
	"STOSB":   OPCODE,
	"STOSD":   OPCODE,
	"STOSW":   OPCODE,
	"STR":     OPCODE,
	"SUB":     OPCODE,
	"TEST":    OPCODE,
	"TIMES":   OPCODE,
	"UD2":     OPCODE,
	"VERR":    OPCODE,
	"VERW":    OPCODE,
	"WAIT":    OPCODE,
	"WBINVD":  OPCODE,
	"WRMSR":   OPCODE,
	"XADD":    OPCODE,
	"XCHG":    OPCODE,
	"XLATB":   OPCODE,
	"XOR":     OPCODE,
}
