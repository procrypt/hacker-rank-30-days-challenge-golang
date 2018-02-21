package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

var mapSize int

func main() {
	fmt.Scan(&mapSize)
	phoneBook := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	for i := 1; i <= mapSize; i++ {
		for scanner.Scan() {
			data := scanner.Text()
			newData := strings.Split(data, " ")
			name := newData[0]
			phone, _ := strconv.Atoi(newData[1])
			phoneBook[name] = phone
			break
		}
	}
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		name :=  scanner.Text()
		_, prs := phoneBook[name]
		if prs {
			fmt.Printf("%v=%v\n", name, phoneBook[name])
		} else {
			fmt.Println("Not found")
		}
	}
}