package decode_test

import (
	"bytes"
	"testing"

	"github.com/WhiCu/sif-go/sif"
	"github.com/WhiCu/sif-go/sif/decode"
	"github.com/stretchr/testify/assert"
)

// TestNewDecoder проверяет инициализацию декодера.
func TestNewDecoder(t *testing.T) {
	r := bytes.NewReader([]byte{})
	d := decode.NewDecoder(r)
	assert.NotNil(t, d, "Декодер не должен быть nil")
}

// TestDecodeHeader проверяет декодирование заголовка.
func TestDecodeHeader(t *testing.T) {
	t.Run("ValidHeader", func(t *testing.T) {
		data := []byte{'S', 'I', 'F', 1, 0, 0, 0, 0} // Корректный заголовок
		reader := bytes.NewReader(data)
		decoder := decode.NewDecoder(reader)

		h, err := decoder.DecodeHeader()
		assert.NoError(t, err, "Ошибка декодирования заголовка")
		assert.Equal(t, sif.SIFSignature, h.Signature, "Неверная сигнатура")
		assert.Equal(t, byte(1), h.Version, "Неверная версия")
	})

	t.Run("InvalidSignature", func(t *testing.T) {
		data := []byte{'X', 'Y', 'Z', 1, 0, 0, 0, 0} // Неверная сигнатура
		reader := bytes.NewReader(data)
		decoder := decode.NewDecoder(reader)

		_, err := decoder.DecodeHeader()
		assert.ErrorIs(t, err, decode.ErrInvalidHeader, "Ожидалась ошибка неверной сигнатуры")
	})

	t.Run("UnexpectedEOF", func(t *testing.T) {
		data := []byte{'S', 'I', 'F'} // Неполный заголовок
		reader := bytes.NewReader(data)
		decoder := decode.NewDecoder(reader)

		_, err := decoder.DecodeHeader()
		assert.ErrorIs(t, err, decode.ErrUnexpectedEOF, "Ожидалась ошибка нехватки данных")
	})
}

// TestDecode проверяет декодирование полного SIF-файла.
func TestDecode(t *testing.T) {
	t.Run("ValidSIF", func(t *testing.T) {
		// Заголовок + тег Info + контент
		data := []byte{
			'S', 'I', 'F', 1, 0, 0, 0, 0, // Заголовок
			0x01, 0x00, 0x00, 0x00, 0x03, 'f', 'o', 'o', // Тег Info (сигнатура 0x01)
			0x00, 0x00, 0x00, 0x00, 0x04, 'd', 'a', 't', 'a', // Контент (сигнатура 0x00)
		}
		var s sif.SIF
		decoder := decode.NewDecoder(bytes.NewReader(data))
		err := decoder.Decode(&s)

		assert.NoError(t, err, "Ошибка декодирования")
		assert.Equal(t, byte(1), s.Header.Version, "Неверная версия заголовка")
		assert.Len(t, s.Tags, 2, "Должен быть два тег")
		assert.Equal(t, []byte("data"), s.Tags[1].Data, "Контент не совпадает")
	})
}

// TestUnmarshal проверяет функции Unmarshal и UnmarshalReader.
func TestUnmarshal(t *testing.T) {
	t.Run("ValidData", func(t *testing.T) {
		data := []byte{
			'S', 'I', 'F', 1, 0, 0, 0, 0,
			0x01, 0x00, 0x00, 0x00, 0x03, 'f', 'o', 'o',
			0x00, 0x00, 0x00, 0x00, 0x04, 'd', 'a', 't', 'a',
		}
		var s sif.SIF
		err := decode.Unmarshal(data, &s)
		assert.NoError(t, err, "Unmarshal вернул ошибку")
		assert.Equal(t, []byte("data"), s.Tags[1].Data, "Контент не совпадает")
	})

	t.Run("EmptyData", func(t *testing.T) {
		var s sif.SIF
		err := decode.Unmarshal([]byte{}, &s)
		assert.ErrorIs(t, err, decode.ErrUnexpectedEOF, "Ожидалась ошибка на пустых данных")
	})
}
