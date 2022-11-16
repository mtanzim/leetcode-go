package main

/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 *
 * https://leetcode.com/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (49.32%)
 * Likes:    12383
 * Dislikes: 690
 * Total Accepted:    1.3M
 * Total Submissions: 2.6M
 * Testcase Example:  '[1,2,2,1]'
 *
 * Given the head of a singly linked list, return true if it is a palindrome or
 * false otherwise.
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,2,2,1]
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: head = [1,2]
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is in the range [1, 10^5].
 * 0 <= Node.val <= 9
 *
 *
 *
 * Follow up: Could you do it in O(n) time and O(1) space?
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
func reverseList(head *ListNode) *ListNode {

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

func isPalindrome(head *ListNode) bool {
	reversedHead := reverseList(head)
	cur := head
	curReversed := reversedHead
	for cur != nil && curReversed != nil {
		if cur.Val != curReversed.Val {
			return false
		}
		cur = cur.Next
		curReversed = curReversed.Next
	}
	return cur == nil && curReversed == nil

}

// @lc code=end
