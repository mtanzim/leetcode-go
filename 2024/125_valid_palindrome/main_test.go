package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "basic",
			args: args{
				s: "cat",
			},
			want: false,
		},
		{
			name: "basic positive",
			args: args{
				s: "mom",
			},
			want: true,
		},
		{
			name: "given",
			args: args{"A man, a plan, a canal: Panama"},
			want: true,
		},
		{
			name: "given 2",
			args: args{"race a car"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
