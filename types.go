// Copyright 2021 The Go Asm Authors
// SPDX-License-Identifier: BSD-3-Clause

package capstone

type CsDetail struct {
	RegsRead       [16]uint16
	RegsReadCount  uint8
	RegsWrite      [20]uint16
	RegsWriteCount uint8
	Groups         [8]uint8
	GroupsCount    uint8
	X86            CsX86
	Arm64          CsArm64
	Arm            CsArm
	M68k           CsM68K
	Mips           CsMips
	Ppc            CsPpc
	Sparc          CsSparc
	Sysz           CsSysz
	Xcore          CsXcore
	Tms320c64x     CsTms320C64X
	M680x          CsM680X
	Evm            CsEvm
	Mos65xx        CsMos65Xx
	_              [4]byte
	Wasm           CsWasm
	Bpf            CsBpf
	Riscv          CsRiscv
}

type CsInsn struct {
	ID       uint32
	Address  uint64
	Size     uint16
	Bytes    [24]uint8
	Mnemonic [32]int8
	Str      [160]int8
	Detail   *CsDetail
}

type CsOptMem struct {
	Malloc    CsMallocT
	Calloc    CsCallocT
	Realloc   CsReallocT
	Free      CsFreeT
	Vsnprintf CsVsnprintfT
}

type CsOptMnem struct {
	ID       uint32
	Mnemonic *int8
}
