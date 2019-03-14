package main

import "testing"

func Test_analyzeLine(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		wants []string
	}{
		// TODO: Add test cases.
		{
			name:  "inContain retrun true",
			args:  args{"[*] -- > second"},
			wants: []string{"Start", "second"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzeLine(tt.args.s)
			for _, tw := range tt.wants {
				if !isContain(stateMachineInfo.States, tw) {
					t.Errorf("analyzeLine() result %v, want %v", stateMachineInfo.States, tw)
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
