package rot13

import "io"

// Reader reads either rot13 encoded or decoded data and produces the opposite of input.
type Reader struct {
	r io.Reader
}

// rot13 encodes or decodes given c with rot13.
func rot13(c byte) byte {
	const rot = 13
	if c >= 'a' && c <= 'z' {
		if c+rot <= 'z' {
			return c + rot
		}
		return c - rot
	} else if c >= 'A' && c <= 'Z' {
		if c+rot <= 'Z' {
			return c + rot
		}
		return c - rot
	} else {
		return c
	}
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	for i, k := range p {
		p[i] = rot13(k)
	}
	return n, err
}

// NewReader creates a new rot13.Reader that reads from given r.
func NewReader(r io.Reader) *Reader {
	return &Reader{
		r: r,
	}
}
