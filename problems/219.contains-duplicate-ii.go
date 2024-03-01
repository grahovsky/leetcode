package problems

/*
 * @lc app=leetcode id=219 lang=golang
 *
 * [219] Contains Duplicate II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	cache := make(map[int]int)

	for i, n := range nums {
		if _, ok := cache[n]; ok {
			if (i - cache[n]) <= k {
				return true
			}
		}
		cache[n] = i
	}

	return false
}

// @lc code=end
