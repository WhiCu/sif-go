package extension_test

import (
	"testing"

	"github.com/WhiCu/sif-go/tag"
	"github.com/WhiCu/sif-go/tag/extension"
	"github.com/stretchr/testify/assert"
)

type numberTestCase struct {
	description string
	input       int32
	expected    *tag.Tag
	shouldError bool
}

type contentTestCase struct {
	description string
	input       []byte
	expected    *tag.Tag
	shouldError bool
}

type stringTestCase struct {
	description string
	input       string
	expected    *tag.Tag
	shouldError bool
}

type directoryTestCase struct {
	description string
	input       []tag.Tag
	expected    *tag.Tag
	shouldError bool
}

func TestTagNum(t *testing.T) {
	cases := []numberTestCase{
		{
			description: "Нормальное число",
			input:       42,
			expected: &tag.Tag{
				Signature: extension.NumSignature,
				Length:    4,
				Data:      tag.Int32ToBytesSlice(42),
			},
			shouldError: false,
		},
		{
			description: "Ноль",
			input:       0,
			expected: &tag.Tag{
				Signature: extension.NumSignature,
				Length:    4,
				Data:      tag.Int32ToBytesSlice(0),
			},
			shouldError: false,
		},
		{
			description: "Отрицательное число",
			input:       -42,
			expected: &tag.Tag{
				Signature: extension.NumSignature,
				Length:    4,
				Data:      tag.Int32ToBytesSlice(-42),
			},
			shouldError: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.description, func(t *testing.T) {
			tg, err := extension.NewNumberTag(tc.input)
			if tc.shouldError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expected.Signature, tg.Signature)
			assert.Equal(t, tc.expected.Length, tg.Length)
			assert.Equal(t, tc.expected.Data, tg.Data)
		})
	}
}

func TestTagCon(t *testing.T) {
	byteCases := []contentTestCase{
		{
			description: "Нормальные данные",
			input:       []byte("test content"),
			expected: &tag.Tag{
				Signature: extension.ContentSignature,
				Length:    12,
				Data:      []byte("test content"),
			},
			shouldError: false,
		},
		{
			description: "Пустые данные",
			input:       []byte{},
			expected: &tag.Tag{
				Signature: extension.ContentSignature,
				Length:    0,
				Data:      []byte{},
			},
			shouldError: false,
		},
		{
			description: "Nil данные",
			input:       nil,
			expected: &tag.Tag{
				Signature: extension.ContentSignature,
				Length:    0,
				Data:      nil,
			},
			shouldError: false,
		},
	}

	stringCases := []stringTestCase{
		{
			description: "Нормальная строка",
			input:       "test content",
			expected: &tag.Tag{
				Signature: extension.ContentSignature,
				Length:    12,
				Data:      []byte("test content"),
			},
			shouldError: false,
		},
		{
			description: "Пустая строка",
			input:       "",
			expected: &tag.Tag{
				Signature: extension.ContentSignature,
				Length:    0,
				Data:      []byte{},
			},
			shouldError: false,
		},
	}

	for _, tc := range byteCases {
		t.Run(tc.description, func(t *testing.T) {
			tg, err := extension.NewContentTag(tc.input)
			if tc.shouldError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expected.Signature, tg.Signature)
			assert.Equal(t, tc.expected.Length, tg.Length)
			assert.Equal(t, tc.expected.Data, tg.Data)
		})
	}

	for _, tc := range stringCases {
		t.Run(tc.description, func(t *testing.T) {
			tg, err := extension.NewContentTagFromString(tc.input)
			if tc.shouldError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expected.Signature, tg.Signature)
			assert.Equal(t, tc.expected.Length, tg.Length)
			assert.Equal(t, tc.expected.Data, tg.Data)
		})
	}
}

func TestTagInf(t *testing.T) {
	byteCases := []contentTestCase{
		{
			description: "Нормальные данные",
			input:       []byte("test info"),
			expected: &tag.Tag{
				Signature: extension.InfoSignature,
				Length:    9,
				Data:      []byte("test info"),
			},
			shouldError: false,
		},
		{
			description: "Пустые данные",
			input:       []byte{},
			expected: &tag.Tag{
				Signature: extension.InfoSignature,
				Length:    0,
				Data:      []byte{},
			},
			shouldError: false,
		},
		{
			description: "Nil данные",
			input:       nil,
			expected: &tag.Tag{
				Signature: extension.InfoSignature,
				Length:    0,
				Data:      nil,
			},
			shouldError: false,
		},
	}

	stringCases := []stringTestCase{
		{
			description: "Нормальная строка",
			input:       "test info",
			expected: &tag.Tag{
				Signature: extension.InfoSignature,
				Length:    9,
				Data:      []byte("test info"),
			},
			shouldError: false,
		},
		{
			description: "Пустая строка",
			input:       "",
			expected: &tag.Tag{
				Signature: extension.InfoSignature,
				Length:    0,
				Data:      []byte{},
			},
			shouldError: false,
		},
	}

	for _, tc := range byteCases {
		t.Run(tc.description, func(t *testing.T) {
			tg, err := extension.NewInfoTag(tc.input)
			if tc.shouldError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expected.Signature, tg.Signature)
			assert.Equal(t, tc.expected.Length, tg.Length)
			assert.Equal(t, tc.expected.Data, tg.Data)
		})
	}

	for _, tc := range stringCases {
		t.Run(tc.description, func(t *testing.T) {
			tg, err := extension.NewInfoTagFromString(tc.input)
			if tc.shouldError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expected.Signature, tg.Signature)
			assert.Equal(t, tc.expected.Length, tg.Length)
			assert.Equal(t, tc.expected.Data, tg.Data)
		})
	}
}

