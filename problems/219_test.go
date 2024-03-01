package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test219(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		res  bool
	}{
		{
			nums: []int{3, 2, 3},
			k:    2,
			res:  true,
		},
		{
			nums: []int{1, 2, 3, 4, 5, 6},
			k:    5,
			res:  false,
		},
		{
			nums: []int{1, 2, 3, 1, 5, 6},
			k:    3,
			res:  true,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 219", func(t *testing.T) {
			require.Equal(t, tc.res, containsNearbyDuplicate(tc.nums, tc.k))
		})
	}
}
