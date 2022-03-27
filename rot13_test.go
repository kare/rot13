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
	if _, err := io.Copy(buf, r); err != nil {
		t.Errorf("error while copying data: %v", err)
	}
	expected := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	if buf.String() != expected {
		t.Errorf("expected '%s', got '%v'", expected, buf.Bytes())
	}
}
