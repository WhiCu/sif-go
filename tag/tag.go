package tag

import (
	"errors"
	"math"
)

const (
	MaxLength = math.MaxInt32
)

var (
	ErrDataTooLong = errors.New("data too long")
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
//   - ContentSignature (1) - тип подписи тега содержимого,
//   - InfoSignature (2) - тип подписи информационного тега,
//   - TypeSignature (4) - тип подписи тега типа.
//
// data - массив байтов, содержащий данные тега.
// TODO: придумай как сделать это через интерфейс Reader
func New(signature byte, data []byte) (tag *Tag, err error) {
	lc := len(data)
	if lc > MaxLength {
		return nil, ErrDataTooLong
	}
	return &Tag{
		Signature: signature,
		Length:    int32(lc),
		Data:      data,
	}, nil
}
func MustNew(signature byte, data []byte) (tag *Tag) {
	lc := len(data)
	if lc > math.MaxInt32 {
		panic(ErrDataTooLong)
	}
	return &Tag{
		Signature: signature,
		Length:    int32(lc),
		Data:      data,
	}
}

// Bytes преобразует структуру Tag в массив байтов.
func (t Tag) Bytes() []byte {
	data := make([]byte, 1+4+int(t.Length))
	data[0] = t.Signature
	lenBytes := Int32ToBytes(t.Length)
	copy(data[1:5], lenBytes[:])
	copy(data[5:], t.Data)
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
