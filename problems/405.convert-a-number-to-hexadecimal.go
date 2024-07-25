package problems

/*
 * @lc app=leetcode id=405 lang=golang
 *
 * [405] Convert a Number to Hexadecimal
 */

// @lc code=start
func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	n, mask, res := uint32(num), "0123456789abcdef", ""

	for n != 0 {
		res = string(mask[n&0xf]) + res
		n >>= 4
	}
	return res
}

// @lc code=end
