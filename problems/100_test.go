package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test100(t *testing.T) {
	tests := []struct {
		tree1  *TreeNode
		tree2  *TreeNode
		result bool
	}{
		{
			tree1:  &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
			tree2:  &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
			result: true,
		},
		{
			tree1:  &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
			tree2:  &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
			result: false,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 100", func(t *testing.T) {
			require.Equal(t, tc.result, isSameTree(tc.tree1, tc.tree2))
		})
	}
}
