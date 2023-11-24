package main

import "fmt"

func main() {
	var n uint
	var temp int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; uint(i) < n; i++ {
		var elem int
		fmt.Scan(&elem)
		arr[i] = elem
	}

	temp = arr[n-1]
	for i := 1; i < int(n); i++ {
		arr[int(n)-i] = arr[int(n)-1-i]
	}
	arr[0] = temp

	for _, elem := range arr {
		fmt.Printf("%d ", elem)
	}
}
