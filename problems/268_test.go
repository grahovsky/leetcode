package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test268(t *testing.T) {
	tests := []struct {
		nums []int
		res  int
	}{
		{
			nums: []int{3, 0, 1},
			res:  2,
		},
		{
			nums: []int{0, 1},
			res:  2,
		},
		{
			nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			res:  8,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 268", func(t *testing.T) {
			require.Equal(t, tc.res, missingNumber(tc.nums))
		})
	}
}
