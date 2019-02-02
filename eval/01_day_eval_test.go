package eval

import (
	"fmt"
	"github.com/hangingman/gosk/lexer"
	"github.com/hangingman/gosk/object"
	"github.com/hangingman/gosk/parser"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
)

// testAsmSource はアセンブラソースとプレイン16進ダンプを受け取りテスト比較する
func testAsmSource(t *testing.T, asmSource string, expectedHex []string) {

	l := lexer.New(asmSource)
	p := parser.New(l)

	// プログラムの解析と評価
	program := p.ParseProgram()
	evaluated := Eval(program)
	// リフレクションで結果をチェック
	assert.Equal(t, "*object.ObjectArray", reflect.TypeOf(evaluated).String())
	// キャストをやる
	objArray, ok := evaluated.(*object.ObjectArray)
	assert.True(t, ok)

	actual := []byte{}
	for _, obj := range *objArray {
		if obj != nil {
			assert.Equal(t, "*object.Binary", reflect.TypeOf(obj).String())
			bin, _ := obj.(*object.Binary)
			actual = append(actual, bin.Value...)
		}
	}

	for i, hex := range getHexdumpFmtArray(actual) {
		assert.Equal(t, expectedHex[i], hex,
			fmt.Sprintf("expectedHex[%d] should be = %s", i, expectedHex[i]))
	}
}

// TestHelloOS1 naskソース１日目(helloos1)のテスト
func TestHelloOS1(t *testing.T) {
	input := `	DB	0xeb, 0x4e, 0x90, 0x48, 0x45, 0x4c, 0x4c, 0x4f
	DB	0x49, 0x50, 0x4c, 0x00, 0x02, 0x01, 0x01, 0x00
	DB	0x02, 0xe0, 0x00, 0x40, 0x0b, 0xf0, 0x09, 0x00
	DB	0x12, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00
	DB	0x40, 0x0b, 0x00, 0x00, 0x00, 0x00, 0x29, 0xff
	DB	0xff, 0xff, 0xff, 0x48, 0x45, 0x4c, 0x4c, 0x4f
	DB	0x2d, 0x4f, 0x53, 0x20, 0x20, 0x20, 0x46, 0x41
	DB	0x54, 0x31, 0x32, 0x20, 0x20, 0x20, 0x00, 0x00
	RESB	16
	DB	0xb8, 0x00, 0x00, 0x8e, 0xd0, 0xbc, 0x00, 0x7c
	DB	0x8e, 0xd8, 0x8e, 0xc0, 0xbe, 0x74, 0x7c, 0x8a
	DB	0x04, 0x83, 0xc6, 0x01, 0x3c, 0x00, 0x74, 0x09
	DB	0xb4, 0x0e, 0xbb, 0x0f, 0x00, 0xcd, 0x10, 0xeb
	DB	0xee, 0xf4, 0xeb, 0xfd, 0x0a, 0x0a, 0x68, 0x65
	DB	0x6c, 0x6c, 0x6f, 0x2c, 0x20, 0x77, 0x6f, 0x72
	DB	0x6c, 0x64, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x00
	RESB	368
	DB	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x55, 0xaa
	DB	0xf0, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00
	RESB	4600
	DB	0xf0, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00
	RESB	1469432`

	// wine nask.exe helloos.nas > helloos.obj
	// hexdump -C helloos.obj > helloos.hex
	// ... generate .hex file little endian
	answer := `00000000  eb 4e 90 48 45 4c 4c 4f  49 50 4c 00 02 01 01 00  |.N.HELLOIPL.....|
00000010  02 e0 00 40 0b f0 09 00  12 00 02 00 00 00 00 00  |...@............|
00000020  40 0b 00 00 00 00 29 ff  ff ff ff 48 45 4c 4c 4f  |@.....)....HELLO|
00000030  2d 4f 53 20 20 20 46 41  54 31 32 20 20 20 00 00  |-OS   FAT12   ..|
00000040  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000050  b8 00 00 8e d0 bc 00 7c  8e d8 8e c0 be 74 7c 8a  |.......|.....t|.|
00000060  04 83 c6 01 3c 00 74 09  b4 0e bb 0f 00 cd 10 eb  |....<.t.........|
00000070  ee f4 eb fd 0a 0a 68 65  6c 6c 6f 2c 20 77 6f 72  |......hello, wor|
00000080  6c 64 0a 00 00 00 00 00  00 00 00 00 00 00 00 00  |ld..............|
00000090  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
*
000001f0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 55 aa  |..............U.|
00000200  f0 ff ff 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000210  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
*
00001400  f0 ff ff 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00001410  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
*
00168000`

	testAsmSource(t, input, strings.Split(answer, "\n"))
}

