package eval

import (
	"encoding/hex"
	"fmt"
	"github.com/hangingman/gosk/object"
	"log"
)

type LabelManagement struct {
	labelBinaryRefMap map[string]*object.Binary // バイナリへの参照を記録し、後で更新する
	labelBytesMap     map[string]int
	opcode            map[string][]byte
	genBytesFns       map[string]func(i int) []byte // 取得した値をどうやってバイト列に戻すか
}

// AddLabelCallback は後でニーモニックが決まるような命令（JMP命令や一部のMOV命令）を処理する
// @param opcode 最終的な機械語用オペコード
// @param ident 使用されるラベル
// @param bin 機械語の格納先コンテナ
// @param from ラベルのあった位置
func (l *LabelManagement) AddLabelCallback(
	opcode []byte,
	ident string,
	bin *object.Binary,
	from int,
	howToGenBytes func(i int) []byte) {

	log.Println(fmt.Sprintf("info: add label %s from %d !!", ident, from))
	l.opcode[ident] = opcode
	l.labelBinaryRefMap[ident] = bin
	l.labelBytesMap[ident] = from
	l.genBytesFns[ident] = howToGenBytes
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
		log.Println(fmt.Sprintf("info: hex style => %s", hex.EncodeToString(int2Byte(to-from))))
		bin.Value = append(bin.Value, opcode...)
		bin.Value = append(bin.Value, l.genBytesFns[ident](to-from)...)
	}

	// TODO: 本当に必要かどうか後で検証
	return 0
}
