// SPDX-License-Identifier: Unlicense OR BSD-3-Clause

package tables

import (
	"encoding/binary"
	"fmt"
)

// Code generated by binarygen from post_src.go. DO NOT EDIT

func ParsePost(src []byte) (Post, int, error) {
	var item Post
	n := 0
	if L := len(src); L < 32 {
		return item, 0, fmt.Errorf("reading Post: "+"EOF: expected length: 32, got %d", L)
	}
	_ = src[31] // early bound checking
	item.version = postVersion(binary.BigEndian.Uint32(src[0:]))
	item.italicAngle = binary.BigEndian.Uint32(src[4:])
	item.UnderlinePosition = int16(binary.BigEndian.Uint16(src[8:]))
	item.UnderlineThickness = int16(binary.BigEndian.Uint16(src[10:]))
	item.isFixedPitch = binary.BigEndian.Uint32(src[12:])
	item.memoryUsage[0] = binary.BigEndian.Uint32(src[16:])
	item.memoryUsage[1] = binary.BigEndian.Uint32(src[20:])
	item.memoryUsage[2] = binary.BigEndian.Uint32(src[24:])
	item.memoryUsage[3] = binary.BigEndian.Uint32(src[28:])
	n += 32

	{
		var (
			read int
			err  error
		)
		switch item.version {
		case postVersion10:
			item.Names, read, err = ParsePostNames10(src[32:])
		case postVersion20:
			item.Names, read, err = ParsePostNames20(src[32:])
		case postVersion30:
			item.Names, read, err = ParsePostNames30(src[32:])
		default:
			err = fmt.Errorf("unsupported PostNamesVersion %d", item.version)
		}
		if err != nil {
			return item, 0, fmt.Errorf("reading Post: %s", err)
		}
		n += read
	}
	return item, n, nil
}

func ParsePostNames10([]byte) (PostNames10, int, error) {
	var item PostNames10
	n := 0
	return item, n, nil
}

func ParsePostNames20(src []byte) (PostNames20, int, error) {
	var item PostNames20
	n := 0
	if L := len(src); L < 2 {
		return item, 0, fmt.Errorf("reading PostNames20: "+"EOF: expected length: 2, got %d", L)
	}
	arrayLengthGlyphNameIndexes := int(binary.BigEndian.Uint16(src[0:]))
	n += 2

	{

		if L := len(src); L < 2+arrayLengthGlyphNameIndexes*2 {
			return item, 0, fmt.Errorf("reading PostNames20: "+"EOF: expected length: %d, got %d", 2+arrayLengthGlyphNameIndexes*2, L)
		}

		item.GlyphNameIndexes = make([]uint16, arrayLengthGlyphNameIndexes) // allocation guarded by the previous check
		for i := range item.GlyphNameIndexes {
			item.GlyphNameIndexes[i] = binary.BigEndian.Uint16(src[2+i*2:])
		}
		n += arrayLengthGlyphNameIndexes * 2
	}
	{

		item.StringData = src[n:]
		n = len(src)
	}
	return item, n, nil
}

func ParsePostNames30([]byte) (PostNames30, int, error) {
	var item PostNames30
	n := 0
	return item, n, nil
}