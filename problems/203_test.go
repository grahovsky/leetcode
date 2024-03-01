package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test203(t *testing.T) {
	tests := []struct {
		head *ListNode
		val  int
		res  *ListNode
	}{
		{
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7}}}},
			val:  6,
			res:  &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 7}}},
		},
		{
			head: nil,
			val:  1,
			res:  nil,
		},
		{
			head: &ListNode{Val: 7, Next: &ListNode{Val: 7, Next: &ListNode{Val: 7, Next: &ListNode{Val: 7}}}},
			val:  7,
			res:  nil,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 203", func(t *testing.T) {
			require.Equal(t, tc.res, removeElements(tc.head, tc.val))
		})
	}
}
