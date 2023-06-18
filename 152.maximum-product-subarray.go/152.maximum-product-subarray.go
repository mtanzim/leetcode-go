package main

import (
	"fmt"
	"math"
	"sort"
)

/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 *
 * https://leetcode.com/problems/maximum-product-subarray/description/
 *
 * algorithms
 * Medium (34.86%)
 * Likes:    16196
 * Dislikes: 493
 * Total Accepted:    1M
 * Total Submissions: 2.9M
 * Testcase Example:  '[2,3,-2,4]'
 *
 * Given an integer array nums, find a subarray that has the largest product,
 * and return the product.
 *
 * The test cases are generated so that the answer will fit in a 32-bit
 * integer.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [2,3,-2,4]
 * Output: 6
 * Explanation: [2,3] has the largest product 6.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [-2,0,-1]
 * Output: 0
 * Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 2 * 10^4
 * -10 <= nums[i] <= 10
 * The product of any prefix or suffix of nums is guarant eed to fit in a 32-bit
 * integer.
 *
 *
 */

// @lc code=start

type cache struct {
	hm map[string]int
}

func newCache() *cache {
	return &cache{make(map[string]int)}
}

func traverse(nums []int, c *cache, start, end int) int {

	if len(nums) == 0 {
		return int(math.Inf(-1))
	}
	key := fmt.Sprintf("%d|%d", start, end)
	if v, ok := c.hm[key]; ok {
		// fmt.Println(fmt.Sprintf("cache hit for key: %s, v: %d", key, v))
		return v
	}
	if len(nums) == 1 {
		res := nums[0]
		c.hm[key] = res
		return res
	}
	if len(nums) == 2 {
		left := nums[0]
		right := nums[1]
		both := left * right
		all := []int{left, right, both}
		sort.Ints(all)
		res := all[len(all)-1]
		c.hm[key] = res
		return res

	}

	maxProd := int(math.Inf(-1))
	curProd := 1
	if cachedPrevProd, ok := c.hm[fmt.Sprintf("%d|%d", start, end-1)]; ok && end < len(nums) {
		curProd = cachedPrevProd * nums[end]
	} else {
		for _, v := range nums {
			curProd *= v
		}
	}

	if curProd > maxProd {
		maxProd = curProd
	}
	for i, v := range nums {
		left := traverse(nums[:i], c, start, start+i)
		if left > maxProd {
			maxProd = left
		}
		right := traverse(nums[i+1:], c, start+i+1, end)
		if right > maxProd {
			maxProd = right
		}
		if v > maxProd {
			maxProd = v
		}

	}
	c.hm[key] = maxProd
	return maxProd
}

// TODO: works but needs DP
func maxProduct(nums []int) int {
	return traverse(nums, newCache(), 0, len(nums)-1)
}

// @lc code=end
