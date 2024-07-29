package benchmarks

import (
	"testing"
	"time"
)

func BenchmarkFunction(b *testing.B, fn func()) {
	for i := 0; i < b.N; i++ {
		start := time.Now()
		fn()
		duration := time.Since(start)
		b.ReportMetric(float64(duration.Nanoseconds()), "ns/op")
	}
}

func BenchmarkMemoryAlloc(b *testing.B, fn func()) {
	for i := 0; i < b.N; i++ {
		b.ReportAllocs()
		fn()
	}
}
