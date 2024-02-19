package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var tail = &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}

func Test160(t *testing.T) {
	tests := []struct {
		headA  *ListNode
		headB  *ListNode
		result *ListNode
	}{
		{
			headA:  &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: tail}},
			headB:  &ListNode{Val: 4, Next: tail},
			result: tail,
		},
		{
			headA:  tail,
			headB:  &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: tail}},
			result: tail,
		},
		{
			headA:  tail,
			headB:  &ListNode{Val: 5, Next: &ListNode{Val: 6}},
			result: nil,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 160", func(t *testing.T) {
			require.Equal(t, tc.result, getIntersectionNode(tc.headA, tc.headB))
		})
	}
}
