// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs capstone_unix.go

package capstone

type CsAcType uint32

const (
	CS_AC_INVALID CsAcType = 0x0
	CS_AC_READ    CsAcType = 0x1
	CS_AC_WRITE   CsAcType = 0x2
)

type CsArch uint32

const (
	CS_ARCH_ARM        CsArch = 0x0
	CS_ARCH_ARM64      CsArch = 0x1
	CS_ARCH_MIPS       CsArch = 0x2
	CS_ARCH_X86        CsArch = 0x3
	CS_ARCH_PPC        CsArch = 0x4
	CS_ARCH_SPARC      CsArch = 0x5
	CS_ARCH_SYSZ       CsArch = 0x6
	CS_ARCH_XCORE      CsArch = 0x7
	CS_ARCH_M68K       CsArch = 0x8
	CS_ARCH_TMS320C64X CsArch = 0x9
	CS_ARCH_M680X      CsArch = 0xa
	CS_ARCH_EVM        CsArch = 0xb
	CS_ARCH_MOS65XX    CsArch = 0xc
	CS_ARCH_WASM       CsArch = 0xd
	CS_ARCH_BPF        CsArch = 0xe
	CS_ARCH_RISCV      CsArch = 0xf
	CS_ARCH_MAX        CsArch = 0x10
	CS_ARCH_ALL        CsArch = 0xffff
)

type CsErr uint32

const (
	CS_ERR_OK        CsErr = 0x0
	CS_ERR_MEM       CsErr = 0x1
	CS_ERR_ARCH      CsErr = 0x2
	CS_ERR_HANDLE    CsErr = 0x3
	CS_ERR_CSH       CsErr = 0x4
	CS_ERR_MODE      CsErr = 0x5
	CS_ERR_OPTION    CsErr = 0x6
	CS_ERR_DETAIL    CsErr = 0x7
	CS_ERR_MEMSETUP  CsErr = 0x8
	CS_ERR_VERSION   CsErr = 0x9
	CS_ERR_DIET      CsErr = 0xa
	CS_ERR_SKIPDATA  CsErr = 0xb
	CS_ERR_X86_ATT   CsErr = 0xc
	CS_ERR_X86_INTEL CsErr = 0xd
	CS_ERR_X86_MASM  CsErr = 0xe
)

type CsGroupType uint32

const (
	CS_GRP_INVALID         CsGroupType = 0x0
	CS_GRP_JUMP            CsGroupType = 0x1
	CS_GRP_CALL            CsGroupType = 0x2
	CS_GRP_RET             CsGroupType = 0x3
	CS_GRP_INT             CsGroupType = 0x4
	CS_GRP_IRET            CsGroupType = 0x5
	CS_GRP_PRIVILEGE       CsGroupType = 0x6
	CS_GRP_BRANCH_RELATIVE CsGroupType = 0x7
)

type CsMode uint32

