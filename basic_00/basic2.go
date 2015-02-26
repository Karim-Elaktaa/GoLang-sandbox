package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time ", time.Now())
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("5 seconds")
	}
}
