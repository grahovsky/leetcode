package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test231(t *testing.T) {
	tests := []struct {
		n   int
		res bool
	}{
		{
			n:   2,
			res: true,
		},
		{
			n:   1024,
			res: true,
		},
		{
			n:   7,
			res: false,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 231", func(t *testing.T) {
			require.Equal(t, tc.res, isPowerOfTwo(tc.n))
		})
	}
}
