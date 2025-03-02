package sif_test

import (
	"bytes"
	"testing"

	"github.com/WhiCu/sif-go/sif"
	"github.com/WhiCu/sif-go/tag"
	"github.com/stretchr/testify/assert"
)

// TestNewSIF проверяет создание SIF с валидными данными.
func TestNewSIF(t *testing.T) {
	content := []byte("test content")
	tags := []*tag.Tag{
		tag.New(tag.ContentSignature, content),
		tag.New(tag.InfoSignature, []byte("meta1")),
	}

	s := sif.New(tags...)
	// assert.NoError(t, err, "Ошибка при создании SIF")
	assert.Equal(t, sif.SIFSignature, s.Header.Signature, "Неверная сигнатура заголовка")
	assert.Len(t, s.Tags, 2, "Должен быть два тег")
	assert.Equal(t, content, s.Tags[0].Data, "Данные контента не совпадают")
}

// TestSIFBytes проверяет сериализацию SIF в байты.
func TestSIFBytes(t *testing.T) {
	content := []byte("data")
	tags := []*tag.Tag{
		tag.New(tag.InfoSignature, []byte("info")),
		tag.New(tag.ContentSignature, content),
	}

	s := sif.New(tags...)
	bs := s.Bytes()

	// Ожидаемые байты заголовка: SIF + версия 1 + резерв [0,0,0,0]
	expectedHeader := []byte{'S', 'I', 'F', 1, 0, 0, 0, 0}
	assert.Equal(t, expectedHeader, bs[:8], "Заголовок сериализован неверно")

	// Проверка наличия тега и контента в байтах
	assert.True(t, bytes.Contains(bs, []byte("info")), "Тег не сериализован")
	assert.True(t, bytes.Contains(bs, []byte("data")), "Контент не сериализован")
}
