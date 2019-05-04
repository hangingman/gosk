package eval

import (
	"strings"
	"testing"
)

const DAY2_ASM_SRC = `; hello-os
; TAB=4

		ORG		0x7c00			; このプログラムがどこに読み込まれるのか

; 以下は標準的なFAT12フォーマットフロッピーディスクのための記述

		JMP		entry
		DB		0x90
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

entry:
		MOV		AX,0			; レジスタ初期化
		MOV		SS,AX
		MOV		SP,0x7c00
		MOV		DS,AX
		MOV		ES,AX

		MOV		SI,msg
putloop:
		MOV		AL,[SI]
		ADD		SI,1			; SIに1を足す
		CMP		AL,0
		JE		fin
		MOV		AH,0x0e			; 一文字表示ファンクション
		MOV		BX,15			; カラーコード
		INT		0x10			; ビデオBIOS呼び出し
		JMP		putloop
fin:
		HLT						; 何かあるまでCPUを停止させる
		JMP		fin				; 無限ループ

msg:
		DB		0x0a, 0x0a		; 改行を2つ
		DB		"hello, world"
		DB		0x0a			; 改行
		DB		0

		RESB	0x7dfe-$		; 0x7dfeまでを0x00で埋める命令

		DB		0x55, 0xaa

; 以下はブートセクタ以外の部分の記述

		DB		0xf0, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00
		RESB	4600
		DB		0xf0, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00
		RESB	1469432
`

// TestHelloOS3 naskソース２日目(helloos3)のテスト
func TestAdd(t *testing.T) {
	input := "ADD		SI,1			; SIに1を足す"
	answer := []string{"00000000  83 c6 01                                          |...|", ""}

	testAsmSource(t, input, answer)
}

// TestHelloOS3 naskソース２日目(helloos3)のテスト
func TestMovDisp(t *testing.T) {
	input := "MOV		AL,[SI]"
	answer := []string{"00000000  8a 04                                             |..|", ""}

	testAsmSource(t, input, answer)
}

// TestHelloOS3 naskソース２日目(helloos3)のテスト
func TestCmp(t *testing.T) {
	input := "CMP		AL,0"
	answer := []string{"00000000  3c 00                                             |<.|", ""}

	testAsmSource(t, input, answer)
}

// TestHelloOS3 naskソース２日目(helloos3)のテスト
func TestDumpHelloOS3(t *testing.T) {
	input := DAY2_ASM_SRC
	testAsmSourceOnlyDump(t, input, []string{""})
}

// TestHelloOS3 naskソース２日目(helloos3)のテスト
func TestHelloOS3(t *testing.T) {
	input := DAY2_ASM_SRC

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
