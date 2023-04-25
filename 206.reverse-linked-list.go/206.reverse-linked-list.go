package main

/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 *
 * https://leetcode.com/problems/reverse-linked-list/description/
 *
 * algorithms
 * Easy (72.20%)
 * Likes:    15528
 * Dislikes: 259
 * Total Accepted:    2.7M
 * Total Submissions: 3.7M
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * Given the head of a singly linked list, reverse the list, and return the
 * reversed list.
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,2,3,4,5]
 * Output: [5,4,3,2,1]
 *
 *
 * Example 2:
 *
 *
 * Input: head = [1,2]
 * Output: [2,1]
 *
 *
 * Example 3:
 *
 *
 * Input: head = []
 * Output: []
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is the range [0, 5000].
 * -5000 <= Node.val <= 5000
 *
 *
 *
 * Follow up: A linked list can be reversed either iteratively or recursively.
 * Could you implement both?
 *
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

func reverseListIterative(head *ListNode) *ListNode {

	temp := []int{}
	for cur := head; cur != nil; cur = cur.Next {
		temp = append(temp, cur.Val)
	}
	if len(temp) == 0 {
		return nil
	}
	newLL := &ListNode{}
	newHead := newLL
	prev := newLL
	for i := len(temp) - 1; i >= 0; i-- {
		newLL.Val = temp[i]
		newLL.Next = &ListNode{}
		prev = newLL
		newLL = newLL.Next
	}
	prev.Next = nil
	return newHead
}

func reverse(head *ListNode) (*ListNode, *ListNode) {

	if head.Next == nil {
		startNode := &ListNode{Val: head.Val}
		return startNode, startNode
	}
	reversedHead, reversedCurPtr := reverse(head.Next)
	reversedCurPtr.Next = &ListNode{Val: head.Val}
	return reversedHead, reversedCurPtr.Next
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	reversed := reverseListIterative(head)
	return reversed

}

// @lc code=end
