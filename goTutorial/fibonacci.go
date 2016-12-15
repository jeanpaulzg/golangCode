package main

import (
	"fmt"
)

func fibonacci(n int, c  []int) []int{
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c[i] = y
		x, y = y, x+y
	}
	return c
	
}

func main() {
	c := make([]int, 9)

	fibonacci(cap(c), c)
	fmt.Println(c)
	for i:=0; i<= cap(c)-1;i++{
		fmt.Println(c[i])
	}
}
