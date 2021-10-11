// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs m68k_unix.go

package capstone

type M68KAddressMode uint32

const (
	M68K_AM_NONE                   M68KAddressMode = 0x0
	M68K_AM_REG_DIRECT_DATA        M68KAddressMode = 0x1
	M68K_AM_REG_DIRECT_ADDR        M68KAddressMode = 0x2
	M68K_AM_REGI_ADDR              M68KAddressMode = 0x3
	M68K_AM_REGI_ADDR_POST_INC     M68KAddressMode = 0x4
	M68K_AM_REGI_ADDR_PRE_DEC      M68KAddressMode = 0x5
	M68K_AM_REGI_ADDR_DISP         M68KAddressMode = 0x6
	M68K_AM_AREGI_INDEX_8_BIT_DISP M68KAddressMode = 0x7
	M68K_AM_AREGI_INDEX_BASE_DISP  M68KAddressMode = 0x8
	M68K_AM_MEMI_POST_INDEX        M68KAddressMode = 0x9
	M68K_AM_MEMI_PRE_INDEX         M68KAddressMode = 0xa
	M68K_AM_PCI_DISP               M68KAddressMode = 0xb
	M68K_AM_PCI_INDEX_8_BIT_DISP   M68KAddressMode = 0xc
	M68K_AM_PCI_INDEX_BASE_DISP    M68KAddressMode = 0xd
	M68K_AM_PC_MEMI_POST_INDEX     M68KAddressMode = 0xe
	M68K_AM_PC_MEMI_PRE_INDEX      M68KAddressMode = 0xf
	M68K_AM_ABSOLUTE_DATA_SHORT    M68KAddressMode = 0x10
	M68K_AM_ABSOLUTE_DATA_LONG     M68KAddressMode = 0x11
	M68K_AM_IMMEDIATE              M68KAddressMode = 0x12
	M68K_AM_BRANCH_DISPLACEMENT    M68KAddressMode = 0x13
)

type M68KCpuSize uint32

const (
	M68K_CPU_SIZE_NONE M68KCpuSize = 0x0
	M68K_CPU_SIZE_BYTE M68KCpuSize = 0x1
	M68K_CPU_SIZE_WORD M68KCpuSize = 0x2
	M68K_CPU_SIZE_LONG M68KCpuSize = 0x4
)

type M68KFpuSize uint32

const (
	M68K_FPU_SIZE_NONE     M68KFpuSize = 0x0
	M68K_FPU_SIZE_SINGLE   M68KFpuSize = 0x4
	M68K_FPU_SIZE_DOUBLE   M68KFpuSize = 0x8
	M68K_FPU_SIZE_EXTENDED M68KFpuSize = 0xc
)

type M68KGroupType uint32

const (
	M68K_GRP_INVALID         M68KGroupType = 0x0
	M68K_GRP_JUMP            M68KGroupType = 0x1
	M68K_GRP_RET             M68KGroupType = 0x3
	M68K_GRP_IRET            M68KGroupType = 0x5
	M68K_GRP_BRANCH_RELATIVE M68KGroupType = 0x7
	M68K_GRP_ENDING          M68KGroupType = 0x8
)

type M68KInsn uint32

