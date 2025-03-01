package decode

import "io"

type Decoder struct {
	r io.Reader
}

func NewDecoder(r io.Reader) *Decoder {
	d := new(Decoder)
	d.r = r
	return d
}
