// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs bpf_unix.go

package capstone

type BpfExtType uint32

const (
	BPF_EXT_INVALID BpfExtType = 0x0
	BPF_EXT_LEN     BpfExtType = 0x1
)

type BpfInsn uint32

const (
	BPF_INS_INVALID BpfInsn = 0x0
	BPF_INS_ADD     BpfInsn = 0x1
	BPF_INS_SUB     BpfInsn = 0x2
	BPF_INS_MUL     BpfInsn = 0x3
	BPF_INS_DIV     BpfInsn = 0x4
	BPF_INS_OR      BpfInsn = 0x5
	BPF_INS_AND     BpfInsn = 0x6
	BPF_INS_LSH     BpfInsn = 0x7
	BPF_INS_RSH     BpfInsn = 0x8
	BPF_INS_NEG     BpfInsn = 0x9
	BPF_INS_MOD     BpfInsn = 0xa
	BPF_INS_XOR     BpfInsn = 0xb
	BPF_INS_MOV     BpfInsn = 0xc
	BPF_INS_ARSH    BpfInsn = 0xd
	BPF_INS_ADD64   BpfInsn = 0xe
	BPF_INS_SUB64   BpfInsn = 0xf
	BPF_INS_MUL64   BpfInsn = 0x10
	BPF_INS_DIV64   BpfInsn = 0x11
	BPF_INS_OR64    BpfInsn = 0x12
	BPF_INS_AND64   BpfInsn = 0x13
	BPF_INS_LSH64   BpfInsn = 0x14
	BPF_INS_RSH64   BpfInsn = 0x15
	BPF_INS_NEG64   BpfInsn = 0x16
	BPF_INS_MOD64   BpfInsn = 0x17
	BPF_INS_XOR64   BpfInsn = 0x18
	BPF_INS_MOV64   BpfInsn = 0x19
	BPF_INS_ARSH64  BpfInsn = 0x1a
	BPF_INS_LE16    BpfInsn = 0x1b
	BPF_INS_LE32    BpfInsn = 0x1c
	BPF_INS_LE64    BpfInsn = 0x1d
	BPF_INS_BE16    BpfInsn = 0x1e
	BPF_INS_BE32    BpfInsn = 0x1f
	BPF_INS_BE64    BpfInsn = 0x20
	BPF_INS_LDW     BpfInsn = 0x21
	BPF_INS_LDH     BpfInsn = 0x22
	BPF_INS_LDB     BpfInsn = 0x23
	BPF_INS_LDDW    BpfInsn = 0x24
	BPF_INS_LDXW    BpfInsn = 0x25
	BPF_INS_LDXH    BpfInsn = 0x26
	BPF_INS_LDXB    BpfInsn = 0x27
	BPF_INS_LDXDW   BpfInsn = 0x28
	BPF_INS_STW     BpfInsn = 0x29
	BPF_INS_STH     BpfInsn = 0x2a
	BPF_INS_STB     BpfInsn = 0x2b
	BPF_INS_STDW    BpfInsn = 0x2c
	BPF_INS_STXW    BpfInsn = 0x2d
	BPF_INS_STXH    BpfInsn = 0x2e
	BPF_INS_STXB    BpfInsn = 0x2f
	BPF_INS_STXDW   BpfInsn = 0x30
	BPF_INS_XADDW   BpfInsn = 0x31
	BPF_INS_XADDDW  BpfInsn = 0x32
	BPF_INS_JMP     BpfInsn = 0x33
	BPF_INS_JEQ     BpfInsn = 0x34
	BPF_INS_JGT     BpfInsn = 0x35
	BPF_INS_JGE     BpfInsn = 0x36
	BPF_INS_JSET    BpfInsn = 0x37
	BPF_INS_JNE     BpfInsn = 0x38
	BPF_INS_JSGT    BpfInsn = 0x39
	BPF_INS_JSGE    BpfInsn = 0x3a
	BPF_INS_CALL    BpfInsn = 0x3b
	BPF_INS_EXIT    BpfInsn = 0x3c
	BPF_INS_JLT     BpfInsn = 0x3d
	BPF_INS_JLE     BpfInsn = 0x3e
	BPF_INS_JSLT    BpfInsn = 0x3f
	BPF_INS_JSLE    BpfInsn = 0x40
	BPF_INS_RET     BpfInsn = 0x41
	BPF_INS_TAX     BpfInsn = 0x42
	BPF_INS_TXA     BpfInsn = 0x43
	BPF_INS_ENDING  BpfInsn = 0x44
	BPF_INS_LD      BpfInsn = 0x21
	BPF_INS_LDX     BpfInsn = 0x25
	BPF_INS_ST      BpfInsn = 0x29
	BPF_INS_STX     BpfInsn = 0x2d
)

type BpfInsnGroup uint32

const (
	BPF_GRP_INVALID BpfInsnGroup = 0x0
	BPF_GRP_LOAD    BpfInsnGroup = 0x1
	BPF_GRP_STORE   BpfInsnGroup = 0x2
	BPF_GRP_ALU     BpfInsnGroup = 0x3
	BPF_GRP_JUMP    BpfInsnGroup = 0x4
	BPF_GRP_CALL    BpfInsnGroup = 0x5
	BPF_GRP_RETURN  BpfInsnGroup = 0x6
	BPF_GRP_MISC    BpfInsnGroup = 0x7
	BPF_GRP_ENDING  BpfInsnGroup = 0x8
)

type BpfOpType uint32

const (
	BPF_OP_INVALID BpfOpType = 0x0
	BPF_OP_REG     BpfOpType = 0x1
	BPF_OP_IMM     BpfOpType = 0x2
	BPF_OP_OFF     BpfOpType = 0x3
	BPF_OP_MEM     BpfOpType = 0x4
	BPF_OP_MMEM    BpfOpType = 0x5
	BPF_OP_MSH     BpfOpType = 0x6
	BPF_OP_EXT     BpfOpType = 0x7
)

type BpfReg uint32

const (
	BPF_REG_INVALID BpfReg = 0x0
	BPF_REG_A       BpfReg = 0x1
	BPF_REG_X       BpfReg = 0x2
	BPF_REG_R0      BpfReg = 0x3
	BPF_REG_R1      BpfReg = 0x4
	BPF_REG_R2      BpfReg = 0x5
	BPF_REG_R3      BpfReg = 0x6
	BPF_REG_R4      BpfReg = 0x7
	BPF_REG_R5      BpfReg = 0x8
	BPF_REG_R6      BpfReg = 0x9
	BPF_REG_R7      BpfReg = 0xa
	BPF_REG_R8      BpfReg = 0xb
	BPF_REG_R9      BpfReg = 0xc
	BPF_REG_R10     BpfReg = 0xd
	BPF_REG_ENDING  BpfReg = 0xe
)

type BpfOpMem struct {
	Base uint32
	Disp uint32
}

type CsBpf struct {
	Count    uint8
	_        [4]byte
	Operands [4]CsBpfOp
}

type CsBpfOp struct {
	Type   uint32
	_      [4]byte
	Reg    uint8
	_      [7]byte
	Access uint8
	_      [7]byte
}
