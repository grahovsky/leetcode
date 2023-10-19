/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Find the Index of the First Occurrence in a String
 */

package problems

// @lc code=start
func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		sub := 0
		for j := 0; j < len(needle); j++ {
			if needle[j] == haystack[i+j] {
				sub++
			}
		}
		if sub == len(needle) {
			return i
		}
	}

	return -1
}

// @lc code=end
