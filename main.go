package main

import (
	"flag"
	"fmt"
	"github.com/hangingman/gosk/repl"
	"os"
	"os/user"
)

const Version = "1.0.0 beta"

func main() {
	var (
		version = flag.Bool("v", false, "バージョンとライセンス情報を表示する")
	)
	// -hオプション用文言
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage:  [--help | -v] source [object/binary] [list]\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *version {
		fmt.Printf("gosk %s\n", Version)
		fmt.Printf("%s", `Copyright (C) 2019 Hiroyuki Nagata.
ライセンス GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Thank you osask project !`)
		os.Exit(0)
	}

	// 引数が無ければREPLモードへ移行
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is yet another assembly gosk!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
