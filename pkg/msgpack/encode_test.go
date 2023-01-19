package msgpack_test

import (
	"fmt"
	"strings"
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
			input:    strings.Repeat("byte", 8),
			expected: []byte{0xd9, 0x20, 0x62, 0x79, 0x74, 0x65, 0x62, 0x79, 0x74, 0x65, 0x62, 0x79, 0x74, 0x65, 0x62, 0x79, 0x74, 0x65, 0x62, 0x79, 0x74, 0x65, 0x62, 0x79, 0x74, 0x65, 0x62, 0x79, 0x74, 0x65, 0x62, 0x79, 0x74, 0x65},
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
		{
			input:    101,
			expected: []byte{0x65},
		},
		{
			input:    144,
			expected: []byte{0xcc, 0x90},
		},
		{
			input:    1208,
			expected: []byte{0xcd, 0x04, 0xb8},
		},
		{
			input:    65599,
			expected: []byte{0xce, 0x00, 0x01, 0x00, 0x3f},
		},
		{
			input:    8589934715,
			expected: []byte{0xcf, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x7b},
		},
		{
			input:    -32,
			expected: []byte{0xe0},
		},
		{
			input:    -39,
			expected: []byte{0xd0, 0xd9},
		},
		{
			input:    -128,
			expected: []byte{0xd1, 0xff, 0x80},
		},
		{
			input:    -32768,
			expected: []byte{0xd2, 0xff, 0xff, 0x80, 0x00},
		},
		{
			input:    -2147483648,
			expected: []byte{0xd3, 0xff, 0xff, 0xff, 0xff, 0x80, 0x00, 0x00, 0x00},
		},
		{
			input:    1.256,
			expected: []byte{0xcb, 0x3f, 0xf4, 0x18, 0x93, 0x74, 0xbc, 0x6a, 0x7f},
		},
		{
			input:    []string{"name", "msgpack"},
			expected: []byte{0x92, 0xa4, 0x6e, 0x61, 0x6d, 0x65, 0xa7, 0x6d, 0x73, 0x67, 0x70, 0x61, 0x63, 0x6b},
		},
		{
			input:    []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2},
			expected: []byte{0xdc, 0x00, 0x10, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x02},
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
		{
			input:    `{"progress": 97.65}`,
			expected: []byte{0x81, 0xa8, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0xcb, 0x40, 0x58, 0x69, 0x99, 0x99, 0x99, 0x99, 0x9a},
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
