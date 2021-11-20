// Copyright 2021 The Go Asm Authors
// SPDX-License-Identifier: BSD-3-Clause

package capstone

import (
	"unsafe"

	_ "github.com/go-asm/capstone/internal/syso"
)

func Version() (major, minor int32) {
	_, _, errno := syscall3(cs_version_trampoline_addr, uintptr(unsafe.Pointer(&major)), uintptr(unsafe.Pointer(&minor)), 0)
	if errno != 0 {
		return 0, 0
	}

	return major, minor
}
