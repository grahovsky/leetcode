package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test217(t *testing.T) {
	tests := []struct {
		nums []int
		res  bool
	}{
		{
			nums: []int{3, 2, 3},
			res:  true,
		},
		{
			nums: []int{1, 2, 3, 4, 5, 6},
			res:  false,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 217", func(t *testing.T) {
			require.Equal(t, tc.res, containsDuplicate(tc.nums))
		})
	}
}
