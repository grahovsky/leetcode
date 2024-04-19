package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test404(t *testing.T) {
	tests := []struct {
		tree *TreeNode
		res  int
	}{
		{
			tree: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
			},
			res: 24,
		},
		{
			tree: &TreeNode{Val: 1},
			res:  0,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 404", func(t *testing.T) {
			require.Equal(t, tc.res, sumOfLeftLeaves(tc.tree))
		})
	}
}
