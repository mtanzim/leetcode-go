package main

/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 *
 * https://leetcode.com/problems/implement-trie-prefix-tree/description/
 *
 * algorithms
 * Medium (60.59%)
 * Likes:    10120
 * Dislikes: 115
 * Total Accepted:    815.1K
 * Total Submissions: 1.3M
 * Testcase Example:  '["Trie","insert","search","search","startsWith","insert","search"]\n' +
  '[[],["apple"],["apple"],["app"],["app"],["app"],["app"]]'
 *
 * A trie (pronounced as "try") or prefix tree is a tree data structure used to
 * efficiently store and retrieve keys in a dataset of strings. There are
 * various applications of this data structure, such as autocomplete and
 * spellchecker.
 *
 * Implement the Trie class:
 *
 *
 * Trie() Initializes the trie object.
 * void insert(String word) Inserts the string word into the trie.
 * boolean search(String word) Returns true if the string word is in the trie
 * (i.e., was inserted before), and false otherwise.
 * boolean startsWith(String prefix) Returns true if there is a previously
 * inserted string word that has the prefix prefix, and false otherwise.
 *
 *
 *
 * Example 1:
 *
 *
 * Input
 * ["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
 * [[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
 * Output
 * [null, null, true, false, true, null, true]
 *
 * Explanation
 * Trie trie = new Trie();
 * trie.insert("apple");
 * trie.search("apple");   // return True
 * trie.search("app");     // return False
 * trie.startsWith("app"); // return True
 * trie.insert("app");
 * trie.search("app");     // return True
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= word.length, prefix.length <= 2000
 * word and prefix consist only of lowercase English letters.
 * At most 3 * 10^4 calls in total will be made to insert, search, and
 * startsWith.
 *
 *
*/

// @lc code=start
const R = 128

type Trie struct {
	parent *Node
}

type Node struct {
	next []*Node
	val  bool
}

func Constructor() Trie {
	// english alphabet
	r := R
	next := make([]*Node, r)
	return Trie{parent: &Node{next: next}}

}

func (this *Trie) Insert(word string) {
	insert(this.parent, rune(word[0]), word, 0)
}

func insert(node *Node, c rune, word string, d int) *Node {
	if node == nil {
		node = &Node{val: true, next: make([]*Node, R)}
	}
	if len(word) == d {
		return node
	}
	nextC := rune(word[d])
	node.val = true
	node.next[nextC] = insert(node.next[c], nextC, word, d+1)
	return node
}

func (this *Trie) Keys() *stringStack {
	q := &stringStack{values: []string{}}
	collect(this.parent, []rune{}, q)
	return q
}

func collect(x *Node, prefix []rune, q *stringStack) {
	if x == nil {
		return
	}
	if x.val {
		q.push(string(prefix))
		// q = append(q, string(prefix))
	}
	for i := 0; i < R; i++ {
		newPrefix := append(prefix, rune(i))
		collect(x.next[i], newPrefix, q)
	}
}

type stringStack struct {
	values []string
}

func (s *stringStack) push(v string) {
	s.values = append(s.values, v)
}

// func (this *Trie) Search(word string) bool {

// }

// func (this *Trie) StartsWith(prefix string) bool {

// }

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
