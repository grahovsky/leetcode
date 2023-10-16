/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

package problems

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		if id, ok := m[n]; ok {
			return []int{id, i}
		}
		m[target-n] = i
	}
	return []int{}
}

// @lc code=end
