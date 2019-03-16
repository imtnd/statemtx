package main

import "testing"

func Test_analyzeLine(t *testing.T) {
	type fields struct {
		States []string
		Events []string
	}
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		wants fields
	}{
		// TODO: Add test cases.
		{
			name:  "analyzeLine analyze state",
			args:  args{"[*] -> second"},
			wants: fields{States: []string{"Start", "second"}, Events: nil},
		},
		{
			name:  "analyzeLine analyze header",
			args:  args{"@startuml{test.png}"},
			wants: fields{States: []string{"Start", "second"}, Events: nil},
		},
		{
			name:  "analyzeLine analyze state and event",
			args:  args{"second -> [*] : event1"},
			wants: fields{States: []string{"Start", "second"}, Events: []string{"event1"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzeLine(tt.args.s)
			for _, ts := range tt.wants.States {
				if !isContain(stateMachineInfo.States, ts) {
					t.Errorf("analyzeLine() result %v, want %v", stateMachineInfo.States, ts)
				}
			}
			for _, te := range tt.wants.Events {
				if !isContain(stateMachineInfo.Events, te) {
					t.Errorf("analyzeLine() result %v, want %v", stateMachineInfo.Events, te)
				}
			}
		})
	}
}

func Test_readUMLFile(t *testing.T) {
	type args struct {
		fname string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			readUMLFile(tt.args.fname)
		})
	}
}

func Test_makeMatrix(t *testing.T) {
	type args struct {
		fname string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			makeMatrix(tt.args.fname)
		})
	}
}
