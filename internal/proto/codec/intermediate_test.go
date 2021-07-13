package codec

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gotd/td/bin"
)

func TestIntermediate(t *testing.T) {
	t.Run("Ok", func(t *testing.T) {
		msg := bytes.Repeat([]byte{1, 2, 3}, 100)
		buf := new(bytes.Buffer)
		if err := writeIntermediate(buf, &bin.Buffer{Buf: msg}); err != nil {
			t.Fatal(err)
		}
		var out bin.Buffer
		if err := readIntermediate(buf, &out, false); err != nil {
			t.Fatal(err)
		}
		require.Equal(t, msg, out.Buf)
	})
	t.Run("BigMessage", func(t *testing.T) {
		codec := Intermediate{}
		t.Run("Read", func(t *testing.T) {
			var b bin.Buffer
			b.PutInt((1024+512)*1024 + 10)

			var out bin.Buffer
			if err := codec.Read(&b, &out); !errors.Is(err, invalidMsgLenErr{}) {
				t.Error(err)
			}
		})
		t.Run("Write", func(t *testing.T) {
			buf := make([]byte, 1024*1024*2)

			if err := codec.Write(nil, &bin.Buffer{Buf: buf}); !errors.Is(err, invalidMsgLenErr{}) {
				t.Error(err)
			}
		})
	})
}
