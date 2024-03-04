package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test228(t *testing.T) {
	tests := []struct {
		nums []int
		res  []string
	}{
		{
			nums: []int{0, 1, 2, 4, 5, 7},
			res:  []string{"0->2", "4->5", "7"},
		},
		{
			nums: []int{0, 2, 3, 4, 6, 8, 9},
			res:  []string{"0", "2->4", "6", "8->9"},
		},
		{
			nums: []int{},
			res:  []string{},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 228", func(t *testing.T) {
			require.Equal(t, tc.res, summaryRanges(tc.nums))
		})
	}
}
