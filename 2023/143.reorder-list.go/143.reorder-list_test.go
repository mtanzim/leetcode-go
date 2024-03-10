package main

import "testing"

func Test_reorderList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple",
			args: args{
				head: &ListNode{Val: 1},
			},
			want: "head -> 1 ",
		},
		{
			name: "first",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
			},
			want: "head -> 1 -> 4 -> 2 -> 3 ",
		},
		{
			name: "first",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			},
			want: "head -> 1 -> 5 -> 2 -> 4 -> 3 ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(tt.args.head)
			got := tt.args.head.String()
			want := tt.want
			if got != want {
				t.Errorf("head.String() = %v, want %v", got, want)
			}
		})
	}
}
