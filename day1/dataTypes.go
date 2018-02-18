package main


import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	// Declare second integer, double, and String variables.
	var (
		j uint64
		e float64
		t string
	)

	// Read and save an integer, double, and String to your variables.
	fmt.Scan(&j)
	fmt.Scan(&e)
	scanner.Scan()
	t = scanner.Text()
	fmt.Println(i+j)
	fmt.Printf("%.1f\n",d+e)
	fmt.Println(s+t)
}

