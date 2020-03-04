// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen_data.go; DO NOT EDIT.

package licenses

import (
	"bytes"
	"crypto/sha256"
)

// exceptionFileTypesMap maps SHAs of normalized licenses (see the exceptionKey function)
// to a list of types for each license.
var exceptionFileTypesMap = map[[sha256.Size]byte][]string{

	// file atlantis
	// from https://github.com/runatlantis/atlantis/blob/master/LICENSE
	[32]uint8{0x1b, 0x7d, 0xa8, 0x36, 0x1, 0x87, 0xe3, 0xd2, 0x42, 0x84, 0xe5, 0x6a, 0xd2, 0x62, 0x1e, 0x73, 0x85, 0x36, 0x6c, 0xd2, 0x5, 0x1, 0xef, 0x6b, 0x12, 0xfe, 0x5, 0x91, 0x70, 0xf6, 0xfd, 0x2f}: []string{"Apache-2.0"},

	// file autoscaler
	// from https://github.com/gardener/autoscaler/blob/v0.1.0/LICENSE
	[32]uint8{0x4c, 0xf8, 0x6a, 0xa5, 0x22, 0x43, 0x74, 0x25, 0xbf, 0x90, 0x7b, 0x52, 0x17, 0x1, 0x6, 0xd8, 0x54, 0x2e, 0x30, 0x1e, 0x11, 0x76, 0xdc, 0x81, 0xc1, 0x8, 0x9, 0x51, 0x84, 0x3e, 0x2c, 0x3b}: []string{"Apache-2.0", "BSD-2-Clause", "BSD-3-Clause", "ISC", "MIT", "MPL-2.0"},

	// file change-case
	// from https://github.com/ku/go-change-case/blob/master/LICENSE
	[32]uint8{0x70, 0x81, 0x4e, 0xf2, 0x69, 0xe, 0xa, 0x2c, 0x61, 0xb5, 0xa0, 0x16, 0x46, 0xa1, 0xe3, 0x9f, 0xbe, 0xf4, 0x0, 0x9f, 0xb5, 0x34, 0xb8, 0xb5, 0xf5, 0x52, 0x65, 0x46, 0xa, 0xb1, 0xba, 0x4e}: []string{"MIT"},

	// file freetype
	// from https://github.com/golang/freetype/blob/master/LICENSE
	[32]uint8{0xfd, 0xf7, 0x59, 0x7, 0xb, 0xae, 0xa0, 0xc7, 0x3e, 0x73, 0x94, 0x9b, 0x11, 0xd7, 0xad, 0x4d, 0x70, 0xff, 0xf3, 0x7c, 0x26, 0xb8, 0x36, 0xe6, 0xa4, 0xcc, 0x53, 0xda, 0x37, 0xe1, 0xc2, 0xb6}: []string{"Freetype"},

	// file frugal
	// from https://github.com/Workiva/frugal/blob/master/LICENSE
	[32]uint8{0x40, 0xd0, 0xe6, 0x26, 0x2e, 0xb8, 0xc5, 0x63, 0x4e, 0x40, 0xff, 0x2f, 0x10, 0x61, 0x41, 0x2f, 0xfa, 0x3d, 0xd7, 0xb2, 0x10, 0x63, 0xff, 0x18, 0xb1, 0xd6, 0x39, 0xb8, 0xa6, 0x58, 0x71, 0x54}: []string{"Apache-2.0"},

	// file hid
	// from https://github.com/mindworks-software/hid/blob/master/LICENSE.md
	[32]uint8{0xe6, 0x32, 0x28, 0xea, 0x40, 0xb6, 0x21, 0x97, 0x7f, 0xa9, 0x4b, 0xd9, 0x1f, 0x39, 0x1c, 0x2f, 0x38, 0xba, 0x48, 0xa6, 0xaf, 0x8d, 0x3b, 0x8c, 0x63, 0x64, 0x1e, 0xad, 0x3, 0x74, 0xe3, 0xa1}: []string{"BSD-3-Clause", "LGPL-2.1"},

	// file mynewt
	// from https://github.com/apache/mynewt-artifact/blob/v0.0.15/LICENSE
	[32]uint8{0x4c, 0xe4, 0xd8, 0x15, 0x11, 0x9d, 0xa8, 0x8a, 0xf2, 0x99, 0x59, 0xe5, 0x71, 0xb1, 0xae, 0xce, 0x3d, 0x1b, 0x73, 0x76, 0x9a, 0x8c, 0xb7, 0x29, 0x95, 0xdb, 0x22, 0x73, 0xad, 0x77, 0x1e, 0x33}: []string{"Apache-2.0"},

	// file newtmgr
	// from https://github.com/apache/mynewt-newtmgr/blob/master/LICENSE
	[32]uint8{0x33, 0x8b, 0xe4, 0xa9, 0xea, 0x5b, 0xf7, 0xd5, 0x2e, 0xc5, 0xc6, 0xaa, 0xd9, 0xe0, 0x3f, 0xb9, 0x9c, 0xe, 0xfd, 0x5a, 0x38, 0x6a, 0x58, 0x6d, 0x18, 0x5a, 0x82, 0xdf, 0x44, 0xd8, 0x6f, 0xca}: []string{"Apache-2.0"},
}

// exceptionKey computes the key for the exceptionFileTypesMap.
func exceptionKey(data []byte) [sha256.Size]byte {
	norm := bytes.Join(bytes.Fields(bytes.ToLower(data)), []byte{' '})
	return sha256.Sum256(norm)
}