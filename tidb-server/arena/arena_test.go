// Copyright 2015 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package arena

import (
	"testing"
)

func TestSimpleArenaAllocator(t *testing.T) {
	arena := NewArenaAllocator(1000)
	slice := arena.AllocBytes(10)
	if arena.off != 10 {
		t.Error("off not match, expect 10 bug got", arena.off)
	}

	if len(slice) != 0 {
		t.Error("slice len not match, expect 0 bug got", len(slice))
	}

	if cap(slice) != 10 {
		t.Error("slice len not match, expect 10 bug got", cap(slice))
	}

	slice = arena.AllocBytes(20)
	if arena.off != 30 {
		t.Error("off not match, expect 30 bug got", arena.off)
	}

	if len(slice) != 0 {
		t.Error("slice len not match, expect 0 bug got", len(slice))
	}

	if cap(slice) != 20 {
		t.Error("slice len not match, expect 20 bug got", cap(slice))
	}

	slice = arena.AllocBytes(1024)

	if arena.off != 30 {
		t.Error("off not match, expect 30 bug got", arena.off)
	}

	if len(slice) != 0 {
		t.Error("slice len not match, expect 0 bug got", len(slice))
	}

	if cap(slice) != 1024 {
		t.Error("slice len not match, expect 1024 bug got", cap(slice))
	}

	slice = arena.AllocBytesWithLen(2, 10)

	if arena.off != 40 {
		t.Error("off not match, expect 40 bug got", arena.off)
	}

	if len(slice) != 2 {
		t.Error("slice len not match, expect 2 bug got", len(slice))
	}

	if cap(slice) != 10 {
		t.Error("slice len not match, expect 10 bug got", cap(slice))
	}

	arena.Reset()
	if arena.off != 0 || cap(arena.arena) != 1000 {
		t.Error("off should be 0")
	}

	if cap(arena.arena) != 1000 {
		t.Error("off should be 0")
	}
}

func TestStdAllocator(t *testing.T) {
	slice := StdAllocator.AllocBytes(20)
	if len(slice) != 0 {
		t.Error("length not match")
	}

	if cap(slice) != 20 {
		t.Error("cap not match")
	}

	slice = StdAllocator.AllocBytesWithLen(10, 20)
	if len(slice) != 10 {
		t.Error("length not match")
	}

	if cap(slice) != 20 {
		t.Error("cap not match")
	}
}
