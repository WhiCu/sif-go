package sif

import (
	"github.com/WhiCu/sif-go/tag"
)

// sif представляет основную структуру SIF-файла.
type SIF struct {
	// Header содержит заголовок SIF.
	Header

	// Tags содержит список тегов SIF.
	Tags []*tag.Tag

	// // Content представляет основное содержимое SIF.
	// Content tag.Tag
}

// New создает новый SIF с указанным содержимым и тегами.
func New(tags ...*tag.Tag) *SIF {
	return &SIF{
		Header: NewHeader(1, [4]byte{}),
		Tags:   tags,
	}
}

func (s *SIF) Add(t *tag.Tag) {
	s.Tags = append(s.Tags, t)
}


// Bytes преобразует структуру SIF в массив байтов.
func (s *SIF) Bytes() []byte {
	headerBytes := s.Header.Bytes()
	totalSize := len(headerBytes)
	for _, t := range s.Tags {
		totalSize += len(t.Bytes())
	}
	data := make([]byte, totalSize)
	offset := 0
	copy(data[offset:], headerBytes)
	offset += len(headerBytes)
	for _, t := range s.Tags {
		tagBytes := t.Bytes()
		copy(data[offset:], tagBytes)
		offset += len(tagBytes)
	}
	return data
}
