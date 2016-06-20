package gohilbert

const maxLevel8 = 4

func IndexToXY8(n, idx uint8) (uint8, uint8, error) {
	if (n > maxLevel8) {
		return 0, 0, ErrLevelOutOfRange
	}

	if (idx > (1 << n) * (1 << n) - 1) {
		return 0, 0, ErrIndexOutOfRange
	}

	idx = idx << (4 - 2 * n)

	i0 := deinterleave8(idx)
	i1 := deinterleave8(idx >> 1)

	t0 := (i0 | i1) ^ 0xF
	t1 := i0 & i1

	prefixT0 := prefixScan8(t0)
	prefixT1 := prefixScan8(t1)

	a := (((i0 ^ 0xF) & prefixT1) | (i0 & prefixT0))

	return (a ^ i1) >> (4 - n), (a ^ i0 ^ i1) >> (4 - n), nil
}

func deinterleave8(x uint8) uint8 {
	x = x & 0x55
	x = (x | (x >> 1)) & 0x33
	x = (x | (x >> 2)) & 0x0F
	return x
}

func prefixScan8(x uint8) uint8 {
	x = (x >> 2) ^ x
	x = (x >> 1) ^ x
	return x
}

func XYToIndex8(n, x, y uint8) (uint8, error) {
	if (n > maxLevel8) {
		return 0, ErrLevelOutOfRange
	}

	var m uint8 = (1 << n) - 1
	switch {
	case x > m:
		return 0, ErrXCoordinateOutOfRange
	case y > m:
		return 0, ErrYCoordinateOutOfRange
	}

	x = x << (4 - n)
	y = y << (4 - n)

	var A, B, C, D uint8

	// Initial prefix scan round, prime with x and y
	{
		a := x ^ y;
		b := 0xF ^ a;
		c := 0xF ^ (x | y)
		d := x & (y ^ 0xF)

		A = a | (b >> 1)
		B = (a >> 1) ^ a

		C = ((c >> 1) ^ (b & (d >> 1))) ^ c
		D = ((a & (c >> 1)) ^ (d >> 1)) ^ d
	}

	// Final round and projection
	{
		a := A
		b := B
		c := C
		d := D

		C ^= ((a & (c >> 2)) ^ (b & (d >> 2)))
		D ^= ((b & (c >> 2)) ^ ((a ^ b) & (d >> 2)))
	}

	// Undo transformation prefix scan
	a := C ^ (C >> 1)
	b := D ^ (D >> 1)

	// Recover index bits
	i0 := x ^ y
	i1 := b | (0xF ^ (i0 | a))

	return ((interleave8(i1) << 1) | interleave8(i0)) >> (8 - 2 * n), nil
}

func interleave8(x uint8) uint8 {
	x = (x | (x << 2)) & 0x33
	x = (x | (x << 1)) & 0x55
	return x
}
