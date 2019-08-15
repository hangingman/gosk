package eval

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/hangingman/gosk/ast"
	"github.com/hangingman/gosk/object"
	"log"
)

type PIMAGE_FILE_HEADER struct {
	machine              uint16
	numberOfSections     uint16
	timeDateStamp        int32
	pointerToSymbolTable int32
	numberOfSymbols      int32
	sizeOfOptionalHeader uint16
	characteristics      uint16
}

func evalFormat(b *object.Binary) {
	buf := new(bytes.Buffer)

	header := PIMAGE_FILE_HEADER{
		machine:              0x14c,
		numberOfSections:     0x0003,
		pointerToSymbolTable: 0x00000000,
		numberOfSymbols:      0x00000000,
	}

	binary.Write(buf, binary.LittleEndian, &header)
	b.Value = append(b.Value, buf.Bytes()...)
}

func evalBits(b *object.Binary) {
}

func evalFile(b *object.Binary) {
}

func evalSettingStatement(stmt *ast.SettingStatement) object.Object {
	tok := stmt.Name.Token.Literal
	val := stmt.Name.Value

	bin := &object.Binary{Value: []byte{}}

	switch {
	case tok == "FORMAT":
		evalFormat(bin)
	case tok == "BITS":
		evalBits(bin)
	case tok == "FILE":
		evalFile(bin)
	}

	log.Println(fmt.Sprintf("info: [%s] name=%s, value=%s", stmt, tok, val))
	stmt.Bin = bin
	return stmt.Bin
}
