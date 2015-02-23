package main

import (
	"fmt"
)

func main(){
	fmt.Println("Basic !")
	basic()
}

func basic(){
	var x int
	x = 5
	y := 4 
	sum, sous := addAndSous(x, y)
	fmt.Println("sum : ", sum ," sous : ", sous)
}

func addAndSous(x, y int)(sum, sous int){
	return x+y, x-y
}