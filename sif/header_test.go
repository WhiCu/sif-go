package sif_test

import (
	"testing"

	"github.com/WhiCu/sif-go/sif"
	"github.com/stretchr/testify/assert"
)

// TestNewHeader проверяет создание заголовка.
func TestNewHeader(t *testing.T) {
	h := sif.NewHeader(2, [4]byte{1, 2, 3, 4})
	assert.Equal(t, sif.SIFSignature, h.Signature, "Неверная сигнатура")
	assert.Equal(t, byte(2), h.Version, "Неверная версия")
	assert.Equal(t, [4]byte{1, 2, 3, 4}, h.Reserve, "Резервные байты не совпадают")
}

// TestHeaderBytes проверяет сериализацию заголовка.
func TestHeaderBytes(t *testing.T) {
	h := sif.NewHeader(1, [4]byte{})
	bytes := h.Bytes()
	expected := []byte{'S', 'I', 'F', 1, 0, 0, 0, 0}
	assert.Equal(t, expected, bytes, "Сериализация заголовка неверна")
}
