package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(1242)
	fmt.Println("Random number ", rand.Intn(10))
}
