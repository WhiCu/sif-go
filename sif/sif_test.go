package sif

import (
	"reflect"
	"testing"

	"github.com/WhiCu/sif-go/tag"
)

func TestNewSIF(t *testing.T) {
	content := []byte{0x01, 0x02, 0x03}
	tag1 := tag.New(tag.InfoSignature, 2, []byte{0x04, 0x05})
	tag2 := tag.New(tag.TypeSignature, 3, []byte{0x06, 0x07, 0x08})

	sifFile := New(content, tag1, tag2)

	if !reflect.DeepEqual(sifFile.Content.Data, content) {
		t.Errorf("Expected content %v, got %v", content, sifFile.Content.Data)
	}

	if len(sifFile.Tags) != 2 {
		t.Errorf("Expected 2 tags, got %v", len(sifFile.Tags))
	}

	if !reflect.DeepEqual(sifFile.Tags[0], tag1) || !reflect.DeepEqual(sifFile.Tags[1], tag2) {
		t.Errorf("Tags do not match the expected values")
	}
}

func TestSIFBytes(t *testing.T) {
	content := []byte{0x01, 0x02, 0x03}
	tag1 := tag.New(tag.InfoSignature, 2, []byte{0x04, 0x05})
	tag2 := tag.New(tag.TypeSignature, 3, []byte{0x06, 0x07, 0x08})

	sifFile := New(content, tag1, tag2)
	bytes := sifFile.Bytes()

	headerBytes := sifFile.Header.Bytes()
	tag1Bytes := tag1.Bytes()
	tag2Bytes := tag2.Bytes()
	contentBytes := sifFile.Content.Bytes()

	expectedBytes := append(headerBytes, tag1Bytes...)
	expectedBytes = append(expectedBytes, tag2Bytes...)
	expectedBytes = append(expectedBytes, contentBytes...)

	if !reflect.DeepEqual(bytes, expectedBytes) {
		t.Errorf("Expected bytes %v, got %v", expectedBytes, bytes)
	}
}
