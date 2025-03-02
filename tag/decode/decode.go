package decode

import (
	"errors"
	"io"

	"github.com/WhiCu/sif-go/tag"
)

var (
	ErrUnexpectedEOF = errors.New("unexpected end of data")
	ErrInvalidHeader = errors.New("invalid header signature")
)

func (d *Decoder) DecodeTag() (t *tag.Tag, err error) {
	defer func() {
		if err == io.ErrUnexpectedEOF {
			err = ErrUnexpectedEOF
		}
	}()
	t = new(tag.Tag)
	// Read Signature (1 byte)
	sig := make([]byte, 1)
	if _, err = io.ReadFull(d.r, sig); err != nil {
		return t, err
	}
	t.Signature = sig[0]

	// Read Length (4 bytes)
	lenBytes := make([]byte, 4)
	if _, err = io.ReadFull(d.r, lenBytes); err != nil {
		return t, err
	}
	t.Length = tag.BytesToInt32([4]byte(lenBytes))

	// Read Data
	data := make([]byte, t.Length)
	if _, err = io.ReadFull(d.r, data); err != nil {
		return t, err
	}
	t.Data = data

	return t, nil
}

func (d *Decoder) DecodeAll() ([]*tag.Tag, error) {
	var tags []*tag.Tag

	for {
		t, err := d.DecodeTag()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}
		tags = append(tags, t)
	}

	return tags, nil
}
