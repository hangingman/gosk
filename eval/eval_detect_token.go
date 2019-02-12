package eval

import (
	"fmt"
	"github.com/hangingman/gosk/token"
	"regexp"
	"strconv"
)

// 8bit, 16bit, 32bitのレジスタ
var r8 = regexp.MustCompile(`AL|CL|DL|BL|AH|CH|DH|BH`)
var r16 = regexp.MustCompile(`AX|BX|CX|DX|SI|DI|BP|SP|IP|FLAGS|CS|SS|DS|ES|FS|GS`)
var r32 = regexp.MustCompile(`EAX|EBX|ECX|EDX|ESI|EDI|EBP|ESP|EIP|EFLAGS`)
var sr = regexp.MustCompile(`ES|CS|SS|DS|FS|GS`)

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

func IsSreg(tok token.Token) bool {
	return IsR(tok, sr)
}

// func IsMoffs8(tok token.Token) bool {
// }

// func IsMoffs16(tok token.Token) bool {
// }

// func IsMoffs32(tok token.Token) bool {
// }
