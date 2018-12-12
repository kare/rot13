package rot13_test

import (
	"io"
	"os"
	"strings"

	"kkn.fi/rot13"
)

func ExampleNewReader() {
	io.Copy(os.Stdout, rot13.NewReader(strings.NewReader("nopqrs\n")))
	io.Copy(os.Stdout, rot13.NewReader(strings.NewReader("abcdef")))
	// Output:
	// abcdef
	// nopqrs
}
