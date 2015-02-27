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
}
