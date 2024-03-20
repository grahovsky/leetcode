package problems

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test232(t *testing.T) {
	tests := []struct {
		res string
	}{
		{
			res: "112false",
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 232", func(t *testing.T) {
			res := ""
			q := ConstructorQ()
			q.Push(1)
			q.Push(2)
			q.Push(3)
			res += strconv.Itoa(q.Peek())
			res += strconv.Itoa(q.Pop())
			res += strconv.Itoa(q.Peek())
			res += strconv.FormatBool(q.Empty())

			require.Equal(t, tc.res, res)
		})
	}
}
