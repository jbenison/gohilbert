package gohilbert_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/jbenison/gohilbert"
)

var testCases64 = []struct {
	n, idx, x, y uint64
}{
	{1, 2, 1, 1},
	{2, 0, 0, 0},
	{3, 5, 3, 0},
	{4, 26, 7, 3},
	{5, 62, 7, 1},
	{6, 79, 0, 11},
	{7, 126, 9, 7},
	{8, 129, 8, 9},
	{9, 595, 16, 29},
	{10, 325, 24, 3},
	{11, 1272, 46, 2},
	{12, 3932, 51, 13},
	{13, 6640, 44, 112},
	{14, 25601, 64, 225},
	{15, 57774, 35, 177},
	{16, 98107, 399, 122},
	{17, 76350, 119, 369},
	{18, 177586, 433, 453},
	{19, 854828, 455, 725},
	{20, 909039, 687, 324},
	{21, 1730209, 861, 2044},
	{22, 3013737, 1286, 1775},
	{23, 1350627, 229, 1724},
	{24, 10081829, 2699, 3944},
	{25, 47684989, 6566, 5849},
	{26, 22667491, 7989, 700},
	{27, 121184025, 11842, 6967},
	{28, 246927173, 9608, 563},
	{29, 373684101, 4900, 30935},
	{30, 466466395, 25171, 14654},
	{31, 380526414, 7733, 31951},
	{32, 726993320, 28625, 28865},
}

func TestIndexToXY64(t *testing.T) {
	for _, tc := range testCases64 {
		x, y, _ := gohilbert.IndexToXY64(tc.n, tc.idx)
		if x != tc.x || y != tc.y {
			t.Errorf("IndexToXY64(%d, %d) failed, want (%d, %d), got (%d, %d)", tc.n, tc.idx, tc.x, tc.y, x, y)
		}
	}
}

func TestXYToIndex64(t *testing.T) {
	for _, tc := range testCases64 {
		idx, _ := gohilbert.XYToIndex64(tc.n, tc.x, tc.y)
		if idx != tc.idx {
			t.Errorf("XYToIndex64(%d, %d, %d) failed, want %d, got %d", tc.n, tc.x, tc.y, tc.idx, idx)
		}
	}
}

func maxIndex64(n uint64) uint64 {
	return (1 << n) * (1 << n) - 1
}

func randomIndex64(n uint64) uint64 {
	max := maxIndex64(n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return uint64(r.Int63n(int64(max)))
}

func TestRandomIndexToXYToIndex64(t *testing.T) {
	var n, idx1, idx2, x, y uint64
	for n = 1; n <= 31; n++ {
		idx1 = randomIndex64(n)
		x, y, _ = gohilbert.IndexToXY64(n, idx1)
		idx2, _ = gohilbert.XYToIndex64(n, x, y)
		if idx1 != idx2 {
			t.Errorf("RandomIndexToXYToIndex64 failed, want %d, got %d", idx1, idx2)
		}
	}
}

func maxXY64(n uint64) uint64 {
	return (1 << n) - 1
}

func randomXY64(n uint64) (uint64, uint64) {
	max := maxXY64(n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return uint64(r.Int63n(int64(max))), uint64(r.Int63n(int64(max)))
}

func TestRandomXYToIndexToXY64(t *testing.T) {
	var n, idx, x1, y1, x2, y2 uint64
	for n = 1; n <= 32; n++ {
		x1, y1 = randomXY64(n)
		idx, _ = gohilbert.XYToIndex64(n, x1, y1)
		x2, y2, _ = gohilbert.IndexToXY64(n, idx)
		if x1 != x2 || y1 != y2 {
			t.Errorf("RandomXYToIndexToXY64 failed at level %d, want (%d, %d), got (%d, %d)", n, x1, y1, x2, y2)
		}
	}
}

var x64_1, y64_1 uint64

func benchmarkIndexToXY64(i uint64, b *testing.B) {
	var x, y uint64
	for n := 0; n < b.N; n++ {
		x, y, _ = gohilbert.IndexToXY64(i, i*i/2)
	}
	x64_1 = x
	y64_1 = y
}

func BenchmarkIndexToXY64_1(b *testing.B) { benchmarkIndexToXY64(1, b) }
func BenchmarkIndexToXY64_2(b *testing.B) { benchmarkIndexToXY64(2, b) }
func BenchmarkIndexToXY64_4(b *testing.B) { benchmarkIndexToXY64(4, b) }
func BenchmarkIndexToXY64_8(b *testing.B) { benchmarkIndexToXY64(8, b) }
func BenchmarkIndexToXY64_16(b *testing.B) { benchmarkIndexToXY64(16, b) }
func BenchmarkIndexToXY64_32(b *testing.B) { benchmarkIndexToXY64(32, b) }

var idx64_2 uint64

func benchmarkXYToIndex64(i uint64, b *testing.B) {
	var idx uint64
	for n := 0; n < b.N; n++ {
		idx, _ = gohilbert.XYToIndex64(i, i, i)
	}
	idx64_2 = idx
}

func BenchmarkXYToIndex64_1(b *testing.B) { benchmarkXYToIndex64(1, b) }
func BenchmarkXYToIndex64_2(b *testing.B) { benchmarkXYToIndex64(2, b) }
func BenchmarkXYToIndex64_4(b *testing.B) { benchmarkXYToIndex64(4, b) }
func BenchmarkXYToIndex64_8(b *testing.B) { benchmarkXYToIndex64(8, b) }
func BenchmarkXYToIndex64_16(b *testing.B) { benchmarkXYToIndex64(16, b) }
func BenchmarkXYToIndex64_32(b *testing.B) { benchmarkXYToIndex64(32, b) }
