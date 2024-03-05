package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test125(t *testing.T) {
	tests := []struct {
		input  string
		result bool
	}{
		{
			input:  "A man, a plan, a canal: Panama",
			result: true,
		},
		{
			input:  "race a car",
			result: false,
		},
		{
			input:  " ",
			result: true,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 125", func(t *testing.T) {
			require.Equal(t, tc.result, isPalindromeString(tc.input))
		})
	}
}
