package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test202(t *testing.T) {
	tests := []struct {
		num int
		res bool
	}{
		{
			num: 19,
			res: true,
		},
		{
			num: 2,
			res: false,
		},
		{
			num: 507,
			res: false,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 202", func(t *testing.T) {
			require.Equal(t, tc.res, isHappy(tc.num))
		})
	}
}
