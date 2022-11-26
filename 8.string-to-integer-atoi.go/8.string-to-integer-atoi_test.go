package main

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "simple",
			args: args{"456"},
			want: 456,
		},
		{
			name: "simple negative",
			args: args{"-456"},
			want: -456,
		},
		{
			name: "whitespace to start",
			args: args{"    -456"},
			want: -456,
		},
		{
			name: "leading zeroes negative",
			args: args{"    -00000456"},
			want: -456,
		},
		{
			name: "leading zeroes positive",
			args: args{"00000456"},
			want: 456,
		},
		{
			name: "discards stuff at the end",
			args: args{"00000456 hello jibber jabber"},
			want: 456,
		},
		{
			name: "failing on leetocde: starts with chars",
			args: args{"words and 987"},
			want: 0,
		},
		{
			name: "failing on leetocde: hits upper bound of 32 bits",
			args: args{"-91283472332"},
			want: -2147483648,
		},
		{
			name: "failing on leetocde: incorrect handling of signs",
			args: args{"+-12"},
			want: 0,
		},
		{
			name: "failing on leetocde: trailing zeroes treated incorrectly",
			args: args{"21474836460"},
			want: 2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
