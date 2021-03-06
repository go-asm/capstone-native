// Code generated by github.com/go-darwin/tools/cmd/mkgodef; DO NOT EDIT.
// Input to cgo -godefs.

//go:build ignore
// +build ignore

package capstone

/*
#include <capstone/wasm.h>
*/
import "C"

type WasmInsn C.enum_wasm_insn

const (
	WASM_INS_UNREACHABLE         WasmInsn = C.WASM_INS_UNREACHABLE
	WASM_INS_NOP                 WasmInsn = C.WASM_INS_NOP
	WASM_INS_BLOCK               WasmInsn = C.WASM_INS_BLOCK
	WASM_INS_LOOP                WasmInsn = C.WASM_INS_LOOP
	WASM_INS_IF                  WasmInsn = C.WASM_INS_IF
	WASM_INS_ELSE                WasmInsn = C.WASM_INS_ELSE
	WASM_INS_END                 WasmInsn = C.WASM_INS_END
	WASM_INS_BR                  WasmInsn = C.WASM_INS_BR
	WASM_INS_BR_IF               WasmInsn = C.WASM_INS_BR_IF
	WASM_INS_BR_TABLE            WasmInsn = C.WASM_INS_BR_TABLE
	WASM_INS_RETURN              WasmInsn = C.WASM_INS_RETURN
	WASM_INS_CALL                WasmInsn = C.WASM_INS_CALL
	WASM_INS_CALL_INDIRECT       WasmInsn = C.WASM_INS_CALL_INDIRECT
	WASM_INS_DROP                WasmInsn = C.WASM_INS_DROP
	WASM_INS_SELECT              WasmInsn = C.WASM_INS_SELECT
	WASM_INS_GET_LOCAL           WasmInsn = C.WASM_INS_GET_LOCAL
	WASM_INS_SET_LOCAL           WasmInsn = C.WASM_INS_SET_LOCAL
	WASM_INS_TEE_LOCAL           WasmInsn = C.WASM_INS_TEE_LOCAL
	WASM_INS_GET_GLOBAL          WasmInsn = C.WASM_INS_GET_GLOBAL
	WASM_INS_SET_GLOBAL          WasmInsn = C.WASM_INS_SET_GLOBAL
	WASM_INS_I32_LOAD            WasmInsn = C.WASM_INS_I32_LOAD
	WASM_INS_I64_LOAD            WasmInsn = C.WASM_INS_I64_LOAD
	WASM_INS_F32_LOAD            WasmInsn = C.WASM_INS_F32_LOAD
	WASM_INS_F64_LOAD            WasmInsn = C.WASM_INS_F64_LOAD
	WASM_INS_I32_LOAD8_S         WasmInsn = C.WASM_INS_I32_LOAD8_S
	WASM_INS_I32_LOAD8_U         WasmInsn = C.WASM_INS_I32_LOAD8_U
	WASM_INS_I32_LOAD16_S        WasmInsn = C.WASM_INS_I32_LOAD16_S
	WASM_INS_I32_LOAD16_U        WasmInsn = C.WASM_INS_I32_LOAD16_U
	WASM_INS_I64_LOAD8_S         WasmInsn = C.WASM_INS_I64_LOAD8_S
	WASM_INS_I64_LOAD8_U         WasmInsn = C.WASM_INS_I64_LOAD8_U
	WASM_INS_I64_LOAD16_S        WasmInsn = C.WASM_INS_I64_LOAD16_S
	WASM_INS_I64_LOAD16_U        WasmInsn = C.WASM_INS_I64_LOAD16_U
	WASM_INS_I64_LOAD32_S        WasmInsn = C.WASM_INS_I64_LOAD32_S
	WASM_INS_I64_LOAD32_U        WasmInsn = C.WASM_INS_I64_LOAD32_U
	WASM_INS_I32_STORE           WasmInsn = C.WASM_INS_I32_STORE
	WASM_INS_I64_STORE           WasmInsn = C.WASM_INS_I64_STORE
	WASM_INS_F32_STORE           WasmInsn = C.WASM_INS_F32_STORE
	WASM_INS_F64_STORE           WasmInsn = C.WASM_INS_F64_STORE
	WASM_INS_I32_STORE8          WasmInsn = C.WASM_INS_I32_STORE8
	WASM_INS_I32_STORE16         WasmInsn = C.WASM_INS_I32_STORE16
	WASM_INS_I64_STORE8          WasmInsn = C.WASM_INS_I64_STORE8
	WASM_INS_I64_STORE16         WasmInsn = C.WASM_INS_I64_STORE16
	WASM_INS_I64_STORE32         WasmInsn = C.WASM_INS_I64_STORE32
	WASM_INS_CURRENT_MEMORY      WasmInsn = C.WASM_INS_CURRENT_MEMORY
	WASM_INS_GROW_MEMORY         WasmInsn = C.WASM_INS_GROW_MEMORY
	WASM_INS_I32_CONST           WasmInsn = C.WASM_INS_I32_CONST
	WASM_INS_I64_CONST           WasmInsn = C.WASM_INS_I64_CONST
	WASM_INS_F32_CONST           WasmInsn = C.WASM_INS_F32_CONST
	WASM_INS_F64_CONST           WasmInsn = C.WASM_INS_F64_CONST
	WASM_INS_I32_EQZ             WasmInsn = C.WASM_INS_I32_EQZ
	WASM_INS_I32_EQ              WasmInsn = C.WASM_INS_I32_EQ
	WASM_INS_I32_NE              WasmInsn = C.WASM_INS_I32_NE
	WASM_INS_I32_LT_S            WasmInsn = C.WASM_INS_I32_LT_S
	WASM_INS_I32_LT_U            WasmInsn = C.WASM_INS_I32_LT_U
	WASM_INS_I32_GT_S            WasmInsn = C.WASM_INS_I32_GT_S
	WASM_INS_I32_GT_U            WasmInsn = C.WASM_INS_I32_GT_U
	WASM_INS_I32_LE_S            WasmInsn = C.WASM_INS_I32_LE_S
	WASM_INS_I32_LE_U            WasmInsn = C.WASM_INS_I32_LE_U
	WASM_INS_I32_GE_S            WasmInsn = C.WASM_INS_I32_GE_S
	WASM_INS_I32_GE_U            WasmInsn = C.WASM_INS_I32_GE_U
	WASM_INS_I64_EQZ             WasmInsn = C.WASM_INS_I64_EQZ
	WASM_INS_I64_EQ              WasmInsn = C.WASM_INS_I64_EQ
	WASM_INS_I64_NE              WasmInsn = C.WASM_INS_I64_NE
	WASM_INS_I64_LT_S            WasmInsn = C.WASM_INS_I64_LT_S
	WASM_INS_I64_LT_U            WasmInsn = C.WASM_INS_I64_LT_U
	WASN_INS_I64_GT_S            WasmInsn = C.WASN_INS_I64_GT_S
	WASM_INS_I64_GT_U            WasmInsn = C.WASM_INS_I64_GT_U
	WASM_INS_I64_LE_S            WasmInsn = C.WASM_INS_I64_LE_S
	WASM_INS_I64_LE_U            WasmInsn = C.WASM_INS_I64_LE_U
	WASM_INS_I64_GE_S            WasmInsn = C.WASM_INS_I64_GE_S
	WASM_INS_I64_GE_U            WasmInsn = C.WASM_INS_I64_GE_U
	WASM_INS_F32_EQ              WasmInsn = C.WASM_INS_F32_EQ
	WASM_INS_F32_NE              WasmInsn = C.WASM_INS_F32_NE
	WASM_INS_F32_LT              WasmInsn = C.WASM_INS_F32_LT
	WASM_INS_F32_GT              WasmInsn = C.WASM_INS_F32_GT
	WASM_INS_F32_LE              WasmInsn = C.WASM_INS_F32_LE
	WASM_INS_F32_GE              WasmInsn = C.WASM_INS_F32_GE
	WASM_INS_F64_EQ              WasmInsn = C.WASM_INS_F64_EQ
	WASM_INS_F64_NE              WasmInsn = C.WASM_INS_F64_NE
	WASM_INS_F64_LT              WasmInsn = C.WASM_INS_F64_LT
	WASM_INS_F64_GT              WasmInsn = C.WASM_INS_F64_GT
	WASM_INS_F64_LE              WasmInsn = C.WASM_INS_F64_LE
	WASM_INS_F64_GE              WasmInsn = C.WASM_INS_F64_GE
	WASM_INS_I32_CLZ             WasmInsn = C.WASM_INS_I32_CLZ
	WASM_INS_I32_CTZ             WasmInsn = C.WASM_INS_I32_CTZ
	WASM_INS_I32_POPCNT          WasmInsn = C.WASM_INS_I32_POPCNT
	WASM_INS_I32_ADD             WasmInsn = C.WASM_INS_I32_ADD
	WASM_INS_I32_SUB             WasmInsn = C.WASM_INS_I32_SUB
	WASM_INS_I32_MUL             WasmInsn = C.WASM_INS_I32_MUL
	WASM_INS_I32_DIV_S           WasmInsn = C.WASM_INS_I32_DIV_S
	WASM_INS_I32_DIV_U           WasmInsn = C.WASM_INS_I32_DIV_U
	WASM_INS_I32_REM_S           WasmInsn = C.WASM_INS_I32_REM_S
	WASM_INS_I32_REM_U           WasmInsn = C.WASM_INS_I32_REM_U
	WASM_INS_I32_AND             WasmInsn = C.WASM_INS_I32_AND
	WASM_INS_I32_OR              WasmInsn = C.WASM_INS_I32_OR
	WASM_INS_I32_XOR             WasmInsn = C.WASM_INS_I32_XOR
	WASM_INS_I32_SHL             WasmInsn = C.WASM_INS_I32_SHL
	WASM_INS_I32_SHR_S           WasmInsn = C.WASM_INS_I32_SHR_S
	WASM_INS_I32_SHR_U           WasmInsn = C.WASM_INS_I32_SHR_U
	WASM_INS_I32_ROTL            WasmInsn = C.WASM_INS_I32_ROTL
	WASM_INS_I32_ROTR            WasmInsn = C.WASM_INS_I32_ROTR
	WASM_INS_I64_CLZ             WasmInsn = C.WASM_INS_I64_CLZ
	WASM_INS_I64_CTZ             WasmInsn = C.WASM_INS_I64_CTZ
	WASM_INS_I64_POPCNT          WasmInsn = C.WASM_INS_I64_POPCNT
	WASM_INS_I64_ADD             WasmInsn = C.WASM_INS_I64_ADD
	WASM_INS_I64_SUB             WasmInsn = C.WASM_INS_I64_SUB
	WASM_INS_I64_MUL             WasmInsn = C.WASM_INS_I64_MUL
	WASM_INS_I64_DIV_S           WasmInsn = C.WASM_INS_I64_DIV_S
	WASM_INS_I64_DIV_U           WasmInsn = C.WASM_INS_I64_DIV_U
	WASM_INS_I64_REM_S           WasmInsn = C.WASM_INS_I64_REM_S
	WASM_INS_I64_REM_U           WasmInsn = C.WASM_INS_I64_REM_U
	WASM_INS_I64_AND             WasmInsn = C.WASM_INS_I64_AND
	WASM_INS_I64_OR              WasmInsn = C.WASM_INS_I64_OR
	WASM_INS_I64_XOR             WasmInsn = C.WASM_INS_I64_XOR
	WASM_INS_I64_SHL             WasmInsn = C.WASM_INS_I64_SHL
	WASM_INS_I64_SHR_S           WasmInsn = C.WASM_INS_I64_SHR_S
	WASM_INS_I64_SHR_U           WasmInsn = C.WASM_INS_I64_SHR_U
	WASM_INS_I64_ROTL            WasmInsn = C.WASM_INS_I64_ROTL
	WASM_INS_I64_ROTR            WasmInsn = C.WASM_INS_I64_ROTR
	WASM_INS_F32_ABS             WasmInsn = C.WASM_INS_F32_ABS
	WASM_INS_F32_NEG             WasmInsn = C.WASM_INS_F32_NEG
	WASM_INS_F32_CEIL            WasmInsn = C.WASM_INS_F32_CEIL
	WASM_INS_F32_FLOOR           WasmInsn = C.WASM_INS_F32_FLOOR
	WASM_INS_F32_TRUNC           WasmInsn = C.WASM_INS_F32_TRUNC
	WASM_INS_F32_NEAREST         WasmInsn = C.WASM_INS_F32_NEAREST
	WASM_INS_F32_SQRT            WasmInsn = C.WASM_INS_F32_SQRT
	WASM_INS_F32_ADD             WasmInsn = C.WASM_INS_F32_ADD
	WASM_INS_F32_SUB             WasmInsn = C.WASM_INS_F32_SUB
	WASM_INS_F32_MUL             WasmInsn = C.WASM_INS_F32_MUL
	WASM_INS_F32_DIV             WasmInsn = C.WASM_INS_F32_DIV
	WASM_INS_F32_MIN             WasmInsn = C.WASM_INS_F32_MIN
	WASM_INS_F32_MAX             WasmInsn = C.WASM_INS_F32_MAX
	WASM_INS_F32_COPYSIGN        WasmInsn = C.WASM_INS_F32_COPYSIGN
	WASM_INS_F64_ABS             WasmInsn = C.WASM_INS_F64_ABS
	WASM_INS_F64_NEG             WasmInsn = C.WASM_INS_F64_NEG
	WASM_INS_F64_CEIL            WasmInsn = C.WASM_INS_F64_CEIL
	WASM_INS_F64_FLOOR           WasmInsn = C.WASM_INS_F64_FLOOR
	WASM_INS_F64_TRUNC           WasmInsn = C.WASM_INS_F64_TRUNC
	WASM_INS_F64_NEAREST         WasmInsn = C.WASM_INS_F64_NEAREST
	WASM_INS_F64_SQRT            WasmInsn = C.WASM_INS_F64_SQRT
	WASM_INS_F64_ADD             WasmInsn = C.WASM_INS_F64_ADD
	WASM_INS_F64_SUB             WasmInsn = C.WASM_INS_F64_SUB
	WASM_INS_F64_MUL             WasmInsn = C.WASM_INS_F64_MUL
	WASM_INS_F64_DIV             WasmInsn = C.WASM_INS_F64_DIV
	WASM_INS_F64_MIN             WasmInsn = C.WASM_INS_F64_MIN
	WASM_INS_F64_MAX             WasmInsn = C.WASM_INS_F64_MAX
	WASM_INS_F64_COPYSIGN        WasmInsn = C.WASM_INS_F64_COPYSIGN
	WASM_INS_I32_WARP_I64        WasmInsn = C.WASM_INS_I32_WARP_I64
	WASP_INS_I32_TRUNC_S_F32     WasmInsn = C.WASP_INS_I32_TRUNC_S_F32
	WASM_INS_I32_TRUNC_U_F32     WasmInsn = C.WASM_INS_I32_TRUNC_U_F32
	WASM_INS_I32_TRUNC_S_F64     WasmInsn = C.WASM_INS_I32_TRUNC_S_F64
	WASM_INS_I32_TRUNC_U_F64     WasmInsn = C.WASM_INS_I32_TRUNC_U_F64
	WASM_INS_I64_EXTEND_S_I32    WasmInsn = C.WASM_INS_I64_EXTEND_S_I32
	WASM_INS_I64_EXTEND_U_I32    WasmInsn = C.WASM_INS_I64_EXTEND_U_I32
	WASM_INS_I64_TRUNC_S_F32     WasmInsn = C.WASM_INS_I64_TRUNC_S_F32
	WASM_INS_I64_TRUNC_U_F32     WasmInsn = C.WASM_INS_I64_TRUNC_U_F32
	WASM_INS_I64_TRUNC_S_F64     WasmInsn = C.WASM_INS_I64_TRUNC_S_F64
	WASM_INS_I64_TRUNC_U_F64     WasmInsn = C.WASM_INS_I64_TRUNC_U_F64
	WASM_INS_F32_CONVERT_S_I32   WasmInsn = C.WASM_INS_F32_CONVERT_S_I32
	WASM_INS_F32_CONVERT_U_I32   WasmInsn = C.WASM_INS_F32_CONVERT_U_I32
	WASM_INS_F32_CONVERT_S_I64   WasmInsn = C.WASM_INS_F32_CONVERT_S_I64
	WASM_INS_F32_CONVERT_U_I64   WasmInsn = C.WASM_INS_F32_CONVERT_U_I64
	WASM_INS_F32_DEMOTE_F64      WasmInsn = C.WASM_INS_F32_DEMOTE_F64
	WASM_INS_F64_CONVERT_S_I32   WasmInsn = C.WASM_INS_F64_CONVERT_S_I32
	WASM_INS_F64_CONVERT_U_I32   WasmInsn = C.WASM_INS_F64_CONVERT_U_I32
	WASM_INS_F64_CONVERT_S_I64   WasmInsn = C.WASM_INS_F64_CONVERT_S_I64
	WASM_INS_F64_CONVERT_U_I64   WasmInsn = C.WASM_INS_F64_CONVERT_U_I64
	WASM_INS_F64_PROMOTE_F32     WasmInsn = C.WASM_INS_F64_PROMOTE_F32
	WASM_INS_I32_REINTERPRET_F32 WasmInsn = C.WASM_INS_I32_REINTERPRET_F32
	WASM_INS_I64_REINTERPRET_F64 WasmInsn = C.WASM_INS_I64_REINTERPRET_F64
	WASM_INS_F32_REINTERPRET_I32 WasmInsn = C.WASM_INS_F32_REINTERPRET_I32
	WASM_INS_F64_REINTERPRET_I64 WasmInsn = C.WASM_INS_F64_REINTERPRET_I64
	WASM_INS_INVALID             WasmInsn = C.WASM_INS_INVALID
	WASM_INS_ENDING              WasmInsn = C.WASM_INS_ENDING
)

