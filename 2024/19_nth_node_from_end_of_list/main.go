package main

/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 *
 * https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (39.83%)
 * Likes:    13793
 * Dislikes: 570
 * Total Accepted:    1.8M
 * Total Submissions: 4.5M
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given the head of a linked list, remove the n^th node from the end of the
 * list and return its head.
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,2,3,4,5], n = 2
 * Output: [1,2,3,5]
 *
 *
 * Example 2:
 *
 *
 * Input: head = [1], n = 1
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: head = [1,2], n = 1
 * Output: [1]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is sz.
 * 1 <= sz <= 30
 * 0 <= Node.val <= 100
 * 1 <= n <= sz
 *
 *
 *
 * Follow up: Could you do this in one pass?
 *
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l, r := head, head

	for i := 1; i < n && r != nil; i++ {
		r = r.Next
	}
	if r.Next == nil {
		head = l.Next
		return head
	}
	for r.Next.Next != nil {
		r = r.Next
		l = l.Next
	}
	l.Next = l.Next.Next
	return head

}
