// +build go1.9

package runtime // import "blitznote.com/src/go.runtime"

import (
	"math/bits"
)

func popcnt(x uint64) uint {
	return uint(bits.OnesCount64(x))
}
