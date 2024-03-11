package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test263(t *testing.T) {
	tests := []struct {
		n   int
		res bool
	}{
		{
			n:   6,
			res: true,
		},
		{
			n:   1,
			res: true,
		},
		{
			n:   14,
			res: false,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 263", func(t *testing.T) {
			require.Equal(t, tc.res, isUgly(tc.n))
		})
	}
}
