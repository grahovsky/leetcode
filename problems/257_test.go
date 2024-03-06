package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test257(t *testing.T) {
	tests := []struct {
		root *TreeNode
		res  []string
	}{
		{
			root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			res:  []string{"1->2", "1->3"},
		},
		{
			root: &TreeNode{Val: 1},
			res:  []string{"1"},
		},
		{
			root: &TreeNode{Val: 1, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 8}}},
			res:  []string{"1->5->8"},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 257", func(t *testing.T) {
			require.Equal(t, tc.res, binaryTreePaths(tc.root))
		})
	}
}
