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

package bitoffset_test

import (
	"testing"

	"github.com/ThingsIXFoundation/bitoffset"
)

func TestUint32(t *testing.T) {
	buf := []byte{0x01, 0x02}
	got := bitoffset.Uint32(buf, 0, 8)
	if got != 1 {
		t.Errorf("got: %d, want %d", got, 1)
	}

	got = bitoffset.Uint32(buf, 8, 8)
	if got != 2 {
		t.Errorf("got: %d, want %d", got, 2)
	}

}

func TestInt32(t *testing.T) {
	buf := []byte{0xFF}
	got := bitoffset.Int32(buf, 0, 8)
	if got != -1 {
		t.Errorf("got: %d, want %d", got, -1)
	}

	buf = []byte{0xFF, 0xFF}
	got = bitoffset.Int32(buf, 0, 16)
	if got != -1 {
		t.Errorf("got: %d, want %d", got, -1)
	}

	buf = []byte{0x01, 0x01}
	got = bitoffset.Int32(buf, 8, 8)
	if got != 1 {
		t.Errorf("got: %d, want %d", got, 1)
	}

	buf = []byte{0x01, 0x01}
	got = bitoffset.Int32(buf, 0, 16)
	if got != 257 {
		t.Errorf("got: %d, want %d", got, 257)
	}

}
