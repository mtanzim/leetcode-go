package main

import (
	"fmt"
	"testing"
)



func TestTrie_Insert(t *testing.T) {
	type fields struct {
		parent *Node
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "basic insert",
			fields: fields{
				parent: Constructor().parent,
			},
			args: args{word: "hello"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Trie{
				parent: tt.fields.parent,
			}
			this.Insert(tt.args.word)
			keys := this.Keys()
			fmt.Println(keys)
		})
	}
}
