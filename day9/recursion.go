package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Println(factorial(num))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := n * factorial(n-1)
	return result
}
