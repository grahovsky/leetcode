package problems

import "strings"

/*
 * @lc app=leetcode id=290 lang=golang
 *
 * [290] Word Pattern
 */

// @lc code=start
func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	runeToWord := make(map[rune]string)
	wordToRune := make(map[string]rune)
	for i, sym := range pattern {
		word := words[i]
		if runeToWord[sym] == "" && wordToRune[word] == 0 {
			runeToWord[sym] = word
			wordToRune[word] = sym
		}
		if sym != wordToRune[word] {
			return false
		}
	}
	return true
}

// @lc code=end
