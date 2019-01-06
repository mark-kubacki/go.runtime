// +build !amd64 !go1.9

package runtime // import "blitznote.com/src/go.runtime"

func popcnt(x uint64) uint {
	x -= (x >> 1) & 0x5555555555555555
	x = (x>>2)&0x3333333333333333 + x&0x3333333333333333
	x += x >> 4
	x &= 0x0f0f0f0f0f0f0f0f
	x *= 0x0101010101010101
	return uint(byte(x >> 56))
}
