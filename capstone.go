// Copyright 2021 The Go Asm Authors
// SPDX-License-Identifier: BSD-3-Clause

package capstone

import (
	"errors"
	"unsafe"

	_ "github.com/go-asm/capstone/internal/syso"
)

// Engine represents a capstone engine.
type Engine struct {
	arch     CsArch
	mode     CsMode
	handle   *Csh
	skipData *CsOptSkipdata
}

// New returns the new Engine with the specified arch and mode.
//go:nosplit
func New(arch CsArch, mode CsMode) (Engine, error) {
	return newEngine(arch, mode)
}

//go:nosplit
func (e *Engine) Arch() CsArch {
	return e.arch
}

//go:nosplit
func (e *Engine) Mode() CsMode {
	return e.mode
}

var cs_close_trampoline_addr uintptr

// Close CS handle. MUST do to release the handle when it is not used anymore.
//
// NOTE: this must be only called when there is no longer usage of Capstone,
// not even access to cs_insn array. The reason is the this API releases some
// cached memory, thus access to any Capstone API after cs_close() might crash
// your application.
//
// In fact,this API invalidate @handle by ZERO out its value (i.e *handle = 0).
//go:nosplit
func (e *Engine) Close() error {
	_, _, errno := syscall3(cs_close_trampoline_addr, uintptr(unsafe.Pointer(e.handle)), 0, 0)
	if errno != 0 {
		return CsErr(errno)
	}
	if e.skipData != nil {
		free(unsafe.Pointer(e.skipData.Mnemonic))
	}

	return nil
}

var cs_version_trampoline_addr uintptr

// Version return combined API version & major and minor version numbers.
//
// major is major number of API version.
// minor is minor number of API version.
//go:nosplit
func Version() (major, minor int32) {
	_, _, errno := syscall3(cs_version_trampoline_addr, uintptr(unsafe.Pointer(&major)), uintptr(unsafe.Pointer(&minor)), 0)
	if errno != 0 {
		return 0, 0
	}

	return major, minor
}

var cs_strerror_trampoline_addr uintptr

// CsStrerror return a string describing given error code.
//go:nosplit
func CsStrerror(code CsErr) byte {
	s, _, errno := syscall3(cs_strerror_trampoline_addr, uintptr(code), 0, 0)
	if errno != 0 {
		return 0
	}

	return byte(s)
}

var (
	ErrOK       = errors.New("ok")
	ErrMem      = errors.New("out-of-memory")
	ErrArch     = errors.New("unsupported architecture")
	ErrHandle   = errors.New("invalid handle")
	ErrCsh      = errors.New("invalid csh argument")
	ErrMode     = errors.New("invalid/unsupported mode")
	ErrOption   = errors.New("invalid/unsupported option")
	ErrDetail   = errors.New("information is unavailable because detail option is OFF")
	ErrMemSetup = errors.New("dynamic memory management uninitialized")
	ErrVersion  = errors.New("unsupported version")
	ErrDiet     = errors.New("access irrelevant data in \"diet\" engine")
	ErrSkipData = errors.New("access irrelevant data for \"data\" instruction in SKIPDATA mode")
	ErrX86ATT   = errors.New("X86 AT&T syntax is unsupported")
	ErrX86Intel = errors.New("X86 Intel syntax is unsupported")
	ErrX86Masm  = errors.New("X86 Masm syntax is unsupported")
)

