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
			want: []string{"", "c", "ca", "cat", "m", "me", "met", "r", "ra", "rat"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor()
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

func TestTrie_Search(t *testing.T) {
	type fields struct {
		parent *Node
	}
	type args struct {
		words      []string
		searchWord string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "basic search",
			fields: fields{
				parent: Constructor().parent,
			},
			args: args{words: []string{"hello"}, searchWord: "hello"},
			want: true,
		},
		{
			name: "search with prefix",
			fields: fields{
				parent: Constructor().parent,
			},
			args: args{words: []string{"hello"}, searchWord: "hell"},
			want: false,
		},
		{
			name: "basic insert 2",
			fields: fields{
				parent: Constructor().parent,
			},
			args: args{words: []string{"cat", "rat", "met"}, searchWord: "fat"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor()
			for _, word := range tt.args.words {
				this.Insert(word)
			}
			got := this.Search(tt.args.searchWord)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("wanted: %v, got: %v ", tt.want, got)
			}
		})
	}
}

func TestTrie_StartsWith(t *testing.T) {
	type fields struct {
		parent *Node
	}
	type args struct {
		words  []string
		prefix string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "basic insert",
			fields: fields{
				parent: Constructor().parent,
			},
			args: args{words: []string{"hello"}, prefix: "hello"},
			want: true,
		},
		{
			name: "basic prefix",
			fields: fields{
				parent: Constructor().parent,
			},
			args: args{words: []string{"hello"}, prefix: "hel"},
			want: true,
		},
		{
			name: "basic prefix single word",
			fields: fields{
				parent: Constructor().parent,
			},
			args: args{words: []string{"hello"}, prefix: "h"},
			want: true,
		},
		{
			name: "basic prefix empty string",
			fields: fields{
				parent: Constructor().parent,
			},
			args: args{words: []string{"hello"}, prefix: ""},
			want: true,
		},
		{
			name: "basic prefix false case",
			fields: fields{
				parent: Constructor().parent,
			},
			args: args{words: []string{"hello"}, prefix: "j"},
			want: false,
		},
		{
			name: "basic insert 2",
			fields: fields{
				parent: Constructor().parent,
			},
			args: args{words: []string{"cat", "rat", "met"}, prefix: "fat"},
			want: false,
		},
		{
			name: "no words in trie",
			fields: fields{
				parent: Constructor().parent,
			},
			args: args{words: []string{}, prefix: "fat"},
			want: false,
		},
		{
			name: "no words in trie",
			fields: fields{
				parent: Constructor().parent,
			},
			args: args{words: []string{"fa"}, prefix: "fat"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor()
			for _, word := range tt.args.words {
				this.Insert(word)
			}
			got := this.StartsWith(tt.args.prefix)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("wanted: %v, got: %v ", tt.want, got)
			}
		})
	}
}
