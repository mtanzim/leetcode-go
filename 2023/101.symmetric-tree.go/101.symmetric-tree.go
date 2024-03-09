package main

import (
	"fmt"
)

/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
 *
 * https://leetcode.com/problems/symmetric-tree/description/
 *
 * algorithms
 * Easy (52.77%)
 * Likes:    11432
 * Dislikes: 260
 * Total Accepted:    1.4M
 * Total Submissions: 2.7M
 * Testcase Example:  '[1,2,2,3,4,4,3]'
 *
 * Given the root of a binary tree, check whether it is a mirror of itself
 * (i.e., symmetric around its center).
 *
 *
 * Example 1:
 *
 *
 * Input: root = [1,2,2,3,4,4,3]
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: root = [1,2,2,null,3,null,3]
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [1, 1000].
 * -100 <= Node.val <= 100
 *
 *
 *
 * Follow up: Could you solve it both recursively and iteratively?
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start

func traverse(node *TreeNode, hm map[string]int, curDepth int, prefix rune) (map[string]int, int) {
	key := fmt.Sprintf("%d%c", curDepth, prefix)
	if node == nil {
		hm[key] = -101
		return hm, curDepth
	}

	hm[key] = node.Val

	_, ld := traverse(node.Left, hm, curDepth+1, 'l')
	_, rd := traverse(node.Right, hm, curDepth+1, 'r')
	var maxDepth int
	if ld > rd {
		maxDepth = ld
	} else {
		maxDepth = rd
	}
	return hm, maxDepth
}
func isSymmetricOld(root *TreeNode) bool {

	lh, lDepth := traverse(root.Left, make(map[string]int), 1, 'l')
	rh, rDepth := traverse(root.Right, make(map[string]int), 1, 'r')

	if len(lh) != len(rh) {
		return false
	}

	if lDepth != rDepth {
		return false
	}

	isSymmetric := true
	for i := 1; i <= lDepth; i++ {

		keyL := fmt.Sprintf("%d%c", i, 'l')
		keyR := fmt.Sprintf("%d%c", i, 'r')
		fmt.Printf("depth: %d \n", i)
		fmt.Printf("lh left: %v \n", lh[keyL])
		fmt.Printf("lh right: %v \n", lh[keyR])
		fmt.Printf("rh left: %v \n", rh[keyL])
		fmt.Printf("rh right: %v \n", rh[keyR])
		if lh[keyL] == rh[keyR] && lh[keyR] == rh[keyL] {
			continue
		}

		return false
	}

	return isSymmetric

}
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return checker(root.Left, root.Right)
}

func checker(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}

	return checker(left.Left, right.Right) && checker(left.Right, right.Left)
}

// @lc code=end
