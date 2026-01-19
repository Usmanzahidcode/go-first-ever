package bench

import "testing"

func sum(n int) int64 {
	var s int64
	for i := 0; i < n; i++ {
		s += int64(i)
	}
	return s
}

// COMMAND: go test ./bench -bench=.
// Another flag is: -cpu=1,2,4 or just 4
func BenchmarkSum1M_Parallel(b *testing.B) {
	const N = 1_000_000

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = sum(N)
		}
	})
}
