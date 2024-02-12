package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test118(t *testing.T) {
	tests := []struct {
		numRows int
		result  [][]int
	}{
		{
			numRows: 5,
			result:  [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
		{
			numRows: 1,
			result:  [][]int{{1}},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 118", func(t *testing.T) {
			require.Equal(t, tc.result, generate(tc.numRows))
		})
	}
}
