package eval

import (
	"strings"
	"testing"
	"time"
)

func TestDumpHarib(t *testing.T) {
	// input := getAsmSource("03_day_harib00f_ipl.nas")
	// testAsmSourceOnlyDump(t, input, []string{""})

	// 実際のテスト
	t.Run("harib00a", testHarib00a)
	t.Run("harib00b", testHarib00b)
	t.Run("harib00c", testHarib00c)
	t.Run("harib00d", testHarib00d)
	t.Run("harib00e", testHarib00e)
	t.Run("harib00f", testHarib00f)
	t.Run("harib00g", testHarib00g)
	// t.Run("harib00h", testHarib00h)
}

// TestHelloOS3 naskソース３日目(harib00a)のテスト
func testHarib00a(t *testing.T) {
	time.Sleep(1 * time.Second)
	input := getAsmSource("03_day_harib00a_ipl.nas")

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

// TestHelloOS3 naskソース３日目(harib00b)のテスト
func testHarib00b(t *testing.T) {
	time.Sleep(1 * time.Second)
	input := getAsmSource("03_day_harib00b_ipl.nas")

	// wine nask.exe ipl.nas ipl.obj
	// hexdump -C ipl.obj > ipl.hex
	// ... generate .hex file little endian
	answer := `00000000  eb 4e 90 48 41 52 49 42  4f 54 45 00 02 01 01 00  |.N.HARIBOTE.....|
00000010  02 e0 00 40 0b f0 09 00  12 00 02 00 00 00 00 00  |...@............|
00000020  40 0b 00 00 00 00 29 ff  ff ff ff 48 41 52 49 42  |@.....)....HARIB|
00000030  4f 54 45 4f 53 20 46 41  54 31 32 20 20 20 00 00  |OTEOS FAT12   ..|
00000040  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000050  b8 00 00 8e d0 bc 00 7c  8e d8 b8 20 08 8e c0 b5  |.......|... ....|
00000060  00 b6 00 b1 02 be 00 00  b4 02 b0 01 bb 00 00 b2  |................|
00000070  00 cd 13 73 10 83 c6 01  83 fe 05 73 0b b4 00 b2  |...s.......s....|
00000080  00 cd 13 eb e3 f4 eb fd  be 9d 7c 8a 04 83 c6 01  |..........|.....|
00000090  3c 00 74 f1 b4 0e bb 0f  00 cd 10 eb ee 0a 0a 6c  |<.t............l|
000000a0  6f 61 64 20 65 72 72 6f  72 0a 00 00 00 00 00 00  |oad error.......|
000000b0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
*
000001f0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 55 aa  |..............U.|
`
	testAsmSource(t, input, strings.Split(answer, "\n"))
}

// TestHelloOS3 naskソース３日目(harib00c)のテスト
func testHarib00c(t *testing.T) {
	time.Sleep(1 * time.Second)
	input := getAsmSource("03_day_harib00c_ipl.nas")

	// wine nask.exe ipl.nas ipl.obj
	// hexdump -C ipl.obj > ipl.hex
	// ... generate .hex file little endian
	answer := `00000000  eb 4e 90 48 41 52 49 42  4f 54 45 00 02 01 01 00  |.N.HARIBOTE.....|
00000010  02 e0 00 40 0b f0 09 00  12 00 02 00 00 00 00 00  |...@............|
00000020  40 0b 00 00 00 00 29 ff  ff ff ff 48 41 52 49 42  |@.....)....HARIB|
00000030  4f 54 45 4f 53 20 46 41  54 31 32 20 20 20 00 00  |OTEOS FAT12   ..|
00000040  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000050  b8 00 00 8e d0 bc 00 7c  8e d8 b8 20 08 8e c0 b5  |.......|... ....|
00000060  00 b6 00 b1 02 be 00 00  b4 02 b0 01 bb 00 00 b2  |................|
00000070  00 cd 13 73 10 83 c6 01  83 fe 05 73 1a b4 00 b2  |...s.......s....|
00000080  00 cd 13 eb e3 8c c0 05  20 00 8e c0 80 c1 01 80  |........ .......|
00000090  f9 12 76 d1 f4 eb fd be  ac 7c 8a 04 83 c6 01 3c  |..v......|.....<|
000000a0  00 74 f1 b4 0e bb 0f 00  cd 10 eb ee 0a 0a 6c 6f  |.t............lo|
000000b0  61 64 20 65 72 72 6f 72  0a 00 00 00 00 00 00 00  |ad error........|
000000c0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
*
000001f0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 55 aa  |..............U.|
`

	testAsmSource(t, input, strings.Split(answer, "\n"))
}

// TestHelloOS3 naskソース３日目(harib00d)のテスト
func testHarib00d(t *testing.T) {
	time.Sleep(1 * time.Second)
	input := getAsmSource("03_day_harib00d_ipl.nas")

	// wine nask.exe ipl.nas ipl.obj
	// hexdump -C ipl.obj > ipl.hex
	// ... generate .hex file little endian
	answer := `00000000  eb 4e 90 48 41 52 49 42  4f 54 45 00 02 01 01 00  |.N.HARIBOTE.....|
00000010  02 e0 00 40 0b f0 09 00  12 00 02 00 00 00 00 00  |...@............|
00000020  40 0b 00 00 00 00 29 ff  ff ff ff 48 41 52 49 42  |@.....)....HARIB|
00000030  4f 54 45 4f 53 20 46 41  54 31 32 20 20 20 00 00  |OTEOS FAT12   ..|
00000040  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000050  b8 00 00 8e d0 bc 00 7c  8e d8 b8 20 08 8e c0 b5  |.......|... ....|
00000060  00 b6 00 b1 02 be 00 00  b4 02 b0 01 bb 00 00 b2  |................|
00000070  00 cd 13 73 10 83 c6 01  83 fe 05 73 2e b4 00 b2  |...s.......s....|
00000080  00 cd 13 eb e3 8c c0 05  20 00 8e c0 80 c1 01 80  |........ .......|
00000090  f9 12 76 d1 b1 01 80 c6  01 80 fe 02 72 c7 b6 00  |..v.........r...|
000000a0  80 c5 01 80 fd 0a 72 bd  f4 eb fd be c0 7c 8a 04  |......r......|..|
000000b0  83 c6 01 3c 00 74 f1 b4  0e bb 0f 00 cd 10 eb ee  |...<.t..........|
000000c0  0a 0a 6c 6f 61 64 20 65  72 72 6f 72 0a 00 00 00  |..load error....|
000000d0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
*
000001f0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 55 aa  |..............U.|
`

	testAsmSource(t, input, strings.Split(answer, "\n"))
}

// TestHelloOS3 naskソース３日目(harib00e)のテスト
func testHarib00e(t *testing.T) {
	time.Sleep(1 * time.Second)
	input := getAsmSource("03_day_harib00e_ipl.nas")

	// wine nask.exe ipl.nas ipl.obj
	// hexdump -C ipl.obj > ipl.hex
	// ... generate .hex file little endian
	answer := `00000000  eb 4e 90 48 41 52 49 42  4f 54 45 00 02 01 01 00  |.N.HARIBOTE.....|
00000010  02 e0 00 40 0b f0 09 00  12 00 02 00 00 00 00 00  |...@............|
00000020  40 0b 00 00 00 00 29 ff  ff ff ff 48 41 52 49 42  |@.....)....HARIB|
00000030  4f 54 45 4f 53 20 46 41  54 31 32 20 20 20 00 00  |OTEOS FAT12   ..|
00000040  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000050  b8 00 00 8e d0 bc 00 7c  8e d8 b8 20 08 8e c0 b5  |.......|... ....|
00000060  00 b6 00 b1 02 be 00 00  b4 02 b0 01 bb 00 00 b2  |................|
00000070  00 cd 13 73 10 83 c6 01  83 fe 05 73 2e b4 00 b2  |...s.......s....|
00000080  00 cd 13 eb e3 8c c0 05  20 00 8e c0 80 c1 01 80  |........ .......|
00000090  f9 12 76 d1 b1 01 80 c6  01 80 fe 02 72 c7 b6 00  |..v.........r...|
000000a0  80 c5 01 80 fd 0a 72 bd  f4 eb fd be c0 7c 8a 04  |......r......|..|
000000b0  83 c6 01 3c 00 74 f1 b4  0e bb 0f 00 cd 10 eb ee  |...<.t..........|
000000c0  0a 0a 6c 6f 61 64 20 65  72 72 6f 72 0a 00 00 00  |..load error....|
000000d0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
*
000001f0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 55 aa  |..............U.|
`

	testAsmSource(t, input, strings.Split(answer, "\n"))
}

// TestHelloOS3 naskソース３日目(harib00f)のテスト
func testHarib00f(t *testing.T) {
	time.Sleep(1 * time.Second)
	input1 := getAsmSource("03_day_harib00f_ipl.nas")

	// wine nask.exe ipl.nas ipl.obj
	// hexdump -C ipl.obj > ipl.hex
	// ... generate .hex file little endian
	answer1 := `00000000  eb 4e 90 48 41 52 49 42  4f 54 45 00 02 01 01 00  |.N.HARIBOTE.....|
00000010  02 e0 00 40 0b f0 09 00  12 00 02 00 00 00 00 00  |...@............|
00000020  40 0b 00 00 00 00 29 ff  ff ff ff 48 41 52 49 42  |@.....)....HARIB|
00000030  4f 54 45 4f 53 20 46 41  54 31 32 20 20 20 00 00  |OTEOS FAT12   ..|
00000040  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000050  b8 00 00 8e d0 bc 00 7c  8e d8 b8 20 08 8e c0 b5  |.......|... ....|
00000060  00 b6 00 b1 02 be 00 00  b4 02 b0 01 bb 00 00 b2  |................|
00000070  00 cd 13 73 10 83 c6 01  83 fe 05 73 2e b4 00 b2  |...s.......s....|
00000080  00 cd 13 eb e3 8c c0 05  20 00 8e c0 80 c1 01 80  |........ .......|
00000090  f9 12 76 d1 b1 01 80 c6  01 80 fe 02 72 c7 b6 00  |..v.........r...|
000000a0  80 c5 01 80 fd 0a 72 bd  e9 55 45 be c3 7c 8a 04  |......r..UE..|..|
000000b0  83 c6 01 3c 00 74 09 b4  0e bb 0f 00 cd 10 eb ee  |...<.t..........|
000000c0  f4 eb fd 0a 0a 6c 6f 61  64 20 65 72 72 6f 72 0a  |.....load error.|
000000d0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
*
000001f0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 55 aa  |..............U.|
`

	testAsmSource(t, input1, strings.Split(answer1, "\n"))

	time.Sleep(1 * time.Second)
	input2 := getAsmSource("03_day_harib00f_haribote.nas")
	answer2 := `00000000  f4 eb fd                                          |...|
`
	testAsmSource(t, input2, strings.Split(answer2, "\n"))

}

// TestHelloOS3 naskソース３日目(harib00g)のテスト
func testHarib00g(t *testing.T) {
	time.Sleep(1 * time.Second)
	input1 := getAsmSource("03_day_harib00g_ipl10.nas")

	// wine nask.exe ipl.nas ipl.obj
	// hexdump -C ipl.obj > ipl.hex
	// ... generate .hex file little endian
	answer1 := `00000000  eb 4e 90 48 41 52 49 42  4f 54 45 00 02 01 01 00  |.N.HARIBOTE.....|
00000010  02 e0 00 40 0b f0 09 00  12 00 02 00 00 00 00 00  |...@............|
00000020  40 0b 00 00 00 00 29 ff  ff ff ff 48 41 52 49 42  |@.....)....HARIB|
00000030  4f 54 45 4f 53 20 46 41  54 31 32 20 20 20 00 00  |OTEOS FAT12   ..|
00000040  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000050  b8 00 00 8e d0 bc 00 7c  8e d8 b8 20 08 8e c0 b5  |.......|... ....|
00000060  00 b6 00 b1 02 be 00 00  b4 02 b0 01 bb 00 00 b2  |................|
00000070  00 cd 13 73 10 83 c6 01  83 fe 05 73 32 b4 00 b2  |...s.......s2...|
00000080  00 cd 13 eb e3 8c c0 05  20 00 8e c0 80 c1 01 80  |........ .......|
00000090  f9 12 76 d1 b1 01 80 c6  01 80 fe 02 72 c7 b6 00  |..v.........r...|
000000a0  80 c5 01 80 fd 0a 72 bd  88 2e f0 0f e9 51 45 be  |......r......QE.|
000000b0  c7 7c 8a 04 83 c6 01 3c  00 74 09 b4 0e bb 0f 00  |.|.....<.t......|
000000c0  cd 10 eb ee f4 eb fd 0a  0a 6c 6f 61 64 20 65 72  |.........load er|
000000d0  72 6f 72 0a 00 00 00 00  00 00 00 00 00 00 00 00  |ror.............|
000000e0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
*
000001f0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 55 aa  |..............U.|
`

	testAsmSource(t, input1, strings.Split(answer1, "\n"))

	time.Sleep(1 * time.Second)
	input2 := getAsmSource("03_day_harib00g_haribote.nas")
	answer2 := `00000000  b0 13 b4 00 cd 10 f4 eb  fd                       |.........|
`
	testAsmSource(t, input2, strings.Split(answer2, "\n"))

}

// TestHelloOS3 naskソース３日目(harib00h)のテスト
func testHarib00h(t *testing.T) {
	time.Sleep(1 * time.Second)
	input1 := getAsmSource("03_day_harib00h_ipl10.nas")

	// wine nask.exe ipl.nas ipl.obj
	// hexdump -C ipl.obj > ipl.hex
	// ... generate .hex file little endian
	answer1 := `00000000  eb 4e 90 48 41 52 49 42  4f 54 45 00 02 01 01 00  |.N.HARIBOTE.....|
00000010  02 e0 00 40 0b f0 09 00  12 00 02 00 00 00 00 00  |...@............|
00000020  40 0b 00 00 00 00 29 ff  ff ff ff 48 41 52 49 42  |@.....)....HARIB|
00000030  4f 54 45 4f 53 20 46 41  54 31 32 20 20 20 00 00  |OTEOS FAT12   ..|
00000040  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000050  b8 00 00 8e d0 bc 00 7c  8e d8 b8 20 08 8e c0 b5  |.......|... ....|
00000060  00 b6 00 b1 02 be 00 00  b4 02 b0 01 bb 00 00 b2  |................|
00000070  00 cd 13 73 10 83 c6 01  83 fe 05 73 32 b4 00 b2  |...s.......s2...|
00000080  00 cd 13 eb e3 8c c0 05  20 00 8e c0 80 c1 01 80  |........ .......|
00000090  f9 12 76 d1 b1 01 80 c6  01 80 fe 02 72 c7 b6 00  |..v.........r...|
000000a0  80 c5 01 80 fd 0a 72 bd  88 2e f0 0f e9 51 45 be  |......r......QE.|
000000b0  c7 7c 8a 04 83 c6 01 3c  00 74 09 b4 0e bb 0f 00  |.|.....<.t......|
000000c0  cd 10 eb ee f4 eb fd 0a  0a 6c 6f 61 64 20 65 72  |.........load er|
000000d0  72 6f 72 0a 00 00 00 00  00 00 00 00 00 00 00 00  |ror.............|
000000e0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
*
000001f0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 55 aa  |..............U.|
`

	testAsmSource(t, input1, strings.Split(answer1, "\n"))

	time.Sleep(1 * time.Second)
	input2 := getAsmSource("03_day_harib00h_haribote.nas")
	answer2 := `00000000  3b 20 68 61 72 69 62 6f  74 65 2d 6f 73 0d 0a 3b  |; haribote-os..;|
00000010  20 54 41 42 3d 34 0d 0a  0d 0a 3b 20 42 4f 4f 54  | TAB=4....; BOOT|
00000020  5f 49 4e 46 4f e9 ab a2  e3 80 8c e8 8f ab 0a 43  |_INFO..........C|
00000030  59 4c 53 09 45 51 55 09  09 30 78 30 66 66 30 09  |YLS.EQU..0x0ff0.|
00000040  09 09 3b 20 e7 b9 9d e6  82 b6 e7 b9 9d e5 8c bb  |..; ............|
00000050  e3 81 9d e7 b9 a7 e3 83  83 e7 b9 a7 e3 82 bd e7  |................|
00000060  b8 ba e7 91 9a e3 82 a3  e3 83 a5 e8 9e b3 e5 a3  |................|
00000070  b9 e2 98 86 e7 b9 a7 0a  4c 45 44 53 09 45 51 55  |........LEDS.EQU|
00000080  09 09 30 78 30 66 66 31  0d 0a 56 4d 4f 44 45 09  |..0x0ff1..VMODE.|
00000090  45 51 55 09 09 30 78 30  66 66 32 09 09 09 3b 20  |EQU..0x0ff2...; |
000000a0  e6 bf b6 e3 82 a4 e8 ac  a8 e3 83 bc e7 b8 ba e3  |................|
000000b0  82 a9 e9 ab a2 e3 80 8c  e7 b8 ba e5 90 b6 ef bd  |................|
000000c0  8b e8 ab a0 e2 92 a6 e3  82 a2 e7 b8 b2 e3 82 86  |................|
000000d0  e3 82 b9 e8 bc 94 e3 83  b3 e7 b9 9d e3 83 a8 e7  |................|
000000e0  b9 a7 e3 82 a9 e7 b9 9d  e3 82 a5 e7 b9 9d e3 82  |................|
000000f0  b7 e7 b8 ba e5 85 b7 e3  82 b7 0a 53 43 52 4e 58  |...........SCRNX|
00000100  09 45 51 55 09 09 30 78  30 66 66 34 09 09 09 3b  |.EQU..0x0ff4...;|
00000110  20 e9 9a 97 e3 80 8d e8  9c 92 e4 b8 9e e3 82 b3  | ...............|
00000120  e3 83 b2 e7 b8 ba e3 83  a7 58 0d 0a 53 43 52 4e  |.........X..SCRN|
00000130  59 09 45 51 55 09 09 30  78 30 66 66 36 09 09 09  |Y.EQU..0x0ff6...|
00000140  3b 20 e9 9a 97 e3 80 8d  e8 9c 92 e4 b8 9e e3 82  |; ..............|
00000150  b3 e3 83 b2 e7 b8 ba e3  83 a7 59 0d 0a 56 52 41  |..........Y..VRA|
00000160  4d 09 45 51 55 09 09 30  78 30 66 66 38 09 09 09  |M.EQU..0x0ff8...|
00000170  3b 20 e7 b9 a7 e3 83 bc  e7 b9 9d e3 82 a5 e7 b9  |; ..............|
00000180  9d e8 bc 94 e3 81 85 e7  b9 9d e3 81 91 e7 b9 9d  |................|
00000190  e8 88 8c e3 83 a3 e7 b9  9d e8 bc 94 e3 81 83 e7  |................|
000001a0  b8 ba e3 83 a7 e9 ab a2  e5 8f a5 e3 82 a1 e7 8b  |................|
000001b0  97 e5 88 86 e8 9d a8 e3  83 bc 0d 0a 0d 0a 09 09  |................|
000001c0  4f 52 47 09 09 30 78 63  32 30 30 09 09 09 3b 20  |ORG..0xc200...; |
000001d0  e7 b8 ba e8 96 99 e7 b9  9d e5 8a b1 ce 9f e7 b9  |................|
000001e0  a7 e3 83 bc e7 b9 9d e3  82 a5 e7 b9 9d e7 b8 ba  |................|
000001f0  e5 be 8c e2 86 90 e7 b8  ba e8 96 99 e2 86 93 e9  |................|
00000200  9a b1 e3 83 a5 e7 b8 ba  e3 82 bd e9 9c 8e e3 82  |................|
00000210  b7 e7 b8 ba e3 82 bb e7  b9 a7 e5 be 8c ef bd 8b  |................|
00000220  e7 b8 ba e3 83 a7 e7 b8  ba 0a 0d 0a 09 09 4d 4f  |..............MO|
00000230  56 09 09 41 4c 2c 30 78  31 33 09 09 09 3b 20 56  |V..AL,0x13...; V|
00000240  47 41 e7 b9 a7 e3 83 bc  e7 b9 9d e3 82 a5 e7 b9  |GA..............|
00000250  9d e8 bc 94 e3 81 85 e7  b9 9d e3 81 91 e7 b9 a7  |................|
00000260  e3 82 b1 e7 b8 b2 32 30  78 32 30 30 78 38 62 69  |......20x200x8bi|
00000270  74 e7 b9 a7 e3 82 a9 e7  b9 9d e3 82 a5 e7 b9 9d  |t...............|
00000280  e3 82 b7 0d 0a 09 09 4d  4f 56 09 09 41 48 2c 30  |.......MOV..AH,0|
00000290  78 30 30 0d 0a 09 09 49  4e 54 09 09 30 78 31 30  |x00....INT..0x10|
000002a0  0d 0a 09 09 4d 4f 56 09  09 42 59 54 45 20 5b 56  |....MOV..BYTE [V|
000002b0  4d 4f 44 45 5d 2c 38 09  3b 20 e9 80 95 e3 82 b5  |MODE],8.; ......|
000002c0  e9 ab b1 e3 80 8c e7 b9  9d e3 80 8c e7 b9 9d e3  |................|
000002d0  82 b7 e7 b9 9d e5 b3 a8  ef bd 92 e7 b9 9d e3 80  |................|
000002e0  82 e7 b9 9d e3 80 8c e7  b8 ba e5 90 b6 ef bd 8b  |................|
000002f0  0d 0a 09 09 4d 4f 56 09  09 57 4f 52 44 20 5b 53  |....MOV..WORD [S|
00000300  43 52 4e 58 5d 2c 33 32  30 0d 0a 09 09 4d 4f 56  |CRNX],320....MOV|
00000310  09 09 57 4f 52 44 20 5b  53 43 52 4e 59 5d 2c 32  |..WORD [SCRNY],2|
00000320  30 30 0d 0a 09 09 4d 4f  56 09 09 44 57 4f 52 44  |00....MOV..DWORD|
00000330  20 5b 56 52 41 4d 5d 2c  30 78 30 30 30 61 30 30  | [VRAM],0x000a00|
00000340  30 30 0d 0a 0d 0a 3b 20  e7 b9 a7 e3 83 a5 e7 b9  |00....; ........|
00000350  9d e3 82 b7 e7 b9 9d e6  87 8a e7 b9 9d e5 b3 a8  |................|
00000360  4c 45 44 e8 bf a5 e3 82  ab e8 ab b7 e4 b9 9d ef  |LED.............|
00000370  bd 92 42 49 4f 53 e7 b8  ba e3 82 a9 e8 ac a8 e5  |..BIOS..........|
00000380  90 b6 e2 88 b4 e7 b8 ba  e3 83 b2 e7 b9 a7 e3 82  |................|
00000390  85 ef bd 89 e7 b8 ba 0a  0d 0a 09 09 4d 4f 56 09  |............MOV.|
000003a0  09 41 48 2c 30 78 30 32  0d 0a 09 09 49 4e 54 09  |.AH,0x02....INT.|
000003b0  09 30 78 31 36 20 09 09  09 3b 20 6b 65 79 62 6f  |.0x16 ...; keybo|
000003c0  61 72 64 20 42 49 4f 53  0d 0a 09 09 4d 4f 56 09  |ard BIOS....MOV.|
000003d0  09 5b 4c 45 44 53 5d 2c  41 4c 0d 0a 0d 0a 66 69  |.[LEDS],AL....fi|
000003e0  6e 3a 0d 0a 09 09 48 4c  54 0d 0a 09 09 4a 4d 50  |n:....HLT....JMP|
000003f0  09 09 66 69 6e 0d 0a                              |..fin..|
000003f7`
	testAsmSource(t, input2, strings.Split(answer2, "\n"))

}
