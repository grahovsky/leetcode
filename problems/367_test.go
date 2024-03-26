package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test367(t *testing.T) {
	tests := []struct {
		num1 int
		res  bool
	}{
		{
			num1: 16,
			res:  true,
		},
		{
			num1: 14,
			res:  false,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 367", func(t *testing.T) {
			require.Equal(t, tc.res, isPerfectSquare(tc.num1))
		})
	}
}
