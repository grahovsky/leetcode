/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

package problems

// @lc code=start
func removeDuplicates(nums []int) int {
	nI := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			nums[nI] = nums[i]
			nI++
		}
	}

	return nI
}

// @lc code=end
