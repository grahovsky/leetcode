/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

package problems

// @lc code=start
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	return append([]int{1}, digits...)
}

// @lc code=end
