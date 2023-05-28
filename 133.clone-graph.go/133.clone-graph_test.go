package main

import (
	"reflect"
	"testing"
)

func Test_cloneGraph(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "single node",
			args: args{
				node: &Node{1, []*Node{}},
			},
			want: &Node{1, []*Node{}},
		},
		{
			name: "single node with neighbors",
			args: args{
				node: &Node{1,
					[]*Node{
						&Node{2, []*Node{&Node{1, []*Node{}}}},
						&Node{4, []*Node{}},
					},
				},
			},
			want: &Node{1, []*Node{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cloneGraph(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cloneGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}
