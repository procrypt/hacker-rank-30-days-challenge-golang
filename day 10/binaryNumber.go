package main

import (
	"fmt"
	"strconv"
)

var num int64

func main() {
	var count, max int
	fmt.Scan(&num)
	binary := strconv.FormatInt(num, 2)

	for _, n := range binary {
		data := string(n)
		if data == "1" {
			count++
			if count > max {
				max = count
			}
		} else {
			count = 0
		}
	}
	fmt.Println(max)
}