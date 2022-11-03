package main

/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
 *
 * https://leetcode.com/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Easy (72.83%)
 * Likes:    10051
 * Dislikes: 476
 * Total Accepted:    1.8M
 * Total Submissions: 2.5M
 * Testcase Example:  '[1,null,2,3]'
 *
 * Given the root of a binary tree, return the inorder traversal of its nodes'
 * values.
 *
 *
 * Example 1:
 *
 *
 * Input: root = [1,null,2,3]
 * Output: [1,3,2]
 *
 *
 * Example 2:
 *
 *
 * Input: root = []
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: root = [1]
 * Output: [1]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [0, 100].
 * -100 <= Node.val <= 100
 *
 *
 *
 * Follow up: Recursive solution is trivial, could you do it iteratively?
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

func inorder(node *TreeNode, lst []int) []int {
	if node == nil {
		return lst
	}
	lst = inorder(node.Left, lst)
	lst = append(lst, node.Val)
	lst = inorder(node.Right, lst)
	return lst
}

func inorderTraversal(root *TreeNode) []int {
	return inorder(root, []int{})
}

// @lc code=end
