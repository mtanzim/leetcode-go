package main

/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 *
 * https://leetcode.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (61.69%)
 * Likes:    16020
 * Dislikes: 1403
 * Total Accepted:    2.8M
 * Total Submissions: 4.5M
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * You are given the heads of two sorted linked lists list1 and list2.
 *
 * Merge the two lists in a one sorted list. The list should be made by
 * splicing together the nodes of the first two lists.
 *
 * Return the head of the merged linked list.
 *
 *
 * Example 1:
 *
 *
 * Input: list1 = [1,2,4], list2 = [1,3,4]
 * Output: [1,1,2,3,4,4]
 *
 *
 * Example 2:
 *
 *
 * Input: list1 = [], list2 = []
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: list1 = [], list2 = [0]
 * Output: [0]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in both lists is in the range [0, 50].
 * -100 <= Node.val <= 100
 * Both list1 and list2 are sorted in non-decreasing order.
 *
 *
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	// corner case when both lists are empty
	if list1 == nil && list2 == nil {
		return nil
	}

	mergedList := &ListNode{}
	head := mergedList
	prev := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			mergedList.Val = list1.Val
			list1 = list1.Next
		} else {
			mergedList.Val = list2.Val
			list2 = list2.Next
		}
		prev = mergedList
		mergedList.Next = &ListNode{}
		mergedList = mergedList.Next
	}
	for list1 != nil {
		mergedList.Val = list1.Val
		list1 = list1.Next
		prev = mergedList
		mergedList.Next = &ListNode{}
		mergedList = mergedList.Next
	}
	for list2 != nil {
		mergedList.Val = list2.Val
		list2 = list2.Next
		prev = mergedList
		mergedList.Next = &ListNode{}
		mergedList = mergedList.Next
	}
	prev.Next = nil

	return head
}

// @lc code=end
