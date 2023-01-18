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
		{
			input:    map[string]string{"name": "msgpack"},
			expected: []byte{0x81, 0xa4, 0x6e, 0x61, 0x6d, 0x65, 0xa7, 0x6d, 0x73, 0x67, 0x70, 0x61, 0x63, 0x6b},
		},
		{
			input: struct {
				Name string `msgpack:"name"`
			}{"msgpack"},
			expected: []byte{0x81, 0xa4, 0x6e, 0x61, 0x6d, 0x65, 0xa7, 0x6d, 0x73, 0x67, 0x70, 0x61, 0x63, 0x6b},
		},
		{
			input:    false,
			expected: []byte{0xc2},
		},
		{
			input:    true,
			expected: []byte{0xc3},
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

func Test_FromJSON(t *testing.T) {
	cases := []struct {
		input    string
		expected []byte
	}{
		{
			input:    `null`,
			expected: []byte{0xc0},
		},
		{
			input:    `"msgpack"`,
			expected: []byte{0xa7, 0x6d, 0x73, 0x67, 0x70, 0x61, 0x63, 0x6b},
		},
		{
			input:    `{"name":"msgpack"}`,
			expected: []byte{0x81, 0xa4, 0x6e, 0x61, 0x6d, 0x65, 0xa7, 0x6d, 0x73, 0x67, 0x70, 0x61, 0x63, 0x6b},
		},
		{
			input:    `{"encoder": {"name": "msgpack"}}`,
			expected: []byte{0x81, 0xa7, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x81, 0xa4, 0x6e, 0x61, 0x6d, 0x65, 0xa7, 0x6d, 0x73, 0x67, 0x70, 0x61, 0x63, 0x6b},
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(fmt.Sprintf("test encode from json: %s", tc.input), func(t *testing.T) {
			t.Parallel()

			data, err := msgpack.FromJSON([]byte(tc.input))
			if err != nil {
				t.Error(err)
			}

			if !cmp.Equal(tc.expected, data) {
				t.Error(cmp.Diff(tc.expected, data))
			}
		})
	}
}
