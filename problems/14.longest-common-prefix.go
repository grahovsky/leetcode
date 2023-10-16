/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

package problems

// @lc code=start
func longestCommonPrefix(strs []string) string {
	pref := []byte{}

loop:
	for i := 0; true; i++ {
		for _, s := range strs {
			if i > len(s)-1 || s[i] != strs[0][i] {
				break loop
			}
		}
		pref = append(pref, strs[0][i])
	}

	return string(pref)
}

// @lc code=end
