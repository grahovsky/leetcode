/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

package problems

// @lc code=start
func romanToInt(s string) int {
	rMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	big := 0
	res := 0

	for i := len(s) - 1; i >= 0; i-- {
		current := s[i]
		val := rMap[current]
		if val >= big {
			big = val
			res += val
		} else {
			res -= val
		}
	}

	return res
}

// @lc code=end
