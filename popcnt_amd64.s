// func popcnt(x uint64) uint
TEXT ·popcnt(SB),$0-16
	MOVQ	x+0(FP), DX
	POPCNTQ	DX, DX
	MOVQ	DX, ret+8(FP)
	RET
