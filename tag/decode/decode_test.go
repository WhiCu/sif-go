package decode_test

import (
	"bytes"
	"testing"

	"github.com/WhiCu/sif-go/tag/decode"
	"github.com/stretchr/testify/assert"
)

// TestNewDecoder проверяет инициализацию декодера.
func TestNewDecoder(t *testing.T) {
	r := bytes.NewReader([]byte{})
	d := decode.NewDecoder(r)
	assert.NotNil(t, d, "Декодер не должен быть nil")
}

// TestDecodeTag проверяет декодирование одного тега.
func TestDecodeTag(t *testing.T) {
	t.Run("ValidTag", func(t *testing.T) {
		// Пример тега: Signature=0x08, Length=4, Data=[0x00, 0x01, 0x02, 0x03]
		data := []byte{0x08, 0x00, 0x00, 0x00, 0x04, 0x00, 0x01, 0x02, 0x03}
		reader := bytes.NewReader(data)
		decoder := decode.NewDecoder(reader)

		result, err := decoder.DecodeTag()
		assert.NoError(t, err, "Ошибка декодирования")
		assert.Equal(t, byte(0x08), result.Signature, "Неверная сигнатура")
		assert.Equal(t, int32(4), result.Length, "Неверная длина данных")
		assert.Equal(t, []byte{0x00, 0x01, 0x02, 0x03}, result.Data, "Данные не совпадают")
	})

	t.Run("UnexpectedEOF", func(t *testing.T) {
		// Неполные данные (не хватает байтов для Length)
		data := []byte{0x08, 0x00}
		reader := bytes.NewReader(data)
		decoder := decode.NewDecoder(reader)

		_, err := decoder.DecodeTag()
		assert.ErrorIs(t, err, decode.ErrUnexpectedEOF, "Ожидалась ошибка нехватки данных")
	})
}

// TestDecodeAll проверяет чтение всех тегов из потока.
func TestDecodeAll(t *testing.T) {
	t.Run("MultipleTags", func(t *testing.T) {
		// Два тега:
		// 1. Signature=0x01, Length=2, Data=[0xAA, 0xBB]
		// 2. Signature=0x02, Length=3, Data=[0xCC, 0xDD, 0xEE]
		data := []byte{
			0x01, 0x00, 0x00, 0x00, 0x02, 0xAA, 0xBB,
			0x02, 0x00, 0x00, 0x00, 0x03, 0xCC, 0xDD, 0xEE,
		}
		reader := bytes.NewReader(data)
		decoder := decode.NewDecoder(reader)

		result, err := decoder.DecodeAll()
		assert.NoError(t, err, "Ошибка декодирования")
		assert.Len(t, result, 2, "Должно быть два тега")

		// Проверка первого тега
		assert.Equal(t, byte(0x01), result[0].Signature)
		assert.Equal(t, []byte{0xAA, 0xBB}, result[0].Data)

		// Проверка второго тега
		assert.Equal(t, byte(0x02), result[1].Signature)
		assert.Equal(t, []byte{0xCC, 0xDD, 0xEE}, result[1].Data)
	})

	t.Run("EmptyData", func(t *testing.T) {
		reader := bytes.NewReader([]byte{})
		decoder := decode.NewDecoder(reader)

		result, err := decoder.DecodeAll()
		assert.NoError(t, err, "Ошибка на пустых данных")
		assert.Empty(t, result, "Список тегов должен быть пустым")
	})
}
