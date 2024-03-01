package problems

/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	cache := make(map[int]struct{})

	for _, n := range nums {
		if _, ok := cache[n]; ok {
			return true
		}
		cache[n] = struct{}{}
	}

	return false
}

// @lc code=end
