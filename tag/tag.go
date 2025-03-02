package tag

import "math"

// tagSingnature представляет тип подписи тега.
type TagSingnature = byte

const (
	// ContentSignature используется для обозначения тега содержимого.
	ContentSignature TagSingnature = 1 << iota
	// InfoSignature используется для обозначения информационного тега.
	InfoSignature
	// typeSignature используется для обозначения тега типа.
	TypeSignature
)

// Tag представляет структуру тега.
type Tag struct {
	// Signature определяет тип тега (например, Content или Info).
	Signature byte
	// Length хранит длину содержимого в виде массива байтов.
	Length int32
	// Data содержит содержимое тега.
	Data []byte
}

// New создает новый объект Tag с указанными параметрами.
//
// signature - тип подписи тега. Он может быть одним из следующих:
//  - ContentSignature (1) - тип подписи тега содержимого,
//  - InfoSignature (2) - тип подписи информационного тега,
//  - TypeSignature (4) - тип подписи тега типа.
//
// data - массив байтов, содержащий данные тега.
func New(signature byte, data []byte) *Tag {
	lc := len(data)
	if lc > math.MaxInt32 {
		//TODO: добавить обработку
		panic("data too long")
	}
	return &Tag{
		Signature: signature,
		Length:    int32(lc),
		Data:      data,
	}
}

// Bytes преобразует структуру Tag в массив байтов.
func (t Tag) Bytes() []byte {
	data := make([]byte, 0)
	data = append(data, t.Signature)
	lenBytes := Int32ToBytes(t.Length)
	data = append(data, lenBytes[:]...)
	data = append(data, t.Data...)
	return data
}

// Int32ToBytes преобразует 32-битное целое число в массив из 4 байтов.
func Int32ToBytes(i int32) [4]byte {
	return [4]byte{
		byte(i >> 24),
		byte(i >> 16),
		byte(i >> 8),
		byte(i),
	}
}

// BytesToInt32 преобразует массив из 4 байтов обратно в 32-битное целое число.
func BytesToInt32(b [4]byte) int32 {
	return int32(b[0])<<24 | int32(b[1])<<16 | int32(b[2])<<8 | int32(b[3])
}
