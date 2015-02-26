package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time ", time.Now())

	//testAfterWithChannel()
	testSleep()

}

func testAfterWithChannel() {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("5 seconds")
	}
}

func testSleep() {
	time.Sleep(2 * time.Second)
	fmt.Println("2 seconds")

}
