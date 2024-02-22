package problems

import "runtime"

/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start

func majorityElement(nums []int) int {
	var ans int
	var count int

	for _, num := range nums {
		if count == 0 {
			ans = num
		}
		if num == ans {
			count++
		} else {
			count--
		}
	}

	return ans
}

// @lc code=end

func majorityElementOld(nums []int) int {
	runtime.GC()

	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
		if m[n] > len(nums)/2 {
			return n
		}
	}

	panic("should exist")
}
