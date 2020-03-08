package runtime

import (
	"runtime"
)

func numcpu() int {
	// As of writing this relies on sysconf(__SC_NPROCESSORS_ONLN=0.x48).
	return runtime.NumCPU()
}
