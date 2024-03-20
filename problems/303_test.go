package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test303(t *testing.T) {
	array := []int{-2, 0, 3, -5, 2, -1}
	tests := []struct {
		left  int
		right int
		res   int
	}{
		{
			left:  0,
			right: 2,
			res:   1,
		},
		{
			left:  2,
			right: 5,
			res:   -1,
		},
		{
			left:  0,
			right: 5,
			res:   -3,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 303", func(t *testing.T) {
			numArray := Constructor(array)
			require.Equal(t, tc.res, numArray.SumRange(tc.left, tc.right))
		})
	}
}
