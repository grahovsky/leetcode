package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test338(t *testing.T) {
	tests := []struct {
		n   int
		res []int
	}{
		{
			n:   2,
			res: []int{0, 1, 1},
		},
		{
			n:   5,
			res: []int{0, 1, 1, 2, 1, 2},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 338", func(t *testing.T) {
			require.Equal(t, tc.res, countBits(tc.n))
		})
	}
}
