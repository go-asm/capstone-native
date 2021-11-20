// Copyright 2021 The Go Asm Authors
// SPDX-License-Identifier: BSD-3-Clause

#include "textflag.h"

GLOBL ·cs_version_trampoline_addr(SB), RODATA, $8
DATA ·cs_version_trampoline_addr(SB)/8, $cs_version_trampoline<>(SB)

TEXT cs_version_trampoline<>(SB), NOSPLIT, $0-0
	JMP cs_version(SB)

GLOBL ·cs_strerror_trampoline_addr(SB), RODATA, $8
DATA ·cs_strerror_trampoline_addr(SB)/8, $cs_strerror_trampoline<>(SB)

TEXT cs_strerror_trampoline<>(SB), NOSPLIT, $0-0
	JMP cs_strerror(SB)
