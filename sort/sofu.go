package main

import (
	"fmt"

	"github.com/raibru/HomeWork/sort/sort"
)

func main() {
	fmt.Println("Run sort algorithm...")
	a := []int{1}
	bs := sort.BubbleSort{}
	err := bs.Sort(a)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

// EOF
