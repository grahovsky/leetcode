package problems

/*
 * @lc app=leetcode id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	index := 0
	for i := 0; i < len(nums1); i++ {
		for j := index; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				nums1[index] = nums1[i]
				nums2[index], nums2[j] = nums2[j], nums2[index]
				index++
				break
			}
		}
	}

	return nums1[:index]
}

// @lc code=end
