/*
 * @lc app=leetcode id=844 lang=golang
 *
 * [844] Backspace String Compare
 */

package problems

import "bytes"

// @lc code=start
func backspaceCompare(s string, t string) bool {
	return trimStringReverse(s) == trimStringReverse(t)
}

func trimStringReverse(s string) string {
	var buffer bytes.Buffer
	sBack := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '#' {
			sBack++
			continue
		}
		if sBack > 0 {
			sBack--
			continue
		}
		buffer.WriteByte(s[i])
	}
	return buffer.String()
}

// @lc code=end
