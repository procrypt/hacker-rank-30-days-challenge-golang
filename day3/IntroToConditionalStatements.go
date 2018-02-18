package main

import "fmt"

var n int

func main() {
	fmt.Scan(&n)
	if n%2 != 0 {
		fmt.Println("Weird")
	}
	if n%2 == 0 {
		for i:=2; i<6; i++ {
			if n == i {
				fmt.Println("Not Weird")
			}
		}
	}
	if n%2 == 0 {
		for i:=6; i<21; i++ {
			if n == i {
				fmt.Println("Weird")
			}
		}
	}
	if n%2 == 0 && n > 20 {
		fmt.Println("Not Weird")
	}
}
