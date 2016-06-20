package gohilbert

const maxLevel32 = 16

func IndexToXY32(n, idx uint32) (uint32,uint32, error) {
	if (n > maxLevel32) {
		return 0, 0, ErrLevelOutOfRange
	}

	if (idx > (1 << n) * (1 << n) - 1) {
		return 0, 0, ErrIndexOutOfRange
	}

	idx = idx << (32 - 2 * n)

	i0 := deinterleave32(idx)
	i1 := deinterleave32(idx >> 1)

	t0 := (i0 | i1) ^ 0xFFFF
	t1 := i0 & i1

	prefixT0 := prefixScan32(t0)
	prefixT1 := prefixScan32(t1)

	a := (((i0 ^ 0xFFFF) & prefixT1) | (i0 & prefixT0))

	return (a ^ i1) >> (16 - n), (a ^ i0 ^ i1) >> (16 - n), nil
}

func deinterleave32(x uint32) uint32{
	x = x & 0x55555555
	x = (x | (x >> 1)) & 0x33333333
	x = (x | (x >> 2)) & 0x0F0F0F0F
	x = (x | (x >> 4)) & 0x00FF00FF
	x = (x | (x >> 8)) & 0x0000FFFF
	return x
}

func prefixScan32(x uint32) uint32 {
	x = (x >> 8) ^ x
	x = (x >> 4) ^ x
	x = (x >> 2) ^ x
	x = (x >> 1) ^ x
	return x
}

func XYToIndex32(n, x, y uint32) (uint32, error) {
	if (n > maxLevel32) {
		return 0, ErrLevelOutOfRange
	}

	var m uint32 = (1 << n) - 1
	switch {
	case x > m:
		return 0, ErrXCoordinateOutOfRange
	case y > m:
		return 0, ErrYCoordinateOutOfRange
	}

	x = x << (16 - n)
	y = y << (16 - n)

	var A, B, C, D uint32

	// Initial prefix scan round, prime with x and y
	{
		a := x ^ y;
		b := 0xFFFF ^ a;
		c := 0xFFFF ^ (x | y)
		d := x & (y ^ 0xFFFF)

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

	{
		a := A
		b := B
		c := C
		d := D

		A = ((a & (a >> 4)) ^ (b & (b >> 4)))
		B = ((a & (b >> 4)) ^ (b & ((a ^ b) >> 4)))

		C ^= ((a & (c >> 4)) ^ (b & (d >> 4)))
		D ^= ((b & (c >> 4)) ^ ((a ^ b) & (d >> 4)))
	}

	// Final round and projection
	{
		a := A
		b := B
		c := C
		d := D

		C ^= ((a & (c >> 8)) ^ (b & (d >> 8)))
		D ^= ((b & (c >> 8)) ^ ((a ^ b) & (d >> 8)))
	}

	// Undo transformation prefix scan
	a := C ^ (C >> 1)
	b := D ^ (D >> 1)

	// Recover index bits
	i0 := x ^ y
	i1 := b | (0xFFFF ^ (i0 | a))

	return ((interleave32(i1) << 1) | interleave32(i0)) >> (32 - 2 * n), nil
}

func interleave32(x uint32) uint32 {
	x = (x | (x << 8)) & 0x00FF00FF
	x = (x | (x << 4)) & 0x0F0F0F0F
	x = (x | (x << 2)) & 0x33333333
	x = (x | (x << 1)) & 0x55555555
	return x
}
