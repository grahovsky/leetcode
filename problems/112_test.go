package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test112(t *testing.T) {
	tests := []struct {
		tree      *TreeNode
		targetSum int
		result    bool
	}{
		{
			tree: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}},
				Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 1}}},
			},
			targetSum: 22,
			result:    true,
		},
		{
			tree: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			targetSum: 5,
			result:    false,
		},
		{
			tree:      nil,
			targetSum: 0,
			result:    false,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 112", func(t *testing.T) {
			require.Equal(t, tc.result, hasPathSum(tc.tree, tc.targetSum))
		})
	}
}
