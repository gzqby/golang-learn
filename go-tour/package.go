package main

import (
	"fmt"
	// "math"
	// "math/rand"
)

func main() {
	fmt.Println(add(10,20));
}

func add(a,b int) (sum, num int) {
	sum = a+b
	num = a-b
	return
}