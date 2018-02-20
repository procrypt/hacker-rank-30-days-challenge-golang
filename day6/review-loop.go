package main

import "fmt"

var TestCase int

func main() {
	fmt.Scan(&TestCase)
	for i:=1; i<=TestCase ; i++ {
			var str string
			fmt.Scan(&str)
			var even, odd string
			for i, k := range str {
				if i%2 == 0 {
					even += string(k)
				}
				if i%2 != 0 {
					odd += string(k)
				}
			}
			fmt.Println(even + " " + odd)
	}
}