package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test222(t *testing.T) {
	tests := []struct {
		tree *TreeNode
		res  int
	}{
		{
			tree: &TreeNode{Left: &TreeNode{}, Right: &TreeNode{}},
			res:  3,
		},
		{
			tree: &TreeNode{},
			res:  1,
		},
		{
			tree: nil,
			res:  0,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 222", func(t *testing.T) {
			require.Equal(t, tc.res, countNodes(tc.tree))
		})
	}
}
