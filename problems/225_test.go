package problems

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test225(t *testing.T) {
	tests := []struct {
		res string
	}{
		{
			res: "332false",
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 225", func(t *testing.T) {
			res := ""
			stack := StackConstructor()
			stack.Push(1)
			stack.Push(2)
			stack.Push(3)
			res += strconv.Itoa(stack.Top())
			res += strconv.Itoa(stack.Pop())
			res += strconv.Itoa(stack.Top())
			res += strconv.FormatBool(stack.Empty())

			require.Equal(t, tc.res, res)
		})
	}
}
