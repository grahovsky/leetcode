package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test387(t *testing.T) {
	tests := []struct {
		s   string
		res int
	}{
		{
			s:   "leetcode",
			res: 0,
		},
		{
			s:   "loveleetcode",
			res: 2,
		},
		{
			s:   "aabb",
			res: -1,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 387", func(t *testing.T) {
			require.Equal(t, tc.res, firstUniqChar(tc.s))
		})
	}
}
