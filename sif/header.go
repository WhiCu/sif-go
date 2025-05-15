package sif

var (
	SIFSignature = [3]byte{'S', 'I', 'F'}
)

// Header представляет структуру заголовка SIF-файла.
type Header struct {
	// Signature хранит идентификатор SIF в виде трех символов.
	Signature [3]byte
	// Version указывает на версию SIF.
	Version byte
	// Reserve зарезервировано для будущего использования.
	Reserve [4]byte
}

// NewHeader создает новый заголовок с указанной версией и резервными данными.
func NewHeader(v byte, r [4]byte) Header {
	return Header{
		Signature: SIFSignature,
		Version:   v,
		Reserve:   r,
	}
}

// Bytes преобразует заголовок в массив байтов.
func (h Header) Bytes() []byte {
	data := make([]byte, 8)
	copy(data[0:3], h.Signature[:])
	data[3] = h.Version
	copy(data[4:8], h.Reserve[:])
	return data
}
