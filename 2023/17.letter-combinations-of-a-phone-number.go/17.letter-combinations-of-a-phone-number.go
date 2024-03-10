package main

/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 *
 * https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (55.42%)
 * Likes:    15029
 * Dislikes: 841
 * Total Accepted:    1.6M
 * Total Submissions: 2.7M
 * Testcase Example:  '"23"'
 *
 * Given a string containing digits from 2-9 inclusive, return all possible
 * letter combinations that the number could represent. Return the answer in
 * any order.
 *
 * A mapping of digits to letters (just like on the telephone buttons) is given
 * below. Note that 1 does not map to any letters.
 *
 *
 * Example 1:
 *
 *
 * Input: digits = "23"
 * Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
 *
 *
 * Example 2:
 *
 *
 * Input: digits = ""
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: digits = "2"
 * Output: ["a","b","c"]
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= digits.length <= 4
 * digits[i] is a digit in the range ['2', '9'].
 *
 *
 */

// @lc code=start

var t9hm = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func splitVal(digit string) []string {
	val := t9hm[digit]
	chars := make([]string, len(val))
	for i := range chars {
		chars[i] = string(val[i])
	}
	return chars
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	if len(digits) == 1 {
		return splitVal(string(digits[0]))
	}
	head := t9hm[string(digits[0])]
	tail := digits[1:]
	tailRes := letterCombinations(tail)

	combinedRes := []string{}
	for _, h := range head {
		for _, v := range tailRes {
			combinedRes = append(combinedRes, string(h)+v)
		}
	}
	return combinedRes
}

// @lc code=end
