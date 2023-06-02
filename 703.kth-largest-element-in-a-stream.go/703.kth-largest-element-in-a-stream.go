package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

/*
 * @lc app=leetcode id=703 lang=golang
 *
 * [703] Kth Largest Element in a Stream
 *
 * https://leetcode.com/problems/kth-largest-element-in-a-stream/description/
 *
 * algorithms
 * Easy (55.50%)
 * Likes:    4695
 * Dislikes: 2805
 * Total Accepted:    412K
 * Total Submissions: 725.3K
 * Testcase Example:  '["KthLargest","add","add","add","add","add"]\n' +
  '[[3,[4,5,8,2]],[3],[5],[10],[9],[4]]'
 *
 * Design a class to find the k^th largest element in a stream. Note that it is
 * the k^th largest element in the sorted order, not the k^th distinct
 * element.
 *
 * Implement KthLargest class:
 *
 *
 * KthLargest(int k, int[] nums) Initializes the object with the integer k and
 * the stream of integers nums.
 * int add(int val) Appends the integer val to the stream and returns the
 * element representing the k^th largest element in the stream.
 *
 *
 *
 * Example 1:
 *
 *
 * Input
 * ["KthLargest", "add", "add", "add", "add", "add"]
 * [[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
 * Output
 * [null, 4, 5, 5, 8, 8]
 *
 * Explanation
 * KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
 * kthLargest.add(3);   // return 4
 * kthLargest.add(5);   // return 5
 * kthLargest.add(10);  // return 5
 * kthLargest.add(9);   // return 8
 * kthLargest.add(4);   // return 8
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= k <= 10^4
 * 0 <= nums.length <= 10^4
 * -10^4 <= nums[i] <= 10^4
 * -10^4 <= val <= 10^4
 * At most 10^4 calls will be made to add.
 * It is guaranteed that there will be at least k elements in the array when
 * you search for the k^th element.
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

type KthLargest struct {
	k         int
	maxSize   int
	maxPq     *PQ[int]
	prevMinPQ *PQ[int]
	nums      []int
}

func Constructor(k int, nums []int) KthLargest {
	maxSize := 2*int(math.Pow10(4)) + 1
	maxPQ := NewMaxPQ[int](maxSize)
	for _, v := range nums {
		maxPQ.Insert(v)
	}
	return KthLargest{
		k:         k,
		maxSize:   maxSize,
		maxPq:     maxPQ,
		nums:      nums,
		prevMinPQ: nil,
	}

}

func (this *KthLargest) Add(val int) int {

	if this.prevMinPQ != nil {
		prevKth, _ := this.prevMinPQ.PeekTop()
		possibleRes := *prevKth
		if val < possibleRes {
			return possibleRes
		}
	}

	this.maxPq.Insert(val)

	temp := make([]int, this.k)
	for i := 0; i < this.k; i++ {
		top, _ := this.maxPq.DelTop()
		temp[i] = *top
	}
	res := temp[this.k-1]
	for _, v := range temp {
		this.maxPq.Insert(v)
	}

	this.prevMinPQ = NewMinPQ[int](this.k)
	for _, v := range temp {
		this.prevMinPQ.Insert(v)
	}
	return res
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
// @lc code=end
