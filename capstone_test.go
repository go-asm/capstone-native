package capstone_test

import (
	"testing"

	"github.com/go-asm/capstone"
)

func TestNew(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		arch       capstone.CsArch
		mode       capstone.CsMode
		wantArch   capstone.CsArch
		wantMode   capstone.CsMode
		wantEngine capstone.Engine
	}{
		"ARCH_X86_MODE_64": {
			arch:     capstone.CS_ARCH_X86,
			mode:     capstone.CS_MODE_64,
			wantArch: capstone.CS_ARCH_X86,
			wantMode: capstone.CS_MODE_64,
		},
		"ARCH_ARM_MODE_32": {
			arch:     capstone.CS_ARCH_ARM,
			mode:     capstone.CS_MODE_32,
			wantArch: capstone.CS_ARCH_ARM,
			wantMode: capstone.CS_MODE_32,
		},
		"ARCH_ARM64_MODE_64": {
			arch:     capstone.CS_ARCH_ARM64,
			mode:     capstone.CS_MODE_64,
			wantArch: capstone.CS_ARCH_ARM64,
			wantMode: capstone.CS_MODE_64,
		},
		"ARCH_BPF_MODE_64": {
			arch:     capstone.CS_ARCH_BPF,
			mode:     capstone.CS_MODE_64,
			wantArch: capstone.CS_ARCH_BPF,
			wantMode: capstone.CS_MODE_64,
		},
		"ARCH_EVM_MODE_64": {
			arch:     capstone.CS_ARCH_EVM,
			mode:     capstone.CS_MODE_64,
			wantArch: capstone.CS_ARCH_EVM,
			wantMode: capstone.CS_MODE_64,
		},
		"ARCH_M680X_MODE_32": {
			arch:     capstone.CS_ARCH_M680X,
			mode:     capstone.CS_MODE_32,
			wantArch: capstone.CS_ARCH_M680X,
			wantMode: capstone.CS_MODE_32,
		},
		"ARCH_M68K_MODE_32": {
			arch:     capstone.CS_ARCH_M68K,
			mode:     capstone.CS_MODE_32,
			wantArch: capstone.CS_ARCH_M68K,
			wantMode: capstone.CS_MODE_32,
		},
		"ARCH_MIPS_MODE_32": {
			arch:     capstone.CS_ARCH_MIPS,
			mode:     capstone.CS_MODE_32,
			wantArch: capstone.CS_ARCH_MIPS,
			wantMode: capstone.CS_MODE_32,
		},
		"ARCH_MOS65XX_MODE_32": {
			arch:     capstone.CS_ARCH_MOS65XX,
			mode:     capstone.CS_MODE_32,
			wantArch: capstone.CS_ARCH_MOS65XX,
			wantMode: capstone.CS_MODE_32,
		},
		"ARCH_PPC_MODE_32": {
			arch:     capstone.CS_ARCH_PPC,
			mode:     capstone.CS_MODE_32,
			wantArch: capstone.CS_ARCH_PPC,
			wantMode: capstone.CS_MODE_32,
		},
		"ARCH_RISCV_MODE_64": {
			arch:     capstone.CS_ARCH_RISCV,
			mode:     capstone.CS_MODE_64,
			wantArch: capstone.CS_ARCH_RISCV,
			wantMode: capstone.CS_MODE_64,
		},
		"ARCH_SPARC_MODE_32": {
			arch:     capstone.CS_ARCH_SPARC,
			mode:     capstone.CS_MODE_32,
			wantArch: capstone.CS_ARCH_SPARC,
			wantMode: capstone.CS_MODE_32,
		},
		"ARCH_SYSZ_MODE_32": {
			arch:     capstone.CS_ARCH_SYSZ,
			mode:     capstone.CS_MODE_32,
			wantArch: capstone.CS_ARCH_SYSZ,
			wantMode: capstone.CS_MODE_32,
		},
		"ARCH_TMS320C64X_MODE_32": {
			arch:     capstone.CS_ARCH_TMS320C64X,
			mode:     capstone.CS_MODE_32,
			wantArch: capstone.CS_ARCH_TMS320C64X,
			wantMode: capstone.CS_MODE_32,
		},
		"ARCH_XCORE_MODE_32": {
			arch:     capstone.CS_ARCH_XCORE,
			mode:     capstone.CS_MODE_32,
			wantArch: capstone.CS_ARCH_XCORE,
			wantMode: capstone.CS_MODE_32,
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			engine, err := capstone.New(tt.arch, tt.mode)
			if err != nil {
				t.Fatalf("capstone.New: %v", err)
			}
			defer engine.Close()

			if got, want := engine.Arch(), tt.wantArch; got != want {
				t.Fatalf("got %v arch but want %v", got, want)
			}
			if got, want := engine.Mode(), tt.wantMode; got != want {
				t.Fatalf("got %v mode but want %v", got, want)
			}
			if got := engine.SkipData(); got != nil {
				t.Fatalf("except got is nil but %v", got)
			}
		})
	}
}

