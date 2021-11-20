// Copyright 2021 The Go Asm Authors
// SPDX-License-Identifier: BSD-3-Clause

package capstone

import (
	"syscall"
	_ "unsafe" // for go:linkname
)

// syscall3 calls a function in libc on behalf of the syscall package.
//
// syscall3 takes a pointer to a struct like:
//  struct {
//   fn    uintptr
//   a1    uintptr
//   a2    uintptr
//   a3    uintptr
//   r1    uintptr
//   r2    uintptr
//   err   uintptr
//  }
//
// syscall3 must be called on the g0 stack with the
// C calling convention (use libcCall).
//
// syscall3 expects a 32-bit result and tests for 32-bit -1
// to decide there was an error.
//go:noescape
//go:linkname syscall3 syscall.Syscall
func syscall3(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err syscall.Errno)

// syscall6 calls a function in libc on behalf of the syscall package.
//
// syscall6 takes a pointer to a struct like:
//  struct {
//   fn    uintptr
//   a1    uintptr
//   a2    uintptr
//   a3    uintptr
//   a4    uintptr
//   a5    uintptr
//   a6    uintptr
//   r1    uintptr
//   r2    uintptr
//   err   uintptr
//  }
//
// syscall6 must be called on the g0 stack with the
// C calling convention (use libcCall).
//
// syscall6 expects a 32-bit result and tests for 32-bit -1
// to decide there was an error.
//go:noescape
//go:linkname syscall6 syscall.Syscall6
func syscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno)

// syscall6x calls a function in libc on behalf of the syscall package.
//
// syscall6x takes a pointer to a struct like:
//  struct {
//   fn    uintptr
//   a1    uintptr
//   a2    uintptr
//   a3    uintptr
//   a4    uintptr
//   a5    uintptr
//   a6    uintptr
//   r1    uintptr
//   r2    uintptr
//   err   uintptr
//  }
//
// syscall6x must be called on the g0 stack with the
// C calling convention (use libcCall).
//
// syscall6x is like syscall6 but expects a 64-bit result
// and tests for 64-bit -1 to decide there was an error.
//go:noescape
//go:linkname syscall6x syscall.Syscall6X
func syscall6x(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno)

// syscall9 calls a function in libc on behalf of the syscall package.
//
// syscall9 takes a pointer to a struct like:
//  struct {
//   fn    uintptr
//   a1    uintptr
//   a2    uintptr
//   a3    uintptr
//   a4    uintptr
//   a5    uintptr
//   a6    uintptr
//   a7    uintptr
//   a8    uintptr
//   a9    uintptr
//   r1    uintptr
//   r2    uintptr
//   err   uintptr
//  }
//
// syscall9 must be called on the g0 stack with the
// C calling convention (use libcCall).
//
// syscall9 expects a 32-bit result and tests for 32-bit -1
// to decide there was an error.
//go:noescape
//go:linkname syscall9 syscall.Syscall9
func syscall9(fn, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2 uintptr, err syscall.Errno)

// syscallptr is like syscallX except that the libc function reports an
// error by returning NULL and setting errno.
//go:noescape
//go:linkname syscallptr syscall.SyscallPtr
func syscallptr(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err syscall.Errno)

// rawsyscall calls a function in libc on behalf of the syscall package.
//go:noescape
//go:linkname rawsyscall syscall.RawSyscall
func rawsyscall(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err syscall.Errno)

// rawsyscall6 calls a function in libc on behalf of the syscall package.
//go:noescape
//go:linkname rawsyscall6 syscall.RawSyscall6
func rawsyscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno)

// goString converts a C string to a Go string.
//go:nosplit
//go:linkname goString runtime.gostringnocopy
func goString(str *byte) string
