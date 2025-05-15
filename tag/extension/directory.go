package extension

import (
	"github.com/WhiCu/sif-go/tag"
)

func NewDirectoryTag(tags ...tag.Tag) (*tag.Tag, error) {
	totalSize := 0
	for _, t := range tags {
		totalSize += len(t.Bytes())
	}
	data := make([]byte, totalSize)
	offset := 0
	for _, t := range tags {
		bytes := t.Bytes()
		copy(data[offset:], bytes)
		offset += len(bytes)
	}
	return tag.New(
		DirectorySignature,
		data,
	)
}

func AddTagToDirectory(dir *tag.Tag, tags ...tag.Tag) error {
	for _, t := range tags {
		bytes := t.Bytes()
		if int(dir.Length)+len(bytes) > tag.MaxLength {
			return tag.ErrDataTooLong
		}
		dir.Length += int32(len(bytes))
		dir.Data = append(dir.Data, bytes...)
	}
	return nil
}
