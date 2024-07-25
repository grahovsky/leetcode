package problems

import "math"

/*
 * @lc app=leetcode id=414 lang=golang
 *
 * [414] Third Maximum Number
 */

// @lc code=start
func thirdMax(nums []int) int {
	first, second, third := math.MinInt, math.MinInt, math.MinInt

	for _, num := range nums {

		if num == first || num == second || num == third {
			continue
		}

		if num > first {
			first, second, third = num, first, second
		} else if num > second {
			second, third = num, second
		} else if num > third {
			third = num
		}
	}

	if third == math.MinInt {
		return first
	}

	return third
}

// @lc code=end
