package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test389(t *testing.T) {
	tests := []struct {
		s   string
		t   string
		res byte
	}{
		{
			s:   "abcd",
			t:   "abcde",
			res: 'e',
		},
		{
			s:   "",
			t:   "y",
			res: 'y',
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 389", func(t *testing.T) {
			require.Equal(t, tc.res, findTheDifference(tc.s, tc.t))
		})
	}
}
