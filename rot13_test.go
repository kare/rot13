package rot13_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"kkn.fi/rot13"
)

func TestReader(t *testing.T) {
	in := strings.NewReader("NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm")
	r := rot13.NewReader(in)
	buf := bytes.NewBuffer(make([]byte, 0))
	io.Copy(buf, r)
	expected := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	if buf.String() != expected {
		t.Fatalf("expected '%s', got '%v'", expected, buf.Bytes())
	}
}
