package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test27(t *testing.T) {
	tests := []struct {
		inputArr  []int
		val       int
		resultArr []int
		result    int
	}{
		{inputArr: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2, resultArr: []int{0, 1, 4, 0, 3}},
		{inputArr: []int{2}, val: 3, resultArr: []int{2}},
		{inputArr: []int{3, 2, 2, 3}, val: 3, resultArr: []int{2, 2}},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 26", func(t *testing.T) {
			result := removeElement(tc.inputArr, tc.val)
			require.Equal(t, result, len(tc.resultArr))
			require.Equal(t, tc.resultArr[:tc.result], tc.inputArr[:tc.result])
		})
	}
}