func TestVersion(t *testing.T) {
	const wantMajor = int32(4)
	const wantMinor = int32(0)

	gotMajor, gotMinor := capstone.Version()
	if gotMajor != wantMajor {
		t.Fatalf("got %d but want %d", gotMajor, wantMajor)
	}
	if gotMinor != wantMinor {
		t.Fatalf("got %d but want %d", gotMinor, wantMinor)
	}
}

func TestError(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		err  capstone.CsErr
		want string
	}{
		"CS_ERR_OK": {
			err:  capstone.CS_ERR_OK,
			want: capstone.ErrOK.Error(),
		},
		"CS_ERR_MEM": {
			err:  capstone.CS_ERR_MEM,
			want: capstone.ErrMem.Error(),
		},
		"CS_ERR_ARCH": {
			err:  capstone.CS_ERR_ARCH,
			want: capstone.ErrArch.Error(),
		},
		"CS_ERR_HANDLE": {
			err:  capstone.CS_ERR_HANDLE,
			want: capstone.ErrHandle.Error(),
		},
		"CS_ERR_MODE": {
			err:  capstone.CS_ERR_MODE,
			want: capstone.ErrMode.Error(),
		},
		"CS_ERR_OPTION": {
			err:  capstone.CS_ERR_OPTION,
			want: capstone.ErrOption.Error(),
		},
		"CS_ERR_DETAIL": {
			err:  capstone.CS_ERR_DETAIL,
			want: capstone.ErrDetail.Error(),
		},
		"CS_ERR_MEMSETUP": {
			err:  capstone.CS_ERR_MEMSETUP,
			want: capstone.ErrMemSetup.Error(),
		},
		"CS_ERR_VERSION": {
			err:  capstone.CS_ERR_VERSION,
			want: capstone.ErrVersion.Error(),
		},
		"CS_ERR_DIET": {
			err:  capstone.CS_ERR_DIET,
			want: capstone.ErrDiet.Error(),
		},
		"CS_ERR_SKIPDATA": {
			err:  capstone.CS_ERR_SKIPDATA,
			want: capstone.ErrSkipData.Error(),
		},
		"CS_ERR_X86_ATT": {
			err:  capstone.CS_ERR_X86_ATT,
			want: capstone.ErrX86ATT.Error(),
		},
		"CS_ERR_X86_INTEL": {
			err:  capstone.CS_ERR_X86_INTEL,
			want: capstone.ErrX86Intel.Error(),
		},
		"CS_ERR_X86_MASM": {
			err:  capstone.CS_ERR_X86_MASM,
			want: capstone.ErrX86Masm.Error(),
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got, want := tt.err.Error(), tt.want; got != want {
				t.Fatalf("got %v but want %v", got, want)
			}
		})
	}
}
