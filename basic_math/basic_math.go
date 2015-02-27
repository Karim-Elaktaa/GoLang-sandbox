package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	rand.Seed(1242)
	fmt.Println("Random number ", rand.Intn(10))

	fmt.Println("Next after 2 ", math.Nextafter(2, 3))

	fmt.Println("PI ", math.Pi)
	fmt.Println("NextAfter PI ", math.Nextafter(math.Pi, 4))

	fmt.Println("Abs PI", math.Abs(-math.Pi))
	fmt.Println("Abs NaN ", math.Abs(math.NaN()))
	fmt.Println("Abs -Inf ", math.Abs(math.Inf(-1)))
	fmt.Println("Abs +Inf ", math.Abs(math.Inf(1)))

}
