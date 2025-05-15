package extension

import "github.com/WhiCu/sif-go/tag"

// NewTypeTag creates a new type tag.
//
// The type tag is used to specify the MIME type of the content.
func NewTypeTag(data []byte) (*tag.Tag, error) {
	return tag.New(
		TypeSignature,
		data,
	)
}

// NewTypeTagFromString creates a new type tag from a string.
func NewTypeTagFromString(s string) (*tag.Tag, error) {
	return NewTypeTag([]byte(s))
} 