package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test144(t *testing.T) {
	tests := []struct {
		tree   *TreeNode
		result []int
	}{
		{
			tree:   &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
			result: []int{1, 2, 3},
		},
		{
			tree:   nil,
			result: []int{},
		},
		{
			tree:   &TreeNode{Val: 1},
			result: []int{1},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 144", func(t *testing.T) {
			require.Equal(t, tc.result, preorderTraversal(tc.tree))
		})
	}
}
