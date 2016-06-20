package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jbenison/gohilbert"
)

// First order      Second order
// 1 | 2            5 |  6 |  9 | 10
// 0 | 3            4 |  7 |  8 | 11
//                  3 |  2 | 13 | 12
//                  0 |  1 | 14 | 15
//
// Third order
// 21 | 22 | 25 | 26 | 37 | 38 | 41 | 42
// 20 | 23 | 24 | 27 | 36 | 39 | 40 | 43
// 19 | 18 | 29 | 28 | 35 | 34 | 45 | 44
// 16 | 17 | 30 | 31 | 32 | 33 | 46 | 47
// 15 | 12 | 11 | 10 | 53 | 52 | 51 | 48
// 14 | 13 |  8 |  9 | 54 | 55 | 50 | 49
//  1 |  2 |  7 |  6 | 57 | 56 | 61 | 62
//  0 |  3 |  4 |  5 | 58 | 59 | 60 | 63
//
func main() {
	/*fmt.Println("Random 32 bit index to x, y\n{level, index, x, y}")
	{
		var n, idx, x, y uint32
		for n = 1; n <= 15; n++ {
			idx = randomIndex32(n)
			x, y, _ = gohilbert.IndexToXY32(n, idx)
			fmt.Printf("{%d, %d, %d, %d},\n", n, idx, x, y)
		}
	}

	fmt.Println("\nRandom 64 bit index to x, y\n{level, index, x, y}")
	{
		var n, idx, x, y uint64
		for n = 1; n <= 31; n++ {
			idx = randomIndex64(n)
			x, y, _ = gohilbert.IndexToXY64(n, idx)
			fmt.Printf("{%d, %d, %d, %d},\n", n, idx, x, y)
		}
	}*/
	/*{
		var n, xy, idx uint32
		for n = 1; n <= 15; n++ {
			fmt.Printf("{%d, %d, %d},\n", n, maxXY32(n), maxIndex32(n))
			xy, _ = randomXY32(n)
			idx = randomIndex32(n)
			fmt.Printf("{%d, %d, %d},\n", n, xy, idx)
		}
	}
	{
		var n, xy, idx uint64
		for n = 1; n <= 31; n++ {
			fmt.Printf("{%d, %d, %d},\n", n, maxXY64(n), maxIndex64(n))
			xy, _ = randomXY64(n)
			idx = randomIndex64(n)
			fmt.Printf("{%d, %d, %d},\n", n, xy, idx)
		}
	}*/


	// Cycles through all
	{
		var n, idx, x, y, max_xy uint32
		for n = 1; n <= 3; n++ {
			max_xy = maxXY32(n)
			for y = 0; y <= max_xy; y++ {
				for x = 0; x <= max_xy; x++ {
					idx, _ = gohilbert.XYToIndex32(n,x,y)
					fmt.Printf("{%d, %d, %d, %d},\n", n, x, y, idx)
				}
			}
		}
	}
	{
		var n, idx, x, y, max_xy uint16
		for n = 1; n <= 3; n++ {
			max_xy = maxXY16(n)
			for y = 0; y <= max_xy; y++ {
				for x = 0; x <= max_xy; x++ {
					idx, _ = gohilbert.XYToIndex16(n,x,y)
					fmt.Printf("{%d, %d, %d, %d},\n", n, x, y, idx)
				}
			}
		}
	}
	{
		var n, idx, x, y, max_xy uint8
		for n = 1; n <= 3; n++ {
			max_xy = maxXY8(n)
			for y = 0; y <= max_xy; y++ {
				for x = 0; x <= max_xy; x++ {
					idx, _ = gohilbert.XYToIndex8(n,x,y)
					fmt.Printf("{%d, %d, %d, %d},\n", n, x, y, idx)
				}
			}
		}
	}
}

func maxIndex64(n uint64) uint64 {
	return (1 << n) * (1 << n) - 1
}

func maxIndex32(n uint32) uint32 {
	return (1 << n) * (1 << n) - 1
}

func maxIndex16(n uint16) uint16 {
	return (1 << n) * (1 << n) - 1
}

func randomIndex64(n uint64) uint64 {
	max := maxIndex64(n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return uint64(r.Int63n(int64(max)))
}

func randomIndex32(n uint32) uint32 {
	max := maxIndex32(n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return uint32(r.Int31n(int32(max)))
}

func maxXY64(n uint64) uint64 {
	return (1 << n) - 1
}

func maxXY32(n uint32) uint32 {
	return (1 << n) - 1
}

func maxXY16(n uint16) uint16 {
	return (1 << n) - 1
}

func maxXY8(n uint8) uint8 {
	return (1 << n) - 1
}

func randomXY64(n uint64) (uint64, uint64) {
	max := maxXY64(n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return uint64(r.Int63n(int64(max))), uint64(r.Int63n(int64(max)))
}

func randomXY32(n uint32) (uint32, uint32) {
	max := maxXY32(n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return uint32(r.Int31n(int32(max))), uint32(r.Int31n(int32(max)))
}
