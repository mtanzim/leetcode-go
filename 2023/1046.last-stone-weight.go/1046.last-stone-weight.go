package main

import (
	"errors"
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode id=1046 lang=golang
 *
 * [1046] Last Stone Weight
 *
 * https://leetcode.com/problems/last-stone-weight/description/
 *
 * algorithms
 * Easy (64.73%)
 * Likes:    5284
 * Dislikes: 91
 * Total Accepted:    482.3K
 * Total Submissions: 737.1K
 * Testcase Example:  '[2,7,4,1,8,1]'
 *
 * You are given an array of integers stones where stones[i] is the weight of
 * the i^th stone.
 *
 * We are playing a game with the stones. On each turn, we choose the heaviest
 * two stones and smash them together. Suppose the heaviest two stones have
 * weights x and y with x <= y. The result of this smash is:
 *
 *
 * If x == y, both stones are destroyed, and
 * If x != y, the stone of weight x is destroyed, and the stone of weight y has
 * new weight y - x.
 *
 *
 * At the end of the game, there is at most one stone left.
 *
 * Return the weight of the last remaining stone. If there are no stones left,
 * return 0.
 *
 *
 * Example 1:
 *
 *
 * Input: stones = [2,7,4,1,8,1]
 * Output: 1
 * Explanation:
 * We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
 * we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
 * we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
 * we combine 1 and 1 to get 0 so the array converts to [1] then that's the
 * value of the last stone.
 *
 *
 * Example 2:
 *
 *
 * Input: stones = [1]
 * Output: 1
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= stones.length <= 30
 * 1 <= stones[i] <= 1000
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

func lastStoneWeight(stones []int) int {
	maxPQ := NewMaxPQ[int](len(stones) + 1)
	for _, v := range stones {
		maxPQ.Insert(v)
	}
	for {
		if maxPQ.Size() < 2 {
			break
		}
		max, _ := maxPQ.DelTop()
		smallerMax, _ := maxPQ.DelTop()
		if *max == *smallerMax {
			continue
		}
		newVal := *max - *smallerMax
		maxPQ.Insert(newVal)
	}
	finalVal, _ := maxPQ.DelTop()
	if finalVal == nil {
		return 0
	}
	return *finalVal
}

// @lc code=end
