package tag

// tagSingnature представляет тип подписи тега.
type tagSingnature = byte

const (
	// ContentSignature используется для обозначения тега содержимого.
	ContentSignature tagSingnature = 0
	// InfoSignature используется для обозначения информационного тега.
	InfoSignature tagSingnature = 1 << (iota - 1)
	// typeSignature используется для обозначения тега типа.
	TypeSignature
)

// Tag представляет структуру тега.
type Tag struct {
	// Signature определяет тип тега (например, Content или Info).
	Signature byte
	// Length хранит длину содержимого в виде массива байтов.
	Length [4]byte
	// Data содержит содержимое тега.
	Data []byte
}

// New создает новый объект Tag с указанными параметрами.
func New(signature byte, length int32, data []byte) Tag {
	return Tag{
		Signature: signature,
		Length:    Int32ToBytes(length),
		Data:      data,
	}
}

// Bytes преобразует структуру Tag в массив байтов.
func (t Tag) Bytes() []byte {
	data := make([]byte, 0)
	data = append(data, t.Signature)
	data = append(data, t.Length[:]...)
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
