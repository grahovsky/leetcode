package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test111(t *testing.T) {
	tests := []struct {
		tree   *TreeNode
		result int
	}{
		{
			tree: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
			},
			result: 2,
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
			result: 2,
		},
		{
			tree:   nil,
			result: 0,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 111", func(t *testing.T) {
			require.Equal(t, tc.result, minDepth(tc.tree))
		})
	}
}
