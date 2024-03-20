package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test292(t *testing.T) {
	tests := []struct {
		n   int
		res bool
	}{
		{
			n:   4,
			res: false,
		},
		{
			n:   1,
			res: true,
		},
		{
			n:   2,
			res: true,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 292", func(t *testing.T) {
			require.Equal(t, tc.res, canWinNim(tc.n))
		})
	}
}
