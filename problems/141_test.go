package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test141(t *testing.T) {
	tests := []struct {
		list   *ListNode
		result bool
	}{
		{
			list:   &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0}}},
			result: true,
		},
		{
			list:   &ListNode{Val: 3},
			result: false,
		},
		{
			list:   &ListNode{},
			result: false,
		},
	}
	for _, tc := range tests {
		tc := tc
		if tc.result == true {
			tc.list.Next = tc.list
		}
		t.Run("test 141", func(t *testing.T) {
			require.Equal(t, tc.result, hasCycle(tc.list))
		})
	}
}
