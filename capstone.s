// Copyright 2021 The Go Asm Authors
// SPDX-License-Identifier: BSD-3-Clause

#include "textflag.h"

GLOBL ·cs_open_trampoline_addr(SB), RODATA, $8
DATA ·cs_open_trampoline_addr(SB)/8, $cs_open_trampoline<>(SB)

TEXT cs_open_trampoline<>(SB), NOSPLIT, $0-0
	JMP cs_open(SB)

GLOBL ·cs_close_trampoline_addr(SB), RODATA, $8
DATA ·cs_close_trampoline_addr(SB)/8, $cs_close_trampoline<>(SB)

TEXT cs_close_trampoline<>(SB), NOSPLIT, $0-0
	JMP cs_close(SB)

GLOBL ·cs_free_trampoline_addr(SB), RODATA, $8
DATA ·cs_free_trampoline_addr(SB)/8, $cs_free_trampoline<>(SB)

TEXT cs_free_trampoline<>(SB), NOSPLIT, $0-0
	JMP cs_free(SB)

GLOBL ·libc_free_trampoline_addr(SB), RODATA, $8
DATA ·libc_free_trampoline_addr(SB)/8, $libc_free_trampoline<>(SB)

TEXT libc_free_trampoline<>(SB), NOSPLIT, $0-0
	JMP free(SB)

GLOBL ·cs_version_trampoline_addr(SB), RODATA, $8
DATA ·cs_version_trampoline_addr(SB)/8, $cs_version_trampoline<>(SB)

TEXT cs_version_trampoline<>(SB), NOSPLIT, $0-0
	JMP cs_version(SB)

GLOBL ·cs_strerror_trampoline_addr(SB), RODATA, $8
DATA ·cs_strerror_trampoline_addr(SB)/8, $cs_strerror_trampoline<>(SB)

TEXT cs_strerror_trampoline<>(SB), NOSPLIT, $0-0
	JMP cs_strerror(SB)