type WasmInsnGroup C.enum_wasm_insn_group

const (
	WASM_GRP_INVALID    WasmInsnGroup = C.WASM_GRP_INVALID
	WASM_GRP_NUMBERIC   WasmInsnGroup = C.WASM_GRP_NUMBERIC
	WASM_GRP_PARAMETRIC WasmInsnGroup = C.WASM_GRP_PARAMETRIC
	WASM_GRP_VARIABLE   WasmInsnGroup = C.WASM_GRP_VARIABLE
	WASM_GRP_MEMORY     WasmInsnGroup = C.WASM_GRP_MEMORY
	WASM_GRP_CONTROL    WasmInsnGroup = C.WASM_GRP_CONTROL
	WASM_GRP_ENDING     WasmInsnGroup = C.WASM_GRP_ENDING
)

type WasmOpType C.enum_wasm_op_type

const (
	WASM_OP_INVALID   WasmOpType = C.WASM_OP_INVALID
	WASM_OP_NONE      WasmOpType = C.WASM_OP_NONE
	WASM_OP_INT7      WasmOpType = C.WASM_OP_INT7
	WASM_OP_VARUINT32 WasmOpType = C.WASM_OP_VARUINT32
	WASM_OP_VARUINT64 WasmOpType = C.WASM_OP_VARUINT64
	WASM_OP_UINT32    WasmOpType = C.WASM_OP_UINT32
	WASM_OP_UINT64    WasmOpType = C.WASM_OP_UINT64
	WASM_OP_IMM       WasmOpType = C.WASM_OP_IMM
	WASM_OP_BRTABLE   WasmOpType = C.WASM_OP_BRTABLE
)

type CsWasm C.struct_cs_wasm

type CsWasmBrtable C.struct_cs_wasm_brtable

type CsWasmOp C.struct_cs_wasm_op
