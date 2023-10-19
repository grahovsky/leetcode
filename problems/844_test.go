package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test844(t *testing.T) {
	tests := []struct {
		s      string
		t      string
		result bool
	}{
		{s: "ab#c", t: "ad#c", result: true},
		{s: "ab##", t: "c#d#", result: true},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 844", func(t *testing.T) {
			result := backspaceCompare(tc.s, tc.t)
			require.Equal(t, tc.result, result)
		})
	}
}
