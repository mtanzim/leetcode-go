package main

/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 *
 * https://leetcode.com/problems/coin-change/description/
 *
 * algorithms
 * Medium (42.31%)
 * Likes:    16380
 * Dislikes: 368
 * Total Accepted:    1.4M
 * Total Submissions: 3.3M
 * Testcase Example:  '[1,2,5]\n11'
 *
 * You are given an integer array coins representing coins of different
 * denominations and an integer amount representing a total amount of money.
 *
 * Return the fewest number of coins that you need to make up that amount. If
 * that amount of money cannot be made up by any combination of the coins,
 * return -1.
 *
 * You may assume that you have an infinite number of each kind of coin.
 *
 *
 * Example 1:
 *
 *
 * Input: coins = [1,2,5], amount = 11
 * Output: 3
 * Explanation: 11 = 5 + 5 + 1
 *
 *
 * Example 2:
 *
 *
 * Input: coins = [2], amount = 3
 * Output: -1
 *
 *
 * Example 3:
 *
 *
 * Input: coins = [1], amount = 0
 * Output: 0
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= coins.length <= 12
 * 1 <= coins[i] <= 2^31 - 1
 * 0 <= amount <= 10^4
 *
 *
 */

// @lc code=start
// TODO: this is broken, come back to it
func traverse(coins, maxCoins []int, amount int) [][]int {
	if len(coins) == 0 {
		return [][]int{}
	}
	sum := 0
	curRes := [][]int{}
	for i, v := range coins {
		maxCount := maxCoins[i]
		sum += maxCount * v
	}
	if sum == amount {
		curRes = append(curRes, maxCoins)
	}
	newMaxCoins := append([]int{}, maxCoins...)
	newCoins := append([]int{}, coins...)
	newMaxCoins[0]--
	if newMaxCoins[0] == -1 {
		newCoins = newCoins[1:]
		newMaxCoins = newMaxCoins[1:]
	}
	return append(curRes, traverse(newCoins, newMaxCoins, amount)...)

}

func coinChange(coins []int, amount int) int {
	maxCoins := make([]int, len(coins))
	for i, v := range coins {
		maxCoins[i] = amount / v
	}

	return len(traverse(coins, maxCoins, amount))
}

// @lc code=end
