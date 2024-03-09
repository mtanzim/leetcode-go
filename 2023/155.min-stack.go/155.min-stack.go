package main

import (
	"fmt"
	"math"
	"strings"
)

/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 *
 * https://leetcode.com/problems/min-stack/description/
 *
 * algorithms
 * Medium (51.64%)
 * Likes:    11301
 * Dislikes: 721
 * Total Accepted:    1.2M
 * Total Submissions: 2.4M
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n' +
  '[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * Design a stack that supports push, pop, top, and retrieving the minimum
 * element in constant time.
 *
 * Implement the MinStack class:
 *
 *
 * MinStack() initializes the stack object.
 * void push(int val) pushes the element val onto the stack.
 * void pop() removes the element on the top of the stack.
 * int top() gets the top element of the stack.
 * int getMin() retrieves the minimum element in the stack.
 *
 *
 * You must implement a solution with O(1) time complexity for each
 * function.
 *
 *
 * Example 1:
 *
 *
 * Input
 * ["MinStack","push","push","push","getMin","pop","top","getMin"]
 * [[],[-2],[0],[-3],[],[],[],[]]
 *
 * Output
 * [null,null,null,null,-3,null,0,-2]
 *
 * Explanation
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin(); // return -3
 * minStack.pop();
 * minStack.top();    // return 0
 * minStack.getMin(); // return -2
 *
 *
 *
 * Constraints:
 *
 *
 * -2^31 <= val <= 2^31 - 1
 * Methods pop, top and getMin operations will always be called on non-empty
 * stacks.
 * At most 3 * 10^4 calls will be made to push, pop, top, and getMin.
 *
 *
*/

// @lc code=start
type LinkedList struct {
	Data int
	next *LinkedList
}

func (this *LinkedList) AddToHead(val int) *LinkedList {
	temp := this
	this = &LinkedList{Data: val}
	this.next = temp
	return this
}

func (this *LinkedList) RemoveHead() *LinkedList {
	this = this.next
	return this
}

type MinStack struct {
	ll  *LinkedList
	min *LinkedList
}

func Constructor() MinStack {
	return MinStack{min: &LinkedList{Data: math.MaxInt}}
}

func (this *MinStack) isEmpty() bool {
	return this.ll == nil
}

func (this *MinStack) Push(val int) {
	this.ll = this.ll.AddToHead(val)
	if val <= this.min.Data {
		this.min = this.min.AddToHead(val)
	}
}

func (this *MinStack) Pop() {
	if this.isEmpty() {
		return
	}
	v := this.Top()
	if v == this.GetMin() {
		this.min = this.min.RemoveHead()
	}
	this.ll = this.ll.RemoveHead()
}

func (this *MinStack) Top() int {
	if this.isEmpty() {
		return 0
	}
	return this.ll.Data
}

func (this *MinStack) GetMin() int {
	if this.isEmpty() {
		return 0
	}
	return this.min.Data
}

func (this *MinStack) String() string {
	var sb strings.Builder
	sb.WriteString("head ")

	cur := this.ll
	for cur != nil {
		sb.WriteString(fmt.Sprintf("-> %d ", cur.Data))
		cur = cur.next
	}
	return sb.String()

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
