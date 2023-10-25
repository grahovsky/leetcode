/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */
package problems

var closing = map[byte]byte{
	')': '(',
	'}': '{',
	']': '[',
}

func isValid(str string) bool {
	s := []byte(str)
	queue := make([]byte, 0)

	for _, ch := range s {
		opening, isClosing := closing[ch]

		if !isClosing { // it's opening
			queue = append(queue, ch)
			continue
		}

		if len(queue) == 0 { // closing, but no opening at all
			return false
		}

		lastSavedOpening := queue[len(queue)-1]
		if lastSavedOpening != opening {
			return false
		}

		queue = queue[:len(queue)-1] // "remove" lastOpening

	}

	return len(queue) == 0
}

// @lc code=end
