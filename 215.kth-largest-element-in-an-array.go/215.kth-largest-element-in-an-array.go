package main

import (
	"errors"
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 *
 * https://leetcode.com/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (65.61%)
 * Likes:    13986
 * Dislikes: 696
 * Total Accepted:    1.8M
 * Total Submissions: 2.6M
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * Given an integer array nums and an integer k, return the k^th largest
 * element in the array.
 *
 * Note that it is the k^th largest element in the sorted order, not the k^th
 * distinct element.
 *
 * You must solve it in O(n) time complexity.
 *
 *
 * Example 1:
 * Input: nums = [3,2,1,5,6,4], k = 2
 * Output: 5
 * Example 2:
 * Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
 * Output: 4
 *
 *
 * Constraints:
 *
 *
 * 1 <= k <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 *
 *
 */

// @lc code=start

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Ordered interface {
	Integer | Float | ~string
}

type PQ[T Ordered] struct {
	keys    []*T
	size    int
	maxSize int
	lessFn  func(keys []*T, i, j int) bool
}

func maxPQless[T Ordered](keys []*T, i, j int) bool {
	return *keys[i] < *keys[j]
}

func minPQless[T Ordered](keys []*T, i, j int) bool {
	return *keys[i] > *keys[j]
}

func NewMaxPQ[T Ordered](maxSize int) *PQ[T] {
	keys := make([]*T, maxSize+1)
	return &PQ[T]{keys: keys, size: 0, maxSize: maxSize, lessFn: maxPQless[T]}
}

func NewMinPQ[T Ordered](maxSize int) *PQ[T] {
	keys := make([]*T, maxSize+1)
	return &PQ[T]{keys: keys, size: 0, maxSize: maxSize, lessFn: minPQless[T]}
}

func (pq *PQ[T]) IsEmpty() bool {
	return pq.size == 0
}

func (pq *PQ[T]) Size() int {
	return pq.size
}

func (pq *PQ[T]) Insert(key T) {
	pq.size++
	pq.keys[pq.size] = &key
	pq.swim(pq.size)
}

func (pq *PQ[T]) swim(k int) {
	for k > 1 && pq.lessFn(pq.keys, k/2, k) {
		pq.exchange(k/2, k)
		k = k / 2
	}
}

func (pq *PQ[T]) PeekTop() (*T, error) {
	if pq.IsEmpty() {
		return nil, errors.New("pq is empty")
	}
	top := pq.keys[1]
	return top, nil
}

func (pq *PQ[T]) DelTop() (*T, error) {
	if pq.IsEmpty() {
		return nil, errors.New("pq is empty")
	}
	top := *pq.keys[1]
	pq.exchange(1, pq.size)
	pq.size--
	pq.keys[pq.size+1] = nil
	pq.sink(1)
	return &top, nil
}

func (pq *PQ[T]) sink(k int) {
	for 2*k <= pq.size {
		leftChild := 2 * k
		rightChild := leftChild + 1
		chosenChild := leftChild
		if leftChild < pq.size && pq.lessFn(pq.keys, leftChild, rightChild) {
			chosenChild = rightChild
		}
		if !pq.lessFn(pq.keys, k, chosenChild) {
			break
		}
		pq.exchange(k, chosenChild)
		k = chosenChild
	}

}

func (pq *PQ[T]) exchange(i, j int) {
	pq.keys[i], pq.keys[j] = pq.keys[j], pq.keys[i]
}

func (pq *PQ[T]) String() string {
	var sb strings.Builder
	for _, v := range pq.keys {
		if v != nil {
			sb.WriteString(fmt.Sprintf("%v ->", *v))
		}
	}
	return sb.String()
}

func findKthLargest(nums []int, k int) int {
	maxPQ := NewMaxPQ[int](len(nums) + 1)
	for _, v := range nums {
		maxPQ.Insert(v)
	}
	for i := 0; i < k-1; i++ {
		maxPQ.DelTop()
	}
	finalVal, _ := maxPQ.DelTop()
	return *finalVal
}

// @lc code=end
