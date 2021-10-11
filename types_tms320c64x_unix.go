// Code generated by github.com/go-darwin/tools/cmd/mkgodef; DO NOT EDIT.
// Input to cgo -godefs.

//go:build ignore
// +build ignore

package capstone

/*
#include <capstone/tms320c64x.h>

typedef struct cs_condition {
	unsigned int reg;
	unsigned int zero;
} cs_condition;

typedef struct cs_funit {
	unsigned int unit;
	unsigned int side;
	unsigned int crosspath;
} cs_funit;

typedef struct csTms320c64x {
	uint8_t op_count;
	cs_tms320c64x_op operands[8]; ///< operands for this instruction.
	cs_condition condition;
	cs_funit funit;
	unsigned int parallel;
} csTms320c64x;
*/
import "C"

type Tms320C64XFunit C.enum_tms320c64x_funit

const (
	TMS320C64X_FUNIT_INVALID Tms320C64XFunit = C.TMS320C64X_FUNIT_INVALID
	TMS320C64X_FUNIT_D       Tms320C64XFunit = C.TMS320C64X_FUNIT_D
	TMS320C64X_FUNIT_L       Tms320C64XFunit = C.TMS320C64X_FUNIT_L
	TMS320C64X_FUNIT_M       Tms320C64XFunit = C.TMS320C64X_FUNIT_M
	TMS320C64X_FUNIT_S       Tms320C64XFunit = C.TMS320C64X_FUNIT_S
	TMS320C64X_FUNIT_NO      Tms320C64XFunit = C.TMS320C64X_FUNIT_NO
)

type Tms320C64XInsn C.enum_tms320c64x_insn

