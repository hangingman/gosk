package eval

import (
	"encoding/hex"
	"fmt"
	"strings"
)

const emptyLine = "00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|"

func isZeroFillLine(hexLine string) bool {
	return strings.HasSuffix(hexLine, emptyLine)
}

// getHexdumpFmtArray バイナリを引数にとって
// `hexdump -C` 形式で出力される文字列を返す
func getHexdumpFmtString(binary []byte) string {
	hexDumpArr := getHexdumpFmtArray(binary)
	return strings.Join(hexDumpArr, "\n")
}

// getHexdumpFmtArray バイナリを引数にとって
// `hexdump -C` 形式で出力される文字列を配列で返す
//
// 以下サンプル
// 00000000  eb 4e 90 48 45 4c 4c 4f  49 50 4c 00 02 01 01 00  |.N.HELLOIPL.....|
// 00000010  02 e0 00 40 0b f0 09 00  12 00 02 00 00 00 00 00  |...@............|
// 00000020  40 0b 00 00 00 00 29 ff  ff ff ff 48 45 4c 4c 4f  |@.....)....HELLO|
func getHexdumpFmtArray(binary []byte) []string {
	binSize := len(binary)
	dumpLines := strings.Split(hex.Dump(binary), "\n")
	hexdumpFmtArr := []string{}

	for i := 0; i < len(dumpLines); i++ {
		line := dumpLines[i]
		peek := i + 1

		if isZeroFillLine(line) && isZeroFillLine(dumpLines[peek]) {
			// 0x00 だけで埋まった行が連続する場合それをskipし '*' で表す
			hexdumpFmtArr = append(hexdumpFmtArr, line)

			for {
				peek++

				if peek == (len(dumpLines) - 1) {
					// 0x00 が最後まで続いている場合はファイルサイズを取得して末尾につける
					hexdumpFmtArr = append(hexdumpFmtArr, "*")
					hexdumpFmtArr = append(hexdumpFmtArr, fmt.Sprintf("%08x", binSize))
					i = peek
					break
				}
				if !isZeroFillLine(dumpLines[peek]) {
					// 0x00 をskip
					hexdumpFmtArr = append(hexdumpFmtArr, "*")
					hexdumpFmtArr = append(hexdumpFmtArr, dumpLines[peek])
					i = peek
					break
				}
			}
		} else {
			// 通常どおり16進の文字列を追加
			hexdumpFmtArr = append(hexdumpFmtArr, line)
		}
	}

	return hexdumpFmtArr
}
