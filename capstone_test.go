package capstone_test

import (
	"testing"

	"github.com/go-asm/capstone"
)

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
