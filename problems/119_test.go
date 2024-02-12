package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test119(t *testing.T) {
	tests := []struct {
		numRows int
		result  []int
	}{
		{
			numRows: 3,
			result:  []int{1, 3, 3, 1},
		},
		{
			numRows: 4,
			result:  []int{1, 4, 6, 4, 1},
		},
		{
			numRows: 0,
			result:  []int{1},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 119", func(t *testing.T) {
			require.Equal(t, tc.result, getRow(tc.numRows))
		})
	}
}
