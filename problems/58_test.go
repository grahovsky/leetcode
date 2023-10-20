package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test58(t *testing.T) {
	tests := []struct {
		input  string
		result int
	}{
		{input: "Hello World", result: 5},
		{input: "   fly me   to   the moon  ", result: 4},
		{input: "a ", result: 1},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 58", func(t *testing.T) {
			result := lengthOfLastWord(tc.input)
			require.Equal(t, tc.result, result)
		})
	}
}
