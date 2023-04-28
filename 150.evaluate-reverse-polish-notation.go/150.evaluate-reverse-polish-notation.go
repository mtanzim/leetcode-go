package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 *
 * https://leetcode.com/problems/evaluate-reverse-polish-notation/description/
 *
 * algorithms
 * Medium (44.04%)
 * Likes:    5705
 * Dislikes: 869
 * Total Accepted:    649.5K
 * Total Submissions: 1.4M
 * Testcase Example:  '["2","1","+","3","*"]'
 *
 * You are given an array of strings tokens that represents an arithmetic
 * expression in a Reverse Polish Notation.
 *
 * Evaluate the expression. Return an integer that represents the value of the
 * expression.
 *
 * Note that:
 *
 *
 * The valid operators are '+', '-', '*', and '/'.
 * Each operand may be an integer or another expression.
 * The division between two integers always truncates toward zero.
 * There will not be any division by zero.
 * The input represents a valid arithmetic expression in a reverse polish
 * notation.
 * The answer and all the intermediate calculations can be represented in a
 * 32-bit integer.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: tokens = ["2","1","+","3","*"]
 * Output: 9
 * Explanation: ((2 + 1) * 3) = 9
 *
 *
 * Example 2:
 *
 *
 * Input: tokens = ["4","13","5","/","+"]
 * Output: 6
 * Explanation: (4 + (13 / 5)) = 6
 *
 *
 * Example 3:
 *
 *
 * Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
 * Output: 22
 * Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
 * = ((10 * (6 / (12 * -11))) + 17) + 5
 * = ((10 * (6 / -132)) + 17) + 5
 * = ((10 * 0) + 17) + 5
 * = (0 + 17) + 5
 * = 17 + 5
 * = 22
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= tokens.length <= 10^4
 * tokens[i] is either an operator: "+", "-", "*", or "/", or an integer in the
 * range [-200, 200].
 *
 *
 */

// @lc code=start
type Operator func(a, b int) int

var operators = map[string]Operator{
	"+": func(a, b int) int {
		return a + b
	},
	"-": func(a, b int) int {
		return a - b
	},
	"*": func(a, b int) int {
		return a * b
	},
	"/": func(a, b int) int {
		return a / b
	},
}

func traverse(curTokens []string) (int, error) {
	if len(curTokens) == 1 {
		return strconv.Atoi(curTokens[0])
	}
	for i, token := range curTokens {
		if fn, ok := operators[token]; ok {
			a, err := strconv.Atoi(curTokens[i-1])
			if err != nil {
				return 0, err
			}
			b, err := strconv.Atoi(curTokens[i-2])
			if err != nil {
				return 0, err
			}
			val := fn(a, b)
			valS := fmt.Sprintf("%d", val)
			updatedTokens := curTokens[:i-2]
			updatedTokens = append(updatedTokens, valS)
			updatedTokens = append(updatedTokens, curTokens[i+1:]...)
			return traverse(updatedTokens)
		}
	}
	return 0, errors.New("function did not execute correctly")
}

func evalRPN(tokens []string) int {
	v, err := traverse(tokens)
	if err != nil {
		fmt.Printf("unexpected error: %v", err)
		return 0
	}
	return v
}

// @lc code=end
