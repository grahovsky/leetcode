package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test278(t *testing.T) {
	tests := []struct {
		num int
		res int
	}{
		{
			num: 6,
			res: 7,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 278", func(t *testing.T) {
			require.Equal(t, tc.res, firstBadVersion(tc.num))
		})
	}
}
