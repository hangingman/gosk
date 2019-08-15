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

type PIMAGE_SECTION_HEADER struct {
	name                 [8]uint8
	misc                 uint32
	virtualAddress       uint32
	sizeOfRawData        uint32
	pointerToRawData     uint32
	pointerToRelocations uint32
	pointerToLinenumbers uint32
	numberOfRelocations  uint16
	numberOfLinenumbers  uint16
	characteristics      uint32
}

func evalFormat(b *object.Binary) {
	log.Println("Process for Portable Executable")
	buf := new(bytes.Buffer)

	// PE file header
	header := PIMAGE_FILE_HEADER{
		machine:              0x14c,
		numberOfSections:     0x0003,
		pointerToSymbolTable: 0x0000008e,
		numberOfSymbols:      0x00000009,
	}
	binary.Write(buf, binary.LittleEndian, &header)

	// PE section header
	text := PIMAGE_SECTION_HEADER{
		//{'.', 't', 'e', 'x', 't', 0, 0, 0 /* name */},
		name: [8]uint8{0x2e, 0x74, 0x65, 0x78, 0x74, 0x00, 0x00, 0x00},
		misc:                 0x00000000, // Misc
		virtualAddress:       0x00000000, // VirtualAddress
		sizeOfRawData:        0x00000002, // SizeOfRawData
		pointerToRawData:     0x0000008c, // PointerToRawData
		pointerToRelocations: 0x0000008e, // PointerToRelocations
		pointerToLinenumbers: 0x00000000, // PointerToLinenumbers
		numberOfRelocations:  0x0000,     // NumberOfRelocations
		numberOfLinenumbers:  0x0000,     // NumberOfLinenumbers
		characteristics:      0x60100020, // +0x38: flags, default_align = 1
	}
	data := PIMAGE_SECTION_HEADER{
		//{ '.', 'd', 'a', 't', 'a', 0, 0, 0 /* name */ },
		name: [8]uint8{0x2e, 0x64, 0x61, 0x74, 0x61, 0x00, 0x00, 0x00},
		misc:                 0x00000000, // Misc
		virtualAddress:       0x00000000, // VirtualAddress
		sizeOfRawData:        0x00000000, // SizeOfRawData
		pointerToRawData:     0x00000000, // PointerToRawData
		pointerToRelocations: 0x00000000, // PointerToRelocations
		pointerToLinenumbers: 0x00000000, // PointerToLinenumbers
		numberOfRelocations:  0x0000,     // NumberOfRelocations
		numberOfLinenumbers:  0x0000,     // NumberOfLinenumbers
		characteristics:      0xc0100040, // +0x38: flags, default_align = 1
	}
	bss := PIMAGE_SECTION_HEADER{
		//{ '.', 'b', 's', 's', 0, 0, 0, 0 /* name */ },
		name: [8]uint8{0x2e, 0x62, 0x73, 0x73, 0x00, 0x00, 0x00, 0x00},
		misc:                 0x00000000, // Misc
		virtualAddress:       0x00000000, // VirtualAddress
		sizeOfRawData:        0x00000000, // SizeOfRawData
		pointerToRawData:     0x00000000, // PointerToRawData
		pointerToRelocations: 0x00000000, // PointerToRelocations
		pointerToLinenumbers: 0x00000000, // PointerToLinenumbers
		numberOfRelocations:  0x0000,     // NumberOfRelocations
		numberOfLinenumbers:  0x0000,     // NumberOfLinenumbers
		characteristics:      0xc0100080, // +0x38: flags, default_align = 1
	}

	binary.Write(buf, binary.LittleEndian, &text)
	binary.Write(buf, binary.LittleEndian, &data)
	binary.Write(buf, binary.LittleEndian, &bss)
	b.Value = append(b.Value, buf.Bytes()...)
	log.Println("Wrote '.text', '.data', '.bss' fields for Portable Executable")
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
