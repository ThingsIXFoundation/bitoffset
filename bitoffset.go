// Copyright 2022 Stichting ThingsIX Foundation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package bitoffset

func Int32(buf []byte, pos uint32, len uint8) int32 {
	ui := int32(Uint32(buf, pos, len))
	if len == 32 || (ui&(1<<(len-1)) == 0) {
		return ui
	}
	return ui | ((^0) << len)
}

func Uint32(buf []byte, pos uint32, len uint8) uint32 {

	var ret uint32 = 0

	for i := pos; i < pos+uint32(len); i++ {
		ret = ret << 1                                  // Shift the return value 1 position because there's more bits to add
		ret |= uint32(((buf[i/8]) >> (7 - i%8)) & 0b01) // get the right byte (buf[i/8])
		// bit 0 is the first in the byte, we have to shift it 7 places over to get it to the last position
		// last mask it with 0b01 to make sure we only get that byte
	}

	return ret
}

func Uint8(buf []byte, pos uint32, len uint8) uint8 {
	var shift_offset uint8 = uint8(pos%8 + uint32(len))
	if shift_offset <= 8 {
		return (buf[pos/8] >> (8 - shift_offset)) & ((1 << len) - 1)
	} else {
		return ((buf[pos/8] << (shift_offset % 8)) | (buf[pos/8+1] >> (16 - shift_offset))) & ((1 << len) - 1)
	}
}

func SetUint8(buf []byte, pos uint32, len uint8, val uint8) {
	// pos and len are in bit position! not bytes!

	// Note that because we are setting a uint8_t we know that we never have to use more than 2 bytes.
	var shift_offset uint8 = uint8(pos%8 + uint32(len))
	var mask uint8 = ((1 << len) - 1)
	val = val & mask // mask value so we don't have any left over bits
	if shift_offset <= 8 {
		buf[pos/8] = (buf[pos/8] & ^(mask << (8 - shift_offset))) | (val << (8 - shift_offset))
	} else {
		buf[pos/8] = (buf[pos/8] & ^(mask >> (shift_offset % 8))) | (val >> (shift_offset % 8))
		buf[pos/8+1] = (buf[pos/8+1] & ^(mask << (16 - shift_offset))) | (val << (16 - shift_offset))
	}

	// pos 0 len 1
	// val =  1 0 1 0 1 0 1 1
	// mask = 0 0 0 0 0 0 0 1
	// val  = 0 0 0 0 0 0 0 1
	// buf  = 1 0 1 0 1 0 1 0
	//
}

func SetUint32(buf []byte, pos uint32, len uint8, val uint32) {
	var mask uint32 = 1 << (len - 1)
	for i := pos; i < pos+uint32(len); i++ {
		if val&mask == 0 {
			buf[i/8] |= 1 << (7 - i%8)
		} else {
			buf[i/8] &= (1 << (7 - i%8))
		}
		mask >>= 1
	}
}
