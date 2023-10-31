/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

package problems

// @lc code=start
func searchInsert(nums []int, target int) int {
	return searchIndex(nums, target, 0)
}

func searchIndex(nums []int, target int, start int) int {
	if len(nums) == 1 {
		if target > nums[0] {
			start++
		}
		return start
	}

	pivot := len(nums) / 2
	if target < nums[pivot] {
		return searchIndex(nums[:pivot], target, start)
	}
	return searchIndex(nums[pivot:], target, start+pivot)
}

// @lc code=end

func SearchInsertNoOLog2(nums []int, target int) int {
	for i, n := range nums {
		if target <= n {
			return i
		}
	}
	return len(nums)
}
