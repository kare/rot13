package rot13_test

import (
	"io"
	"log"
	"os"
	"strings"

	"kkn.fi/rot13"
)

func newReader(s string) io.Reader {
	in := strings.NewReader(s)
	return rot13.NewReader(in)
}

func ExampleNewReader() {
	if _, err := io.Copy(os.Stdout, newReader("nopqrs\n")); err != nil {
		log.Fatalf("error while connecting stdin to stdout: %v", err)
	}
	if _, err := io.Copy(os.Stdout, newReader("abcdef")); err != nil {
		log.Fatalf("error while connecting stdin to stdout: %v", err)
	}
	// Output:
	// abcdef
	// nopqrs
}
