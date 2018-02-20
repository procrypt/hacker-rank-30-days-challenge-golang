package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

var arraySize int

func main() {

	fmt.Scan(&arraySize)
	array := make([]int, arraySize)
	arrayElements :=  bufio.NewScanner(os.Stdin)
	arrayElements.Scan()
	data := arrayElements.Text()
	newData := strings.Split(data," ")

	size := len(array)
	for _, k:= range newData {
		array[size-1],_ = strconv.Atoi(k)
		size--
	}
	for _, k := range array {
		fmt.Printf("%d ", k)
	}
}
