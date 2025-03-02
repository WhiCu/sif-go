package extension_test

import (
	"testing"

	"github.com/WhiCu/sif-go/tag"
	"github.com/WhiCu/sif-go/tag/extension"
	"github.com/stretchr/testify/assert"
)

func TestTagNum(t *testing.T) {
	t.Run("Test_tagnum", func(t *testing.T) {
		tg := extension.NewNumberTag(42)

		assert.Equal(t, extension.NumSigature, tg.Signature, "Неверная сигнатура")
		assert.Equal(t, int32(4), tg.Length, "Неверная длина данных")
		assert.Equal(t, 4, len(tg.Data), "Невераня длина слайса c данными")
		assert.Equal(t, tag.Int32ToBytes(42), [4]byte(tg.Data), "Данные не совпадают")
	})
}

func TestTagCon(t *testing.T) {
	t.Run("Test_tagcontent", func(t *testing.T) {
		data := []byte("test content")
		tg := extension.NewContentTag(data)

		assert.Equal(t, tag.ContentSignature, tg.Signature, "Неверная сигнатура")
		assert.Equal(t, int32(len(data)), tg.Length, "Неверная длина данных")
		assert.Equal(t, data, tg.Data, "Данные не совпадают")
	})

	t.Run("Test_tagcontentfromstring", func(t *testing.T) {
		data := "test content"
		tg := extension.NewContentTagFromString(data)

		assert.Equal(t, tag.ContentSignature, tg.Signature, "Неверная сигнатура")
		assert.Equal(t, int32(len(data)), tg.Length, "Неверная длина данных")
		assert.Equal(t, []byte(data), tg.Data, "Данные не совпадают")
	})
}

// TestEdgeCases_TagCon проверяет пограничные условия для тега Content
func TestEdgeCases_TagCon(t *testing.T) {
	t.Run("Пустые данные", func(t *testing.T) {
		tg := extension.NewContentTag([]byte{})
		assert.Equal(t, tag.ContentSignature, tg.Signature, "Неверная сигнатура")
		assert.Equal(t, int32(0), tg.Length, "Длина должна быть 0")
		assert.Empty(t, tg.Data, "Данные должны быть пустыми")
	})

	t.Run("Nil данные", func(t *testing.T) {
		tg := extension.NewContentTag(nil)
		assert.Equal(t, tag.ContentSignature, tg.Signature, "Неверная сигнатура")
		assert.Equal(t, int32(0), tg.Length, "Длина должна быть 0 для nil")
		assert.Nil(t, tg.Data, "Данные должны быть nil")
	})
}

func TestTagInf(t *testing.T) {
	t.Run("Test_tagcontent", func(t *testing.T) {
		data := []byte("test content")
		tg := extension.NewInfoTag(data)

		assert.Equal(t, tag.InfoSignature, tg.Signature, "Неверная сигнатура")
		assert.Equal(t, int32(len(data)), tg.Length, "Неверная длина данных")
		assert.Equal(t, data, tg.Data, "Данные не совпадают")
	})

	t.Run("Test_tagcontentfromstring", func(t *testing.T) {
		data := "test content"
		tg := extension.NewInfoTagFromString(data)

		assert.Equal(t, tag.InfoSignature, tg.Signature, "Неверная сигнатура")
		assert.Equal(t, int32(len(data)), tg.Length, "Неверная длина данных")
		assert.Equal(t, []byte(data), tg.Data, "Данные не совпадают")
	})
}

// TestEdgeCases_TagInf проверяет пограничные условия для тега Info
func TestEdgeCases_TagInf(t *testing.T) {
	t.Run("Пустые данные", func(t *testing.T) {
		tg := extension.NewInfoTag([]byte{})
		assert.Equal(t, tag.InfoSignature, tg.Signature, "Неверная сигнатура")
		assert.Equal(t, int32(0), tg.Length, "Длина должна быть 0")
		assert.Empty(t, tg.Data, "Данные должны быть пустыми")
	})

	t.Run("Nil данные", func(t *testing.T) {
		tg := extension.NewInfoTag(nil)
		assert.Equal(t, tag.InfoSignature, tg.Signature, "Неверная сигнатура")
		assert.Equal(t, int32(0), tg.Length, "Длина должна быть 0 для nil")
		assert.Nil(t, tg.Data, "Данные должны быть nil")
	})
}
