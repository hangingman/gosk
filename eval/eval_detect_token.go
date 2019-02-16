package eval

import (
	"fmt"
	"github.com/hangingman/gosk/token"
	"log"
	"regexp"
	"strconv"
)

// 8bit, 16bit, 32bitのレジスタ
var r8 = regexp.MustCompile(`AL|CL|DL|BL|AH|CH|DH|BH`)
var r16 = regexp.MustCompile(`AX|CX|DX|BX|SP|BP|SI|DI|IP|FLAGS|CS|SS|DS|ES|FS|GS`)
var r32 = regexp.MustCompile(`EAX|ECX|EDX|EBX|ESP|EBP|ESI|EDI|EIP|EFLAGS`)
var sr = regexp.MustCompile(`ES|CS|SS|DS|FS|GS`)

// 8bit, 16bit, 32bitのレジスタとレジスタコードの対応
var r8CodeMap = map[string]int{
	"AL": 0, "CL": 1, "DL": 2, "BL": 3, "AH": 4, "CH": 5, "DH": 6, "BH": 7,
}
var r16CodeMap = map[string]int{
	"AX": 0, "CX": 1, "DX": 2, "BX": 3, "SP": 4, "BP": 5, "SI": 6, "DI": 7,
}
var r32CodeMap = map[string]int{
	"EAX": 0, "ECX": 1, "EDX": 2, "EBX": 3, "ESP": 4, "EBP": 5, "ESI": 6, "EDI": 7,
}
var sregCodeMap = map[string]int{
	"ES": 0, "CS": 1, "SS": 2, "DS": 3, "FS": 4, "GS": 5,
}

func IsR(tok token.Token, re *regexp.Regexp) bool {
	return tok.Type == token.REGISTER && re.MatchString(tok.Literal)
}

func IsR8(tok token.Token) bool {
	return IsR(tok, r8)
}

func IsR16(tok token.Token) bool {
	return IsR(tok, r16)
}

func IsR32(tok token.Token) bool {
	return IsR(tok, r32)
}

func IsSreg(tok token.Token) bool {
	log.Println(fmt.Sprintf("info: %s !!!!!!!!!!!!!", tok))
	log.Println(fmt.Sprintf("info: %v !!!!!!!!!!!!!",
		tok.Type == token.SEG_REGISTER,
	))
	log.Println(fmt.Sprintf("info: %v !!!!!!!!!!!!!",
		sr.MatchString(tok.Literal),
	))
	return tok.Type == token.SEG_REGISTER && sr.MatchString(tok.Literal)
}

func IsImm8(tok token.Token) bool {
	if tok.Type == token.INT {
		_, err := strconv.ParseInt(tok.Literal, 10, 8)
		return IsNil(err)
	}
	if isByteHex(tok) {
		return true
	}
	return false
}

func IsImm16(tok token.Token) bool {
	if tok.Type == token.INT {
		_, err := strconv.ParseInt(tok.Literal, 10, 16)
		return IsNil(err)
	}
	if isWordHex(tok) {
		return true
	}
	return false
}

func IsImm32(tok token.Token) bool {
	if tok.Type == token.INT {
		_, err := strconv.ParseInt(tok.Literal, 10, 32)
		return IsNil(err)
	}
	if isDwordHex(tok) {
		return true
	}
	return false
}

// func IsRm8(tok token.Token) bool {
// }

// func IsRm16(tok token.Token) bool {
// }

// func IsRm32(tok token.Token) bool {
// }

// func IsMoffs8(tok token.Token) bool {
// }

// func IsMoffs16(tok token.Token) bool {
// }

// func IsMoffs32(tok token.Token) bool {
// }