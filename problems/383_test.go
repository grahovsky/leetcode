package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test383(t *testing.T) {
	tests := []struct {
		a   string
		b   string
		res bool
	}{
		{
			a:   "a",
			b:   "b",
			res: false,
		},
		{
			a:   "aa",
			b:   "ab",
			res: false,
		},
		{
			a:   "aa",
			b:   "aab",
			res: true,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 383", func(t *testing.T) {
			require.Equal(t, tc.res, canConstruct(tc.a, tc.b))
		})
	}
}