const (
	TMS320C64X_INS_INVALID   Tms320C64XInsn = C.TMS320C64X_INS_INVALID
	TMS320C64X_INS_ABS       Tms320C64XInsn = C.TMS320C64X_INS_ABS
	TMS320C64X_INS_ABS2      Tms320C64XInsn = C.TMS320C64X_INS_ABS2
	TMS320C64X_INS_ADD       Tms320C64XInsn = C.TMS320C64X_INS_ADD
	TMS320C64X_INS_ADD2      Tms320C64XInsn = C.TMS320C64X_INS_ADD2
	TMS320C64X_INS_ADD4      Tms320C64XInsn = C.TMS320C64X_INS_ADD4
	TMS320C64X_INS_ADDAB     Tms320C64XInsn = C.TMS320C64X_INS_ADDAB
	TMS320C64X_INS_ADDAD     Tms320C64XInsn = C.TMS320C64X_INS_ADDAD
	TMS320C64X_INS_ADDAH     Tms320C64XInsn = C.TMS320C64X_INS_ADDAH
	TMS320C64X_INS_ADDAW     Tms320C64XInsn = C.TMS320C64X_INS_ADDAW
	TMS320C64X_INS_ADDK      Tms320C64XInsn = C.TMS320C64X_INS_ADDK
	TMS320C64X_INS_ADDKPC    Tms320C64XInsn = C.TMS320C64X_INS_ADDKPC
	TMS320C64X_INS_ADDU      Tms320C64XInsn = C.TMS320C64X_INS_ADDU
	TMS320C64X_INS_AND       Tms320C64XInsn = C.TMS320C64X_INS_AND
	TMS320C64X_INS_ANDN      Tms320C64XInsn = C.TMS320C64X_INS_ANDN
	TMS320C64X_INS_AVG2      Tms320C64XInsn = C.TMS320C64X_INS_AVG2
	TMS320C64X_INS_AVGU4     Tms320C64XInsn = C.TMS320C64X_INS_AVGU4
	TMS320C64X_INS_B         Tms320C64XInsn = C.TMS320C64X_INS_B
	TMS320C64X_INS_BDEC      Tms320C64XInsn = C.TMS320C64X_INS_BDEC
	TMS320C64X_INS_BITC4     Tms320C64XInsn = C.TMS320C64X_INS_BITC4
	TMS320C64X_INS_BNOP      Tms320C64XInsn = C.TMS320C64X_INS_BNOP
	TMS320C64X_INS_BPOS      Tms320C64XInsn = C.TMS320C64X_INS_BPOS
	TMS320C64X_INS_CLR       Tms320C64XInsn = C.TMS320C64X_INS_CLR
	TMS320C64X_INS_CMPEQ     Tms320C64XInsn = C.TMS320C64X_INS_CMPEQ
	TMS320C64X_INS_CMPEQ2    Tms320C64XInsn = C.TMS320C64X_INS_CMPEQ2
	TMS320C64X_INS_CMPEQ4    Tms320C64XInsn = C.TMS320C64X_INS_CMPEQ4
	TMS320C64X_INS_CMPGT     Tms320C64XInsn = C.TMS320C64X_INS_CMPGT
	TMS320C64X_INS_CMPGT2    Tms320C64XInsn = C.TMS320C64X_INS_CMPGT2
	TMS320C64X_INS_CMPGTU4   Tms320C64XInsn = C.TMS320C64X_INS_CMPGTU4
	TMS320C64X_INS_CMPLT     Tms320C64XInsn = C.TMS320C64X_INS_CMPLT
	TMS320C64X_INS_CMPLTU    Tms320C64XInsn = C.TMS320C64X_INS_CMPLTU
	TMS320C64X_INS_DEAL      Tms320C64XInsn = C.TMS320C64X_INS_DEAL
	TMS320C64X_INS_DOTP2     Tms320C64XInsn = C.TMS320C64X_INS_DOTP2
	TMS320C64X_INS_DOTPN2    Tms320C64XInsn = C.TMS320C64X_INS_DOTPN2
	TMS320C64X_INS_DOTPNRSU2 Tms320C64XInsn = C.TMS320C64X_INS_DOTPNRSU2
	TMS320C64X_INS_DOTPRSU2  Tms320C64XInsn = C.TMS320C64X_INS_DOTPRSU2
	TMS320C64X_INS_DOTPSU4   Tms320C64XInsn = C.TMS320C64X_INS_DOTPSU4
	TMS320C64X_INS_DOTPU4    Tms320C64XInsn = C.TMS320C64X_INS_DOTPU4
	TMS320C64X_INS_EXT       Tms320C64XInsn = C.TMS320C64X_INS_EXT
	TMS320C64X_INS_EXTU      Tms320C64XInsn = C.TMS320C64X_INS_EXTU
	TMS320C64X_INS_GMPGTU    Tms320C64XInsn = C.TMS320C64X_INS_GMPGTU
	TMS320C64X_INS_GMPY4     Tms320C64XInsn = C.TMS320C64X_INS_GMPY4
	TMS320C64X_INS_LDB       Tms320C64XInsn = C.TMS320C64X_INS_LDB
	TMS320C64X_INS_LDBU      Tms320C64XInsn = C.TMS320C64X_INS_LDBU
	TMS320C64X_INS_LDDW      Tms320C64XInsn = C.TMS320C64X_INS_LDDW
	TMS320C64X_INS_LDH       Tms320C64XInsn = C.TMS320C64X_INS_LDH
	TMS320C64X_INS_LDHU      Tms320C64XInsn = C.TMS320C64X_INS_LDHU
	TMS320C64X_INS_LDNDW     Tms320C64XInsn = C.TMS320C64X_INS_LDNDW
	TMS320C64X_INS_LDNW      Tms320C64XInsn = C.TMS320C64X_INS_LDNW
	TMS320C64X_INS_LDW       Tms320C64XInsn = C.TMS320C64X_INS_LDW
	TMS320C64X_INS_LMBD      Tms320C64XInsn = C.TMS320C64X_INS_LMBD
	TMS320C64X_INS_MAX2      Tms320C64XInsn = C.TMS320C64X_INS_MAX2
	TMS320C64X_INS_MAXU4     Tms320C64XInsn = C.TMS320C64X_INS_MAXU4
	TMS320C64X_INS_MIN2      Tms320C64XInsn = C.TMS320C64X_INS_MIN2
	TMS320C64X_INS_MINU4     Tms320C64XInsn = C.TMS320C64X_INS_MINU4
	TMS320C64X_INS_MPY       Tms320C64XInsn = C.TMS320C64X_INS_MPY
	TMS320C64X_INS_MPY2      Tms320C64XInsn = C.TMS320C64X_INS_MPY2
	TMS320C64X_INS_MPYH      Tms320C64XInsn = C.TMS320C64X_INS_MPYH
	TMS320C64X_INS_MPYHI     Tms320C64XInsn = C.TMS320C64X_INS_MPYHI
	TMS320C64X_INS_MPYHIR    Tms320C64XInsn = C.TMS320C64X_INS_MPYHIR
	TMS320C64X_INS_MPYHL     Tms320C64XInsn = C.TMS320C64X_INS_MPYHL
	TMS320C64X_INS_MPYHLU    Tms320C64XInsn = C.TMS320C64X_INS_MPYHLU
	TMS320C64X_INS_MPYHSLU   Tms320C64XInsn = C.TMS320C64X_INS_MPYHSLU
	TMS320C64X_INS_MPYHSU    Tms320C64XInsn = C.TMS320C64X_INS_MPYHSU
	TMS320C64X_INS_MPYHU     Tms320C64XInsn = C.TMS320C64X_INS_MPYHU
	TMS320C64X_INS_MPYHULS   Tms320C64XInsn = C.TMS320C64X_INS_MPYHULS
	TMS320C64X_INS_MPYHUS    Tms320C64XInsn = C.TMS320C64X_INS_MPYHUS
	TMS320C64X_INS_MPYLH     Tms320C64XInsn = C.TMS320C64X_INS_MPYLH
	TMS320C64X_INS_MPYLHU    Tms320C64XInsn = C.TMS320C64X_INS_MPYLHU
	TMS320C64X_INS_MPYLI     Tms320C64XInsn = C.TMS320C64X_INS_MPYLI
	TMS320C64X_INS_MPYLIR    Tms320C64XInsn = C.TMS320C64X_INS_MPYLIR
	TMS320C64X_INS_MPYLSHU   Tms320C64XInsn = C.TMS320C64X_INS_MPYLSHU
	TMS320C64X_INS_MPYLUHS   Tms320C64XInsn = C.TMS320C64X_INS_MPYLUHS
	TMS320C64X_INS_MPYSU     Tms320C64XInsn = C.TMS320C64X_INS_MPYSU
	TMS320C64X_INS_MPYSU4    Tms320C64XInsn = C.TMS320C64X_INS_MPYSU4
	TMS320C64X_INS_MPYU      Tms320C64XInsn = C.TMS320C64X_INS_MPYU
	TMS320C64X_INS_MPYU4     Tms320C64XInsn = C.TMS320C64X_INS_MPYU4
	TMS320C64X_INS_MPYUS     Tms320C64XInsn = C.TMS320C64X_INS_MPYUS
	TMS320C64X_INS_MVC       Tms320C64XInsn = C.TMS320C64X_INS_MVC
	TMS320C64X_INS_MVD       Tms320C64XInsn = C.TMS320C64X_INS_MVD
	TMS320C64X_INS_MVK       Tms320C64XInsn = C.TMS320C64X_INS_MVK
	TMS320C64X_INS_MVKL      Tms320C64XInsn = C.TMS320C64X_INS_MVKL
	TMS320C64X_INS_MVKLH     Tms320C64XInsn = C.TMS320C64X_INS_MVKLH
	TMS320C64X_INS_NOP       Tms320C64XInsn = C.TMS320C64X_INS_NOP
	TMS320C64X_INS_NORM      Tms320C64XInsn = C.TMS320C64X_INS_NORM
	TMS320C64X_INS_OR        Tms320C64XInsn = C.TMS320C64X_INS_OR
	TMS320C64X_INS_PACK2     Tms320C64XInsn = C.TMS320C64X_INS_PACK2
	TMS320C64X_INS_PACKH2    Tms320C64XInsn = C.TMS320C64X_INS_PACKH2
	TMS320C64X_INS_PACKH4    Tms320C64XInsn = C.TMS320C64X_INS_PACKH4
	TMS320C64X_INS_PACKHL2   Tms320C64XInsn = C.TMS320C64X_INS_PACKHL2
	TMS320C64X_INS_PACKL4    Tms320C64XInsn = C.TMS320C64X_INS_PACKL4
	TMS320C64X_INS_PACKLH2   Tms320C64XInsn = C.TMS320C64X_INS_PACKLH2
	TMS320C64X_INS_ROTL      Tms320C64XInsn = C.TMS320C64X_INS_ROTL
	TMS320C64X_INS_SADD      Tms320C64XInsn = C.TMS320C64X_INS_SADD
	TMS320C64X_INS_SADD2     Tms320C64XInsn = C.TMS320C64X_INS_SADD2
	TMS320C64X_INS_SADDU4    Tms320C64XInsn = C.TMS320C64X_INS_SADDU4
	TMS320C64X_INS_SADDUS2   Tms320C64XInsn = C.TMS320C64X_INS_SADDUS2
	TMS320C64X_INS_SAT       Tms320C64XInsn = C.TMS320C64X_INS_SAT
	TMS320C64X_INS_SET       Tms320C64XInsn = C.TMS320C64X_INS_SET
	TMS320C64X_INS_SHFL      Tms320C64XInsn = C.TMS320C64X_INS_SHFL
	TMS320C64X_INS_SHL       Tms320C64XInsn = C.TMS320C64X_INS_SHL
	TMS320C64X_INS_SHLMB     Tms320C64XInsn = C.TMS320C64X_INS_SHLMB
	TMS320C64X_INS_SHR       Tms320C64XInsn = C.TMS320C64X_INS_SHR
	TMS320C64X_INS_SHR2      Tms320C64XInsn = C.TMS320C64X_INS_SHR2
	TMS320C64X_INS_SHRMB     Tms320C64XInsn = C.TMS320C64X_INS_SHRMB
	TMS320C64X_INS_SHRU      Tms320C64XInsn = C.TMS320C64X_INS_SHRU
	TMS320C64X_INS_SHRU2     Tms320C64XInsn = C.TMS320C64X_INS_SHRU2
	TMS320C64X_INS_SMPY      Tms320C64XInsn = C.TMS320C64X_INS_SMPY
	TMS320C64X_INS_SMPY2     Tms320C64XInsn = C.TMS320C64X_INS_SMPY2
	TMS320C64X_INS_SMPYH     Tms320C64XInsn = C.TMS320C64X_INS_SMPYH
	TMS320C64X_INS_SMPYHL    Tms320C64XInsn = C.TMS320C64X_INS_SMPYHL
	TMS320C64X_INS_SMPYLH    Tms320C64XInsn = C.TMS320C64X_INS_SMPYLH
	TMS320C64X_INS_SPACK2    Tms320C64XInsn = C.TMS320C64X_INS_SPACK2
	TMS320C64X_INS_SPACKU4   Tms320C64XInsn = C.TMS320C64X_INS_SPACKU4
	TMS320C64X_INS_SSHL      Tms320C64XInsn = C.TMS320C64X_INS_SSHL
	TMS320C64X_INS_SSHVL     Tms320C64XInsn = C.TMS320C64X_INS_SSHVL
	TMS320C64X_INS_SSHVR     Tms320C64XInsn = C.TMS320C64X_INS_SSHVR
	TMS320C64X_INS_SSUB      Tms320C64XInsn = C.TMS320C64X_INS_SSUB
	TMS320C64X_INS_STB       Tms320C64XInsn = C.TMS320C64X_INS_STB
	TMS320C64X_INS_STDW      Tms320C64XInsn = C.TMS320C64X_INS_STDW
	TMS320C64X_INS_STH       Tms320C64XInsn = C.TMS320C64X_INS_STH
	TMS320C64X_INS_STNDW     Tms320C64XInsn = C.TMS320C64X_INS_STNDW
	TMS320C64X_INS_STNW      Tms320C64XInsn = C.TMS320C64X_INS_STNW
	TMS320C64X_INS_STW       Tms320C64XInsn = C.TMS320C64X_INS_STW
	TMS320C64X_INS_SUB       Tms320C64XInsn = C.TMS320C64X_INS_SUB
	TMS320C64X_INS_SUB2      Tms320C64XInsn = C.TMS320C64X_INS_SUB2
	TMS320C64X_INS_SUB4      Tms320C64XInsn = C.TMS320C64X_INS_SUB4
	TMS320C64X_INS_SUBAB     Tms320C64XInsn = C.TMS320C64X_INS_SUBAB
	TMS320C64X_INS_SUBABS4   Tms320C64XInsn = C.TMS320C64X_INS_SUBABS4
	TMS320C64X_INS_SUBAH     Tms320C64XInsn = C.TMS320C64X_INS_SUBAH
	TMS320C64X_INS_SUBAW     Tms320C64XInsn = C.TMS320C64X_INS_SUBAW
	TMS320C64X_INS_SUBC      Tms320C64XInsn = C.TMS320C64X_INS_SUBC
	TMS320C64X_INS_SUBU      Tms320C64XInsn = C.TMS320C64X_INS_SUBU
	TMS320C64X_INS_SWAP4     Tms320C64XInsn = C.TMS320C64X_INS_SWAP4
	TMS320C64X_INS_UNPKHU4   Tms320C64XInsn = C.TMS320C64X_INS_UNPKHU4
	TMS320C64X_INS_UNPKLU4   Tms320C64XInsn = C.TMS320C64X_INS_UNPKLU4
	TMS320C64X_INS_XOR       Tms320C64XInsn = C.TMS320C64X_INS_XOR
	TMS320C64X_INS_XPND2     Tms320C64XInsn = C.TMS320C64X_INS_XPND2
	TMS320C64X_INS_XPND4     Tms320C64XInsn = C.TMS320C64X_INS_XPND4
	TMS320C64X_INS_IDLE      Tms320C64XInsn = C.TMS320C64X_INS_IDLE
	TMS320C64X_INS_MV        Tms320C64XInsn = C.TMS320C64X_INS_MV
	TMS320C64X_INS_NEG       Tms320C64XInsn = C.TMS320C64X_INS_NEG
	TMS320C64X_INS_NOT       Tms320C64XInsn = C.TMS320C64X_INS_NOT
	TMS320C64X_INS_SWAP2     Tms320C64XInsn = C.TMS320C64X_INS_SWAP2
	TMS320C64X_INS_ZERO      Tms320C64XInsn = C.TMS320C64X_INS_ZERO
	TMS320C64X_INS_ENDING    Tms320C64XInsn = C.TMS320C64X_INS_ENDING
)

