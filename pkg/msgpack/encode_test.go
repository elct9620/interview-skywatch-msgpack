package msgpack_test

import (
	"fmt"
	"testing"

	"github.com/elct9620/skywatch-msgpack/pkg/msgpack"
	"github.com/google/go-cmp/cmp"
)

func Test_Marshal(t *testing.T) {
	cases := []struct {
		input    any
		expected []byte
	}{
		{
			input:    nil,
			expected: []byte{0xc0},
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("test encode %v", tc.input), func(t *testing.T) {
			t.Parallel()

			data, err := msgpack.Marahal(tc.input)
			if err != nil {
				t.Error(err)
			}

			if !cmp.Equal(tc.expected, data) {
				t.Error(cmp.Diff(tc.expected, data))
			}
		})
	}
}
