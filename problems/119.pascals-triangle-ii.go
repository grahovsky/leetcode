package problems

/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 */

// @lc code=start
func getRow(rowIndex int) []int {
	rowIndex++
	res := make([][]int, rowIndex)
	for rowN := 0; rowN < rowIndex; rowN++ {
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
	return res[rowIndex-1]
}

// @lc code=end
