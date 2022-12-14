package main

import "testing"

func Test_numIslands(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			want: 1,
			name: "basic",
			args: args{
				grid: [][]byte{
					[]byte{'1', '1', '1', '1', '0'},
					[]byte{'1', '1', '0', '1', '0'},
					[]byte{'1', '1', '0', '0', '0'},
					[]byte{'0', '0', '0', '0', '0'},
				},
			},
		},

		{
			want: 3,
			name: "basic2",
			args: args{
				grid: [][]byte{
					[]byte{'1', '1', '0', '0', '0'},
					[]byte{'1', '1', '0', '0', '0'},
					[]byte{'0', '0', '1', '0', '0'},
					[]byte{'0', '0', '0', '1', '1'},
				},
			},
		},

		{
			want: 1,
			name: "failing on leetcode",
			args: args{
				grid: [][]byte{
					[]byte{'1', '1', '1'},
					[]byte{'0', '1', '0'},
					[]byte{'1', '1', '1'},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands(tt.args.grid); got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
