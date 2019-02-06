package eval

import (
	"fmt"
	"github.com/hangingman/gosk/object"
	"log"
)

type LabelManagement struct {
	labelBinaryRefMap map[string]*object.Binary // バイナリへの参照を記録し、後で更新する
	labelBytesMap     map[string]int
	opcode            map[string][]byte
}

func (l *LabelManagement) AddLabelCallback(opcode []byte, ident string, bin *object.Binary, from int) {
	log.Println(fmt.Sprintf("info: add label %s from %d !!", ident, from))
	l.opcode[ident] = opcode
	l.labelBinaryRefMap[ident] = bin
	l.labelBytesMap[ident] = from
}

func (l *LabelManagement) RemoveLabelCallback(ident string) {
	log.Println(fmt.Sprintf("info: remove label %s !!", ident))
}

func (l *LabelManagement) Emit(ident string, to int) {
	log.Println(fmt.Sprintf("info: emit label %s to %d !!", ident, to))
	opcode, opcodeOk := l.opcode[ident]
	bin, binOk := l.labelBinaryRefMap[ident]
	from, fromOk := l.labelBytesMap[ident]
	if opcodeOk && binOk && fromOk {
		bin.Value = append(bin.Value, opcode...)
		bin.Value = append(bin.Value, int2Byte(to-from)...)
	}
}
