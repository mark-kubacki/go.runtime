package runtime // import "blitznote.com/src/go.runtime"

import (
	"runtime"
	"testing"
)

func TestNumCPU(t *testing.T) {
	e := runtime.NumCPU()
	g := NumCPU()
	if e != g {
		t.Errorf("NumCPU = %v, expected %v", g, e)
	}
}
