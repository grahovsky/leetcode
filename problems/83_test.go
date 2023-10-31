package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test83(t *testing.T) {
	tests := []struct {
		input1 *ListNode
		result *ListNode
	}{
		{
			input1: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
			result: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
		},
		{
			input1: &ListNode{
				Val:  1,
				Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}},
			},
			result: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 70", func(t *testing.T) {
			result := deleteDuplicates(tc.input1)
			require.Equal(t, tc.result, result)
		})
	}
}
