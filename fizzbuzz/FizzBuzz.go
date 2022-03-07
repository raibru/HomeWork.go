package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Run Fizz-Buzz...")
	result := runFizzBuzz(1, 100)
	fmt.Printf("Iterate %d times\n", result)
	fmt.Println("...finish")
}

func runFizzBuzz(min int, max int) int {
	i := min
	for ; i <= max; i++ {
		out := ApplyFizzBuzz(i)
		fmt.Println(out)
	}
	return i - min
}

// ApplyFizzBuzz value return Fizz, Buzz, FizzBuzz or number as string
func ApplyFizzBuzz(num int) string {
	out := ""
	if num%3 == 0 {
		out += "Fizz"
	}
	if num%5 == 0 {
		out += "Buzz"
	}
	if out == "" {
		out = strconv.Itoa(num)
	}
	return out
}
