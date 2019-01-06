// +build !go1.9

package runtime // import "blitznote.com/src/go.runtime"

//go:noescape
func popcnt(x uint64) uint
