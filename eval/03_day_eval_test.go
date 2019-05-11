package eval

import (
	"strings"
	"testing"
	"time"
)

func TestDumpHarib(t *testing.T) {
	input := getAsmSource("03_day_harib00b_ipl.nas")
	testAsmSourceOnlyDump(t, input, []string{""})

	// 実際のテスト
	t.Run("harib00a", testHarib00a)
	t.Run("harib00b", testHarib00b)
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
