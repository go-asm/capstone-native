// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs wasm_unix.go

package capstone

type WasmInsn uint32

const (
	WASM_INS_UNREACHABLE         WasmInsn = 0x0
	WASM_INS_NOP                 WasmInsn = 0x1
	WASM_INS_BLOCK               WasmInsn = 0x2
	WASM_INS_LOOP                WasmInsn = 0x3
	WASM_INS_IF                  WasmInsn = 0x4
	WASM_INS_ELSE                WasmInsn = 0x5
	WASM_INS_END                 WasmInsn = 0xb
	WASM_INS_BR                  WasmInsn = 0xc
	WASM_INS_BR_IF               WasmInsn = 0xd
	WASM_INS_BR_TABLE            WasmInsn = 0xe
	WASM_INS_RETURN              WasmInsn = 0xf
	WASM_INS_CALL                WasmInsn = 0x10
	WASM_INS_CALL_INDIRECT       WasmInsn = 0x11
	WASM_INS_DROP                WasmInsn = 0x1a
	WASM_INS_SELECT              WasmInsn = 0x1b
	WASM_INS_GET_LOCAL           WasmInsn = 0x20
	WASM_INS_SET_LOCAL           WasmInsn = 0x21
	WASM_INS_TEE_LOCAL           WasmInsn = 0x22
	WASM_INS_GET_GLOBAL          WasmInsn = 0x23
	WASM_INS_SET_GLOBAL          WasmInsn = 0x24
	WASM_INS_I32_LOAD            WasmInsn = 0x28
	WASM_INS_I64_LOAD            WasmInsn = 0x29
	WASM_INS_F32_LOAD            WasmInsn = 0x2a
	WASM_INS_F64_LOAD            WasmInsn = 0x2b
	WASM_INS_I32_LOAD8_S         WasmInsn = 0x2c
	WASM_INS_I32_LOAD8_U         WasmInsn = 0x2d
	WASM_INS_I32_LOAD16_S        WasmInsn = 0x2e
	WASM_INS_I32_LOAD16_U        WasmInsn = 0x2f
	WASM_INS_I64_LOAD8_S         WasmInsn = 0x30
	WASM_INS_I64_LOAD8_U         WasmInsn = 0x31
	WASM_INS_I64_LOAD16_S        WasmInsn = 0x32
	WASM_INS_I64_LOAD16_U        WasmInsn = 0x33
	WASM_INS_I64_LOAD32_S        WasmInsn = 0x34
	WASM_INS_I64_LOAD32_U        WasmInsn = 0x35
	WASM_INS_I32_STORE           WasmInsn = 0x36
	WASM_INS_I64_STORE           WasmInsn = 0x37
	WASM_INS_F32_STORE           WasmInsn = 0x38
	WASM_INS_F64_STORE           WasmInsn = 0x39
	WASM_INS_I32_STORE8          WasmInsn = 0x3a
	WASM_INS_I32_STORE16         WasmInsn = 0x3b
	WASM_INS_I64_STORE8          WasmInsn = 0x3c
	WASM_INS_I64_STORE16         WasmInsn = 0x3d
	WASM_INS_I64_STORE32         WasmInsn = 0x3e
	WASM_INS_CURRENT_MEMORY      WasmInsn = 0x3f
	WASM_INS_GROW_MEMORY         WasmInsn = 0x40
	WASM_INS_I32_CONST           WasmInsn = 0x41
	WASM_INS_I64_CONST           WasmInsn = 0x42
	WASM_INS_F32_CONST           WasmInsn = 0x43
	WASM_INS_F64_CONST           WasmInsn = 0x44
	WASM_INS_I32_EQZ             WasmInsn = 0x45
	WASM_INS_I32_EQ              WasmInsn = 0x46
	WASM_INS_I32_NE              WasmInsn = 0x47
	WASM_INS_I32_LT_S            WasmInsn = 0x48
	WASM_INS_I32_LT_U            WasmInsn = 0x49
	WASM_INS_I32_GT_S            WasmInsn = 0x4a
	WASM_INS_I32_GT_U            WasmInsn = 0x4b
	WASM_INS_I32_LE_S            WasmInsn = 0x4c
	WASM_INS_I32_LE_U            WasmInsn = 0x4d
	WASM_INS_I32_GE_S            WasmInsn = 0x4e
	WASM_INS_I32_GE_U            WasmInsn = 0x4f
	WASM_INS_I64_EQZ             WasmInsn = 0x50
	WASM_INS_I64_EQ              WasmInsn = 0x51
	WASM_INS_I64_NE              WasmInsn = 0x52
	WASM_INS_I64_LT_S            WasmInsn = 0x53
	WASM_INS_I64_LT_U            WasmInsn = 0x54
	WASN_INS_I64_GT_S            WasmInsn = 0x55
	WASM_INS_I64_GT_U            WasmInsn = 0x56
	WASM_INS_I64_LE_S            WasmInsn = 0x57
	WASM_INS_I64_LE_U            WasmInsn = 0x58
	WASM_INS_I64_GE_S            WasmInsn = 0x59
	WASM_INS_I64_GE_U            WasmInsn = 0x5a
	WASM_INS_F32_EQ              WasmInsn = 0x5b
	WASM_INS_F32_NE              WasmInsn = 0x5c
	WASM_INS_F32_LT              WasmInsn = 0x5d
	WASM_INS_F32_GT              WasmInsn = 0x5e
	WASM_INS_F32_LE              WasmInsn = 0x5f
	WASM_INS_F32_GE              WasmInsn = 0x60
	WASM_INS_F64_EQ              WasmInsn = 0x61
	WASM_INS_F64_NE              WasmInsn = 0x62
	WASM_INS_F64_LT              WasmInsn = 0x63
	WASM_INS_F64_GT              WasmInsn = 0x64
	WASM_INS_F64_LE              WasmInsn = 0x65
	WASM_INS_F64_GE              WasmInsn = 0x66
	WASM_INS_I32_CLZ             WasmInsn = 0x67
	WASM_INS_I32_CTZ             WasmInsn = 0x68
	WASM_INS_I32_POPCNT          WasmInsn = 0x69
	WASM_INS_I32_ADD             WasmInsn = 0x6a
	WASM_INS_I32_SUB             WasmInsn = 0x6b
	WASM_INS_I32_MUL             WasmInsn = 0x6c
	WASM_INS_I32_DIV_S           WasmInsn = 0x6d
	WASM_INS_I32_DIV_U           WasmInsn = 0x6e
	WASM_INS_I32_REM_S           WasmInsn = 0x6f
	WASM_INS_I32_REM_U           WasmInsn = 0x70
	WASM_INS_I32_AND             WasmInsn = 0x71
	WASM_INS_I32_OR              WasmInsn = 0x72
	WASM_INS_I32_XOR             WasmInsn = 0x73
	WASM_INS_I32_SHL             WasmInsn = 0x74
	WASM_INS_I32_SHR_S           WasmInsn = 0x75
	WASM_INS_I32_SHR_U           WasmInsn = 0x76
	WASM_INS_I32_ROTL            WasmInsn = 0x77
	WASM_INS_I32_ROTR            WasmInsn = 0x78
	WASM_INS_I64_CLZ             WasmInsn = 0x79
	WASM_INS_I64_CTZ             WasmInsn = 0x7a
	WASM_INS_I64_POPCNT          WasmInsn = 0x7b
	WASM_INS_I64_ADD             WasmInsn = 0x7c
	WASM_INS_I64_SUB             WasmInsn = 0x7d
	WASM_INS_I64_MUL             WasmInsn = 0x7e
	WASM_INS_I64_DIV_S           WasmInsn = 0x7f
	WASM_INS_I64_DIV_U           WasmInsn = 0x80
	WASM_INS_I64_REM_S           WasmInsn = 0x81
	WASM_INS_I64_REM_U           WasmInsn = 0x82
	WASM_INS_I64_AND             WasmInsn = 0x83
	WASM_INS_I64_OR              WasmInsn = 0x84
	WASM_INS_I64_XOR             WasmInsn = 0x85
	WASM_INS_I64_SHL             WasmInsn = 0x86
	WASM_INS_I64_SHR_S           WasmInsn = 0x87
	WASM_INS_I64_SHR_U           WasmInsn = 0x88
	WASM_INS_I64_ROTL            WasmInsn = 0x89
	WASM_INS_I64_ROTR            WasmInsn = 0x8a
	WASM_INS_F32_ABS             WasmInsn = 0x8b
	WASM_INS_F32_NEG             WasmInsn = 0x8c
	WASM_INS_F32_CEIL            WasmInsn = 0x8d
	WASM_INS_F32_FLOOR           WasmInsn = 0x8e
	WASM_INS_F32_TRUNC           WasmInsn = 0x8f
	WASM_INS_F32_NEAREST         WasmInsn = 0x90
	WASM_INS_F32_SQRT            WasmInsn = 0x91
	WASM_INS_F32_ADD             WasmInsn = 0x92
	WASM_INS_F32_SUB             WasmInsn = 0x93
	WASM_INS_F32_MUL             WasmInsn = 0x94
	WASM_INS_F32_DIV             WasmInsn = 0x95
	WASM_INS_F32_MIN             WasmInsn = 0x96
	WASM_INS_F32_MAX             WasmInsn = 0x97
	WASM_INS_F32_COPYSIGN        WasmInsn = 0x98
	WASM_INS_F64_ABS             WasmInsn = 0x99
	WASM_INS_F64_NEG             WasmInsn = 0x9a
	WASM_INS_F64_CEIL            WasmInsn = 0x9b
	WASM_INS_F64_FLOOR           WasmInsn = 0x9c
	WASM_INS_F64_TRUNC           WasmInsn = 0x9d
	WASM_INS_F64_NEAREST         WasmInsn = 0x9e
	WASM_INS_F64_SQRT            WasmInsn = 0x9f
	WASM_INS_F64_ADD             WasmInsn = 0xa0
	WASM_INS_F64_SUB             WasmInsn = 0xa1
	WASM_INS_F64_MUL             WasmInsn = 0xa2
	WASM_INS_F64_DIV             WasmInsn = 0xa3
	WASM_INS_F64_MIN             WasmInsn = 0xa4
	WASM_INS_F64_MAX             WasmInsn = 0xa5
	WASM_INS_F64_COPYSIGN        WasmInsn = 0xa6
	WASM_INS_I32_WARP_I64        WasmInsn = 0xa7
	WASP_INS_I32_TRUNC_S_F32     WasmInsn = 0xa8
	WASM_INS_I32_TRUNC_U_F32     WasmInsn = 0xa9
	WASM_INS_I32_TRUNC_S_F64     WasmInsn = 0xaa
	WASM_INS_I32_TRUNC_U_F64     WasmInsn = 0xab
	WASM_INS_I64_EXTEND_S_I32    WasmInsn = 0xac
	WASM_INS_I64_EXTEND_U_I32    WasmInsn = 0xad
	WASM_INS_I64_TRUNC_S_F32     WasmInsn = 0xae
	WASM_INS_I64_TRUNC_U_F32     WasmInsn = 0xaf
	WASM_INS_I64_TRUNC_S_F64     WasmInsn = 0xb0
	WASM_INS_I64_TRUNC_U_F64     WasmInsn = 0xb1
	WASM_INS_F32_CONVERT_S_I32   WasmInsn = 0xb2
	WASM_INS_F32_CONVERT_U_I32   WasmInsn = 0xb3
	WASM_INS_F32_CONVERT_S_I64   WasmInsn = 0xb4
	WASM_INS_F32_CONVERT_U_I64   WasmInsn = 0xb5
	WASM_INS_F32_DEMOTE_F64      WasmInsn = 0xb6
	WASM_INS_F64_CONVERT_S_I32   WasmInsn = 0xb7
	WASM_INS_F64_CONVERT_U_I32   WasmInsn = 0xb8
	WASM_INS_F64_CONVERT_S_I64   WasmInsn = 0xb9
	WASM_INS_F64_CONVERT_U_I64   WasmInsn = 0xba
	WASM_INS_F64_PROMOTE_F32     WasmInsn = 0xbb
	WASM_INS_I32_REINTERPRET_F32 WasmInsn = 0xbc
	WASM_INS_I64_REINTERPRET_F64 WasmInsn = 0xbd
	WASM_INS_F32_REINTERPRET_I32 WasmInsn = 0xbe
	WASM_INS_F64_REINTERPRET_I64 WasmInsn = 0xbf
	WASM_INS_INVALID             WasmInsn = 0x200
	WASM_INS_ENDING              WasmInsn = 0x201
)

