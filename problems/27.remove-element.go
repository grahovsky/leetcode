/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

package problems

// @lc code=start
func removeElement(nums []int, val int) int {
	k := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k], nums[i] = nums[i], nums[k]
			k++
		}
	}

	return k
}

func RemoveElement_old(nums []int, val int) int {
	k := len(nums)
	for i := 0; i < k; i++ {
		if nums[i] == val {
			k--
			for nums[k] == val && k > i {
				k--
			}
			nums[i], nums[k] = nums[k], nums[i]
		}
	}

	return k
}

// @lc code=end
