package main

import "sort"

/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 *
 * https://leetcode.com/problems/validate-binary-search-tree/description/
 *
 * algorithms
 * Medium (31.66%)
 * Likes:    14539
 * Dislikes: 1182
 * Total Accepted:    1.9M
 * Total Submissions: 6M
 * Testcase Example:  '[2,1,3]'
 *
 * Given the root of a binary tree, determine if it is a valid binary search
 * tree (BST).
 *
 * A valid BST is defined as follows:
 *
 *
 * The left subtree of a node contains only nodes with keys less than the
 * node's key.
 * The right subtree of a node contains only nodes with keys greater than the
 * node's key.
 * Both the left and right subtrees must also be binary search trees.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: root = [2,1,3]
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: root = [5,1,4,null,null,3,6]
 * Output: false
 * Explanation: The root node's value is 5 but its right child's value is
 * 4.
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [1, 10^4].
 * -2^31 <= Node.val <= 2^31 - 1
 *
 *
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
type stack struct {
	values []int
}

func (s *stack) push(v int) {
	s.values = append(s.values, v)
}

func (s *stack) at(i int) int {
	return s.values[i]
}

func traverse(root *TreeNode, s *stack) {
	if root == nil {
		return
	}
	traverse(root.Left, s)
	s.push(root.Val)
	traverse(root.Right, s)
}

func hasDups(vs []int) bool {
	set := make(map[int]bool)
	for _, v := range vs {
		if _, ok := set[v]; ok {
			return true
		}
		set[v] = true
	}
	return false
}

func isValidBST(root *TreeNode) bool {
	empty := []int{}
	s := &stack{empty}
	traverse(root, s)
	return sort.IntsAreSorted(s.values) && !hasDups(s.values)
}

// @lc code=end
