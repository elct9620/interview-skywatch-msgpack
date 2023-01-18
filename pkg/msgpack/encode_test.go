package msgpack_test

import (
	"testing"

	"github.com/elct9620/skywatch-msgpack/pkg/msgpack"
	"github.com/google/go-cmp/cmp"
)

func Test_Marshal(t *testing.T) {
	data, _ := msgpack.Marahal(nil)
	if !cmp.Equal([]byte{0xc0}, data) {
		t.Error(cmp.Diff([]byte{0xc0}, data))
	}
}
