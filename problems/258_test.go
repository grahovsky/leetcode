package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test258(t *testing.T) {
	tests := []struct {
		n   int
		res int
	}{
		{
			n:   38,
			res: 2,
		},
		{
			n:   0,
			res: 0,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 258", func(t *testing.T) {
			require.Equal(t, tc.res, addDigits(tc.n))
		})
	}
}
