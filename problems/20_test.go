package problems

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test20(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{input: "()", expected: true},
		{input: "{}[]()", expected: true},
		{input: "{)", expected: false},
		{input: "{([[[]]])}", expected: true},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("test %s", tc.input), func(t *testing.T) {
			result := isValid(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}
