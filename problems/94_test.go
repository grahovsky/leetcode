package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test94(t *testing.T) {
	tests := []struct {
		input  *TreeNode
		result []int
	}{
		{
			input:  &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
			result: []int{1, 3, 2},
		},
		{
			input:  nil,
			result: []int{},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 94", func(t *testing.T) {
			require.Equal(t, tc.result, inorderTraversal(tc.input))
		})
	}
}
