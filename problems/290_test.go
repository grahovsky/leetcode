package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test290(t *testing.T) {
	tests := []struct {
		pattern string
		s       string
		res     bool
	}{
		{
			pattern: "abba",
			s:       "dog cat cat dog",
			res:     true,
		},
		{
			pattern: "abba",
			s:       "dog cat cat fish",
			res:     false,
		},
		{
			pattern: "aaaa",
			s:       "dog cat cat dog",
			res:     false,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 290", func(t *testing.T) {
			require.Equal(t, tc.res, wordPattern(tc.pattern, tc.s))
		})
	}
}
