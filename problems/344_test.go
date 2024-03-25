package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test344(t *testing.T) {
	tests := []struct {
		s   []byte
		res []byte
	}{
		{
			s:   []byte{'h', 'e', 'l', 'l', 'o'},
			res: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			s:   []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			res: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 344", func(t *testing.T) {
			reverseString(tc.s)
			require.Equal(t, tc.res, tc.s)
		})
	}
}
