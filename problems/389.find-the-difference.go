package problems

/*
 * @lc app=leetcode id=389 lang=golang
 *
 * [389] Find the Difference
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	var m byte
	for i := range s {
		m += t[i]
		m -= s[i]
	}
	return m + t[len(t)-1]
}

// @lc code=end
