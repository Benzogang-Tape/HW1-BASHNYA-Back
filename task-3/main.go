package main

import "fmt"

const shift_value int = 1

func main() {
	var n uint
	fmt.Scan(&n)
	arr := arr_init(n)

	ans := cicle_slide(arr)
	ans()

	print_arr(arr)
}

func cicle_slide(arr []int) func() {
	return func() {
		for j := 0; j < shift_value; j++ {
			temp := arr[len(arr)-1]
			for i := 1; i < len(arr); i++ {
				arr[len(arr)-i] = arr[len(arr)-1-i]
			}
			arr[0] = temp
		}
	}
}

func arr_init(n uint) []int {
	arr := make([]int, n)
	var elem int
	for i := 0; i < int(n); i++ {
		fmt.Scan(&elem)
		arr[i] = elem
	}
	return arr
}

func print_arr(arr []int) {
	for _, elem := range arr {
		fmt.Printf("%d ", elem)
	}
}
