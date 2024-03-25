package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test342(t *testing.T) {
	tests := []struct {
		n   int
		res bool
	}{
		{
			n:   16,
			res: true,
		},
		{
			n:   5,
			res: false,
		},
		{
			n:   1,
			res: true,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 342", func(t *testing.T) {
			require.Equal(t, tc.res, isPowerOfFour(tc.n))
		})
	}
}
