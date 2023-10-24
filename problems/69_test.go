package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test69(t *testing.T) {
	tests := []struct {
		input1 int
		result int
	}{
		{input1: 4, result: 2},
		{input1: 8, result: 2},
		{input1: 10000, result: 100},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 69", func(t *testing.T) {
			result := mySqrt(tc.input1)
			require.Equal(t, tc.result, result)
		})
	}
}
