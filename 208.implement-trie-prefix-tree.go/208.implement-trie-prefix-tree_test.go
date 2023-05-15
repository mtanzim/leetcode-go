package main

import (
	"reflect"
	"testing"
)

func TestTrie_Insert(t *testing.T) {
	type fields struct {
		parent *Node
	}
	type args struct {
		words []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "basic insert",
			fields: fields{
				parent: Constructor().parent,
			},
			args: args{words: []string{"hello"}},
			want: []string{"", "h", "he", "hel", "hell", "hello"},
		},
		{
			name: "basic insert 2",
			fields: fields{
				parent: Constructor().parent,
			},
			args: args{words: []string{"cat", "rat", "met"}},
			want: []string{"", "c", "ca", "cat", "m", "me", "met" ,"r", "ra", "rat"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Trie{
				parent: tt.fields.parent,
			}
			for _, word := range tt.args.words {
				this.Insert(word)
			}
			keys := this.Keys()
			if !reflect.DeepEqual(tt.want, keys.values) {
				t.Errorf("wanted: %v, got: %v ", tt.want, keys.values)
			}
		})
	}
}
