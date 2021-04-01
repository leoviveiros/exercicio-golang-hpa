package main

import (
	"crypto/sha512"
	"strings"
)

// O algorítmo do exemplo em PHP não estava estressando a CPU suficientemente, mesmo com um count maior
func CpuStresser() {
	const count = 23

	builder := strings.Builder{}
	sha := sha512.New()

	builder.WriteString("aBcDeFgHiJkLMnOpQrStUvWxYz")

	for i := 0; i < count; i++ {
		sum := sha.Sum([]byte(builder.String()))
		builder.Write(sum)
	}

	// PHP example
	/* x := 0.0001
	const count = 10000000

	for i := 0; i <= count; i++ {
		x += math.Sqrt(x)
	} */
}