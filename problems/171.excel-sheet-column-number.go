package problems

import (
	"math"
)

/*
 * @lc app=leetcode id=171 lang=golang
 *
 * [171] Excel Sheet Column Number
 */

// @lc code=start
func titleToNumber(columnTitle string) int {
	res := 0

	for i, r := range columnTitle {
		res = res + (int(r)-64)*int(math.Pow(26.0, float64(len(columnTitle)-1-i)))
	}

	return res
}

// @lc code=end
