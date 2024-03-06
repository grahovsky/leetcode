package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test242(t *testing.T) {
	tests := []struct {
		s   string
		t   string
		res bool
	}{
		{
			s:   "anagram",
			t:   "nagaram",
			res: true,
		},
		{
			s:   "rat",
			t:   "car",
			res: false,
		},
		{
			s:   "rat",
			t:   "tarr",
			res: false,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 242", func(t *testing.T) {
			require.Equal(t, tc.res, isAnagram(tc.s, tc.t))
		})
	}
}
