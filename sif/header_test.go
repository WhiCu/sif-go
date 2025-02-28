package sif

import (
	"reflect"
	"testing"
)

func TestNewHeader(t *testing.T) {
	reserve := [3]byte{0x00, 0x01, 0x02}
	header := NewHeader(1, reserve)

	if header.Version != 1 {
		t.Errorf("Expected version %v, got %v", 1, header.Version)
	}

	if !reflect.DeepEqual(header.Reserve, reserve) {
		t.Errorf("Expected reserve %v, got %v", reserve, header.Reserve)
	}

	expectedSignature := [3]byte{'S', 'I', 'F'}
	if !reflect.DeepEqual(header.Signature, expectedSignature) {
		t.Errorf("Expected signature %v, got %v", expectedSignature, header.Signature)
	}
}

func TestHeaderBytes(t *testing.T) {
	reserve := [3]byte{0x00, 0x01, 0x02}
	header := NewHeader(1, reserve)

	bytes := header.Bytes()
	expectedLength := 3 + 1 + 3 // 3 bytes signature + 1 byte version + 3 bytes reserve

	if len(bytes) != expectedLength {
		t.Errorf("Expected byte slice length %v, got %v", expectedLength, len(bytes))
	}

	expectedBytes := append([]byte{'S', 'I', 'F'}, 1)
	expectedBytes = append(expectedBytes, reserve[:]...)
	if !reflect.DeepEqual(bytes, expectedBytes) {
		t.Errorf("Expected bytes %v, got %v", expectedBytes, bytes)
	}
}
