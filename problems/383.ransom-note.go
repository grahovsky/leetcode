package problems

/*
 * @lc app=leetcode id=383 lang=golang
 *
 * [383] Ransom Note
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	magMap := make(map[rune]int)
	for _, ch := range magazine {
		magMap[ch]++
	}

	for _, ch := range ransomNote {
		magMap[ch]--
		if magMap[ch] < 0 {
			return false
		}
	}

	return true
}

// @lc code=end
