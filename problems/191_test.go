package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test191(t *testing.T) {
	tests := []struct {
		num uint32
		res int
	}{
		{
			num: 0b00000010100101000001111010011100,
			res: 12,
		},
		{
			num: 0b11111111111111111111111111111101,
			res: 31,
		},
		{
			num: 0b00000000000000000000000000001011,
			res: 3,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 190", func(t *testing.T) {
			require.Equal(t, tc.res, hammingWeight(tc.num))
		})
	}
}
