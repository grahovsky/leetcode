/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */

package problems

// @lc code=start
func lengthOfLastWord(s string) int {
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			count++
		} else if count > 0 {
			return count
		}
	}
	return count
}

// @lc code=end

func LengthOfLastWordOld(s string) int {
	space := true
	endSpaces := 0

	for i := 0; i < len(s); i++ {
		if s[len(s)-i-1] == ' ' {
			if !space {
				return i - endSpaces
			}
			endSpaces++
		} else {
			space = false
		}
	}
	return len(s) - endSpaces
}