type WasmInsnGroup uint32

const (
	WASM_GRP_INVALID    WasmInsnGroup = 0x0
	WASM_GRP_NUMBERIC   WasmInsnGroup = 0x8
	WASM_GRP_PARAMETRIC WasmInsnGroup = 0x9
	WASM_GRP_VARIABLE   WasmInsnGroup = 0xa
	WASM_GRP_MEMORY     WasmInsnGroup = 0xb
	WASM_GRP_CONTROL    WasmInsnGroup = 0xc
	WASM_GRP_ENDING     WasmInsnGroup = 0xd
)

type WasmOpType uint32

const (
	WASM_OP_INVALID   WasmOpType = 0x0
	WASM_OP_NONE      WasmOpType = 0x1
	WASM_OP_INT7      WasmOpType = 0x2
	WASM_OP_VARUINT32 WasmOpType = 0x3
	WASM_OP_VARUINT64 WasmOpType = 0x4
	WASM_OP_UINT32    WasmOpType = 0x5
	WASM_OP_UINT64    WasmOpType = 0x6
	WASM_OP_IMM       WasmOpType = 0x7
	WASM_OP_BRTABLE   WasmOpType = 0x8
)

type CsWasm struct {
	Count    uint8
	_        [4]byte
	Operands [2]CsWasmOp
}

type CsWasmBrtable struct {
	Length  uint32
	Address uint64
	Target  uint32
	_       [4]byte
}

type CsWasmOp struct {
	Type uint32
	Size uint32
	Int7 int8
	_    [23]byte
}
