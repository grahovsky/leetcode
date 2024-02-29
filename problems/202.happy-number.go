package problems

/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 */

// @lc code=start
func isHappy(n int) bool {
	// xxx -> .. > x < 10
	// example 9999 -> 324 -> 29 -> 85 -> 89 -> 145 -> 42 -> 20 -> 4
	if n < 10 {
		// 1 ->> happy
		// 7 ->> 49 -> 97 -> 130 -> 10 -> 1   happy
		// other -> x -> xx ->.. -> x
		return n == 1 || n == 7
	}

	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}

	if sum == 1 {
		return true
	}

	return isHappy(sum)
}

// @lc code=end
