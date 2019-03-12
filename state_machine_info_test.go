package main

import (
	"reflect"
	"testing"
)

func Test_isContain(t *testing.T) {
	type args struct {
		s   []string
		key string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "inContain retrun true",
			args: args{[]string{"hoge", "fuga"}, "hoge"},
			want: true,
		},
		{
			name: "inContain retrun false",
			args: args{[]string{"hoge", "huga"}, "piyo"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isContain(tt.args.s, tt.args.key); got != tt.want {
				t.Errorf("isContain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStateMachine_addState(t *testing.T) {
	type fields struct {
		States []string
		Events []string
	}
	type args struct {
		state string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		// TODO: Add test cases.
		{
			name:   "add Start",
			fields: fields{nil, nil},
			args:   args{"[*]"},
			want:   []string{"Start"},
		},
		{
			name:   "add state",
			fields: fields{[]string{"Start"}, []string{""}},
			args:   args{"hoge"},
			want:   []string{"Start", "hoge"},
		},
		{
			name:   "add End",
			fields: fields{[]string{"Start", "hoge"}, []string{""}},
			args:   args{"[*]"},
			want:   []string{"Start", "hoge", "End"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &StateMachine{
				States: tt.fields.States,
				Events: tt.fields.Events,
			}
			if got := sm.addState(tt.args.state); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StateMachine.addState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStateMachine_addEvent(t *testing.T) {
	type fields struct {
		States []string
		Events []string
	}
	type args struct {
		event string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &StateMachine{
				States: tt.fields.States,
				Events: tt.fields.Events,
			}
			if got := sm.addEvent(tt.args.event); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StateMachine.addEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}
