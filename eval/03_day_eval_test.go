package eval

import (
	"strings"
	"testing"
)

const DAY3_HARIB00A_ASM_SRC = `; haribote-ipl
; TAB=4

		ORG		0x7c00			; このプログラムがどこに読み込まれるのか

; 以下は標準的なFAT12フォーマットフロッピーディスクのための記述

		JMP		entry
		DB		0x90
		DB		"HARIBOTE"		; ブートセクタの名前を自由に書いてよい（8バイト）
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
		DB		"HARIBOTEOS "	; ディスクの名前（11バイト）
		DB		"FAT12   "		; フォーマットの名前（8バイト）
		RESB	18				; とりあえず18バイトあけておく

; プログラム本体

entry:
		MOV		AX,0			; レジスタ初期化
		MOV		SS,AX
		MOV		SP,0x7c00
		MOV		DS,AX

; ディスクを読む

		MOV		AX,0x0820
		MOV		ES,AX
		MOV		CH,0			; シリンダ0
		MOV		DH,0			; ヘッド0
		MOV		CL,2			; セクタ2

		MOV		AH,0x02			; AH=0x02 : ディスク読み込み
		MOV		AL,1			; 1セクタ
		MOV		BX,0
		MOV		DL,0x00			; Aドライブ
		INT		0x13			; ディスクBIOS呼び出し
		JC		error

; 読み終わったけどとりあえずやることないので寝る

fin:
		HLT						; 何かあるまでCPUを停止させる
		JMP		fin				; 無限ループ

error:
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
msg:
		DB		0x0a, 0x0a		; 改行を2つ
		DB		"load error"
		DB		0x0a			; 改行
		DB		0

		RESB	0x7dfe-$		; 0x7dfeまでを0x00で埋める命令

		DB		0x55, 0xaa
`

// TestHelloOS3 naskソース２日目(harib00a)のテスト
func TestDumpHarib00a(t *testing.T) {
	input := DAY3_HARIB00A_ASM_SRC
	testAsmSourceOnlyDump(t, input, []string{""})
}

// TestHelloOS3 naskソース３日目(harib00a)のテスト
func TestHarib00a(t *testing.T) {
	input := DAY3_HARIB00A_ASM_SRC

	// wine nask.exe ipl.nas ipl.obj
	// hexdump -C ipl.obj > ipl.hex
	// ... generate .hex file little endian
	answer := `00000000  eb 4e 90 48 41 52 49 42  4f 54 45 00 02 01 01 00  |.N.HARIBOTE.....|
00000010  02 e0 00 40 0b f0 09 00  12 00 02 00 00 00 00 00  |...@............|
00000020  40 0b 00 00 00 00 29 ff  ff ff ff 48 41 52 49 42  |@.....)....HARIB|
00000030  4f 54 45 4f 53 20 46 41  54 31 32 20 20 20 00 00  |OTEOS FAT12   ..|
00000040  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000050  b8 00 00 8e d0 bc 00 7c  8e d8 b8 20 08 8e c0 b5  |.......|... ....|
00000060  00 b6 00 b1 02 b4 02 b0  01 bb 00 00 b2 00 cd 13  |................|
00000070  72 03 f4 eb fd be 8a 7c  8a 04 83 c6 01 3c 00 74  |r......|.....<.t|
00000080  f1 b4 0e bb 0f 00 cd 10  eb ee 0a 0a 6c 6f 61 64  |............load|
00000090  20 65 72 72 6f 72 0a 00  00 00 00 00 00 00 00 00  | error..........|
000000a0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
*
000001f0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 55 aa  |..............U.|
` // 00000200 ← hexdumpだと左の行が出るが、Goのdumpでは出ない。動作に支障はないので無視する。

	testAsmSource(t, input, strings.Split(answer, "\n"))
}
