/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

package problems

// @lc code=start
func isPalindromeNum(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	if x < 9 {
		return true
	}
	reverse := 0
	for x > reverse {
		reverse = reverse*10 + x%10
		x /= 10
	}
	return x == reverse || x == reverse/10
}

// @lc code=end
