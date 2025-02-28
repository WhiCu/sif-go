package tag

import (
	"reflect"
	"testing"
)

func TestNewTag(t *testing.T) {
	data := []byte{0x01, 0x02, 0x03}
	tag := New(ContentSignature, 3, data)

	if tag.Signature != ContentSignature {
		t.Errorf("Expected signature %v, got %v", ContentSignature, tag.Signature)
	}

	if !reflect.DeepEqual(tag.Data, data) {
		t.Errorf("Expected data %v, got %v", data, tag.Data)
	}

	length := BytesToInt32(tag.Length)
	if length != 3 {
		t.Errorf("Expected length %v, got %v", 3, length)
	}
}

func TestBytesConversion(t *testing.T) {
	data := []byte{0x01, 0x02, 0x03}
	tag := New(InfoSignature, 3, data)

	bytes := tag.Bytes()
	expectedLength := 1 + 4 + len(data) // 1 byte signature + 4 bytes length + data length

	if len(bytes) != expectedLength {
		t.Errorf("Expected byte slice length %v, got %v", expectedLength, len(bytes))
	}

	if bytes[0] != InfoSignature {
		t.Errorf("Expected signature %v, got %v", InfoSignature, bytes[0])
	}

	length := BytesToInt32([4]byte{bytes[1], bytes[2], bytes[3], bytes[4]})
	if length != 3 {
		t.Errorf("Expected length %v, got %v", 3, length)
	}

	if !reflect.DeepEqual(bytes[5:], data) {
		t.Errorf("Expected data %v, got %v", data, bytes[5:])
	}
}

func TestInt32ToBytesAndBack(t *testing.T) {
	value := int32(123456)
	bytes := Int32ToBytes(value)
	result := BytesToInt32(bytes)

	if result != value {
		t.Errorf("Expected %v, got %v", value, result)
	}
}