func (e CsErr) Error() string {
	switch e {
	case CS_ERR_OK: //< No error: everything was fine
		return ErrOK.Error()
	case CS_ERR_MEM: //< Out-Of-Memory error: cs_open(), cs_disasm(), cs_disasm_iter()
		return ErrMem.Error()
	case CS_ERR_ARCH: //< Unsupported architecture: cs_open()
		return ErrArch.Error()
	case CS_ERR_HANDLE: //< Invalid handle: cs_op_count(), cs_op_index()
		return ErrHandle.Error()
	case CS_ERR_CSH: //< Invalid csh argument: cs_close(), cs_errno(), cs_option()
		return ErrCsh.Error()
	case CS_ERR_MODE: //< Invalid/unsupported mode: cs_open()
		return ErrMode.Error()
	case CS_ERR_OPTION: //< Invalid/unsupported option: cs_option()
		return ErrOption.Error()
	case CS_ERR_DETAIL: //< Information is unavailable because detail option is OFF
		return ErrDetail.Error()
	case CS_ERR_MEMSETUP: //< Dynamic memory management uninitialized (see CS_OPT_MEM)
		return ErrMemSetup.Error()
	case CS_ERR_VERSION: //< Unsupported version (bindings)
		return ErrVersion.Error()
	case CS_ERR_DIET: //< Access irrelevant data in "diet" engine
		return ErrDiet.Error()
	case CS_ERR_SKIPDATA: //< Access irrelevant data for "data" instruction in SKIPDATA mode
		return ErrSkipData.Error()
	case CS_ERR_X86_ATT: //< X86 AT&T syntax is unsupported (opt-out at compile time)
		return ErrX86ATT.Error()
	case CS_ERR_X86_INTEL: //< X86 Intel syntax is unsupported (opt-out at compile time)
		return ErrX86Intel.Error()
	case CS_ERR_X86_MASM: //< X86 Masm syntax is unsupported (opt-out at compile time)
		return ErrX86Masm.Error()
	default:
		b := CsStrerror(e)
		s := goString(&b)
		return s
	}
}

// Errno report the last error number when some API function fail.
// Like glibc's errno, cs_errno might not retain its old value once accessed.
func Errno(err error) error {
	switch err {
	case CS_ERR_OK:
		return ErrOK
	case CS_ERR_MEM:
		return ErrMem
	case CS_ERR_ARCH:
		return ErrArch
	case CS_ERR_HANDLE:
		return ErrHandle
	case CS_ERR_CSH:
		return ErrCsh
	case CS_ERR_MODE:
		return ErrMode
	case CS_ERR_OPTION:
		return ErrOption
	case CS_ERR_DETAIL:
		return ErrDetail
	case CS_ERR_MEMSETUP:
		return ErrMemSetup
	case CS_ERR_VERSION:
		return ErrVersion
	case CS_ERR_DIET:
		return ErrDiet
	case CS_ERR_SKIPDATA:
		return ErrSkipData
	case CS_ERR_X86_ATT:
		return ErrX86ATT
	case CS_ERR_X86_INTEL:
		return ErrX86Intel
	case CS_ERR_X86_MASM:
		return ErrX86Masm
	}

	return err
}

//sys func CsDisasm(handle int64, code byte, codeSize int64, address uint64, count int64, insn Cs_insn **) uint64
//sys func CsDisasmIter(handle int64, code Unsigned char **, size int64, address uint64, insn Cs_insn) int32
//sys func CsErrno(handle int64) Cs_err
//sys func CsGroupName(handle int64, groupId uint32) int8
//sys func CsInsnGroup(handle int64, insn Cs_insn, groupId uint32) int32
//sys func CsInsnName(handle int64, insnId uint32) int8
//sys func CsMalloc(handle int64) Cs_insn
//sys func CsOpCount(handle int64, insn Cs_insn, opType uint32) int32
//sys func CsOpIndex(handle int64, insn Cs_insn, opType uint32, position uint32) int32
//sys func CsOption(handle int64, type Cs_opt_type, value int64) Cs_err
//sys func CsRegName(handle int64, regId uint32) int8
//sys func CsRegRead(handle int64, insn Cs_insn, regId uint32) int32
//sys func CsRegWrite(handle int64, insn Cs_insn, regId uint32) int32
//sys func CsRegsAccess(handle int64, insn Cs_insn, regsRead Unsigned short [64], regsReadCount byte, regsWrite Unsigned short [64], regsWriteCount byte) Cs_err
//sys func CsStrerror(code Cs_err) int8
//sys func CsSupport(query int32) int32
