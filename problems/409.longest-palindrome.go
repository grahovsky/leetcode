package problems

/*
 * @lc app=leetcode id=409 lang=golang
 *
 * [409] Longest Palindrome
 */

// @lc code=start
func longestPalindrome(s string) int {
	lets := make(map[rune]int)
	for _, l := range s {
		lets[l]++
	}

	pairs, one := 0, 0
	for _, count := range lets {
		pairs += count
		if count%2 == 1 {
			pairs--
			one = 1
		}
	}

	return pairs + one
}

// @lc code=end
