package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test104(t *testing.T) {
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
			result: 3,
		},
		{
			tree: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			},
			result: 2,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 100", func(t *testing.T) {
			require.Equal(t, tc.result, maxDepth(tc.tree))
		})
	}
}
