package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test392(t *testing.T) {
	tests := []struct {
		s   string
		t   string
		res bool
	}{
		{
			s:   "abc",
			t:   "ahbgdc",
			res: true,
		},
		{
			s:   "axc",
			t:   "ahbgdc",
			res: false,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 392", func(t *testing.T) {
			require.Equal(t, tc.res, isSubsequence(tc.s, tc.t))
		})
	}
}
