package main

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "testing",
			args: args{"babad"},
			want: "bab",
		},
		{
			name: "testing",
			args: args{"cbbd"},
			want: "bb",
		},
		{
			name: "fails because I only get the latest instaces of chars",
			args: args{"ccc"},
			want: "ccc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
