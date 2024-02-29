package problems

/*
 * @lc app=leetcode id=191 lang=golang
 *
 * [191] Number of 1 Bits
 */

// @lc code=start
func hammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		if num&0b1 == 1 {
			res++
		}
		num >>= 1
	}
	return res
}

// @lc code=end
