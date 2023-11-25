package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Printf("Answer is: %f", math.Sqrt(math.Pow(float64(a), 2)+math.Pow(float64(b), 2)))
}
