package problems

/*
 * @lc app=leetcode id=303 lang=golang
 *
 * [303] Range Sum Query - Immutable
 */

// @lc code=start
type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	numArray := NumArray{sums: make([]int, len(nums))}

	sum := 0
	for i, num := range nums {
		sum += num
		numArray.sums[i] = sum
	}

	return numArray
}

func (numArray *NumArray) SumRange(left int, right int) int {
	leftSum := 0
	if left > 0 {
		leftSum = numArray.sums[left-1]
	}
	return numArray.sums[right] - leftSum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// @lc code=end
