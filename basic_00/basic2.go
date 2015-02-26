package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time ", time.Now())

	//testAfterWithChannel()
	//testSleep()
	//testParseDuration()
	testSince()
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

func testSince() {
	var date = time.Date(1992, 3, 24, 1, 1, 1, 1, time.Local)
	var dur = time.Since(date)
	fmt.Println("Duration since ", date, "= ", dur)
	fmt.Println("Hours since 24 march 1992 = ", dur.Hours())
	fmt.Println("date.Weekday = ", date.Weekday())

}
