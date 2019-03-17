package main

import (
	"testing"

	"github.com/tealeg/xlsx"
)

func Test_makeXlsx(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			makeXlsx()
		})
	}
}

func Test_addStateNameLine(t *testing.T) {
	type args struct {
		s *xlsx.Sheet
	}
	tests := []struct {
		name  string
		args  args
		pre   []string
		wants []string
	}{
		// TODO: Add test cases.
		{
			name:  "addStateNameLine test add Start and second",
			args:  args{s: new(xlsx.Sheet)},
			pre:   []string{"Start", "second"},
			wants: []string{"Event \\ State", "0:Start", "1:second"},
		},
	}
	for _, tt := range tests {
		stateMachineInfo = StateMachine{}
		for _, p := range tt.pre {
			stateMachineInfo.addState(p)
		}
		t.Run(tt.name, func(t *testing.T) {
			result := addStateNameLine(tt.args.s)
			for i, want := range tt.wants {
				cell := result.Cell(0, i)
				if cell.Value != want {
					t.Errorf("%v result %v, want %v", tt.name, cell.Value, want)
				}
			}
		})
	}
}

func Test_addEventNameLine(t *testing.T) {
	type args struct {
		s *xlsx.Sheet
	}
	tests := []struct {
		name  string
		args  args
		pre   []string
		wants []string
	}{
		// TODO: Add test cases.
		{
			name:  "addEventNameLine test add event1",
			args:  args{s: new(xlsx.Sheet)},
			pre:   []string{"event1"},
			wants: []string{"event1"},
		},
	}
	for _, tt := range tests {
		stateMachineInfo = StateMachine{}
		for _, p := range tt.pre {
			stateMachineInfo.addEvent(p)
		}
		t.Run(tt.name, func(t *testing.T) {
			result := addEventNameLine(tt.args.s)
			for i, want := range tt.wants {
				cell := result.Cell(i+1, 0)
				if cell.Value != want {
					t.Errorf("%v result %v, want %v", tt.name, cell.Value, want)
				}
			}
		})
	}
}
