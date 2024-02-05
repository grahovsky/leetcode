package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test108(t *testing.T) {
	tests := []struct {
		sortedArray []int
		result      *TreeNode
	}{
		{
			sortedArray: []int{-10, -3, 0, 5, 9},
			result: &TreeNode{
				Val:   0,
				Left:  &TreeNode{Val: -3, Left: &TreeNode{Val: -10}},
				Right: &TreeNode{Val: 9, Left: &TreeNode{Val: 5}},
			},
		},
		{
			sortedArray: []int{1, 3},
			result: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 1},
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 108", func(t *testing.T) {
			require.Equal(t, tc.result, sortedArrayToBST(tc.sortedArray))
		})
	}
}
