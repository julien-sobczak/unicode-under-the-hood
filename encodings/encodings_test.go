package main

import (
	"bytes"
	"testing"
)

func TestEncodeUTF32BE(t *testing.T) {
	res := EncodeUTF32BE([]uint32{
		0x0068,  // Latin Small Letter H
		0x0065,  // Latin Small Letter E
		0x0079,  // Latin Small Letter Y
		0x1F64C, // Person Raising Both Hands In Celebration Emoji
	})
	expected := []byte{
		0x00, 0x00, 0xFE, 0xFF, // BOM
		0x00, 0x00, 0x00, 0x68, // h
		0x00, 0x00, 0x00, 0x65, // e
		0x00, 0x00, 0x00, 0x79, // y
		0x00, 0x01, 0xF6, 0x4C, // emoji
	}
	if !bytes.Equal(res, expected) {
		t.Errorf("encoding error")
	}
}

func TestEncodeUTF16BE(t *testing.T) {
	res := EncodeUTF16BE([]uint32{
		0x0068,  // Latin Small Letter H
		0x0065,  // Latin Small Letter E
		0x0079,  // Latin Small Letter Y
		0x1F64C, // Person Raising Both Hands In Celebration Emoji
	})
	expected := []byte{
		0xFE, 0xFF, // BOM
		0x00, 0x68, // h
		0x00, 0x65, // e
		0x00, 0x79, // y
		0x3D, 0x4C, // emoji
	}
	if !bytes.Equal(res, expected) {
		t.Errorf("encoding error")
	}
}

func TestEncodeUTF8(t *testing.T) {
	res := EncodeUTF8([]uint32{
		0x0068,  // Latin Small Letter H
		0x0065,  // Latin Small Letter E
		0x0079,  // Latin Small Letter Y
		0x1F64C, // Person Raising Both Hands In Celebration Emoji
	})
	expected := []byte{
		0x68,                   // h
		0x65,                   // e
		0x79,                   // y
		0xF0, 0x9F, 0x99, 0x8C, // emoji
	}
	if !bytes.Equal(res, expected) {
		t.Errorf("encoding error")
	}
}
