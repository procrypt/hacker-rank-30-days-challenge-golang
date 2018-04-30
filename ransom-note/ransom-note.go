package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func makeMap(s string) map[string]int {
	a := make(map[string]int)
	b := strings.Split(s, " ")
	for i := range b {
		_,ok := a[string(b[i])]
		if ok {
			a[string(b[i])]++
		} else {
			a[string(b[i])] = 1
		}
	}
	return a
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	count := 0
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	i := scanner.Text()
	a := strings.Split(i, " ")
	intNote, _  := strconv.Atoi(a[1])

	scanner.Scan()
	b := scanner.Text()

	scanner.Scan()
	c := scanner.Text()

	magazine := makeMap(b)
	note := makeMap(c)

	for k,v := range note {
		i,ok := magazine[k]
		if ok {
			if i >= v {
				count++
			}
		}
	}

	if count == intNote {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}