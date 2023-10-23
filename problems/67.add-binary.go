/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */

package problems

import "math/big"

// @lc code=start
func addBinary(a string, b string) string {
	aInt, bInt := new(big.Int), new(big.Int)
	aInt.SetString(a, 2)
	bInt.SetString(b, 2)

	return new(big.Int).Add(aInt, bInt).Text(2)
}

// @lc code=end
