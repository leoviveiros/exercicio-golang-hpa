package main

import (
	"testing"
	"time"
)

func TestCpuStresser(t *testing.T) {
	start := time.Now()

	CpuStresser()

	end := time.Now()
	elapsed := end.Sub(start)

	println("Elapsed time:", elapsed)

	if (elapsed < 100) {
		t.Errorf("cpu not stressed, elapsed %d", elapsed)
	}
}