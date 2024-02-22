package problems

import "runtime"

/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
func majorityElement(nums []int) int {
	runtime.GC()

	m := make(map[int]int)
	for _, n := range nums {
		m[n] += 1
		if m[n] > len(nums)/2 {
			return n
		}
	}

	panic("should exist")
}

// @lc code=end
