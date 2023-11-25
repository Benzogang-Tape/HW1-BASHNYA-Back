package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type dictionary map[int]struct{}

func (dict dictionary) add(elem int) {
	dict[elem] = struct{}{}
}

func main() {
	set := dictionary{}
	str := bufio.NewScanner(os.Stdin)
	str.Scan()
	res_str := str.Text()
	slc := strings.Split(res_str, " ")

	for _, value := range slc {
		ans, _ := strconv.Atoi(value)
		set.add(ans)
	}
	fmt.Println(len(set))
}
