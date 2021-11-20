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
