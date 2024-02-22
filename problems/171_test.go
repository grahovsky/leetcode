package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test171(t *testing.T) {
	tests := []struct {
		title  string
		number int
	}{
		{
			title:  "AB",
			number: 28,
		},
		{
			title:  "A",
			number: 1,
		},
		{
			title:  "ZY",
			number: 701,
		},
		{
			title:  "AZ",
			number: 52,
		},
		{
			title:  "NTP",
			number: 10000,
		},
		{
			title:  "FXSHRXW",
			number: 2147483647,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 171", func(t *testing.T) {
			require.Equal(t, tc.number, titleToNumber(tc.title))
		})
	}
}
