package main

import "fmt"

func main() {
	var n uint
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; uint(i) < n; i++ {
		var elem int
		fmt.Scan(&elem)
		arr[i] = elem
	}

	ans := cicle_slide(arr)
	ans()

	for _, elem := range arr {
		fmt.Printf("%d ", elem)
	}
}

func cicle_slide(arr []int) func() {
	return func() {
		temp := arr[len(arr)-1]
		for i := 1; i < len(arr); i++ {
			arr[len(arr)-i] = arr[len(arr)-1-i]
		}
		arr[0] = temp
	}
}
