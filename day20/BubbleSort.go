package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func BubbleSort(item []int) {
	NumberOfSwaps := 0
	for i:=0; i<len(item); i++ {
		for j:=0; j<len(item)-1; j++ {
			if item[j] > item[j+1] {
				item[j], item[j+1] = item[j+1], item[j]
				NumberOfSwaps++
			}
		}
	}
	fmt.Printf("Array is sorted in %d swaps.\n", NumberOfSwaps)
	fmt.Printf("First Element: %d\n", item[0])
	fmt.Printf("Last Element: %d\n", item[len(item)-1])
}
func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var a int
	var data []string
	var item []int
	fmt.Scan(&a)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		data = strings.Split(text, " ")
		for i := range data {
			n,_ := strconv.Atoi(data[i])
			item = append(item, n)
		}
	}
	BubbleSort(item)
}