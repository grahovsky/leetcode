package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test28(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		result   int
	}{
		{haystack: "sadbutsad", needle: "sad", result: 0},
		{haystack: "leetcode", needle: "leeto", result: -1},
		{haystack: "a", needle: "a", result: 0},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 28", func(t *testing.T) {
			result := strStr(tc.haystack, tc.needle)
			require.Equal(t, tc.result, result)
		})
	}
}
