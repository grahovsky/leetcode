package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test21(t *testing.T) {
	input1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	input2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	returned := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  1,
			Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4}}}},
		},
	}

	tests := []struct {
		input1   *ListNode
		input2   *ListNode
		returned *ListNode
	}{
		{input1: input1, input2: input2, returned: returned},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 21", func(t *testing.T) {
			result := mergeTwoLists(tc.input1, tc.input2)
			require.Equal(t, tc.returned, result)
		})
	}
}
