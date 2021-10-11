// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs types_tms320c64x_unix.go

package capstone

type Tms320C64XFunit uint32

const (
	TMS320C64X_FUNIT_INVALID Tms320C64XFunit = 0x0
	TMS320C64X_FUNIT_D       Tms320C64XFunit = 0x1
	TMS320C64X_FUNIT_L       Tms320C64XFunit = 0x2
	TMS320C64X_FUNIT_M       Tms320C64XFunit = 0x3
	TMS320C64X_FUNIT_S       Tms320C64XFunit = 0x4
	TMS320C64X_FUNIT_NO      Tms320C64XFunit = 0x5
)

type Tms320C64XInsn uint32

const (
	TMS320C64X_INS_INVALID   Tms320C64XInsn = 0x0
	TMS320C64X_INS_ABS       Tms320C64XInsn = 0x1
	TMS320C64X_INS_ABS2      Tms320C64XInsn = 0x2
	TMS320C64X_INS_ADD       Tms320C64XInsn = 0x3
	TMS320C64X_INS_ADD2      Tms320C64XInsn = 0x4
	TMS320C64X_INS_ADD4      Tms320C64XInsn = 0x5
	TMS320C64X_INS_ADDAB     Tms320C64XInsn = 0x6
	TMS320C64X_INS_ADDAD     Tms320C64XInsn = 0x7
	TMS320C64X_INS_ADDAH     Tms320C64XInsn = 0x8
	TMS320C64X_INS_ADDAW     Tms320C64XInsn = 0x9
	TMS320C64X_INS_ADDK      Tms320C64XInsn = 0xa
	TMS320C64X_INS_ADDKPC    Tms320C64XInsn = 0xb
	TMS320C64X_INS_ADDU      Tms320C64XInsn = 0xc
	TMS320C64X_INS_AND       Tms320C64XInsn = 0xd
	TMS320C64X_INS_ANDN      Tms320C64XInsn = 0xe
	TMS320C64X_INS_AVG2      Tms320C64XInsn = 0xf
	TMS320C64X_INS_AVGU4     Tms320C64XInsn = 0x10
	TMS320C64X_INS_B         Tms320C64XInsn = 0x11
	TMS320C64X_INS_BDEC      Tms320C64XInsn = 0x12
	TMS320C64X_INS_BITC4     Tms320C64XInsn = 0x13
	TMS320C64X_INS_BNOP      Tms320C64XInsn = 0x14
	TMS320C64X_INS_BPOS      Tms320C64XInsn = 0x15
	TMS320C64X_INS_CLR       Tms320C64XInsn = 0x16
	TMS320C64X_INS_CMPEQ     Tms320C64XInsn = 0x17
	TMS320C64X_INS_CMPEQ2    Tms320C64XInsn = 0x18
	TMS320C64X_INS_CMPEQ4    Tms320C64XInsn = 0x19
	TMS320C64X_INS_CMPGT     Tms320C64XInsn = 0x1a
	TMS320C64X_INS_CMPGT2    Tms320C64XInsn = 0x1b
	TMS320C64X_INS_CMPGTU4   Tms320C64XInsn = 0x1c
	TMS320C64X_INS_CMPLT     Tms320C64XInsn = 0x1d
	TMS320C64X_INS_CMPLTU    Tms320C64XInsn = 0x1e
	TMS320C64X_INS_DEAL      Tms320C64XInsn = 0x1f
	TMS320C64X_INS_DOTP2     Tms320C64XInsn = 0x20
	TMS320C64X_INS_DOTPN2    Tms320C64XInsn = 0x21
	TMS320C64X_INS_DOTPNRSU2 Tms320C64XInsn = 0x22
	TMS320C64X_INS_DOTPRSU2  Tms320C64XInsn = 0x23
	TMS320C64X_INS_DOTPSU4   Tms320C64XInsn = 0x24
	TMS320C64X_INS_DOTPU4    Tms320C64XInsn = 0x25
	TMS320C64X_INS_EXT       Tms320C64XInsn = 0x26
	TMS320C64X_INS_EXTU      Tms320C64XInsn = 0x27
	TMS320C64X_INS_GMPGTU    Tms320C64XInsn = 0x28
	TMS320C64X_INS_GMPY4     Tms320C64XInsn = 0x29
	TMS320C64X_INS_LDB       Tms320C64XInsn = 0x2a
	TMS320C64X_INS_LDBU      Tms320C64XInsn = 0x2b
	TMS320C64X_INS_LDDW      Tms320C64XInsn = 0x2c
	TMS320C64X_INS_LDH       Tms320C64XInsn = 0x2d
	TMS320C64X_INS_LDHU      Tms320C64XInsn = 0x2e
	TMS320C64X_INS_LDNDW     Tms320C64XInsn = 0x2f
	TMS320C64X_INS_LDNW      Tms320C64XInsn = 0x30
	TMS320C64X_INS_LDW       Tms320C64XInsn = 0x31
	TMS320C64X_INS_LMBD      Tms320C64XInsn = 0x32
	TMS320C64X_INS_MAX2      Tms320C64XInsn = 0x33
	TMS320C64X_INS_MAXU4     Tms320C64XInsn = 0x34
	TMS320C64X_INS_MIN2      Tms320C64XInsn = 0x35
	TMS320C64X_INS_MINU4     Tms320C64XInsn = 0x36
	TMS320C64X_INS_MPY       Tms320C64XInsn = 0x37
	TMS320C64X_INS_MPY2      Tms320C64XInsn = 0x38
	TMS320C64X_INS_MPYH      Tms320C64XInsn = 0x39
	TMS320C64X_INS_MPYHI     Tms320C64XInsn = 0x3a
	TMS320C64X_INS_MPYHIR    Tms320C64XInsn = 0x3b
	TMS320C64X_INS_MPYHL     Tms320C64XInsn = 0x3c
	TMS320C64X_INS_MPYHLU    Tms320C64XInsn = 0x3d
	TMS320C64X_INS_MPYHSLU   Tms320C64XInsn = 0x3e
	TMS320C64X_INS_MPYHSU    Tms320C64XInsn = 0x3f
	TMS320C64X_INS_MPYHU     Tms320C64XInsn = 0x40
	TMS320C64X_INS_MPYHULS   Tms320C64XInsn = 0x41
	TMS320C64X_INS_MPYHUS    Tms320C64XInsn = 0x42
	TMS320C64X_INS_MPYLH     Tms320C64XInsn = 0x43
	TMS320C64X_INS_MPYLHU    Tms320C64XInsn = 0x44
	TMS320C64X_INS_MPYLI     Tms320C64XInsn = 0x45
	TMS320C64X_INS_MPYLIR    Tms320C64XInsn = 0x46
	TMS320C64X_INS_MPYLSHU   Tms320C64XInsn = 0x47
	TMS320C64X_INS_MPYLUHS   Tms320C64XInsn = 0x48
	TMS320C64X_INS_MPYSU     Tms320C64XInsn = 0x49
	TMS320C64X_INS_MPYSU4    Tms320C64XInsn = 0x4a
	TMS320C64X_INS_MPYU      Tms320C64XInsn = 0x4b
	TMS320C64X_INS_MPYU4     Tms320C64XInsn = 0x4c
	TMS320C64X_INS_MPYUS     Tms320C64XInsn = 0x4d
	TMS320C64X_INS_MVC       Tms320C64XInsn = 0x4e
	TMS320C64X_INS_MVD       Tms320C64XInsn = 0x4f
	TMS320C64X_INS_MVK       Tms320C64XInsn = 0x50
	TMS320C64X_INS_MVKL      Tms320C64XInsn = 0x51
	TMS320C64X_INS_MVKLH     Tms320C64XInsn = 0x52
	TMS320C64X_INS_NOP       Tms320C64XInsn = 0x53
	TMS320C64X_INS_NORM      Tms320C64XInsn = 0x54
	TMS320C64X_INS_OR        Tms320C64XInsn = 0x55
	TMS320C64X_INS_PACK2     Tms320C64XInsn = 0x56
	TMS320C64X_INS_PACKH2    Tms320C64XInsn = 0x57
	TMS320C64X_INS_PACKH4    Tms320C64XInsn = 0x58
	TMS320C64X_INS_PACKHL2   Tms320C64XInsn = 0x59
	TMS320C64X_INS_PACKL4    Tms320C64XInsn = 0x5a
	TMS320C64X_INS_PACKLH2   Tms320C64XInsn = 0x5b
	TMS320C64X_INS_ROTL      Tms320C64XInsn = 0x5c
	TMS320C64X_INS_SADD      Tms320C64XInsn = 0x5d
	TMS320C64X_INS_SADD2     Tms320C64XInsn = 0x5e
	TMS320C64X_INS_SADDU4    Tms320C64XInsn = 0x5f
	TMS320C64X_INS_SADDUS2   Tms320C64XInsn = 0x60
	TMS320C64X_INS_SAT       Tms320C64XInsn = 0x61
	TMS320C64X_INS_SET       Tms320C64XInsn = 0x62
	TMS320C64X_INS_SHFL      Tms320C64XInsn = 0x63
	TMS320C64X_INS_SHL       Tms320C64XInsn = 0x64
	TMS320C64X_INS_SHLMB     Tms320C64XInsn = 0x65
	TMS320C64X_INS_SHR       Tms320C64XInsn = 0x66
	TMS320C64X_INS_SHR2      Tms320C64XInsn = 0x67
	TMS320C64X_INS_SHRMB     Tms320C64XInsn = 0x68
	TMS320C64X_INS_SHRU      Tms320C64XInsn = 0x69
	TMS320C64X_INS_SHRU2     Tms320C64XInsn = 0x6a
	TMS320C64X_INS_SMPY      Tms320C64XInsn = 0x6b
	TMS320C64X_INS_SMPY2     Tms320C64XInsn = 0x6c
	TMS320C64X_INS_SMPYH     Tms320C64XInsn = 0x6d
	TMS320C64X_INS_SMPYHL    Tms320C64XInsn = 0x6e
	TMS320C64X_INS_SMPYLH    Tms320C64XInsn = 0x6f
	TMS320C64X_INS_SPACK2    Tms320C64XInsn = 0x70
	TMS320C64X_INS_SPACKU4   Tms320C64XInsn = 0x71
	TMS320C64X_INS_SSHL      Tms320C64XInsn = 0x72
	TMS320C64X_INS_SSHVL     Tms320C64XInsn = 0x73
	TMS320C64X_INS_SSHVR     Tms320C64XInsn = 0x74
	TMS320C64X_INS_SSUB      Tms320C64XInsn = 0x75
	TMS320C64X_INS_STB       Tms320C64XInsn = 0x76
	TMS320C64X_INS_STDW      Tms320C64XInsn = 0x77
	TMS320C64X_INS_STH       Tms320C64XInsn = 0x78
	TMS320C64X_INS_STNDW     Tms320C64XInsn = 0x79
	TMS320C64X_INS_STNW      Tms320C64XInsn = 0x7a
	TMS320C64X_INS_STW       Tms320C64XInsn = 0x7b
	TMS320C64X_INS_SUB       Tms320C64XInsn = 0x7c
	TMS320C64X_INS_SUB2      Tms320C64XInsn = 0x7d
	TMS320C64X_INS_SUB4      Tms320C64XInsn = 0x7e
	TMS320C64X_INS_SUBAB     Tms320C64XInsn = 0x7f
	TMS320C64X_INS_SUBABS4   Tms320C64XInsn = 0x80
	TMS320C64X_INS_SUBAH     Tms320C64XInsn = 0x81
	TMS320C64X_INS_SUBAW     Tms320C64XInsn = 0x82
	TMS320C64X_INS_SUBC      Tms320C64XInsn = 0x83
	TMS320C64X_INS_SUBU      Tms320C64XInsn = 0x84
	TMS320C64X_INS_SWAP4     Tms320C64XInsn = 0x85
	TMS320C64X_INS_UNPKHU4   Tms320C64XInsn = 0x86
	TMS320C64X_INS_UNPKLU4   Tms320C64XInsn = 0x87
	TMS320C64X_INS_XOR       Tms320C64XInsn = 0x88
	TMS320C64X_INS_XPND2     Tms320C64XInsn = 0x89
	TMS320C64X_INS_XPND4     Tms320C64XInsn = 0x8a
	TMS320C64X_INS_IDLE      Tms320C64XInsn = 0x8b
	TMS320C64X_INS_MV        Tms320C64XInsn = 0x8c
	TMS320C64X_INS_NEG       Tms320C64XInsn = 0x8d
	TMS320C64X_INS_NOT       Tms320C64XInsn = 0x8e
	TMS320C64X_INS_SWAP2     Tms320C64XInsn = 0x8f
	TMS320C64X_INS_ZERO      Tms320C64XInsn = 0x90
	TMS320C64X_INS_ENDING    Tms320C64XInsn = 0x91
)

