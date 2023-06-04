package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		k    int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testing",
			args: args{
				k:    3,
				nums: []int{4, 5, 8, 2},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kthLargest := Constructor(tt.args.k, tt.args.nums)
			got := kthLargest.Add(3)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}
