package problems

/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 */

// @lc code=start
func singleNumber(nums []int) int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	return xor
}

// @lc code=end
