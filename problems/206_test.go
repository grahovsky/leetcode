package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test206(t *testing.T) {
	tests := []struct {
		list    *ListNode
		reverse *ListNode
	}{
		{
			list: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
		},
	}
	for _, tc := range tests {
		tc := tc

		t.Run("test 206", func(t *testing.T) {
			require.Equal(t, reverseList(tc.list), reverseList(tc.list))
		})
	}
}
