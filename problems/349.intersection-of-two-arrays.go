package problems

/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{})
	interset := []int{}

	for _, num := range nums1 {
		_, exist := set[num]
		if exist {
			continue
		}
		for _, numN := range nums2 {
			if num == numN {
				interset = append(interset, num)
				set[num] = struct{}{}
				break
			}
		}
	}
	return interset
}

// @lc code=end