type Tms320C64XInsnGroup C.enum_tms320c64x_insn_group

const (
	TMS320C64X_GRP_INVALID  Tms320C64XInsnGroup = C.TMS320C64X_GRP_INVALID
	TMS320C64X_GRP_JUMP     Tms320C64XInsnGroup = C.TMS320C64X_GRP_JUMP
	TMS320C64X_GRP_FUNIT_D  Tms320C64XInsnGroup = C.TMS320C64X_GRP_FUNIT_D
	TMS320C64X_GRP_FUNIT_L  Tms320C64XInsnGroup = C.TMS320C64X_GRP_FUNIT_L
	TMS320C64X_GRP_FUNIT_M  Tms320C64XInsnGroup = C.TMS320C64X_GRP_FUNIT_M
	TMS320C64X_GRP_FUNIT_S  Tms320C64XInsnGroup = C.TMS320C64X_GRP_FUNIT_S
	TMS320C64X_GRP_FUNIT_NO Tms320C64XInsnGroup = C.TMS320C64X_GRP_FUNIT_NO
	TMS320C64X_GRP_ENDING   Tms320C64XInsnGroup = C.TMS320C64X_GRP_ENDING
)

type Tms320C64XMemDir C.enum_tms320c64x_mem_dir

const (
	TMS320C64X_MEM_DIR_INVALID Tms320C64XMemDir = C.TMS320C64X_MEM_DIR_INVALID
	TMS320C64X_MEM_DIR_FW      Tms320C64XMemDir = C.TMS320C64X_MEM_DIR_FW
	TMS320C64X_MEM_DIR_BW      Tms320C64XMemDir = C.TMS320C64X_MEM_DIR_BW
)

