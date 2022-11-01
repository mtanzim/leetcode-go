package main

/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 *
 * https://leetcode.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (63.08%)
 * Likes:    11368
 * Dislikes: 214
 * Total Accepted:    1.6M
 * Total Submissions: 2.5M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given the root of a binary tree, return the level order traversal of its
 * nodes' values. (i.e., from left to right, level by level).
 *
 *
 * Example 1:
 *
 *
 * Input: root = [3,9,20,null,null,15,7]
 * Output: [[3],[9,20],[15,7]]
 *
 *
 * Example 2:
 *
 *
 * Input: root = [1]
 * Output: [[1]]
 *
 *
 * Example 3:
 *
 *
 * Input: root = []
 * Output: []
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [0, 2000].
 * -1000 <= Node.val <= 1000
 *
 *
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
func traverse(node *TreeNode, curMap map[int][]int, curDepth int) map[int][]int {
	if node == nil {
		return curMap
	}
	if curArr, ok := curMap[curDepth]; ok {
		curMap[curDepth] = append(curArr, node.Val)
	} else {
		curMap[curDepth] = []int{node.Val}
	}
	traverse(node.Left, curMap, curDepth+1)
	traverse(node.Right, curMap, curDepth+1)
	return curMap

}
func levelOrder(root *TreeNode) [][]int {
	hm := traverse(root, make(map[int][]int), 0)
	res := make([][]int, len(hm))
	for i := 0; i < len(hm); i++ {
		res[i] = hm[i]
	}
	return res
}

// @lc code=end
