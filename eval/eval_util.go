package eval

import (
	"encoding/binary"
	"encoding/hex"
	"github.com/hangingman/gosk/token"
	"reflect"
	"strconv"
)

func IsNil(x interface{}) bool {
	return x == nil || reflect.ValueOf(x).IsNil()
}

func IsNotNil(x interface{}) bool {
	return !IsNil(x)
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

func isByteHex(tok token.Token) bool {
	return tok.Type == token.HEX_LIT && len(tok.Literal) == 4
}

func isWordHex(tok token.Token) bool {
	return tok.Type == token.HEX_LIT && len(tok.Literal) == 6
}

func isDwordHex(tok token.Token) bool {
	return tok.Type == token.HEX_LIT && len(tok.Literal) == 8
}

func imm8ToByte(tok token.Token) []byte {
	if tok.Type == token.HEX_LIT {
		bs, _ := hex.DecodeString(string([]rune(tok.Literal)[2:]))
		return bs
	}
	if tok.Type == token.INT {
		v, _ := strconv.Atoi(tok.Literal)
		return int2Byte(v)
	}
	return nil
}

func imm16ToWord(tok token.Token) []byte {
	if tok.Type == token.HEX_LIT {
		bs, _ := hex.DecodeString(string([]rune(tok.Literal)[2:]))
		return bs
	}
	if tok.Type == token.INT {
		v, _ := strconv.Atoi(tok.Literal)
		return int2Word(v)
	}
	return nil
}

func imm32ToDword(tok token.Token) []byte {
	if tok.Type == token.HEX_LIT {
		bs, _ := hex.DecodeString(string([]rune(tok.Literal)[2:]))
		return bs
	}
	if tok.Type == token.INT {
		v, _ := strconv.Atoi(tok.Literal)
		return int2Dword(v)
	}
	return nil
}
