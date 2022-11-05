package main

import "testing"

func Test_exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "single positive",
			want: true,
			args: args{
				word: "a",
				board: [][]byte{
					[]byte{'a'},
				},
			},
		},

		{
			name: "single negative",
			want: false,
			args: args{
				word: "b",
				board: [][]byte{
					[]byte{'a'},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
