package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time ", time.Now())

	//testAfterWithChannel()
	//testSleep()
	testParseDuration()
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

func testParseDuration() {
	var d, e = time.ParseDuration("300ms")
	fmt.Println("Parse = ", d.Nanoseconds())
	fmt.Println("Error = ", e)

}
