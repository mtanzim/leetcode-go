package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "basic",
			args: args{121},
			want: true,
		},
		{
			name: "negative",
			args: args{-121},
			want: false,
		},
		{
			name: "short",
			args: args{12},
			want: false,
		},
		{
			name: "not negative",
			args: args{123},
			want: false,
		},
		{
			name: "shorter",
			args: args{1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
