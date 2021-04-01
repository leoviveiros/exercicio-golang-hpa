package main

import "math"

func CpuStresser() {
	x := 0.0001
	const count = 10000000

	for i := 0; i <= count; i++ {
		x += math.Sqrt(x)
	}
}