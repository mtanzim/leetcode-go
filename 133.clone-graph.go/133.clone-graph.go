package main

import "fmt"

/*
 * @lc app=leetcode id=133 lang=golang
 *
 * [133] Clone Graph
 *
 * https://leetcode.com/problems/clone-graph/description/
 *
 * algorithms
 * Medium (50.35%)
 * Likes:    8274
 * Dislikes: 3319
 * Total Accepted:    986.4K
 * Total Submissions: 1.8M
 * Testcase Example:  '[[2,4],[1,3],[2,4],[1,3]]'
 *
 * Given a reference of a node in a connected undirected graph.
 *
 * Return a deep copy (clone) of the graph.
 *
 * Each node in the graph contains a value (int) and a list (List[Node]) of its
 * neighbors.
 *
 *
 * class Node {
 * ⁠   public int val;
 * ⁠   public List<Node> neighbors;
 * }
 *
 *
 *
 *
 * Test case format:
 *
 * For simplicity, each node's value is the same as the node's index
 * (1-indexed). For example, the first node with val == 1, the second node with
 * val == 2, and so on. The graph is represented in the test case using an
 * adjacency list.
 *
 * An adjacency list is a collection of unordered lists used to represent a
 * finite graph. Each list describes the set of neighbors of a node in the
 * graph.
 *
 * The given node will always be the first node with val = 1. You must return
 * the copy of the given node as a reference to the cloned graph.
 *
 *
 * Example 1:
 *
 *
 * Input: adjList = [[2,4],[1,3],[2,4],[1,3]]
 * Output: [[2,4],[1,3],[2,4],[1,3]]
 * Explanation: There are 4 nodes in the graph.
 * 1st node (val = 1)'s neighbors are 2nd node (val = 2) and 4th node (val =
 * 4).
 * 2nd node (val = 2)'s neighbors are 1st node (val = 1) and 3rd node (val =
 * 3).
 * 3rd node (val = 3)'s neighbors are 2nd node (val = 2) and 4th node (val =
 * 4).
 * 4th node (val = 4)'s neighbors are 1st node (val = 1) and 3rd node (val =
 * 3).
 *
 *
 * Example 2:
 *
 *
 * Input: adjList = [[]]
 * Output: [[]]
 * Explanation: Note that the input contains one empty list. The graph consists
 * of only one node with val = 1 and it does not have any neighbors.
 *
 *
 * Example 3:
 *
 *
 * Input: adjList = []
 * Output: []
 * Explanation: This an empty graph, it does not have any nodes.
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the graph is in the range [0, 100].
 * 1 <= Node.val <= 100
 * Node.val is unique for each node.
 * There are no repeated edges and no self-loops in the graph.
 * The Graph is connected and all nodes can be visited starting from the given
 * node.
 *
 *
 */

// Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

// func (n Node) String() string {
// 	var sb strings.Builder
// 	sb.WriteString(fmt.Sprintf("value: %d", n.Val))
// 	for _, v := range n.Neighbors {
// 		sb.WriteString(fmt.Sprintf("neighbor: %d", v.Val))
// 	}
// 	return sb.String()
// }

// @lc code=start

func dfs(node *Node, marked map[int]bool) *Node {
	if node == nil {
		return node
	}
	marked[node.Val] = true
	newNeighbors := []*Node{}

	// fmt.Println(fmt.Sprintf("node: %d", node.Val))
	for _, n := range node.Neighbors {
		fmt.Println(fmt.Sprintf("neighbor:%d, from:%d", n.Val, node.Val))
		if !marked[n.Val] {
			nn := dfs(n, marked)
			newNeighbors = append(newNeighbors, nn)
		}
	}
	newNode := &Node{
		Val:       node.Val,
		Neighbors: newNeighbors,
	}
	return newNode
}

func cloneGraph(node *Node) *Node {
	return dfs(node, make(map[int]bool))
}

// @lc code=end
