package decode

import (
	"bufio"
	"bytes"
	"errors"
	"io"

	"github.com/WhiCu/sif-go/sif"
	"github.com/WhiCu/sif-go/tag/decode"
)

var (
	ErrUnexpectedEOF = errors.New("unexpected end of data")
	ErrInvalidHeader = errors.New("invalid header signature")
)

func UnmarshalReader(data io.Reader, s *sif.SIF) (err error) {
	err = NewDecoder(bufio.NewReader(data)).Decode(s)
	return err
}

func Unmarshal(data []byte, s *sif.SIF) (err error) {
	return UnmarshalReader(bytes.NewReader(data), s)
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
	// Считываем данные
	buf := bufio.NewReader(d.r)
	// Считываем теги
	decoder := decode.NewDecoder(buf)
	tags, err := decoder.DecodeAll()
	if err != nil {
		return err
	}
	s.Tags = tags
	return nil
}
