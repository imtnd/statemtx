package main

// StateMachine is state machine information
type StateMachine struct {
	States []string
	Events []string
}

func isContain(s []string, key string) bool {
	for _, v := range s {
		if v == key {
			return true
		}
	}
	return false
}

func (sm *StateMachine) addState(state string) []string {
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

func (sm *StateMachine) addEvent(event string) []string {
	if isContain(sm.Events, event) != true {
		sm.Events = append(sm.Events, event)
	}
	return sm.Events
}
