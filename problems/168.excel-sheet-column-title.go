package problems

/*
 * @lc app=leetcode id=168 lang=golang
 *
 * [168] Excel Sheet Column Title
 */

// @lc code=start
func convertToTitle(columnNumber int) string {
	column := func(n int) string {
		return string(byte(rune(n + 64)))
	}

	title := ""
	for columnNumber > 0 {
		if columnNumber <= 26 {
			title = column(columnNumber) + title
			break
		}

		diff := columnNumber % 26
		columnNumber /= 26
		if diff == 0 {
			title = "Z" + title
			columnNumber--
		} else {
			title = column(diff) + title
		}
	}

	return title
}

// @lc code=end
