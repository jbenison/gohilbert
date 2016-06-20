package gohilbert_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/jbenison/gohilbert"
)

var testCases32 = []struct {
	n, idx, x, y uint32
}{
	{1, 2, 1, 1},
	{2, 4, 0, 2},
	{3, 14, 0, 2},
	{4, 15, 3, 0},
	{5, 27, 3, 6},
	{6, 45, 5, 6},
	{7, 203, 5, 12},
	{8, 353, 28, 5},
	{9, 297, 7, 22},
	{10, 1534, 14, 48},
	{11, 2969, 37, 48},
	{12, 1532, 15, 49},
	{13, 5258, 11, 107},
	{14, 15028, 2, 68},
	{15, 7629, 41, 82},
	{16, 49064, 241, 129},
}

func TestIndexToXY32(t *testing.T) {
	for _, tc := range testCases32 {
		x, y, _ := gohilbert.IndexToXY32(tc.n, tc.idx)
		if x != tc.x || y != tc.y {
			t.Errorf("IndexToXY32(%d, %d) failed, want (%d, %d), got (%d, %d)", tc.n, tc.idx, tc.x, tc.y, x, y)
		}
	}
}

func TestXYToIndex32(t *testing.T) {
	for _, tc := range testCases32 {
		idx, _ := gohilbert.XYToIndex32(tc.n, tc.x, tc.y)
		if idx != tc.idx {
			t.Errorf("XYToIndex32(%d, %d, %d) failed, want %d, got %d", tc.n, tc.x, tc.y, tc.idx, idx)
		}
	}
}

func maxIndex32(n uint32) uint32 {
	return (1 << n) * (1 << n) - 1
}

func randomIndex32(n uint32) uint32 {
	max := maxIndex32(n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return uint32(r.Int31n(int32(max)))
}

func TestRandomIndexToXYToIndex32(t *testing.T) {
	var n, idx1, idx2, x, y uint32
	for n = 1; n <= 15; n++ {
		idx1 = randomIndex32(n)
		x, y, _ = gohilbert.IndexToXY32(n, idx1)
		idx2, _ = gohilbert.XYToIndex32(n, x, y)
		if idx1 != idx2 {
			t.Errorf("RandomIndexToXYToIndex32 failed, want %d, got %d", idx1, idx2)
		}
	}
}

func maxXY32(n uint32) uint32 {
	return (1 << n) - 1
}

func randomXY32(n uint32) (uint32, uint32) {
	max := maxXY32(n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return uint32(r.Int31n(int32(max))), uint32(r.Int31n(int32(max)))
}

func TestRandomXYToIndexToXY32(t *testing.T) {
	var n, idx, x1, y1, x2, y2 uint32
	for n = 1; n <= 16; n++ {
		x1, y1 = randomXY32(n)
		idx, _ = gohilbert.XYToIndex32(n, x1, y1)
		x2, y2, _ = gohilbert.IndexToXY32(n, idx)
		if x1 != x2 || y1 != y2 {
			t.Errorf("RandomXYToIndexToXY32 failed at level %d, want (%d, %d), got (%d, %d)", n, x1, y1, x2, y2)
		}
	}
}

var x32_1, y32_1 uint32

func benchmarkIndexToXY32(i uint32, b *testing.B) {
	var x, y uint32
	for n := 0; n < b.N; n++ {
		x, y, _ = gohilbert.IndexToXY32(i, i * i / 2)
	}
	x32_1 = x
	y32_1 = y
}

func BenchmarkIndexToXY32_1(b *testing.B) {
	benchmarkIndexToXY32(1, b)
}
func BenchmarkIndexToXY32_2(b *testing.B) {
	benchmarkIndexToXY32(2, b)
}
func BenchmarkIndexToXY32_4(b *testing.B) {
	benchmarkIndexToXY32(4, b)
}
func BenchmarkIndexToXY32_8(b *testing.B) {
	benchmarkIndexToXY32(8, b)
}
func BenchmarkIndexToXY32_16(b *testing.B) {
	benchmarkIndexToXY32(16, b)
}

var idx32_2 uint32

func benchmarkXYToIndex32(i uint32, b *testing.B) {
	var idx uint32
	for n := 0; n < b.N; n++ {
		idx, _ = gohilbert.XYToIndex32(i, i, i)
	}
	idx32_2 = idx
}

func BenchmarkXYToIndex32_1(b *testing.B) {
	benchmarkXYToIndex32(1, b)
}
func BenchmarkXYToIndex32_2(b *testing.B) {
	benchmarkXYToIndex32(2, b)
}
func BenchmarkXYToIndex32_4(b *testing.B) {
	benchmarkXYToIndex32(4, b)
}
func BenchmarkXYToIndex32_8(b *testing.B) {
	benchmarkXYToIndex32(8, b)
}
func BenchmarkXYToIndex32_16(b *testing.B) {
	benchmarkXYToIndex32(16, b)
}
