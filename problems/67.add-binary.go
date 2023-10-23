/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */

package problems

import "fmt"

// @lc code=start
func addBinary(a, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	iB := len(b) - 1
	result := ""
	var sum, shift byte

	for iA := len(a) - 1; iA >= 0; iA-- {
		sum = (a[iA] - 48) // '1' to 1, '0' to 0
		if shift > 0 {
			sum += shift
			shift--
		}
		if iB >= 0 {
			sum += (b[iB] - 48)
			iB--
		}

		if sum > 1 {
			shift++
		}

		result = fmt.Sprintf("%b", sum%2) + result
	}

	if shift == 0 {
		return result
	}
	return "1" + result
}

// @lc code=end

// func addBinary_big(a string, b string) string {
// 	aInt, bInt := new(big.Int), new(big.Int)
// 	aInt.SetString(a, 2)
// 	bInt.SetString(b, 2)

// 	return new(big.Int).Add(aInt, bInt).Text(2)
// }
