package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test101(t *testing.T) {
	tests := []struct {
		tree   *TreeNode
		result bool
	}{
		{
			tree: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}},
				Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
			},
			result: true,
		},
		{
			tree: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}},
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}},
			},
			result: false,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 100", func(t *testing.T) {
			require.Equal(t, tc.result, isSymmetric(tc.tree))
		})
	}
}
