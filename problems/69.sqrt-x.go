/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

package problems

// @lc code=start
func mySqrt(x int) int {
	left, right := 0, x+1

	for left < right {
		mid := left + (right-left)/2

		if mid*mid > x {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left - 1
}

// @lc code=end

func MySqrt_old(x int) int {
	for i := 1; true; i++ {
		if i*i > x {
			return i - 1
		}
	}
	return x
}
