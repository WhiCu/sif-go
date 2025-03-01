package sif

import (
	"github.com/WhiCu/sif-go/tag"
)

// sif представляет основную структуру SIF-файла.
type SIF struct {
	// Header содержит заголовок SIF.
	Header

	// Tags содержит список тегов SIF.
	Tags []tag.Tag

	// // Content представляет основное содержимое SIF.
	// Content tag.Tag
}

// New создает новый SIF с указанным содержимым и тегами.
func New(tags ...tag.Tag) (*SIF, error) {
	// lc := len(Content)
	// if lc > lenInt32 {
	// 	return nil, ErrContentTooLong
	// }
	return &SIF{
		Header: NewHeader(1, [4]byte{}),
		// Content: tag.New(
		// 	tag.ContentSignature,
		// 	int32(lc),
		// 	Content),
		Tags: tags,
	}, nil
}

// Bytes преобразует структуру SIF в массив байтов.
func (s *SIF) Bytes() []byte {
	data := make([]byte, 0)

	// Добавление байтов заголовка.
	data = append(data, s.Header.Bytes()...)

	// Добавление байтов каждого тега.
	for _, t := range s.Tags {
		data = append(data, t.Bytes()...)
	}

	// // Добавление байтов основного содержимого.
	// data = append(data, s.Content.Bytes()...)
	return data
}
