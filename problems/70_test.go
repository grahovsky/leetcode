package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test70(t *testing.T) {
	tests := []struct {
		input1 int
		result int
	}{
		{input1: 2, result: 2},
		{input1: 3, result: 3},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 70", func(t *testing.T) {
			result := climbStairs(tc.input1)
			require.Equal(t, tc.result, result)
		})
	}
}
