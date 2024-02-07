package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test110(t *testing.T) {
	tests := []struct {
		tree   *TreeNode
		result bool
	}{
		{
			tree: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
			},
			result: true,
		},
		{
			tree: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 4},
						Right: &TreeNode{Val: 4},
					},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{Val: 2},
			},
			result: false,
		},
		{
			tree:   nil,
			result: true,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 110", func(t *testing.T) {
			require.Equal(t, tc.result, isBalanced(tc.tree))
		})
	}
}
