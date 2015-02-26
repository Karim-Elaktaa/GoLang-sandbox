package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time ", time.Now())

	testAfterWithChannel()
}

func testAfterWithChannel() {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("5 seconds")
	}
}
