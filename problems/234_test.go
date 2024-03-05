package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test234(t *testing.T) {
	tests := []struct {
		head *ListNode
		res  bool
	}{
		{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  2,
							Next: &ListNode{Val: 1},
						},
					},
				},
			},
			res: true,
		},
		{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  1,
							Next: &ListNode{Val: 1},
						},
					},
				},
			},
			res: false,
		},
		{
			head: &ListNode{
				Val: 1,
			},
			res: true,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 234", func(t *testing.T) {
			require.Equal(t, tc.res, isPalindrome(tc.head))
		})
	}
}
