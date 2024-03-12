package main

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "given",
			args: args{
				n:    2,
				head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}},
		},
		{
			name: "single elem",
			args: args{
				n:    1,
				head: &ListNode{Val: 1},
			},
			want: nil,
		},

		{
			name: "double elem",
			args: args{
				n:    1,
				head: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			},
			want: &ListNode{Val: 1},
		},

		{
			name: "double elem remove 1",
			args: args{
				n:    2,
				head: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			},
			want: &ListNode{Val: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
