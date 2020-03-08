// +build solaris darwin dragonfly freebsd illumos netbsd openbsd

package runtime

import (
	_ "unsafe" // required to use //go:linkname
)

//go:noescape
//go:linkname numcpu runtime.getncpu
func numcpu() int32
