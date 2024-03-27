package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test374(t *testing.T) {
	tests := []struct {
		n   int
		res int
	}{
		{
			n:   10,
			res: 6,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 373", func(t *testing.T) {
			require.Equal(t, tc.res, guessNumber(tc.n))
		})
	}
}
