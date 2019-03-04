package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type State = string
type Event = string

type StateMachine struct {
	States []State
	Events []Event
}

var stateMachineInfo StateMachine

func isContain(s []string, key string) bool {
	for _, v := range s {
		if v == key {
			return true
		}
	}
	return false
}

func (sm *StateMachine) addState(state State) []State {
	if isContain(sm.States, state) == true {
		return sm.States
	}
	sm.States = append(sm.States, state)
	return sm.States
}

func (sm *StateMachine) addEvent(event Event) []Event {
	if isContain(sm.Events, event) == true {
		return sm.Events
	}
	sm.Events = append(sm.Events, event)
	return sm.Events
}
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
	fmt.Println(stateMachineInfo.States)
	fmt.Println(stateMachineInfo.Events)
}

func makeMatrix(fname string) {
	fmt.Println("hoge")
	readUMLFile(fname)

}
