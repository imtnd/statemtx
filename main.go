package main

import (
	"flag"
	"fmt"
)

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
	makeMatrix(arg)

}
