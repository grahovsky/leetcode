package problems

/*
 * @lc app=leetcode id=367 lang=golang
 *
 * [367] Valid Perfect Square
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}

	low, high := 0, num/2
	for low <= high {
		mid := low + (high-low)/2

		if mid*mid == num {
			return true
		}

		if mid*mid > num {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return false
}

// @lc code=end
