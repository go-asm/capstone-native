// Copyright 2021 The Go Asm Authors
// SPDX-License-Identifier: BSD-3-Clause

package capstone

import "unsafe"

var cs_open_trampoline_addr uintptr

//go:nosplit
func newEngine(arch CsArch, mode CsMode) (Engine, error) {
	var handle Csh
	_, _, errno := syscall3(cs_open_trampoline_addr, uintptr(arch), uintptr(mode), uintptr(unsafe.Pointer(&handle)))
	if errno != 0 {
		return Engine{}, CsErr(errno)
	}

	e := Engine{
		arch:   arch,
		mode:   mode,
		handle: &handle,
	}

	return e, nil
}

var libc_free_trampoline_addr uintptr

//go:cgo_import_dynamic libc_free free "/usr/lib/libSystem.B.dylib"

//go:nosplit
func free(p unsafe.Pointer) error {
	_, _, errno := syscall3(libc_free_trampoline_addr, uintptr(unsafe.Pointer(&p)), 0, 0)
	if errno != 0 {
		return errno
	}

	return nil
}
