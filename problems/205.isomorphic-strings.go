package problems

/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	words1 := make(map[byte]byte)
	words2 := make(map[byte]byte)

	for i := range t {
		if words1[s[i]] == 0 {
			words1[s[i]] = t[i]
		}
		if words2[t[i]] == 0 {
			words2[t[i]] = s[i]
		}
		if words1[s[i]] != t[i] || words2[t[i]] != s[i] {
			return false
		}
	}

	return true
}

// @lc code=end
