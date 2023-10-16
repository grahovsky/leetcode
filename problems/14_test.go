package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{input: []string{"flow", "flower", "flight"}, expected: "fl"},
		{input: []string{"dog", "racecar", "car"}, expected: ""},
		{input: []string{""}, expected: ""},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.expected, func(t *testing.T) {
			result := longestCommonPrefix(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}
