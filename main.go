package main

import (
	"flag"
	"fmt"
	"path/filepath"
)

var outputFileName string
var stateMachineInfo StateMachine

func usage() {
	fmt.Println("usage: ")

}

func main() {
	flag.Parse()
	args := flag.NArg()
	if args != 1 {
		usage()
	}

	arg := flag.Arg(0)

	outputFileName = filepath.Base(arg[:len(arg)-len(filepath.Ext(arg))]) + ".xlsx"
	makeMatrix(arg)

}
