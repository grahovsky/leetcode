package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test326(t *testing.T) {
	tests := []struct {
		n   int
		res bool
	}{
		{
			n:   27,
			res: true,
		},
		{
			n:   0,
			res: false,
		},
		{
			n:   -1,
			res: false,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 326", func(t *testing.T) {
			require.Equal(t, tc.res, isPowerOfThree(tc.n))
		})
	}
}
