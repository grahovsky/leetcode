package problems

/*
 * @lc app=leetcode id=231 lang=golang
 *
 * [231] Power of Two
 */

// @lc code=start
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}

	return n&(n-1) == 0
}

// @lc code=end

/* explain
1 = 2^0 = 1 = 0b0001
2 = 2^1 = 2 = 0b0010
4 = 2^2 = 4 = 0b0100
8 = 2^3 = 8 = 0b1000

check for 1xxx
1000 (n)
&
0111 (n-1)
=
0000
*/
