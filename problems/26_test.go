package problems

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test26(t *testing.T) {
	tests := []struct {
		input    []int
		returned int
		changed  []int
	}{
		{input: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, returned: 5, changed: []int{0, 1, 2, 3, 4}},
		{input: []int{1, 1, 2}, returned: 2, changed: []int{1, 2}},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("test %d", tc.returned), func(t *testing.T) {
			result := removeDuplicates(tc.input)
			require.Equal(t, tc.returned, result)
			require.Equal(t, tc.changed, tc.input[:tc.returned])
		})
	}
}
