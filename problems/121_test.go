package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test121(t *testing.T) {
	tests := []struct {
		prices []int
		result int
	}{
		{
			prices: []int{7, 6, 4, 3, 1},
			result: 0,
		},
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			result: 5,
		},
		{
			prices: []int{100, 200, 1, 90},
			result: 100,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 119", func(t *testing.T) {
			require.Equal(t, tc.result, maxProfit(tc.prices))
		})
	}
}
