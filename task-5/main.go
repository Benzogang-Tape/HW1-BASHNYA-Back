package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	str := bufio.NewScanner(os.Stdin)
	str.Scan()
	res_str := str.Text()
	res_str = strings.Replace(res_str, "1", "one", -1)
	fmt.Println(res_str)
}

//ans := strings.Split(res_str, "")
/*
for i, value := range res_str {
		if value == 49 {
			res_str = strings.Replace(res_str, "1", "one", 1)
		}
		fmt.Println(i, value)
	}*/
