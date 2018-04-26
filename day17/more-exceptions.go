package main

import "fmt"

type calculator struct {
	n int
	p int
}

func(c *calculator) power() int {
	if c.n > 0 && c.p > 0 {
		return c.n*c.p
	} else {
		fmt.Println("n and p should be non-negative")
		return 0
	}
}

func New() *calculator {
	return &calculator{}
}

func main() {
	a := New()
	a.n = 3
	a.p = -3
	fmt.Println(a.power())
}
