package gohilbert

const maxLevel16 = 8

func IndexToXY16(n, idx uint16) (uint16, uint16, error) {
	if (n > maxLevel16) {
		return 0, 0, ErrLevelOutOfRange
	}

	if (idx > (1 << n) * (1 << n) - 1) {
		return 0, 0, ErrIndexOutOfRange
	}

	idx = idx << (16 - 2 * n)

	i0 := deinterleave16(idx)
	i1 := deinterleave16(idx >> 1)

	t0 := (i0 | i1) ^ 0xFF
	t1 := i0 & i1

	prefixT0 := prefixScan16(t0)
	prefixT1 := prefixScan16(t1)

	a := (((i0 ^ 0xFF) & prefixT1) | (i0 & prefixT0))

	return (a ^ i1) >> (8 - n), (a ^ i0 ^ i1) >> (8 - n), nil
}

func deinterleave16(x uint16) uint16 {
	x = x & 0x5555
	x = (x | (x >> 1)) & 0x3333
	x = (x | (x >> 2)) & 0x0F0F
	x = (x | (x >> 4)) & 0x00FF
	return x
}

func prefixScan16(x uint16) uint16 {
	x = (x >> 4) ^ x
	x = (x >> 2) ^ x
	x = (x >> 1) ^ x
	return x
}

func XYToIndex16(n, x, y uint16) (uint16, error) {
	if (n > maxLevel16) {
		return 0, ErrLevelOutOfRange
	}

	var m uint16 = (1 << n) - 1
	switch {
	case x > m:
		return 0, ErrXCoordinateOutOfRange
	case y > m:
		return 0, ErrYCoordinateOutOfRange
	}

	x = x << (8 - n)
	y = y << (8 - n)

	var A, B, C, D uint16

	// Initial prefix scan round, prime with x and y
	{
		a := x ^ y;
		b := 0xFF ^ a;
		c := 0xFF ^ (x | y)
		d := x & (y ^ 0xFF)

		A = a | (b >> 1)
		B = (a >> 1) ^ a

		C = ((c >> 1) ^ (b & (d >> 1))) ^ c
		D = ((a & (c >> 1)) ^ (d >> 1)) ^ d
	}

	{
		a := A
		b := B
		c := C
		d := D

		A = ((a & (a >> 2)) ^ (b & (b >> 2)))
		B = ((a & (b >> 2)) ^ (b & ((a ^ b) >> 2)))

		C ^= ((a & (c >> 2)) ^ (b & (d >> 2)));
		D ^= ((b & (c >> 2)) ^ ((a ^ b) & (d >> 2)))
	}

	// Final round and projection
	{
		a := A
		b := B
		c := C
		d := D

		C ^= ((a & (c >> 4)) ^ (b & (d >> 4)))
		D ^= ((b & (c >> 4)) ^ ((a ^ b) & (d >> 4)))
	}

	// Undo transformation prefix scan
	a := C ^ (C >> 1)
	b := D ^ (D >> 1)

	// Recover index bits
	i0 := x ^ y
	i1 := b | (0xFF ^ (i0 | a))

	return ((interleave16(i1) << 1) | interleave16(i0)) >> (16 - 2 * n), nil
}

func interleave16(x uint16) uint16 {
	x = (x | (x << 4)) & 0x0F0F
	x = (x | (x << 2)) & 0x3333
	x = (x | (x << 1)) & 0x5555
	return x
}
