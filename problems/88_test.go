package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test88(t *testing.T) {
	tests := []struct {
		nums1  []int
		m      int
		nums2  []int
		n      int
		result []int
	}{
		{nums1: []int{1, 2, 3, 0, 0, 0}, m: 3, nums2: []int{2, 5, 6}, n: 3, result: []int{1, 2, 2, 3, 5, 6}},
		{nums1: []int{0}, m: 0, nums2: []int{1}, n: 1, result: []int{1}},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 88", func(t *testing.T) {
			merge(tc.nums1, tc.m, tc.nums2, tc.n)
			require.Equal(t, tc.result, tc.nums1)
		})
	}
}
