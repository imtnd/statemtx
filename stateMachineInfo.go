package main

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
	if state == "[*]" {
		if isContain(sm.States, "Start") != true {
			sm.addState("Start")
		} else {
			sm.addState("End")
		}
	} else if isContain(sm.States, state) != true {
		sm.States = append(sm.States, state)
	}
	return sm.States
}

func (sm *StateMachine) addEvent(event Event) []Event {
	if isContain(sm.Events, event) == true {
		return sm.Events
	}
	sm.Events = append(sm.Events, event)
	return sm.Events
}
