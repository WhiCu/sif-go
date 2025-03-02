package tag_test

import (
	"testing"

	"github.com/WhiCu/sif-go/tag"
	"github.com/stretchr/testify/assert"
)

// TestNewTag проверяет создание тега с разными параметрами.
func TestNewTag(t *testing.T) {
	t.Run("ValidContentTag", func(t *testing.T) {
		data := []byte{0x01, 0x02, 0x03}
		tg := tag.New(tag.ContentSignature, data)

		assert.Equal(t, tag.ContentSignature, tg.Signature, "Неверная сигнатура")
		assert.Equal(t, int32(3), tg.Length, "Неверная длина данных")
		assert.Equal(t, data, tg.Data, "Данные не совпадают")
	})

	t.Run("EmptyData", func(t *testing.T) {
		tg := tag.New(tag.InfoSignature, []byte{})
		assert.Equal(t, int32(0), tg.Length, "Длина должна быть 0")
		assert.Empty(t, tg.Data, "Данные должны быть пустыми")
	})
}

// TestBytes проверяет сериализацию тега в байты.
func TestBytes(t *testing.T) {
	t.Run("SimpleTag", func(t *testing.T) {
		data := []byte{0xAA, 0xBB}
		tg := tag.New(0x08, data)

		expected := []byte{0x08, 0x00, 0x00, 0x00, 0x02, 0xAA, 0xBB}
		assert.Equal(t, expected, tg.Bytes(), "Сериализация неверна")
	})

	t.Run("ZeroLength", func(t *testing.T) {
		tg := tag.New(0x01, []byte{})
		expected := []byte{0x01, 0x00, 0x00, 0x00, 0x00}
		assert.Equal(t, expected, tg.Bytes(), "Пустые данные не обработаны")
	})
}

// TestInt32ToBytes проверяет конвертацию int32 в байты.
func TestInt32ToBytes(t *testing.T) {
	testCases := []struct {
		name     string
		input    int32
		expected [4]byte
	}{
		{"Zero", 0, [4]byte{0x00, 0x00, 0x00, 0x00}},
		{"MaxInt32", 2147483647, [4]byte{0x7F, 0xFF, 0xFF, 0xFF}},
		{"Negative", -1, [4]byte{0xFF, 0xFF, 0xFF, 0xFF}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tag.Int32ToBytes(tc.input)
			assert.Equal(t, tc.expected, result, "Конвертация неверна")
		})
	}
}

// TestBytesToInt32 проверяет конвертацию байтов в int32.
func TestBytesToInt32(t *testing.T) {
	testCases := []struct {
		name     string
		input    [4]byte
		expected int32
	}{
		{"Zero", [4]byte{0x00, 0x00, 0x00, 0x00}, 0},
		{"MaxInt32", [4]byte{0x7F, 0xFF, 0xFF, 0xFF}, 2147483647},
		{"Negative", [4]byte{0xFF, 0xFF, 0xFF, 0xFF}, -1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tag.BytesToInt32(tc.input)
			assert.Equal(t, tc.expected, result, "Конвертация неверна")
		})
	}
}