const (
	CS_MODE_LITTLE_ENDIAN         CsMode = 0x0
	CS_MODE_ARM                   CsMode = 0x0
	CS_MODE_16                    CsMode = 0x2
	CS_MODE_32                    CsMode = 0x4
	CS_MODE_64                    CsMode = 0x8
	CS_MODE_THUMB                 CsMode = 0x10
	CS_MODE_MCLASS                CsMode = 0x20
	CS_MODE_V8                    CsMode = 0x40
	CS_MODE_MICRO                 CsMode = 0x10
	CS_MODE_MIPS3                 CsMode = 0x20
	CS_MODE_MIPS32R6              CsMode = 0x40
	CS_MODE_MIPS2                 CsMode = 0x80
	CS_MODE_V9                    CsMode = 0x10
	CS_MODE_QPX                   CsMode = 0x10
	CS_MODE_SPE                   CsMode = 0x20
	CS_MODE_BOOKE                 CsMode = 0x40
	CS_MODE_M68K_000              CsMode = 0x2
	CS_MODE_M68K_010              CsMode = 0x4
	CS_MODE_M68K_020              CsMode = 0x8
	CS_MODE_M68K_030              CsMode = 0x10
	CS_MODE_M68K_040              CsMode = 0x20
	CS_MODE_M68K_060              CsMode = 0x40
	CS_MODE_BIG_ENDIAN            CsMode = 0x80000000
	CS_MODE_MIPS32                CsMode = 0x4
	CS_MODE_MIPS64                CsMode = 0x8
	CS_MODE_M680X_6301            CsMode = 0x2
	CS_MODE_M680X_6309            CsMode = 0x4
	CS_MODE_M680X_6800            CsMode = 0x8
	CS_MODE_M680X_6801            CsMode = 0x10
	CS_MODE_M680X_6805            CsMode = 0x20
	CS_MODE_M680X_6808            CsMode = 0x40
	CS_MODE_M680X_6809            CsMode = 0x80
	CS_MODE_M680X_6811            CsMode = 0x100
	CS_MODE_M680X_CPU12           CsMode = 0x200
	CS_MODE_M680X_HCS08           CsMode = 0x400
	CS_MODE_BPF_CLASSIC           CsMode = 0x0
	CS_MODE_BPF_EXTENDED          CsMode = 0x1
	CS_MODE_RISCV32               CsMode = 0x1
	CS_MODE_RISCV64               CsMode = 0x2
	CS_MODE_RISCVC                CsMode = 0x4
	CS_MODE_MOS65XX_6502          CsMode = 0x2
	CS_MODE_MOS65XX_65C02         CsMode = 0x4
	CS_MODE_MOS65XX_W65C02        CsMode = 0x8
	CS_MODE_MOS65XX_65816         CsMode = 0x10
	CS_MODE_MOS65XX_65816_LONG_M  CsMode = 0x20
	CS_MODE_MOS65XX_65816_LONG_X  CsMode = 0x40
	CS_MODE_MOS65XX_65816_LONG_MX CsMode = 0x60
)

type CsOpType uint32

const (
	CS_OP_INVALID CsOpType = 0x0
	CS_OP_REG     CsOpType = 0x1
	CS_OP_IMM     CsOpType = 0x2
	CS_OP_MEM     CsOpType = 0x3
	CS_OP_FP      CsOpType = 0x4
)

type CsOptType uint32

const (
	CS_OPT_INVALID        CsOptType = 0x0
	CS_OPT_SYNTAX         CsOptType = 0x1
	CS_OPT_DETAIL         CsOptType = 0x2
	CS_OPT_MODE           CsOptType = 0x3
	CS_OPT_MEM            CsOptType = 0x4
	CS_OPT_SKIPDATA       CsOptType = 0x5
	CS_OPT_SKIPDATA_SETUP CsOptType = 0x6
	CS_OPT_MNEMONIC       CsOptType = 0x7
	CS_OPT_UNSIGNED       CsOptType = 0x8
)

type CsOptValue uint32

const (
	CS_OPT_OFF              CsOptValue = 0x0
	CS_OPT_ON               CsOptValue = 0x3
	CS_OPT_SYNTAX_DEFAULT   CsOptValue = 0x0
	CS_OPT_SYNTAX_INTEL     CsOptValue = 0x1
	CS_OPT_SYNTAX_ATT       CsOptValue = 0x2
	CS_OPT_SYNTAX_NOREGNAME CsOptValue = 0x3
	CS_OPT_SYNTAX_MASM      CsOptValue = 0x4
	CS_OPT_SYNTAX_MOTOROLA  CsOptValue = 0x5
)

type CsOptSkipdata struct {
	Mnemonic *int8
	Callback *[0]byte
	Data     *byte
}

type CsCallocT *[0]byte

type CsFreeT *[0]byte

type CsMallocT *[0]byte

type CsReallocT *[0]byte

type CsRegs [64]uint16

type CsSkipdataCbT *[0]byte

type CsVsnprintfT *[0]byte

type Csh uint64

const CS_API_MAJOR = 0x5

const CS_API_MINOR = 0x0

const CS_NEXT_VERSION = 0x5

const CS_VERSION_MAJOR = 0x5

const CS_VERSION_MINOR = 0x0

const CS_VERSION_EXTRA = 0x0

const CS_MNEMONIC_SIZE = 0x20
