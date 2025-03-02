package extension

import "github.com/WhiCu/sif-go/tag"

func NewInfoTag(data []byte) *tag.Tag {
	return tag.New(
		tag.InfoSignature,
		data)
}

func NewInfoTagFromString(data string) *tag.Tag {
	return NewInfoTag([]byte(data))
}
