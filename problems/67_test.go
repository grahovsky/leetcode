package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test67(t *testing.T) {
	tests := []struct {
		input1 string
		input2 string
		result string
	}{
		{input1: "11", input2: "1", result: "100"},
		{input1: "1010", input2: "1011", result: "10101"},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 67", func(t *testing.T) {
			result := addBinary(tc.input1, tc.input2)
			require.Equal(t, tc.result, result)
		})
	}
}
