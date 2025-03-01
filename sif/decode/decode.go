package decode

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/WhiCu/sif-go/sif"
	"github.com/WhiCu/sif-go/tag/decode"
)

var (
	ErrUnexpectedEOF = errors.New("unexpected end of data")
	ErrInvalidHeader = errors.New("invalid header signature")
)

func UnmarshalReader(data io.Reader, s *sif.SIF) error {
	d, err := io.ReadAll(data)
	if err != nil {
		return err
	}
	err = NewDecoder(bytes.NewReader(d)).Decode(s)
	return err
}

func Unmarshal(data []byte, s *sif.SIF) error {
	err := NewDecoder(bytes.NewReader(data)).Decode(s)
	return err
}

func (d *Decoder) DecodeHeader() (h sif.Header, err error) {
	defer func() {
		if err == io.ErrUnexpectedEOF || err == io.EOF {
			err = ErrUnexpectedEOF
		}
	}()
	buf := make([]byte, 8)

	n, err := io.ReadFull(d.r, buf)
	if err != nil {
		return h, err
	}
	if n < 7 {
		return h, ErrUnexpectedEOF
	}

	copy(h.Signature[:], buf[0:3])
	h.Version = buf[3]
	copy(h.Reserve[:], buf[4:8])

	if !bytes.Equal(h.Signature[:], sif.SIFSignature[:]) {
		return h, ErrInvalidHeader
	}

	return h, nil
}

func (d *Decoder) Decode(s *sif.SIF) (err error) {
	defer func() {
		if err == io.ErrUnexpectedEOF {
			err = ErrUnexpectedEOF
		}
	}()

	s.Header, err = d.DecodeHeader()
	if err != nil {
		return err
	}
	fmt.Println("1")
	// Считываем данные
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(d.r); err != nil {
		return err
	}

	// Считываем теги
	decoder := decode.NewDecoder(buf)
	tags, err := decoder.DecodeAll()
	if err != nil {
		return err
	}
	s.Tags = tags
	return nil
}
