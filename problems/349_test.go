package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test349(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		res   []int
	}{
		{
			nums1: []int{1, 2, 2, 1},
			nums2: []int{2, 2},
			res:   []int{2},
		},
		{
			nums1: []int{4, 9, 5},
			nums2: []int{9, 4, 9, 8, 4},
			res:   []int{4, 9},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 349", func(t *testing.T) {
			require.Equal(t, tc.res, intersection(tc.nums1, tc.nums2))
		})
	}
}
