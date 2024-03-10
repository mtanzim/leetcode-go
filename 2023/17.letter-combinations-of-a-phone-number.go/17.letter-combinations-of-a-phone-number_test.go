package main

import (
	"reflect"
	"testing"
)

func Test_traverse(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "base case 0",
			args: args{""},
			want: []string{},
		},
		{
			name: "base case 1",
			args: args{"2"},
			want: []string{"a", "b", "c"},
		},
		{
			name: "base case 2 digits",
			args: args{"23"},
			want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name: "an actual test",
			args: args{"234"},
			want: []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCombinations(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("traverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
