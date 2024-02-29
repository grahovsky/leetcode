package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test190(t *testing.T) {
	tests := []struct {
		num uint32
		res uint32
	}{
		{
			num: 0b00000010100101000001111010011100,
			res: 0b00111001011110000010100101000000,
		},
		{
			num: 0b11111111111111111111111111111101,
			res: 0b10111111111111111111111111111111,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 190", func(t *testing.T) {
			require.Equal(t, tc.res, reverseBits(tc.num))
		})
	}
}
