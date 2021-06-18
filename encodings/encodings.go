package main

import (
	"bytes"
	"encoding/binary"
)

func EncodeUTF32BE(codepoints []uint32) []byte {
	buf := new(bytes.Buffer)

	// BOM (optional)
	binary.Write(buf, binary.BigEndian, uint32(0xFEFF))

	for _, codepoint := range codepoints {
		// Each codepoint is written as unit32
		binary.Write(buf, binary.BigEndian, codepoint)
	}

	return buf.Bytes()
}

func EncodeUTF16BE(codepoints []uint32) []byte {
	// Code is inspired by Go official implementation of module unicode/utf16
	// https://github.com/golang/go/blob/go1.16/src/unicode/utf16/utf16.go

	buf := new(bytes.Buffer)

	// BOM
	binary.Write(buf, binary.BigEndian, uint16(0xFEFF))
	for _, v := range codepoints {
		switch {
		case v < 0x10000:
			// Code points in the Basic Multilingual Plane (BMP)
			// are written as such in uint16 as they can safely
			// be stored in two bytes.
			binary.Write(buf, binary.BigEndian, uint16(v))
		case 0x10000 <= v:
			// Code points in Supplementary Planes are encoded
			// as two 16-bit code units called a surrogate pair.

			// 0x10000 is subtracted from the code point,
			// leaving a 20-bit number in the hex number range 0x00000–0xFFFFF
			r := v - 0x10000

			// The high ten bits (in the range 0x000–0x3FF) are added to 0xD800
			// to give the first 16-bit code unit or high surrogate,
			// which will be in the range 0xD800–0xDBFF.
			r1 := 0xd800 + (r>>10)&0x3ff
			binary.Write(buf, binary.BigEndian, uint8(r1))

			// The low ten bits (also in the range 0x000–0x3FF) are added
			// to 0xDC00 to give the second 16-bit code unit or low surrogate,
			// which will be in the range 0xDC00–0xDFFF.
			r2 := 0xdc00 + r&0x3ff
			binary.Write(buf, binary.BigEndian, uint8(r2))
		}
	}

	return buf.Bytes()
}

func EncodeUTF8(codepoints []uint32) []byte {
	// Code is inspired by Go official implementation of module unicode/utf8
	// https://github.com/golang/go/blob/go1.16/src/unicode/utf8/utf8.go

	buf := new(bytes.Buffer)

	// Note: The Unicode Standard neither requires nor recommends
	// the use of the BOM for UTF-8.

	for _, r := range codepoints {
		switch i := uint32(r); {

		// 1 byte for ASCII characters
		case int(r) <= 0x007F: // 127
			buf.WriteByte(byte(r)) // 0xxxxxxx

		// 2 bytes for most Latin scripts
		case i <= 0x07FF: // 2047
			buf.WriteByte(0b11000000 | byte(r>>6))         // 110xxxxx
			buf.WriteByte(0b10000000 | byte(r)&0b00111111) // 10xxxxxx

		// 3 bytes for the rest of the BMP
		case i <= 0xFFFF: // 65535
			buf.WriteByte(0b11100000 | byte(r>>12))           // 1110xxxx
			buf.WriteByte(0b10000000 | byte(r>>6)&0b00111111) // 10xxxxxx
			buf.WriteByte(0b10000000 | byte(r)&0b00111111)    // 10xxxxxx

		// 4 bytes for other planes and most emojis
		default:
			buf.WriteByte(0b11110000 | byte(r>>18))            // 11110xxx
			buf.WriteByte(0b10000000 | byte(r>>12)&0b00111111) // 10xxxxxx
			buf.WriteByte(0b10000000 | byte(r>>6)&0b00111111)  // 10xxxxxx
			buf.WriteByte(0b10000000 | byte(r)&0b00111111)     // 10xxxxxx
		}
	}

	return buf.Bytes()
}
