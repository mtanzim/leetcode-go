package main

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode id=143 lang=golang
 *
 * [143] Reorder List
 *
 * https://leetcode.com/problems/reorder-list/description/
 *
 * algorithms
 * Medium (50.73%)
 * Likes:    8584
 * Dislikes: 290
 * Total Accepted:    660.3K
 * Total Submissions: 1.3M
 * Testcase Example:  '[1,2,3,4]'
 *
 * You are given the head of a singly linked-list. The list can be represented
 * as:
 *
 *
 * L0 → L1 → … → Ln - 1 → Ln
 *
 *
 * Reorder the list to be on the following form:
 *
 *
 * L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
 *
 *
 * You may not modify the values in the list's nodes. Only nodes themselves may
 * be changed.
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,2,3,4]
 * Output: [1,4,2,3]
 *
 *
 * Example 2:
 *
 *
 * Input: head = [1,2,3,4,5]
 * Output: [1,5,2,4,3]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is in the range [1, 5 * 10^4].
 * 1 <= Node.val <= 1000
 *
 *
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func (this *ListNode) String() string {
	var sb strings.Builder
	sb.WriteString("head ")

	cur := this
	for cur != nil {
		sb.WriteString(fmt.Sprintf("-> %d ", cur.Val))
		cur = cur.Next
	}
	return sb.String()
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

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

func reorderList(head *ListNode) {
	reversed := reverseListIterative(head)
	listLen := 0
	cur := head
	for cur != nil {
		listLen++
		cur = cur.Next
	}

	curOrig := head
	curReversed := reversed
	orderedLL := &ListNode{}
	curOrdered := orderedLL
	prev := curOrdered
	for i := 0; i < listLen; i++ {
		if i%2 == 0 {
			curOrdered.Val = curOrig.Val
			curOrig = curOrig.Next

		} else {
			curOrdered.Val = curReversed.Val
			curReversed = curReversed.Next
		}
		prev = curOrdered
		curOrdered.Next = &ListNode{}
		curOrdered = curOrdered.Next
	}
	prev.Next = nil
	head.Next = orderedLL.Next
}

// @lc code=end
