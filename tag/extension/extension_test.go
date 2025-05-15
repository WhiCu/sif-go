package extension_test

import (
	"testing"

	"github.com/WhiCu/sif-go/tag"
	"github.com/WhiCu/sif-go/tag/extension"
	"github.com/stretchr/testify/assert"
)

func TestTagNum(t *testing.T) {
	t.Run("Test_tagnum", func(t *testing.T) {
		tg, _ := extension.NewNumberTag(42)

		assert.Equal(t, extension.NumSignature, tg.Signature, "Неверная сигнатура")
		assert.Equal(t, int32(4), tg.Length, "Неверная длина данных")
		assert.Equal(t, 4, len(tg.Data), "Невераня длина слайса c данными")
		assert.Equal(t, tag.Int32ToBytes(42), [4]byte(tg.Data), "Данные не совпадают")
	})
}

func TestTagCon(t *testing.T) {
	t.Run("Test_tagcontent", func(t *testing.T) {
		data := []byte("test content")
		tg, _ := extension.NewContentTag(data)

		assert.Equal(t, extension.ContentSignature, tg.Signature, "Неверная сигнатура")
		assert.Equal(t, int32(len(data)), tg.Length, "Неверная длина данных")
		assert.Equal(t, data, tg.Data, "Данные не совпадают")
	})

	t.Run("Test_tagcontentfromstring", func(t *testing.T) {
		data := "test content"
		tg, _ := extension.NewContentTagFromString(data)

		assert.Equal(t, extension.ContentSignature, tg.Signature, "Неверная сигнатура")
		assert.Equal(t, int32(len(data)), tg.Length, "Неверная длина данных")
		assert.Equal(t, []byte(data), tg.Data, "Данные не совпадают")
	})
}

// TestEdgeCases_TagCon проверяет пограничные условия для тега Content
func TestEdgeCases_TagCon(t *testing.T) {
	t.Run("Пустые данные", func(t *testing.T) {
		tg, _ := extension.NewContentTag([]byte{})
		assert.Equal(t, extension.ContentSignature, tg.Signature, "Неверная сигнатура")
		assert.Equal(t, int32(0), tg.Length, "Длина должна быть 0")
		assert.Empty(t, tg.Data, "Данные должны быть пустыми")
	})

	t.Run("Nil данные", func(t *testing.T) {
		tg, _ := extension.NewContentTag(nil)
		assert.Equal(t, extension.ContentSignature, tg.Signature, "Неверная сигнатура")
		assert.Equal(t, int32(0), tg.Length, "Длина должна быть 0 для nil")
		assert.Nil(t, tg.Data, "Данные должны быть nil")
	})
}

func TestTagInf(t *testing.T) {
	t.Run("Test_tagcontent", func(t *testing.T) {
		data := []byte("test content")
		tg, _ := extension.NewInfoTag(data)

		assert.Equal(t, extension.InfoSignature, tg.Signature, "Неверная сигнатура")
		assert.Equal(t, int32(len(data)), tg.Length, "Неверная длина данных")
		assert.Equal(t, data, tg.Data, "Данные не совпадают")
	})

	t.Run("Test_tagcontentfromstring", func(t *testing.T) {
		data := "test content"
		tg, _ := extension.NewInfoTagFromString(data)

		assert.Equal(t, extension.InfoSignature, tg.Signature, "Неверная сигнатура")
		assert.Equal(t, int32(len(data)), tg.Length, "Неверная длина данных")
		assert.Equal(t, []byte(data), tg.Data, "Данные не совпадают")
	})
}

// TestEdgeCases_TagInf проверяет пограничные условия для тега Info
func TestEdgeCases_TagInf(t *testing.T) {
	t.Run("Пустые данные", func(t *testing.T) {
		tg, _ := extension.NewInfoTag([]byte{})
		assert.Equal(t, extension.InfoSignature, tg.Signature, "Неверная сигнатура")
		assert.Equal(t, int32(0), tg.Length, "Длина должна быть 0")
		assert.Empty(t, tg.Data, "Данные должны быть пустыми")
	})

	t.Run("Nil данные", func(t *testing.T) {
		tg, _ := extension.NewInfoTag(nil)
		assert.Equal(t, extension.InfoSignature, tg.Signature, "Неверная сигнатура")
		assert.Equal(t, int32(0), tg.Length, "Длина должна быть 0 для nil")
		assert.Nil(t, tg.Data, "Данные должны быть nil")
	})
}