const (
	M68K_INS_INVALID   M68KInsn = 0x0
	M68K_INS_ABCD      M68KInsn = 0x1
	M68K_INS_ADD       M68KInsn = 0x2
	M68K_INS_ADDA      M68KInsn = 0x3
	M68K_INS_ADDI      M68KInsn = 0x4
	M68K_INS_ADDQ      M68KInsn = 0x5
	M68K_INS_ADDX      M68KInsn = 0x6
	M68K_INS_AND       M68KInsn = 0x7
	M68K_INS_ANDI      M68KInsn = 0x8
	M68K_INS_ASL       M68KInsn = 0x9
	M68K_INS_ASR       M68KInsn = 0xa
	M68K_INS_BHS       M68KInsn = 0xb
	M68K_INS_BLO       M68KInsn = 0xc
	M68K_INS_BHI       M68KInsn = 0xd
	M68K_INS_BLS       M68KInsn = 0xe
	M68K_INS_BCC       M68KInsn = 0xf
	M68K_INS_BCS       M68KInsn = 0x10
	M68K_INS_BNE       M68KInsn = 0x11
	M68K_INS_BEQ       M68KInsn = 0x12
	M68K_INS_BVC       M68KInsn = 0x13
	M68K_INS_BVS       M68KInsn = 0x14
	M68K_INS_BPL       M68KInsn = 0x15
	M68K_INS_BMI       M68KInsn = 0x16
	M68K_INS_BGE       M68KInsn = 0x17
	M68K_INS_BLT       M68KInsn = 0x18
	M68K_INS_BGT       M68KInsn = 0x19
	M68K_INS_BLE       M68KInsn = 0x1a
	M68K_INS_BRA       M68KInsn = 0x1b
	M68K_INS_BSR       M68KInsn = 0x1c
	M68K_INS_BCHG      M68KInsn = 0x1d
	M68K_INS_BCLR      M68KInsn = 0x1e
	M68K_INS_BSET      M68KInsn = 0x1f
	M68K_INS_BTST      M68KInsn = 0x20
	M68K_INS_BFCHG     M68KInsn = 0x21
	M68K_INS_BFCLR     M68KInsn = 0x22
	M68K_INS_BFEXTS    M68KInsn = 0x23
	M68K_INS_BFEXTU    M68KInsn = 0x24
	M68K_INS_BFFFO     M68KInsn = 0x25
	M68K_INS_BFINS     M68KInsn = 0x26
	M68K_INS_BFSET     M68KInsn = 0x27
	M68K_INS_BFTST     M68KInsn = 0x28
	M68K_INS_BKPT      M68KInsn = 0x29
	M68K_INS_CALLM     M68KInsn = 0x2a
	M68K_INS_CAS       M68KInsn = 0x2b
	M68K_INS_CAS2      M68KInsn = 0x2c
	M68K_INS_CHK       M68KInsn = 0x2d
	M68K_INS_CHK2      M68KInsn = 0x2e
	M68K_INS_CLR       M68KInsn = 0x2f
	M68K_INS_CMP       M68KInsn = 0x30
	M68K_INS_CMPA      M68KInsn = 0x31
	M68K_INS_CMPI      M68KInsn = 0x32
	M68K_INS_CMPM      M68KInsn = 0x33
	M68K_INS_CMP2      M68KInsn = 0x34
	M68K_INS_CINVL     M68KInsn = 0x35
	M68K_INS_CINVP     M68KInsn = 0x36
	M68K_INS_CINVA     M68KInsn = 0x37
	M68K_INS_CPUSHL    M68KInsn = 0x38
	M68K_INS_CPUSHP    M68KInsn = 0x39
	M68K_INS_CPUSHA    M68KInsn = 0x3a
	M68K_INS_DBT       M68KInsn = 0x3b
	M68K_INS_DBF       M68KInsn = 0x3c
	M68K_INS_DBHI      M68KInsn = 0x3d
	M68K_INS_DBLS      M68KInsn = 0x3e
	M68K_INS_DBCC      M68KInsn = 0x3f
	M68K_INS_DBCS      M68KInsn = 0x40
	M68K_INS_DBNE      M68KInsn = 0x41
	M68K_INS_DBEQ      M68KInsn = 0x42
	M68K_INS_DBVC      M68KInsn = 0x43
	M68K_INS_DBVS      M68KInsn = 0x44
	M68K_INS_DBPL      M68KInsn = 0x45
	M68K_INS_DBMI      M68KInsn = 0x46
	M68K_INS_DBGE      M68KInsn = 0x47
	M68K_INS_DBLT      M68KInsn = 0x48
	M68K_INS_DBGT      M68KInsn = 0x49
	M68K_INS_DBLE      M68KInsn = 0x4a
	M68K_INS_DBRA      M68KInsn = 0x4b
	M68K_INS_DIVS      M68KInsn = 0x4c
	M68K_INS_DIVSL     M68KInsn = 0x4d
	M68K_INS_DIVU      M68KInsn = 0x4e
	M68K_INS_DIVUL     M68KInsn = 0x4f
	M68K_INS_EOR       M68KInsn = 0x50
	M68K_INS_EORI      M68KInsn = 0x51
	M68K_INS_EXG       M68KInsn = 0x52
	M68K_INS_EXT       M68KInsn = 0x53
	M68K_INS_EXTB      M68KInsn = 0x54
	M68K_INS_FABS      M68KInsn = 0x55
	M68K_INS_FSABS     M68KInsn = 0x56
	M68K_INS_FDABS     M68KInsn = 0x57
	M68K_INS_FACOS     M68KInsn = 0x58
	M68K_INS_FADD      M68KInsn = 0x59
	M68K_INS_FSADD     M68KInsn = 0x5a
	M68K_INS_FDADD     M68KInsn = 0x5b
	M68K_INS_FASIN     M68KInsn = 0x5c
	M68K_INS_FATAN     M68KInsn = 0x5d
	M68K_INS_FATANH    M68KInsn = 0x5e
	M68K_INS_FBF       M68KInsn = 0x5f
	M68K_INS_FBEQ      M68KInsn = 0x60
	M68K_INS_FBOGT     M68KInsn = 0x61
	M68K_INS_FBOGE     M68KInsn = 0x62
	M68K_INS_FBOLT     M68KInsn = 0x63
	M68K_INS_FBOLE     M68KInsn = 0x64
	M68K_INS_FBOGL     M68KInsn = 0x65
	M68K_INS_FBOR      M68KInsn = 0x66
	M68K_INS_FBUN      M68KInsn = 0x67
	M68K_INS_FBUEQ     M68KInsn = 0x68
	M68K_INS_FBUGT     M68KInsn = 0x69
	M68K_INS_FBUGE     M68KInsn = 0x6a
	M68K_INS_FBULT     M68KInsn = 0x6b
	M68K_INS_FBULE     M68KInsn = 0x6c
	M68K_INS_FBNE      M68KInsn = 0x6d
	M68K_INS_FBT       M68KInsn = 0x6e
	M68K_INS_FBSF      M68KInsn = 0x6f
	M68K_INS_FBSEQ     M68KInsn = 0x70
	M68K_INS_FBGT      M68KInsn = 0x71
	M68K_INS_FBGE      M68KInsn = 0x72
	M68K_INS_FBLT      M68KInsn = 0x73
	M68K_INS_FBLE      M68KInsn = 0x74
	M68K_INS_FBGL      M68KInsn = 0x75
	M68K_INS_FBGLE     M68KInsn = 0x76
	M68K_INS_FBNGLE    M68KInsn = 0x77
	M68K_INS_FBNGL     M68KInsn = 0x78
	M68K_INS_FBNLE     M68KInsn = 0x79
	M68K_INS_FBNLT     M68KInsn = 0x7a
	M68K_INS_FBNGE     M68KInsn = 0x7b
	M68K_INS_FBNGT     M68KInsn = 0x7c
	M68K_INS_FBSNE     M68KInsn = 0x7d
	M68K_INS_FBST      M68KInsn = 0x7e
	M68K_INS_FCMP      M68KInsn = 0x7f
	M68K_INS_FCOS      M68KInsn = 0x80
	M68K_INS_FCOSH     M68KInsn = 0x81
	M68K_INS_FDBF      M68KInsn = 0x82
	M68K_INS_FDBEQ     M68KInsn = 0x83
	M68K_INS_FDBOGT    M68KInsn = 0x84
	M68K_INS_FDBOGE    M68KInsn = 0x85
	M68K_INS_FDBOLT    M68KInsn = 0x86
	M68K_INS_FDBOLE    M68KInsn = 0x87
	M68K_INS_FDBOGL    M68KInsn = 0x88
	M68K_INS_FDBOR     M68KInsn = 0x89
	M68K_INS_FDBUN     M68KInsn = 0x8a
	M68K_INS_FDBUEQ    M68KInsn = 0x8b
	M68K_INS_FDBUGT    M68KInsn = 0x8c
	M68K_INS_FDBUGE    M68KInsn = 0x8d
	M68K_INS_FDBULT    M68KInsn = 0x8e
	M68K_INS_FDBULE    M68KInsn = 0x8f
	M68K_INS_FDBNE     M68KInsn = 0x90
	M68K_INS_FDBT      M68KInsn = 0x91
	M68K_INS_FDBSF     M68KInsn = 0x92
	M68K_INS_FDBSEQ    M68KInsn = 0x93
	M68K_INS_FDBGT     M68KInsn = 0x94
	M68K_INS_FDBGE     M68KInsn = 0x95
	M68K_INS_FDBLT     M68KInsn = 0x96
	M68K_INS_FDBLE     M68KInsn = 0x97
	M68K_INS_FDBGL     M68KInsn = 0x98
	M68K_INS_FDBGLE    M68KInsn = 0x99
	M68K_INS_FDBNGLE   M68KInsn = 0x9a
	M68K_INS_FDBNGL    M68KInsn = 0x9b
	M68K_INS_FDBNLE    M68KInsn = 0x9c
	M68K_INS_FDBNLT    M68KInsn = 0x9d
	M68K_INS_FDBNGE    M68KInsn = 0x9e
	M68K_INS_FDBNGT    M68KInsn = 0x9f
	M68K_INS_FDBSNE    M68KInsn = 0xa0
	M68K_INS_FDBST     M68KInsn = 0xa1
	M68K_INS_FDIV      M68KInsn = 0xa2
	M68K_INS_FSDIV     M68KInsn = 0xa3
	M68K_INS_FDDIV     M68KInsn = 0xa4
	M68K_INS_FETOX     M68KInsn = 0xa5
	M68K_INS_FETOXM1   M68KInsn = 0xa6
	M68K_INS_FGETEXP   M68KInsn = 0xa7
	M68K_INS_FGETMAN   M68KInsn = 0xa8
	M68K_INS_FINT      M68KInsn = 0xa9
	M68K_INS_FINTRZ    M68KInsn = 0xaa
	M68K_INS_FLOG10    M68KInsn = 0xab
	M68K_INS_FLOG2     M68KInsn = 0xac
	M68K_INS_FLOGN     M68KInsn = 0xad
	M68K_INS_FLOGNP1   M68KInsn = 0xae
	M68K_INS_FMOD      M68KInsn = 0xaf
	M68K_INS_FMOVE     M68KInsn = 0xb0
	M68K_INS_FSMOVE    M68KInsn = 0xb1
	M68K_INS_FDMOVE    M68KInsn = 0xb2
	M68K_INS_FMOVECR   M68KInsn = 0xb3
	M68K_INS_FMOVEM    M68KInsn = 0xb4
	M68K_INS_FMUL      M68KInsn = 0xb5
	M68K_INS_FSMUL     M68KInsn = 0xb6
	M68K_INS_FDMUL     M68KInsn = 0xb7
	M68K_INS_FNEG      M68KInsn = 0xb8
	M68K_INS_FSNEG     M68KInsn = 0xb9
	M68K_INS_FDNEG     M68KInsn = 0xba
	M68K_INS_FNOP      M68KInsn = 0xbb
	M68K_INS_FREM      M68KInsn = 0xbc
	M68K_INS_FRESTORE  M68KInsn = 0xbd
	M68K_INS_FSAVE     M68KInsn = 0xbe
	M68K_INS_FSCALE    M68KInsn = 0xbf
	M68K_INS_FSGLDIV   M68KInsn = 0xc0
	M68K_INS_FSGLMUL   M68KInsn = 0xc1
	M68K_INS_FSIN      M68KInsn = 0xc2
	M68K_INS_FSINCOS   M68KInsn = 0xc3
	M68K_INS_FSINH     M68KInsn = 0xc4
	M68K_INS_FSQRT     M68KInsn = 0xc5
	M68K_INS_FSSQRT    M68KInsn = 0xc6
	M68K_INS_FDSQRT    M68KInsn = 0xc7
	M68K_INS_FSF       M68KInsn = 0xc8
	M68K_INS_FSBEQ     M68KInsn = 0xc9
	M68K_INS_FSOGT     M68KInsn = 0xca
	M68K_INS_FSOGE     M68KInsn = 0xcb
	M68K_INS_FSOLT     M68KInsn = 0xcc
	M68K_INS_FSOLE     M68KInsn = 0xcd
	M68K_INS_FSOGL     M68KInsn = 0xce
	M68K_INS_FSOR      M68KInsn = 0xcf
	M68K_INS_FSUN      M68KInsn = 0xd0
	M68K_INS_FSUEQ     M68KInsn = 0xd1
	M68K_INS_FSUGT     M68KInsn = 0xd2
	M68K_INS_FSUGE     M68KInsn = 0xd3
	M68K_INS_FSULT     M68KInsn = 0xd4
	M68K_INS_FSULE     M68KInsn = 0xd5
	M68K_INS_FSNE      M68KInsn = 0xd6
	M68K_INS_FST       M68KInsn = 0xd7
	M68K_INS_FSSF      M68KInsn = 0xd8
	M68K_INS_FSSEQ     M68KInsn = 0xd9
	M68K_INS_FSGT      M68KInsn = 0xda
	M68K_INS_FSGE      M68KInsn = 0xdb
	M68K_INS_FSLT      M68KInsn = 0xdc
	M68K_INS_FSLE      M68KInsn = 0xdd
	M68K_INS_FSGL      M68KInsn = 0xde
	M68K_INS_FSGLE     M68KInsn = 0xdf
	M68K_INS_FSNGLE    M68KInsn = 0xe0
	M68K_INS_FSNGL     M68KInsn = 0xe1
	M68K_INS_FSNLE     M68KInsn = 0xe2
	M68K_INS_FSNLT     M68KInsn = 0xe3
	M68K_INS_FSNGE     M68KInsn = 0xe4
	M68K_INS_FSNGT     M68KInsn = 0xe5
	M68K_INS_FSSNE     M68KInsn = 0xe6
	M68K_INS_FSST      M68KInsn = 0xe7
	M68K_INS_FSUB      M68KInsn = 0xe8
	M68K_INS_FSSUB     M68KInsn = 0xe9
	M68K_INS_FDSUB     M68KInsn = 0xea
	M68K_INS_FTAN      M68KInsn = 0xeb
	M68K_INS_FTANH     M68KInsn = 0xec
	M68K_INS_FTENTOX   M68KInsn = 0xed
	M68K_INS_FTRAPF    M68KInsn = 0xee
	M68K_INS_FTRAPEQ   M68KInsn = 0xef
	M68K_INS_FTRAPOGT  M68KInsn = 0xf0
	M68K_INS_FTRAPOGE  M68KInsn = 0xf1
	M68K_INS_FTRAPOLT  M68KInsn = 0xf2
	M68K_INS_FTRAPOLE  M68KInsn = 0xf3
	M68K_INS_FTRAPOGL  M68KInsn = 0xf4
	M68K_INS_FTRAPOR   M68KInsn = 0xf5
	M68K_INS_FTRAPUN   M68KInsn = 0xf6
	M68K_INS_FTRAPUEQ  M68KInsn = 0xf7
	M68K_INS_FTRAPUGT  M68KInsn = 0xf8
	M68K_INS_FTRAPUGE  M68KInsn = 0xf9
	M68K_INS_FTRAPULT  M68KInsn = 0xfa
	M68K_INS_FTRAPULE  M68KInsn = 0xfb
	M68K_INS_FTRAPNE   M68KInsn = 0xfc
	M68K_INS_FTRAPT    M68KInsn = 0xfd
	M68K_INS_FTRAPSF   M68KInsn = 0xfe
	M68K_INS_FTRAPSEQ  M68KInsn = 0xff
	M68K_INS_FTRAPGT   M68KInsn = 0x100
	M68K_INS_FTRAPGE   M68KInsn = 0x101
	M68K_INS_FTRAPLT   M68KInsn = 0x102
	M68K_INS_FTRAPLE   M68KInsn = 0x103
	M68K_INS_FTRAPGL   M68KInsn = 0x104
	M68K_INS_FTRAPGLE  M68KInsn = 0x105
	M68K_INS_FTRAPNGLE M68KInsn = 0x106
	M68K_INS_FTRAPNGL  M68KInsn = 0x107
	M68K_INS_FTRAPNLE  M68KInsn = 0x108
	M68K_INS_FTRAPNLT  M68KInsn = 0x109
	M68K_INS_FTRAPNGE  M68KInsn = 0x10a
	M68K_INS_FTRAPNGT  M68KInsn = 0x10b
	M68K_INS_FTRAPSNE  M68KInsn = 0x10c
	M68K_INS_FTRAPST   M68KInsn = 0x10d
	M68K_INS_FTST      M68KInsn = 0x10e
	M68K_INS_FTWOTOX   M68KInsn = 0x10f
	M68K_INS_HALT      M68KInsn = 0x110
	M68K_INS_ILLEGAL   M68KInsn = 0x111
	M68K_INS_JMP       M68KInsn = 0x112
	M68K_INS_JSR       M68KInsn = 0x113
	M68K_INS_LEA       M68KInsn = 0x114
	M68K_INS_LINK      M68KInsn = 0x115
	M68K_INS_LPSTOP    M68KInsn = 0x116
	M68K_INS_LSL       M68KInsn = 0x117
	M68K_INS_LSR       M68KInsn = 0x118
	M68K_INS_MOVE      M68KInsn = 0x119
	M68K_INS_MOVEA     M68KInsn = 0x11a
	M68K_INS_MOVEC     M68KInsn = 0x11b
	M68K_INS_MOVEM     M68KInsn = 0x11c
	M68K_INS_MOVEP     M68KInsn = 0x11d
	M68K_INS_MOVEQ     M68KInsn = 0x11e
	M68K_INS_MOVES     M68KInsn = 0x11f
	M68K_INS_MOVE16    M68KInsn = 0x120
	M68K_INS_MULS      M68KInsn = 0x121
	M68K_INS_MULU      M68KInsn = 0x122
	M68K_INS_NBCD      M68KInsn = 0x123
	M68K_INS_NEG       M68KInsn = 0x124
	M68K_INS_NEGX      M68KInsn = 0x125
	M68K_INS_NOP       M68KInsn = 0x126
	M68K_INS_NOT       M68KInsn = 0x127
	M68K_INS_OR        M68KInsn = 0x128
	M68K_INS_ORI       M68KInsn = 0x129
	M68K_INS_PACK      M68KInsn = 0x12a
	M68K_INS_PEA       M68KInsn = 0x12b
	M68K_INS_PFLUSH    M68KInsn = 0x12c
	M68K_INS_PFLUSHA   M68KInsn = 0x12d
	M68K_INS_PFLUSHAN  M68KInsn = 0x12e
	M68K_INS_PFLUSHN   M68KInsn = 0x12f
	M68K_INS_PLOADR    M68KInsn = 0x130
	M68K_INS_PLOADW    M68KInsn = 0x131
	M68K_INS_PLPAR     M68KInsn = 0x132
	M68K_INS_PLPAW     M68KInsn = 0x133
	M68K_INS_PMOVE     M68KInsn = 0x134
	M68K_INS_PMOVEFD   M68KInsn = 0x135
	M68K_INS_PTESTR    M68KInsn = 0x136
	M68K_INS_PTESTW    M68KInsn = 0x137
	M68K_INS_PULSE     M68KInsn = 0x138
	M68K_INS_REMS      M68KInsn = 0x139
	M68K_INS_REMU      M68KInsn = 0x13a
	M68K_INS_RESET     M68KInsn = 0x13b
	M68K_INS_ROL       M68KInsn = 0x13c
	M68K_INS_ROR       M68KInsn = 0x13d
	M68K_INS_ROXL      M68KInsn = 0x13e
	M68K_INS_ROXR      M68KInsn = 0x13f
	M68K_INS_RTD       M68KInsn = 0x140
	M68K_INS_RTE       M68KInsn = 0x141
	M68K_INS_RTM       M68KInsn = 0x142
	M68K_INS_RTR       M68KInsn = 0x143
	M68K_INS_RTS       M68KInsn = 0x144
	M68K_INS_SBCD      M68KInsn = 0x145
	M68K_INS_ST        M68KInsn = 0x146
	M68K_INS_SF        M68KInsn = 0x147
	M68K_INS_SHI       M68KInsn = 0x148
	M68K_INS_SLS       M68KInsn = 0x149
	M68K_INS_SCC       M68KInsn = 0x14a
	M68K_INS_SHS       M68KInsn = 0x14b
	M68K_INS_SCS       M68KInsn = 0x14c
	M68K_INS_SLO       M68KInsn = 0x14d
	M68K_INS_SNE       M68KInsn = 0x14e
	M68K_INS_SEQ       M68KInsn = 0x14f
	M68K_INS_SVC       M68KInsn = 0x150
	M68K_INS_SVS       M68KInsn = 0x151
	M68K_INS_SPL       M68KInsn = 0x152
	M68K_INS_SMI       M68KInsn = 0x153
	M68K_INS_SGE       M68KInsn = 0x154
	M68K_INS_SLT       M68KInsn = 0x155
	M68K_INS_SGT       M68KInsn = 0x156
	M68K_INS_SLE       M68KInsn = 0x157
	M68K_INS_STOP      M68KInsn = 0x158
	M68K_INS_SUB       M68KInsn = 0x159
	M68K_INS_SUBA      M68KInsn = 0x15a
	M68K_INS_SUBI      M68KInsn = 0x15b
	M68K_INS_SUBQ      M68KInsn = 0x15c
	M68K_INS_SUBX      M68KInsn = 0x15d
	M68K_INS_SWAP      M68KInsn = 0x15e
	M68K_INS_TAS       M68KInsn = 0x15f
	M68K_INS_TRAP      M68KInsn = 0x160
	M68K_INS_TRAPV     M68KInsn = 0x161
	M68K_INS_TRAPT     M68KInsn = 0x162
	M68K_INS_TRAPF     M68KInsn = 0x163
	M68K_INS_TRAPHI    M68KInsn = 0x164
	M68K_INS_TRAPLS    M68KInsn = 0x165
	M68K_INS_TRAPCC    M68KInsn = 0x166
	M68K_INS_TRAPHS    M68KInsn = 0x167
	M68K_INS_TRAPCS    M68KInsn = 0x168
	M68K_INS_TRAPLO    M68KInsn = 0x169
	M68K_INS_TRAPNE    M68KInsn = 0x16a
	M68K_INS_TRAPEQ    M68KInsn = 0x16b
	M68K_INS_TRAPVC    M68KInsn = 0x16c
	M68K_INS_TRAPVS    M68KInsn = 0x16d
	M68K_INS_TRAPPL    M68KInsn = 0x16e
	M68K_INS_TRAPMI    M68KInsn = 0x16f
	M68K_INS_TRAPGE    M68KInsn = 0x170
	M68K_INS_TRAPLT    M68KInsn = 0x171
	M68K_INS_TRAPGT    M68KInsn = 0x172
	M68K_INS_TRAPLE    M68KInsn = 0x173
	M68K_INS_TST       M68KInsn = 0x174
	M68K_INS_UNLK      M68KInsn = 0x175
	M68K_INS_UNPK      M68KInsn = 0x176
	M68K_INS_ENDING    M68KInsn = 0x177
)

