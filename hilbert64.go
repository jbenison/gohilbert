package gohilbert

const maxLevel64 = 32

func IndexToXY64(n, idx uint64) (uint64, uint64, error) {
	if (n > maxLevel64) {
		return 0, 0, ErrLevelOutOfRange
	}

	if (idx > (1 << n) * (1 << n) - 1) {
		return 0, 0, ErrIndexOutOfRange
	}

	idx = idx << (64 - 2 * n)

	i0 := deinterleave64(idx)
	i1 := deinterleave64(idx >> 1)

	t0 := (i0 | i1) ^ 0xFFFFFFFF
	t1 := i0 & i1

	prefixT0 := prefixScan64(t0)
	prefixT1 := prefixScan64(t1)

	a := (((i0 ^ 0xFFFFFFFF) & prefixT1) | (i0 & prefixT0));

	return (a ^ i1) >> (32 - n), (a ^ i0 ^ i1) >> (32 - n), nil
}

func prefixScan64(x uint64) uint64 {
	x = (x >> 16) ^ x
	x = (x >> 8) ^ x
	x = (x >> 4) ^ x
	x = (x >> 2) ^ x
	x = (x >> 1) ^ x
	return x
}

func deinterleave64(x uint64) uint64{
	x = x & 0x5555555555555555
	x = (x | (x >> 1)) & 0x3333333333333333
	x = (x | (x >> 2)) & 0x0F0F0F0F0F0F0F0F
	x = (x | (x >> 4)) & 0x00FF00FF00FF00FF
	x = (x | (x >> 8)) & 0x0000FFFF0000FFFF
	x = (x | (x >> 16)) & 0x00000000FFFFFFFF
	return x
}

func XYToIndex64(n, x, y uint64) (uint64, error) {
	if (n > maxLevel64) {
		return 0, ErrLevelOutOfRange
	}

	var m uint64 = (1 << n) - 1
	switch {
	case x > m:
		return 0, ErrXCoordinateOutOfRange
	case y > m:
		return 0, ErrYCoordinateOutOfRange
	}

	x = x << (32 - n)
	y = y << (32 - n)

	var A, B, C, D uint64

	// Initial prefix scan round, prime with x and y
	{
		a := x ^ y
		b := 0xFFFFFFFF ^ a
		c := 0xFFFFFFFF ^ (x | y)
		d := x & (y ^ 0xFFFFFFFF)

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

		C ^= ((a & (c >> 2)) ^ (b & (d >> 2)))
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

	{
		a := A
		b := B
		c := C
		d := D

		A = ((a & (a >> 8)) ^ (b & (b >> 8)))
		B = ((a & (b >> 8)) ^ (b & ((a ^ b) >> 8)))

		C ^= ((a & (c >> 8)) ^ (b & (d >> 8)))
		D ^= ((b & (c >> 8)) ^ ((a ^ b) & (d >> 8)))
	}

	// Final round and projection
	{
		a := A
		b := B
		c := C
		d := D

		C ^= ((a & (c >> 16)) ^ (b & (d >> 16)))
		D ^= ((b & (c >> 16)) ^ ((a ^ b) & (d >> 16)))
	}

	// Undo transformation prefix scan
	a := C ^ (C >> 1)
	b := D ^ (D >> 1)

	// Recover index bits
	i0 := x ^ y;
	i1 := b | (0xFFFFFFFF ^ (i0 | a))

	return ((interleave64(i1) << 1) | interleave64(i0)) >> (64 - 2 * n), nil
}

func interleave64(x uint64) uint64 {
	x = (x | (x << 16)) & 0x0000FFFF0000FFFF
	x = (x | (x << 8)) & 0x00FF00FF00FF00FF
	x = (x | (x << 4)) & 0x0F0F0F0F0F0F0F0F
	x = (x | (x << 2)) & 0x3333333333333333
	x = (x | (x << 1)) & 0x5555555555555555
	return x
}
