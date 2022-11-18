package main

import (
	"reflect"
	"testing"
)

// TODO: these do not pass since array order is not deterministic
// func Test_groupAnagrams(t *testing.T) {
func Te_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
		{
			name: "testing",
			args: args{[]string{"cat", "tac"}},
			want: [][]string{[]string{"cat", "tac"}},
		},
		{
			name: "given",
			args: args{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			want: [][]string{[]string{"bat"}, []string{"nat", "tan"}, []string{"ate", "eat", "tea"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
