package extension

import (
	"github.com/WhiCu/sif-go/tag"
)

// NewContentTag creates a new content tag.
//
// If the length of the data exceeds the maximum value that can be represented
// by an int32 (0xFFFFFFFF), the length is set to 0xFFFFFFFF and the data is
// truncated.
func NewContentTag(data []byte) *tag.Tag {
	return tag.New(
		tag.ContentSignature,
		data,
	)
}

func NewContentTagFromString(s string) *tag.Tag {
	return NewContentTag([]byte(s))
}
