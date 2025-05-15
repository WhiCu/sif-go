package extension

import "github.com/WhiCu/sif-go/tag"

func NewInfoTag(data []byte) (*tag.Tag, error) {
	return tag.New(
		InfoSignature,
		data)
}

func NewInfoTagFromString(data string) (*tag.Tag, error) {
	return NewInfoTag([]byte(data))
}
