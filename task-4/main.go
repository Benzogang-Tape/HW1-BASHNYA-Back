package main

import "fmt"

func main() {
	var n uint
	fmt.Scan(&n)
	matrix := matrix_init(n)
	print_matrix(matrix)

	if symmetry_check(matrix) {
		fmt.Println("\nyes")
	} else {
		(fmt.Println("\nno"))
	}
}

func print_matrix(matrix [][]float64) {
	fmt.Println()
	length := len(matrix)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			fmt.Printf("%.1f	", matrix[i][j])
		}
		fmt.Println()
	}
}

func matrix_init(n uint) [][]float64 {
	matrix := make([][]float64, n)
	var elem float64
	for i := 0; i < int(n); i++ {
		matrix[i] = make([]float64, n)
		for j := 0; j < int(n); j++ {
			fmt.Scan(&elem)
			matrix[i][j] = elem
		}
	}
	return matrix
}

func symmetry_check(matrix [][]float64) bool {
	length := len(matrix)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if matrix[i][j] != matrix[j][i] {
				return false
			}
		}
	}
	return true
}