type Tms320C64XMemDisp C.enum_tms320c64x_mem_disp

const (
	TMS320C64X_MEM_DISP_INVALID  Tms320C64XMemDisp = C.TMS320C64X_MEM_DISP_INVALID
	TMS320C64X_MEM_DISP_CONSTANT Tms320C64XMemDisp = C.TMS320C64X_MEM_DISP_CONSTANT
	TMS320C64X_MEM_DISP_REGISTER Tms320C64XMemDisp = C.TMS320C64X_MEM_DISP_REGISTER
)

type Tms320C64XMemMod C.enum_tms320c64x_mem_mod

const (
	TMS320C64X_MEM_MOD_INVALID Tms320C64XMemMod = C.TMS320C64X_MEM_MOD_INVALID
	TMS320C64X_MEM_MOD_NO      Tms320C64XMemMod = C.TMS320C64X_MEM_MOD_NO
	TMS320C64X_MEM_MOD_PRE     Tms320C64XMemMod = C.TMS320C64X_MEM_MOD_PRE
	TMS320C64X_MEM_MOD_POST    Tms320C64XMemMod = C.TMS320C64X_MEM_MOD_POST
)

type Tms320C64XOpType C.enum_tms320c64x_op_type

const (
	TMS320C64X_OP_INVALID Tms320C64XOpType = C.TMS320C64X_OP_INVALID
	TMS320C64X_OP_REG     Tms320C64XOpType = C.TMS320C64X_OP_REG
	TMS320C64X_OP_IMM     Tms320C64XOpType = C.TMS320C64X_OP_IMM
	TMS320C64X_OP_MEM     Tms320C64XOpType = C.TMS320C64X_OP_MEM
	TMS320C64X_OP_REGPAIR Tms320C64XOpType = C.TMS320C64X_OP_REGPAIR
)

