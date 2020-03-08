// +build linux windows plan9

package runtime

import (
	_ "unsafe" // required to use //go:linkname
)

//go:noescape
//go:linkname numcpu runtime.getproccount
func numcpu() int32
