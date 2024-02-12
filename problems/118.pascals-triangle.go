package problems

/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */

// @lc code=start
func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for rowN := 0; rowN < numRows; rowN++ {
		row := make([]int, rowN+1)
		row[0] = 1
		for colN := 1; colN <= rowN; colN++ {
			if colN < rowN {
				row[colN] = res[rowN-1][colN-1] + res[rowN-1][colN]
			} else {
				row[colN] = 1
			}
		}
		res[rowN] = row
	}
	return res
}

// @lc code=end
