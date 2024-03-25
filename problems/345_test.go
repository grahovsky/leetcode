package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test345(t *testing.T) {
	tests := []struct {
		s   string
		res string
	}{
		{
			s:   "hello",
			res: "holle",
		},
		{
			s:   "leetcode",
			res: "leotcede",
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 345", func(t *testing.T) {
			require.Equal(t, tc.res, reverseVowels(tc.s))
		})
	}
}
