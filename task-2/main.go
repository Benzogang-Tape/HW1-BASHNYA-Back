package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	check_triangle(a, b, c)
}

func check_triangle(a, b, c float64) {
	if (a < b+c) && (b < a+c) && (c < a+b) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
