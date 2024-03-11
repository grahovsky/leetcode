package problems

/*
 * @lc app=leetcode id=258 lang=golang
 *
 * [258] Add Digits
 */

// @lc code=start
func addDigits(num int) int {
	res := 0
	for num > 0 {
		res += num % 10
		num /= 10
	}

	if res > 9 {
		res = addDigits(res)
	}
	return res

	/* math solution
	   if num == 0 {
	       return 0
	   } else if (num % 9) == 0 {
	       return 9
	   }
	   return num % 9
	*/
}

// @lc code=end