func TestTagDir(t *testing.T) {
	cases := []directoryTestCase{
		{
			description: "Пустая директория",
			input:       []tag.Tag{},
			expected: &tag.Tag{
				Signature: extension.DirectorySignature,
				Length:    0,
				Data:      []byte{},
			},
			shouldError: false,
		},
		{
			description: "Директория с одним тегом",
			input: []tag.Tag{
				{
					Signature: extension.NumSignature,
					Length:    4,
					Data:      tag.Int32ToBytesSlice(42),
				},
			},
			expected: &tag.Tag{
				Signature: extension.DirectorySignature,
				Length:    9,
				Data:      append([]byte{extension.NumSignature, 0, 0, 0, 4}, tag.Int32ToBytesSlice(42)...),
			},
			shouldError: false,
		},
		{
			description: "Директория с несколькими тегами",
			input: []tag.Tag{
				{
					Signature: extension.NumSignature,
					Length:    4,
					Data:      tag.Int32ToBytesSlice(42),
				},
				{
					Signature: extension.ContentSignature,
					Length:    5,
					Data:      []byte("hello"),
				},
			},
			expected: &tag.Tag{
				Signature: extension.DirectorySignature,
				Length:    19,
				Data: append(
					append([]byte{extension.NumSignature, 0, 0, 0, 4}, tag.Int32ToBytesSlice(42)...),
					append([]byte{extension.ContentSignature, 0, 0, 0, 5}, []byte("hello")...)...,
				),
			},
			shouldError: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.description, func(t *testing.T) {
			tg, err := extension.NewDirectoryTag(tc.input...)
			if tc.shouldError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expected.Signature, tg.Signature)
			assert.Equal(t, tc.expected.Length, tg.Length)
			assert.Equal(t, tc.expected.Data, tg.Data)
		})
	}
}

func TestAddTagToDirectory(t *testing.T) {
	dir, err := extension.NewDirectoryTag()
	assert.NoError(t, err)

	// Добавляем тег в пустую директорию
	numTag, err := extension.NewNumberTag(42)
	assert.NoError(t, err)
	err = extension.AddTagToDirectory(dir, *numTag)
	assert.NoError(t, err)
	assert.Equal(t, int32(9), dir.Length)
	assert.Equal(t, append([]byte{extension.NumSignature, 0, 0, 0, 4}, tag.Int32ToBytesSlice(42)...), dir.Data)

	// Добавляем еще один тег
	contentTag, err := extension.NewContentTagFromString("hello")
	assert.NoError(t, err)
	err = extension.AddTagToDirectory(dir, *contentTag)
	assert.NoError(t, err)
	assert.Equal(t, int32(19), dir.Length)
	assert.Equal(t, append(
		append([]byte{extension.NumSignature, 0, 0, 0, 4}, tag.Int32ToBytesSlice(42)...),
		append([]byte{extension.ContentSignature, 0, 0, 0, 5}, []byte("hello")...)...,
	), dir.Data)
}

type typeTestCase struct {
	description string
	input       []byte
	expected    *tag.Tag
	shouldError bool
}

func TestTagType(t *testing.T) {
	cases := []typeTestCase{
		{
			description: "Нормальные данные типа",
			input:       []byte("text/plain"),
			expected: &tag.Tag{
				Signature: extension.TypeSignature,
				Length:    10,
				Data:      []byte("text/plain"),
			},
			shouldError: false,
		},
		{
			description: "Пустые данные типа",
			input:       []byte{},
			expected: &tag.Tag{
				Signature: extension.TypeSignature,
				Length:    0,
				Data:      []byte{},
			},
			shouldError: false,
		},
		{
			description: "Nil данные типа",
			input:       nil,
			expected: &tag.Tag{
				Signature: extension.TypeSignature,
				Length:    0,
				Data:      nil,
			},
			shouldError: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.description, func(t *testing.T) {
			tg, err := extension.NewTypeTag(tc.input)
			if tc.shouldError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expected.Signature, tg.Signature)
			assert.Equal(t, tc.expected.Length, tg.Length)
			assert.Equal(t, tc.expected.Data, tg.Data)
		})
	}
}
