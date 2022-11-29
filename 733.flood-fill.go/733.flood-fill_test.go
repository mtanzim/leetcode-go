package main

import (
	"reflect"
	"testing"
)

func Test_floodFill(t *testing.T) {
	type args struct {
		image [][]int
		sr    int
		sc    int
		color int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "given test case",
			args: args{
				image: [][]int{[]int{1,1,1}, []int{1,1,0},[]int{1,0,1}},
				sr: 1,
				sc: 1,
				color: 2,
			},
			want: [][]int{[]int{2,2,2}, []int{2,2,0},[]int{2,0,1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := floodFill(tt.args.image, tt.args.sr, tt.args.sc, tt.args.color); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("floodFill() = %v, want %v", got, tt.want)
			}
		})
	}
}
