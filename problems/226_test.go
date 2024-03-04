package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test226(t *testing.T) {
	tests := []struct {
		tree *TreeNode
		res  *TreeNode
	}{
		{
			tree: &TreeNode{
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			res: &TreeNode{
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 2},
			},
		},
		{
			tree: &TreeNode{
				Left:  &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
				Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 5}},
			},
			res: &TreeNode{
				Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 5}},
				Right: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
			},
		},
		{
			tree: &TreeNode{},
			res:  &TreeNode{},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 226", func(t *testing.T) {
			require.Equal(t, tc.res, invertTree(tc.tree))
		})
	}
}
