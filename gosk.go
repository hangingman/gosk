package main

import (
	"bufio"
	"fmt"
	"os"
)

type Scan struct {
	line     int
	lineHead int
}

func (self *Scan) Init() {
	self.line = 1     // 現在の行
	self.lineHead = 0 // 行の先頭文字位置
}

func (self *Scan) Err(s int) {
	fmt.Printf("\n!!error!!%d\n", s)

}

func main() {
        var fp *os.File
        var err error

        if len(os.Args) < 2 {
                fp  = os.Stdin
        } else {
                fp, err = os.Open(os.Args[1])
                if err != nil {
			panic(err)
                }
                defer fp.Close()
        }

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		// 解析対象文字の設定
		parser := &Parser{Buffer: scanner.Text()}
		parser.Init()                // parser初期化
		parser.s.Init()              // SCAN構造体初期化
		err := parser.Parse()        // 解析

		if err != nil {
			fmt.Println(err)
		} else {
			parser.Execute() // アクション処理
		}
	}
}
