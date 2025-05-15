package extension

import "github.com/WhiCu/sif-go/tag"

func NewNumberTag(num int32) (*tag.Tag, error) {
	numBytes := tag.Int32ToBytes(num)
	return tag.New(
		NumSignature,
		numBytes[:],
	)
}
