package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64) {
	z = float64(1)
	for i := 0; i<10;i++ {
		z = z - (z*z-x)/2*z
	}
	return
}

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(math.Sqrt(4))
}