type Tms320C64XInsnGroup uint32

const (
	TMS320C64X_GRP_INVALID  Tms320C64XInsnGroup = 0x0
	TMS320C64X_GRP_JUMP     Tms320C64XInsnGroup = 0x1
	TMS320C64X_GRP_FUNIT_D  Tms320C64XInsnGroup = 0x80
	TMS320C64X_GRP_FUNIT_L  Tms320C64XInsnGroup = 0x81
	TMS320C64X_GRP_FUNIT_M  Tms320C64XInsnGroup = 0x82
	TMS320C64X_GRP_FUNIT_S  Tms320C64XInsnGroup = 0x83
	TMS320C64X_GRP_FUNIT_NO Tms320C64XInsnGroup = 0x84
	TMS320C64X_GRP_ENDING   Tms320C64XInsnGroup = 0x85
)

type Tms320C64XMemDir uint32

const (
	TMS320C64X_MEM_DIR_INVALID Tms320C64XMemDir = 0x0
	TMS320C64X_MEM_DIR_FW      Tms320C64XMemDir = 0x1
	TMS320C64X_MEM_DIR_BW      Tms320C64XMemDir = 0x2
)

type Tms320C64XMemDisp uint32

const (
	TMS320C64X_MEM_DISP_INVALID  Tms320C64XMemDisp = 0x0
	TMS320C64X_MEM_DISP_CONSTANT Tms320C64XMemDisp = 0x1
	TMS320C64X_MEM_DISP_REGISTER Tms320C64XMemDisp = 0x2
)

