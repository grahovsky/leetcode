package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test169(t *testing.T) {
	tests := []struct {
		nums []int
		res  int
	}{
		{
			nums: []int{3, 2, 3},
			res:  3,
		},
		{
			nums: []int{2, 2, 1, 1, 1, 2, 2},
			res:  2,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 169", func(t *testing.T) {
			require.Equal(t, tc.res, majorityElement(tc.nums))
		})
	}
}
