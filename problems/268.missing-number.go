package problems

/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 */

// @lc code=start
func missingNumber(nums []int) int {
	values, calc := len(nums), 0

	for i, v := range nums {
		values += i
		calc += v
	}

	return values - calc
}

// @lc code=end