type Tms320C64XMemMod uint32

const (
	TMS320C64X_MEM_MOD_INVALID Tms320C64XMemMod = 0x0
	TMS320C64X_MEM_MOD_NO      Tms320C64XMemMod = 0x1
	TMS320C64X_MEM_MOD_PRE     Tms320C64XMemMod = 0x2
	TMS320C64X_MEM_MOD_POST    Tms320C64XMemMod = 0x3
)

type Tms320C64XOpType uint32

const (
	TMS320C64X_OP_INVALID Tms320C64XOpType = 0x0
	TMS320C64X_OP_REG     Tms320C64XOpType = 0x1
	TMS320C64X_OP_IMM     Tms320C64XOpType = 0x2
	TMS320C64X_OP_MEM     Tms320C64XOpType = 0x3
	TMS320C64X_OP_REGPAIR Tms320C64XOpType = 0x40
)

type Tms320C64XReg uint32

const (
	TMS320C64X_REG_INVALID Tms320C64XReg = 0x0
	TMS320C64X_REG_AMR     Tms320C64XReg = 0x1
	TMS320C64X_REG_CSR     Tms320C64XReg = 0x2
	TMS320C64X_REG_DIER    Tms320C64XReg = 0x3
	TMS320C64X_REG_DNUM    Tms320C64XReg = 0x4
	TMS320C64X_REG_ECR     Tms320C64XReg = 0x5
	TMS320C64X_REG_GFPGFR  Tms320C64XReg = 0x6
	TMS320C64X_REG_GPLYA   Tms320C64XReg = 0x7
	TMS320C64X_REG_GPLYB   Tms320C64XReg = 0x8
	TMS320C64X_REG_ICR     Tms320C64XReg = 0x9
	TMS320C64X_REG_IER     Tms320C64XReg = 0xa
	TMS320C64X_REG_IERR    Tms320C64XReg = 0xb
	TMS320C64X_REG_ILC     Tms320C64XReg = 0xc
	TMS320C64X_REG_IRP     Tms320C64XReg = 0xd
	TMS320C64X_REG_ISR     Tms320C64XReg = 0xe
	TMS320C64X_REG_ISTP    Tms320C64XReg = 0xf
	TMS320C64X_REG_ITSR    Tms320C64XReg = 0x10
	TMS320C64X_REG_NRP     Tms320C64XReg = 0x11
	TMS320C64X_REG_NTSR    Tms320C64XReg = 0x12
	TMS320C64X_REG_REP     Tms320C64XReg = 0x13
	TMS320C64X_REG_RILC    Tms320C64XReg = 0x14
	TMS320C64X_REG_SSR     Tms320C64XReg = 0x15
	TMS320C64X_REG_TSCH    Tms320C64XReg = 0x16
	TMS320C64X_REG_TSCL    Tms320C64XReg = 0x17
	TMS320C64X_REG_TSR     Tms320C64XReg = 0x18
	TMS320C64X_REG_A0      Tms320C64XReg = 0x19
	TMS320C64X_REG_A1      Tms320C64XReg = 0x1a
	TMS320C64X_REG_A2      Tms320C64XReg = 0x1b
	TMS320C64X_REG_A3      Tms320C64XReg = 0x1c
	TMS320C64X_REG_A4      Tms320C64XReg = 0x1d
	TMS320C64X_REG_A5      Tms320C64XReg = 0x1e
	TMS320C64X_REG_A6      Tms320C64XReg = 0x1f
	TMS320C64X_REG_A7      Tms320C64XReg = 0x20
	TMS320C64X_REG_A8      Tms320C64XReg = 0x21
	TMS320C64X_REG_A9      Tms320C64XReg = 0x22
	TMS320C64X_REG_A10     Tms320C64XReg = 0x23
	TMS320C64X_REG_A11     Tms320C64XReg = 0x24
	TMS320C64X_REG_A12     Tms320C64XReg = 0x25
	TMS320C64X_REG_A13     Tms320C64XReg = 0x26
	TMS320C64X_REG_A14     Tms320C64XReg = 0x27
	TMS320C64X_REG_A15     Tms320C64XReg = 0x28
	TMS320C64X_REG_A16     Tms320C64XReg = 0x29
	TMS320C64X_REG_A17     Tms320C64XReg = 0x2a
	TMS320C64X_REG_A18     Tms320C64XReg = 0x2b
	TMS320C64X_REG_A19     Tms320C64XReg = 0x2c
	TMS320C64X_REG_A20     Tms320C64XReg = 0x2d
	TMS320C64X_REG_A21     Tms320C64XReg = 0x2e
	TMS320C64X_REG_A22     Tms320C64XReg = 0x2f
	TMS320C64X_REG_A23     Tms320C64XReg = 0x30
	TMS320C64X_REG_A24     Tms320C64XReg = 0x31
	TMS320C64X_REG_A25     Tms320C64XReg = 0x32
	TMS320C64X_REG_A26     Tms320C64XReg = 0x33
	TMS320C64X_REG_A27     Tms320C64XReg = 0x34
	TMS320C64X_REG_A28     Tms320C64XReg = 0x35
	TMS320C64X_REG_A29     Tms320C64XReg = 0x36
	TMS320C64X_REG_A30     Tms320C64XReg = 0x37
	TMS320C64X_REG_A31     Tms320C64XReg = 0x38
	TMS320C64X_REG_B0      Tms320C64XReg = 0x39
	TMS320C64X_REG_B1      Tms320C64XReg = 0x3a
	TMS320C64X_REG_B2      Tms320C64XReg = 0x3b
	TMS320C64X_REG_B3      Tms320C64XReg = 0x3c
	TMS320C64X_REG_B4      Tms320C64XReg = 0x3d
	TMS320C64X_REG_B5      Tms320C64XReg = 0x3e
	TMS320C64X_REG_B6      Tms320C64XReg = 0x3f
	TMS320C64X_REG_B7      Tms320C64XReg = 0x40
	TMS320C64X_REG_B8      Tms320C64XReg = 0x41
	TMS320C64X_REG_B9      Tms320C64XReg = 0x42
	TMS320C64X_REG_B10     Tms320C64XReg = 0x43
	TMS320C64X_REG_B11     Tms320C64XReg = 0x44
	TMS320C64X_REG_B12     Tms320C64XReg = 0x45
	TMS320C64X_REG_B13     Tms320C64XReg = 0x46
	TMS320C64X_REG_B14     Tms320C64XReg = 0x47
	TMS320C64X_REG_B15     Tms320C64XReg = 0x48
	TMS320C64X_REG_B16     Tms320C64XReg = 0x49
	TMS320C64X_REG_B17     Tms320C64XReg = 0x4a
	TMS320C64X_REG_B18     Tms320C64XReg = 0x4b
	TMS320C64X_REG_B19     Tms320C64XReg = 0x4c
	TMS320C64X_REG_B20     Tms320C64XReg = 0x4d
	TMS320C64X_REG_B21     Tms320C64XReg = 0x4e
	TMS320C64X_REG_B22     Tms320C64XReg = 0x4f
	TMS320C64X_REG_B23     Tms320C64XReg = 0x50
	TMS320C64X_REG_B24     Tms320C64XReg = 0x51
	TMS320C64X_REG_B25     Tms320C64XReg = 0x52
	TMS320C64X_REG_B26     Tms320C64XReg = 0x53
	TMS320C64X_REG_B27     Tms320C64XReg = 0x54
	TMS320C64X_REG_B28     Tms320C64XReg = 0x55
	TMS320C64X_REG_B29     Tms320C64XReg = 0x56
	TMS320C64X_REG_B30     Tms320C64XReg = 0x57
	TMS320C64X_REG_B31     Tms320C64XReg = 0x58
	TMS320C64X_REG_PCE1    Tms320C64XReg = 0x59
	TMS320C64X_REG_ENDING  Tms320C64XReg = 0x5a
	TMS320C64X_REG_EFR     Tms320C64XReg = 0x5
	TMS320C64X_REG_IFR     Tms320C64XReg = 0xe
)

type CsTms320C64X struct {
	Count     uint8
	Operands  [8]CsTms320C64XOp
	Condition CsCondition
	Funit     CsFunit
	Parallel  uint32
}

type CsTms320C64XOp struct {
	Type uint32
	Reg  uint32
	_    [24]byte
}

type CsCondition struct {
	Reg  uint32
	Zero uint32
}

type CsFunit struct {
	Unit      uint32
	Side      uint32
	Crosspath uint32
}

type Tms320C64XOpMem struct {
	Base      uint32
	Disp      uint32
	Unit      uint32
	Scaled    uint32
	Disptype  uint32
	Direction uint32
	Modify    uint32
}