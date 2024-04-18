package problems

/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 */

// @lc code=start
func firstUniqChar(s string) int {
	cache := make(map[rune]int)
	for _, r := range s {
		cache[r]++
	}

	for i, r := range s {
		if cache[r] == 1 {
			return i
		}
	}

	return -1
}

// @lc code=end
