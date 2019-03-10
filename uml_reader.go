package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func analyzeLine(s string) {
	if !strings.Contains(s, ">") {
		return
	}
	state1 := strings.Split(s, "-")[0]
	state1 = strings.TrimSpace(state1)
	stateMachineInfo.addState(state1)

	state2 := strings.Split(s, ">")[1]
	state2 = strings.Split(state2, ":")[0]
	state2 = strings.TrimSpace(state2)
	stateMachineInfo.addState(state2)

	if !strings.Contains(s, ":") {
		return
	}
	event := strings.SplitN(s, ":", 2)[1]
	event = strings.TrimSpace(event)
	stateMachineInfo.addEvent(event)
}

func readUMLFile(fname string) {
	f, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		analyzeLine(scanner.Text())
	}
}

func makeMatrix(fname string) {
	readUMLFile(fname)
	makeXlsx()
}
