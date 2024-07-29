package runner

import (
	"sync"
	"testing"
)

func RunTests(t *testing.T, tests []func(t *testing.T)) {
	for _, test := range tests {
		test(t)
	}
}

func RunTestsParallel(t *testing.T, tests []func(t *testing.T)) {
	var wg sync.WaitGroup
	for _, test := range tests {
		wg.Add(1)
		go func(test func(t *testing.T)) {
			defer wg.Done()
			test(t)
		}(test)
	}
	wg.Wait()
}

func RunBenchmarks(b *testing.B, benchmarks []func(b *testing.B)) {
	for _, benchmark := range benchmarks {
		benchmark(b)
	}
}
