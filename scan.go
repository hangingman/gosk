package main

import (
    "fmt"
)

type Scan struct {
    line       int
    lineHead   int
}

func (self *Scan) Init() {
    self.line     = 1      // 現在の行
    self.lineHead = 0      // 行の先頭文字位置
}

func (self *Scan) Err(s int) {
    fmt.Printf("\n!!error!!%d\n", s)
}


func main() {
    const s string = "line001\nint a = 10\n"  // 解析対象文字列
    
    parser := &Parser{Buffer: s}  // 解析対象文字の設定
    parser.Init()                 // parser初期化
    parser.s.Init()               // SCAN構造体初期化
    err := parser.Parse()         // 解析
    if err != nil {
        fmt.Println(err)
    } else {
        parser.Execute()          // アクション処理
    }
}
