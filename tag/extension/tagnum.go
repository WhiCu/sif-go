package extension

import "github.com/WhiCu/sif-go/tag"

var (
	NumSigature tag.TagSingnature = 8
)

func NewNumberTag(num int32) *tag.Tag {
	numBytes := tag.Int32ToBytes(num)
	return tag.New(
		NumSigature,
		numBytes[:],
	)
}
