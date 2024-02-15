package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test136(t *testing.T) {
	tests := []struct {
		nums   []int
		result int
	}{
		{
			nums:   []int{2, 2, 1},
			result: 1,
		},
		{
			nums:   []int{4, 1, 2, 1, 2},
			result: 4,
		},
		{
			nums:   []int{1},
			result: 1,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 136", func(t *testing.T) {
			require.Equal(t, tc.result, singleNumber(tc.nums))
		})
	}
}
