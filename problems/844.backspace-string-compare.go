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
		} else {
			if sBack == 0 {
				buffer.WriteByte(s[i])
			} else {
				sBack--
			}
		}
	}
	return buffer.String()
}

// @lc code=end
