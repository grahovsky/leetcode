/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

package problems

// @lc code=start
func mySqrt(x int) int {
	for i := 1; true; i++ {
		if i*i > x {
			return i - 1
		}
	}
	return x
}

// @lc code=end
