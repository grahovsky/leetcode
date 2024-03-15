package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test283(t *testing.T) {
	tests := []struct {
		nums []int
		res  []int
	}{
		{
			nums: []int{0, 1, 0, 3, 12},
			res:  []int{1, 3, 12, 0, 0},
		},
		{
			nums: []int{0},
			res:  []int{0},
		},
		{
			nums: []int{0, 1, 0, 2, 0},
			res:  []int{1, 2, 0, 0, 0},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 283", func(t *testing.T) {
			moveZeroes(tc.nums)
			require.Equal(t, tc.res, tc.nums)
		})
	}
}
