package main

/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
 *
 * https://leetcode.com/problems/maximum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (72.99%)
 * Likes:    8827
 * Dislikes: 143
 * Total Accepted:    2M
 * Total Submissions: 2.7M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given the root of a binary tree, return its maximum depth.
 *
 * A binary tree's maximum depth is the number of nodes along the longest path
 * from the root node down to the farthest leaf node.
 *
 *
 * Example 1:
 *
 *
 * Input: root = [3,9,20,null,null,15,7]
 * Output: 3
 *
 *
 * Example 2:
 *
 *
 * Input: root = [1,null,2]
 * Output: 2
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [0, 10^4].
 * -100 <= Node.val <= 100
 *
 *
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
func traverse(node *TreeNode, curHeight int) int {
	if node == nil {
		return curHeight
	}

	var lh, rh int
	if node.Left != nil {
		lh = traverse(node.Left, curHeight+1)
	}
	if node.Right != nil {
		rh = traverse(node.Right, curHeight+1)
	}
	if node.Left == nil && node.Right == nil {
		return curHeight + 1
	}
	if rh > lh {
		return rh
	}
	return lh
}

func maxDepth(root *TreeNode) int {
	return traverse(root, 0)
}

// @lc code=end
