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

// AddLabelCallback は後でニーモニックが決まるような命令（JMP命令や一部のMOV命令）を処理する
// @param opcode 最終的な機械語用オペコード
// @param ident 使用されるラベル
// @param bin 機械語の格納先コンテナ
// @param from ラベルのあった位置
func (l *LabelManagement) AddLabelCallback(opcode []byte, ident string, bin *object.Binary, from int) {
	log.Println(fmt.Sprintf("info: add label %s from %d !!", ident, from))
	l.opcode[ident] = opcode
	l.labelBinaryRefMap[ident] = bin
	l.labelBytesMap[ident] = from
}

func (l *LabelManagement) RemoveLabelCallback(ident string) {
	log.Println(fmt.Sprintf("info: remove label %s !!", ident))
}

// Emit はAddLabelCallbackを使用後にラベルが見つかったときのコールバック関数
// コールバックとは書いたが、呼び出すのは自分自身
// @param ident 使用されるラベル
// @param from ラベルのあった位置
// @return 増えたバイトサイズ
func (l *LabelManagement) Emit(ident string, to int) int {
	opcode, opcodeOk := l.opcode[ident]
	bin, binOk := l.labelBinaryRefMap[ident]
	from, fromOk := l.labelBytesMap[ident]

	if opcodeOk && binOk && fromOk {
		log.Println(fmt.Sprintf("info: from=%d, to=%d", from, to))
		log.Println(fmt.Sprintf("info: emit label %s to %d !!", ident, to-from))
		bin.Value = append(bin.Value, opcode...)
		bin.Value = append(bin.Value, int2Byte(to-from)...)
	}

	// TODO: 本当に必要かどうか後で検証
	return 0
}
