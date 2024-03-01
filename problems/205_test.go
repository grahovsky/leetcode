package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test205(t *testing.T) {
	tests := []struct {
		s   string
		t   string
		res bool
	}{
		{
			s:   "egg",
			t:   "add",
			res: true,
		},
		{
			s:   "foo",
			t:   "bat",
			res: false,
		},
		{
			s:   "paper",
			t:   "title",
			res: true,
		},
		{
			s:   "abad",
			t:   "adcd",
			res: false,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 205", func(t *testing.T) {
			require.Equal(t, tc.res, isIsomorphic(tc.s, tc.t))
		})
	}
}
