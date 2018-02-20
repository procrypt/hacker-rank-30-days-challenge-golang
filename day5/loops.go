package main

import (
	"fmt"
	"strconv"
)

var num int

func main() {
	fmt.Scan(&num)
	a := strconv.Itoa(num)
	for i:=1; i<11; i++ {
		result := strconv.Itoa(num * i)
		b := strconv.Itoa(i)
		fmt.Println(a + " x " + b + " = " + result)
	}
}