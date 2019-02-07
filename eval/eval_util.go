package eval

import (
	"encoding/binary"
	"reflect"
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
