package main

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "single element case, no carry",
			args: args{
				l1: &ListNode{2, nil},
				l2: &ListNode{3, nil},
			},
			want: &ListNode{5, nil},
		},
		{
			name: "single element case, with carry",
			args: args{
				l1: &ListNode{8, nil},
				l2: &ListNode{4, nil},
			},
			want: &ListNode{2, &ListNode{1, nil}},
		},
		{
			name: "single element case, with carry on boundary",
			args: args{
				l1: &ListNode{6, nil},
				l2: &ListNode{4, nil},
			},
			want: &ListNode{0, &ListNode{1, nil}},
		},
		{
			name: "multiple elements, lls are the same size",
			args: args{
				l1: &ListNode{2, &ListNode{4, &ListNode{3, nil}}},
				l2: &ListNode{5, &ListNode{6, &ListNode{4, nil}}},

			},
			want: &ListNode{7, &ListNode{0, &ListNode{8, nil}}},
		},
		{
			name: "multiple elements, lls are the different sizes",
			args: args{
				l1: &ListNode{3, &ListNode{2, &ListNode{1, nil}}},
				l2: &ListNode{4, nil},

			},
			want: &ListNode{7, &ListNode{2, &ListNode{1, nil}}},
		},


	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
