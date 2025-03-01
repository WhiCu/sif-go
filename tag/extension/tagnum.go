package extension

import "github.com/WhiCu/sif-go/tag"

var (
	numSigature tag.TagSingnature = 8
)

func New(num int32) tag.Tag {
	numBytes := tag.Int32ToBytes(num)
	return tag.New(
		numSigature,
		4,
		numBytes[:],
	)
}