// TestHelloOS2 naskソース１日目(helloos2)のテスト
func TestHelloOS2(t *testing.T) {
	input := `; hello-os
; TAB=4

; 以下は標準的なFAT12フォーマットフロッピーディスクのための記述

		DB		0xeb, 0x4e, 0x90
		DB		"HELLOIPL"		; ブートセクタの名前を自由に書いてよい（8バイト）
		DW		512				; 1セクタの大きさ（512にしなければいけない）
		DB		1				; クラスタの大きさ（1セクタにしなければいけない）
		DW		1				; FATがどこから始まるか（普通は1セクタ目からにする）
		DB		2				; FATの個数（2にしなければいけない）
		DW		224				; ルートディレクトリ領域の大きさ（普通は224エントリにする）
		DW		2880			; このドライブの大きさ（2880セクタにしなければいけない）
		DB		0xf0			; メディアのタイプ（0xf0にしなければいけない）
		DW		9				; FAT領域の長さ（9セクタにしなければいけない）
		DW		18				; 1トラックにいくつのセクタがあるか（18にしなければいけない）
		DW		2				; ヘッドの数（2にしなければいけない）
		DD		0				; パーティションを使ってないのでここは必ず0
		DD		2880			; このドライブ大きさをもう一度書く
		DB		0,0,0x29		; よくわからないけどこの値にしておくといいらしい
		DD		0xffffffff		; たぶんボリュームシリアル番号
		DB		"HELLO-OS   "	; ディスクの名前（11バイト）
		DB		"FAT12   "		; フォーマットの名前（8バイト）
		RESB	18				; とりあえず18バイトあけておく

; プログラム本体

		DB		0xb8, 0x00, 0x00, 0x8e, 0xd0, 0xbc, 0x00, 0x7c
		DB		0x8e, 0xd8, 0x8e, 0xc0, 0xbe, 0x74, 0x7c, 0x8a
		DB		0x04, 0x83, 0xc6, 0x01, 0x3c, 0x00, 0x74, 0x09
		DB		0xb4, 0x0e, 0xbb, 0x0f, 0x00, 0xcd, 0x10, 0xeb
		DB		0xee, 0xf4, 0xeb, 0xfd

; メッセージ部分

		DB		0x0a, 0x0a		; 改行を2つ
		DB		"hello, world"
		DB		0x0a			; 改行
		DB		0

		RESB	0x1fe-$			; 0x001feまでを0x00で埋める命令

		DB		0x55, 0xaa

; 以下はブートセクタ以外の部分の記述

		DB		0xf0, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00
		RESB	4600
		DB		0xf0, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00
		RESB	1469432`

	// wine nask.exe helloos.nas > helloos.obj
	// hexdump -C helloos.obj > helloos.hex
	// ... generate .hex file little endian
	answer := `00000000  eb 4e 90 48 45 4c 4c 4f  49 50 4c 00 02 01 01 00  |.N.HELLOIPL.....|
00000010  02 e0 00 40 0b f0 09 00  12 00 02 00 00 00 00 00  |...@............|
00000020  40 0b 00 00 00 00 29 ff  ff ff ff 48 45 4c 4c 4f  |@.....)....HELLO|
00000030  2d 4f 53 20 20 20 46 41  54 31 32 20 20 20 00 00  |-OS   FAT12   ..|
00000040  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000050  b8 00 00 8e d0 bc 00 7c  8e d8 8e c0 be 74 7c 8a  |.......|.....t|.|
00000060  04 83 c6 01 3c 00 74 09  b4 0e bb 0f 00 cd 10 eb  |....<.t.........|
00000070  ee f4 eb fd 0a 0a 68 65  6c 6c 6f 2c 20 77 6f 72  |......hello, wor|
00000080  6c 64 0a 00 00 00 00 00  00 00 00 00 00 00 00 00  |ld..............|
00000090  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
*
000001f0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 55 aa  |..............U.|
00000200  f0 ff ff 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000210  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
*
00001400  f0 ff ff 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00001410  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
*
00168000`

	testAsmSource(t, input, strings.Split(answer, "\n"))
}