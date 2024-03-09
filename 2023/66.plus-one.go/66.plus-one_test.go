package main

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "basic1",
			args: args{[]int{4, 3, 2, 1}},
			want: []int{4, 3, 2, 2},
		},
		{
			name: "with carry over",
			args: args{[]int{4, 3, 2, 9}},
			want: []int{4, 3, 3, 0},
		},
		{
			name: "carry over to the beginning",
			args: args{[]int{9}},
			want: []int{1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
