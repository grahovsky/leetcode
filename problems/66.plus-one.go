/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

package problems

// @lc code=start
func plusOne(digits []int) []int {
	incr := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += incr
		incr--
		if digits[i] == 10 {
			digits[i] = 0
			incr++
		}
		if incr == 0 {
			return digits
		}
	}

	return append([]int{1}, digits...)
}

// @lc code=end