type Tms320C64XReg C.enum_tms320c64x_reg

const (
	TMS320C64X_REG_INVALID Tms320C64XReg = C.TMS320C64X_REG_INVALID
	TMS320C64X_REG_AMR     Tms320C64XReg = C.TMS320C64X_REG_AMR
	TMS320C64X_REG_CSR     Tms320C64XReg = C.TMS320C64X_REG_CSR
	TMS320C64X_REG_DIER    Tms320C64XReg = C.TMS320C64X_REG_DIER
	TMS320C64X_REG_DNUM    Tms320C64XReg = C.TMS320C64X_REG_DNUM
	TMS320C64X_REG_ECR     Tms320C64XReg = C.TMS320C64X_REG_ECR
	TMS320C64X_REG_GFPGFR  Tms320C64XReg = C.TMS320C64X_REG_GFPGFR
	TMS320C64X_REG_GPLYA   Tms320C64XReg = C.TMS320C64X_REG_GPLYA
	TMS320C64X_REG_GPLYB   Tms320C64XReg = C.TMS320C64X_REG_GPLYB
	TMS320C64X_REG_ICR     Tms320C64XReg = C.TMS320C64X_REG_ICR
	TMS320C64X_REG_IER     Tms320C64XReg = C.TMS320C64X_REG_IER
	TMS320C64X_REG_IERR    Tms320C64XReg = C.TMS320C64X_REG_IERR
	TMS320C64X_REG_ILC     Tms320C64XReg = C.TMS320C64X_REG_ILC
	TMS320C64X_REG_IRP     Tms320C64XReg = C.TMS320C64X_REG_IRP
	TMS320C64X_REG_ISR     Tms320C64XReg = C.TMS320C64X_REG_ISR
	TMS320C64X_REG_ISTP    Tms320C64XReg = C.TMS320C64X_REG_ISTP
	TMS320C64X_REG_ITSR    Tms320C64XReg = C.TMS320C64X_REG_ITSR
	TMS320C64X_REG_NRP     Tms320C64XReg = C.TMS320C64X_REG_NRP
	TMS320C64X_REG_NTSR    Tms320C64XReg = C.TMS320C64X_REG_NTSR
	TMS320C64X_REG_REP     Tms320C64XReg = C.TMS320C64X_REG_REP
	TMS320C64X_REG_RILC    Tms320C64XReg = C.TMS320C64X_REG_RILC
	TMS320C64X_REG_SSR     Tms320C64XReg = C.TMS320C64X_REG_SSR
	TMS320C64X_REG_TSCH    Tms320C64XReg = C.TMS320C64X_REG_TSCH
	TMS320C64X_REG_TSCL    Tms320C64XReg = C.TMS320C64X_REG_TSCL
	TMS320C64X_REG_TSR     Tms320C64XReg = C.TMS320C64X_REG_TSR
	TMS320C64X_REG_A0      Tms320C64XReg = C.TMS320C64X_REG_A0
	TMS320C64X_REG_A1      Tms320C64XReg = C.TMS320C64X_REG_A1
	TMS320C64X_REG_A2      Tms320C64XReg = C.TMS320C64X_REG_A2
	TMS320C64X_REG_A3      Tms320C64XReg = C.TMS320C64X_REG_A3
	TMS320C64X_REG_A4      Tms320C64XReg = C.TMS320C64X_REG_A4
	TMS320C64X_REG_A5      Tms320C64XReg = C.TMS320C64X_REG_A5
	TMS320C64X_REG_A6      Tms320C64XReg = C.TMS320C64X_REG_A6
	TMS320C64X_REG_A7      Tms320C64XReg = C.TMS320C64X_REG_A7
	TMS320C64X_REG_A8      Tms320C64XReg = C.TMS320C64X_REG_A8
	TMS320C64X_REG_A9      Tms320C64XReg = C.TMS320C64X_REG_A9
	TMS320C64X_REG_A10     Tms320C64XReg = C.TMS320C64X_REG_A10
	TMS320C64X_REG_A11     Tms320C64XReg = C.TMS320C64X_REG_A11
	TMS320C64X_REG_A12     Tms320C64XReg = C.TMS320C64X_REG_A12
	TMS320C64X_REG_A13     Tms320C64XReg = C.TMS320C64X_REG_A13
	TMS320C64X_REG_A14     Tms320C64XReg = C.TMS320C64X_REG_A14
	TMS320C64X_REG_A15     Tms320C64XReg = C.TMS320C64X_REG_A15
	TMS320C64X_REG_A16     Tms320C64XReg = C.TMS320C64X_REG_A16
	TMS320C64X_REG_A17     Tms320C64XReg = C.TMS320C64X_REG_A17
	TMS320C64X_REG_A18     Tms320C64XReg = C.TMS320C64X_REG_A18
	TMS320C64X_REG_A19     Tms320C64XReg = C.TMS320C64X_REG_A19
	TMS320C64X_REG_A20     Tms320C64XReg = C.TMS320C64X_REG_A20
	TMS320C64X_REG_A21     Tms320C64XReg = C.TMS320C64X_REG_A21
	TMS320C64X_REG_A22     Tms320C64XReg = C.TMS320C64X_REG_A22
	TMS320C64X_REG_A23     Tms320C64XReg = C.TMS320C64X_REG_A23
	TMS320C64X_REG_A24     Tms320C64XReg = C.TMS320C64X_REG_A24
	TMS320C64X_REG_A25     Tms320C64XReg = C.TMS320C64X_REG_A25
	TMS320C64X_REG_A26     Tms320C64XReg = C.TMS320C64X_REG_A26
	TMS320C64X_REG_A27     Tms320C64XReg = C.TMS320C64X_REG_A27
	TMS320C64X_REG_A28     Tms320C64XReg = C.TMS320C64X_REG_A28
	TMS320C64X_REG_A29     Tms320C64XReg = C.TMS320C64X_REG_A29
	TMS320C64X_REG_A30     Tms320C64XReg = C.TMS320C64X_REG_A30
	TMS320C64X_REG_A31     Tms320C64XReg = C.TMS320C64X_REG_A31
	TMS320C64X_REG_B0      Tms320C64XReg = C.TMS320C64X_REG_B0
	TMS320C64X_REG_B1      Tms320C64XReg = C.TMS320C64X_REG_B1
	TMS320C64X_REG_B2      Tms320C64XReg = C.TMS320C64X_REG_B2
	TMS320C64X_REG_B3      Tms320C64XReg = C.TMS320C64X_REG_B3
	TMS320C64X_REG_B4      Tms320C64XReg = C.TMS320C64X_REG_B4
	TMS320C64X_REG_B5      Tms320C64XReg = C.TMS320C64X_REG_B5
	TMS320C64X_REG_B6      Tms320C64XReg = C.TMS320C64X_REG_B6
	TMS320C64X_REG_B7      Tms320C64XReg = C.TMS320C64X_REG_B7
	TMS320C64X_REG_B8      Tms320C64XReg = C.TMS320C64X_REG_B8
	TMS320C64X_REG_B9      Tms320C64XReg = C.TMS320C64X_REG_B9
	TMS320C64X_REG_B10     Tms320C64XReg = C.TMS320C64X_REG_B10
	TMS320C64X_REG_B11     Tms320C64XReg = C.TMS320C64X_REG_B11
	TMS320C64X_REG_B12     Tms320C64XReg = C.TMS320C64X_REG_B12
	TMS320C64X_REG_B13     Tms320C64XReg = C.TMS320C64X_REG_B13
	TMS320C64X_REG_B14     Tms320C64XReg = C.TMS320C64X_REG_B14
	TMS320C64X_REG_B15     Tms320C64XReg = C.TMS320C64X_REG_B15
	TMS320C64X_REG_B16     Tms320C64XReg = C.TMS320C64X_REG_B16
	TMS320C64X_REG_B17     Tms320C64XReg = C.TMS320C64X_REG_B17
	TMS320C64X_REG_B18     Tms320C64XReg = C.TMS320C64X_REG_B18
	TMS320C64X_REG_B19     Tms320C64XReg = C.TMS320C64X_REG_B19
	TMS320C64X_REG_B20     Tms320C64XReg = C.TMS320C64X_REG_B20
	TMS320C64X_REG_B21     Tms320C64XReg = C.TMS320C64X_REG_B21
	TMS320C64X_REG_B22     Tms320C64XReg = C.TMS320C64X_REG_B22
	TMS320C64X_REG_B23     Tms320C64XReg = C.TMS320C64X_REG_B23
	TMS320C64X_REG_B24     Tms320C64XReg = C.TMS320C64X_REG_B24
	TMS320C64X_REG_B25     Tms320C64XReg = C.TMS320C64X_REG_B25
	TMS320C64X_REG_B26     Tms320C64XReg = C.TMS320C64X_REG_B26
	TMS320C64X_REG_B27     Tms320C64XReg = C.TMS320C64X_REG_B27
	TMS320C64X_REG_B28     Tms320C64XReg = C.TMS320C64X_REG_B28
	TMS320C64X_REG_B29     Tms320C64XReg = C.TMS320C64X_REG_B29
	TMS320C64X_REG_B30     Tms320C64XReg = C.TMS320C64X_REG_B30
	TMS320C64X_REG_B31     Tms320C64XReg = C.TMS320C64X_REG_B31
	TMS320C64X_REG_PCE1    Tms320C64XReg = C.TMS320C64X_REG_PCE1
	TMS320C64X_REG_ENDING  Tms320C64XReg = C.TMS320C64X_REG_ENDING
	TMS320C64X_REG_EFR     Tms320C64XReg = C.TMS320C64X_REG_EFR
	TMS320C64X_REG_IFR     Tms320C64XReg = C.TMS320C64X_REG_IFR
)

type CsTms320C64X C.csTms320c64x

type CsTms320C64XOp C.struct_cs_tms320c64x_op

type CsCondition C.cs_condition

type CsFunit C.cs_funit

type Tms320C64XOpMem C.struct_tms320c64x_op_mem
