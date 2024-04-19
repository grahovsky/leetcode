package problems

import "fmt"

/*
 * @lc app=leetcode id=401 lang=golang
 *
 * [401] Binary Watch
 */

// @lc code=start
func readBinaryWatch(turnedOn int) []string {
	s := []string{}
	for h := 0; h < 12; h++ {
		bh := onesCount(h)
		for m := 0; m < 60; m++ {
			bm := onesCount(m)
			if bh+bm == turnedOn {
				s = append(s, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}

	return s
}

func onesCount(n int) int {
	ret := 0
	for n != 0 {
		ret += n & 0b1
		n >>= 1
	}
	return ret
}

// @lc code=end
