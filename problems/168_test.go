package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test168(t *testing.T) {
	tests := []struct {
		columnNumber int
		title        string
	}{
		{
			columnNumber: 28,
			title:        "AB",
		},
		{
			columnNumber: 1,
			title:        "A",
		},
		{
			columnNumber: 701,
			title:        "ZY",
		},
		{
			columnNumber: 52,
			title:        "AZ",
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 168", func(t *testing.T) {
			require.Equal(t, tc.title, convertToTitle(tc.columnNumber))
		})
	}
}
