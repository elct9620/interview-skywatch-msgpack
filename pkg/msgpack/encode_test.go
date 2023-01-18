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
		{
			input:    "msgpack",
			expected: []byte{0xa7, 0x6d, 0x73, 0x67, 0x70, 0x61, 0x63, 0x6b},
		},
	}

	for _, tc := range cases {
		tc := tc

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
