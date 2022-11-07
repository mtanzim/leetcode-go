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
		{
			name: "double positive",
			want: true,
			args: args{
				word: "ab",
				board: [][]byte{
					[]byte{'a', 'b', 'c', 'e'},
					[]byte{'s', 'f', 'c', 's'},
				},
			},
		},
		{
			name: "double negative",
			want: false,
			args: args{
				word: "ae",
				board: [][]byte{
					[]byte{'a', 'b', 'c', 'e'},
					[]byte{'s', 'f', 'c', 's'},
				},
			},
		},

		{
			name: "failing on leetcode",
			want: true,
			args: args{
				word: "see",
				board: [][]byte{
					[]byte{'a', 'b', 'c', 'e'},
					[]byte{'s', 'f', 'c', 's'},
					[]byte{'a', 'd', 'e', 'e'},
				},
			},
		},

		{
			name: "failing on leetcode again",
			want: true,
			args: args{
				word: "aab",
				board: [][]byte{
					[]byte{'c', 'a', 'a'},
					[]byte{'a', 'a', 'a'},
					[]byte{'b', 'c', 'd'},
				},
			},
		},
		{
			name: "failing on leetcode again again!!",
			want: true,
			args: args{
				word: "acdb",
				board: [][]byte{
					[]byte{'a', 'b'},
					[]byte{'c', 'd'},
				},
			},
		},
		{
			name: "getting closer",
			want: true,
			args: args{
				word: "abceseeefs",
				board: [][]byte{
					[]byte{'a', 'b', 'c', 'e'},
					[]byte{'s', 'f', 'e', 's'},
					[]byte{'a', 'd', 'e', 'e'},
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
