package problems

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=228 lang=golang
 *
 * [228] Summary Ranges
 */

// @lc code=start
func summaryRanges(nums []int) []string {
	result := []string{}
	builder := strings.Builder{}

	for i := 0; i < len(nums); i++ {
		if builder.Len() == 0 {
			builder.WriteString(strconv.Itoa(nums[i]))
			continue
		}

		if nums[i] > nums[i-1]+1 {
			end := strconv.Itoa(nums[i-1])
			if builder.String() != end {
				builder.WriteString("->")
				builder.WriteString(end)
			}

			result = append(result, builder.String())
			builder.Reset()
			builder.WriteString(strconv.Itoa(nums[i]))
		}
	}

	if builder.Len() > 0 {
		end := strconv.Itoa(nums[len(nums)-1])
		if builder.String() != end {
			builder.WriteString("->")
			builder.WriteString(end)
		}
		result = append(result, builder.String())
	}
	return result
}

// @lc code=end
