package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test401(t *testing.T) {
	tests := []struct {
		n   int
		res []string
	}{
		{
			n:   1,
			res: []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"},
		},
		{
			n:   9,
			res: []string{},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 401", func(t *testing.T) {
			require.Equal(t, tc.res, readBinaryWatch(tc.n))
		})
	}
}
