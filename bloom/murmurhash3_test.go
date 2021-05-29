// Copyright (c) 2013, 2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package bloom_test

import (
	"testing"

	"github.com/btcsuite/btcutil/v2/bloom"
)

// TestMurmurHash3 ensure the MurmurHash3 function produces the correct hash
// when given various seeds and data.
func TestMurmurHash3(t *testing.T) {
	var tests = []struct {
		seed uint32
		data []byte
		out  uint32
	}{
		{0x00000000, []byte{}, 0x00000000},
		{0xfba4c795, []byte{}, 0x6a396f08},
		{0xffffffff, []byte{}, 0x81f16f39},
		{0x00000000, []byte{0x00}, 0x514e28b7},
		{0xfba4c795, []byte{0x00}, 0xea3f0b17},
		{0x00000000, []byte{0xff}, 0xfd6cf10d},
		{0x00000000, []byte{0x00, 0x11}, 0x16c6b7ab},
		{0x00000000, []byte{0x00, 0x11, 0x22}, 0x8eb51c3d},
		{0x00000000, []byte{0x00, 0x11, 0x22, 0x33}, 0xb4471bf8},
		{0x00000000, []byte{0x00, 0x11, 0x22, 0x33, 0x44}, 0xe2301fa8},
		{0x00000000, []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55}, 0xfc2e4a15},
		{0x00000000, []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66}, 0xb074502c},
		{0x00000000, []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77}, 0x8034d2a0},
		{0x00000000, []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88}, 0xb4698def},
	}

	for i, test := range tests {
		result := bloom.MurmurHash3(test.seed, test.data)
		if result != test.out {
			t.Errorf("MurmurHash3 test #%d failed: got %v want %v\n",
				i, result, test.out)
			continue
		}
	}
}
