package problems

/*
 * @lc app=leetcode id=374 lang=golang
 *
 * [374] Guess Number Higher or Lower
 */

// @lc code=start
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	left, right, mid := 1, n, 0
	for left <= right {
		mid = (left + right) / 2
		if pre := guess(mid); pre == -1 {
			right = mid - 1
		} else if pre == 1 {
			left = mid + 1
		} else {
			break
		}
	}

	return mid
}

// @lc code=end

func guess(n int) int {
	if n < 6 {
		return 1
	} else if n > 6 {
		return -1
	}
	return 0
}
