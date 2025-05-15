package extension

// tagSingnature представляет тип подписи тега.
type TagSingnature = byte

const (
	// ContentSignature используется для обозначения тега содержимого.
	ContentSignature TagSingnature = 1 << iota
	// InfoSignature используется для обозначения информационного тега.
	InfoSignature
	// typeSignature используется для обозначения тега типа.
	TypeSignature
	// DirectorySignature используется для обозначения тега директория.
	DirectorySignature
	// NumSignature используется для обозначения тега номера.
	NumSignature
)