type M68KOpBrDispSize uint32

const (
	M68K_OP_BR_DISP_SIZE_INVALID M68KOpBrDispSize = 0x0
	M68K_OP_BR_DISP_SIZE_BYTE    M68KOpBrDispSize = 0x1
	M68K_OP_BR_DISP_SIZE_WORD    M68KOpBrDispSize = 0x2
	M68K_OP_BR_DISP_SIZE_LONG    M68KOpBrDispSize = 0x4
)

type M68KOpType uint32

const (
	M68K_OP_INVALID   M68KOpType = 0x0
	M68K_OP_REG       M68KOpType = 0x1
	M68K_OP_IMM       M68KOpType = 0x2
	M68K_OP_MEM       M68KOpType = 0x3
	M68K_OP_FP_SINGLE M68KOpType = 0x4
	M68K_OP_FP_DOUBLE M68KOpType = 0x5
	M68K_OP_REG_BITS  M68KOpType = 0x6
	M68K_OP_REG_PAIR  M68KOpType = 0x7
	M68K_OP_BR_DISP   M68KOpType = 0x8
)

type M68KReg uint32

const (
	M68K_REG_INVALID M68KReg = 0x0
	M68K_REG_D0      M68KReg = 0x1
	M68K_REG_D1      M68KReg = 0x2
	M68K_REG_D2      M68KReg = 0x3
	M68K_REG_D3      M68KReg = 0x4
	M68K_REG_D4      M68KReg = 0x5
	M68K_REG_D5      M68KReg = 0x6
	M68K_REG_D6      M68KReg = 0x7
	M68K_REG_D7      M68KReg = 0x8
	M68K_REG_A0      M68KReg = 0x9
	M68K_REG_A1      M68KReg = 0xa
	M68K_REG_A2      M68KReg = 0xb
	M68K_REG_A3      M68KReg = 0xc
	M68K_REG_A4      M68KReg = 0xd
	M68K_REG_A5      M68KReg = 0xe
	M68K_REG_A6      M68KReg = 0xf
	M68K_REG_A7      M68KReg = 0x10
	M68K_REG_FP0     M68KReg = 0x11
	M68K_REG_FP1     M68KReg = 0x12
	M68K_REG_FP2     M68KReg = 0x13
	M68K_REG_FP3     M68KReg = 0x14
	M68K_REG_FP4     M68KReg = 0x15
	M68K_REG_FP5     M68KReg = 0x16
	M68K_REG_FP6     M68KReg = 0x17
	M68K_REG_FP7     M68KReg = 0x18
	M68K_REG_PC      M68KReg = 0x19
	M68K_REG_SR      M68KReg = 0x1a
	M68K_REG_CCR     M68KReg = 0x1b
	M68K_REG_SFC     M68KReg = 0x1c
	M68K_REG_DFC     M68KReg = 0x1d
	M68K_REG_USP     M68KReg = 0x1e
	M68K_REG_VBR     M68KReg = 0x1f
	M68K_REG_CACR    M68KReg = 0x20
	M68K_REG_CAAR    M68KReg = 0x21
	M68K_REG_MSP     M68KReg = 0x22
	M68K_REG_ISP     M68KReg = 0x23
	M68K_REG_TC      M68KReg = 0x24
	M68K_REG_ITT0    M68KReg = 0x25
	M68K_REG_ITT1    M68KReg = 0x26
	M68K_REG_DTT0    M68KReg = 0x27
	M68K_REG_DTT1    M68KReg = 0x28
	M68K_REG_MMUSR   M68KReg = 0x29
	M68K_REG_URP     M68KReg = 0x2a
	M68K_REG_SRP     M68KReg = 0x2b
	M68K_REG_FPCR    M68KReg = 0x2c
	M68K_REG_FPSR    M68KReg = 0x2d
	M68K_REG_FPIAR   M68KReg = 0x2e
	M68K_REG_ENDING  M68KReg = 0x2f
)

type M68KSizeType uint32

const (
	M68K_SIZE_TYPE_INVALID M68KSizeType = 0x0
	M68K_SIZE_TYPE_CPU     M68KSizeType = 0x1
	M68K_SIZE_TYPE_FPU     M68KSizeType = 0x2
)

type CsM68K struct {
	Operands [4]CsM68KOp
	Size     M68KOpSize
	Count    uint8
	_        [7]byte
}

type CsM68KOp struct {
	Imm           uint64
	Mem           M68KOpMem
	Br_disp       M68KOpBrDisp
	Register_bits uint32
	Type          uint32
	Address_mode  uint32
}

type CsM68KOpRegPair struct {
	X_0 uint32
	X_1 uint32
}

type M68KOpBrDisp struct {
	Disp int32
	Size uint8
	_    [3]byte
}

type M68KOpMem struct {
	Base_reg    uint32
	Index_reg   uint32
	In_base_reg uint32
	In_disp     uint32
	Out_disp    uint32
	Disp        int16
	Scale       uint8
	Bitfield    uint8
	Width       uint8
	Offset      uint8
	Index_size  uint8
	_           [1]byte
}

type M68KOpSize struct {
	Type uint32
	Size uint32
}

const M68K_OPERAND_COUNT = 0x4
