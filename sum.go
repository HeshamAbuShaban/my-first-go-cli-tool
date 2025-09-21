package main

import (
  "fmt"
  "strconv"
)


func cmdSum(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: mycli sum <a> <b>")
		return
	}
	a, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid number:", args[0])
		return
	}
	b, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Invalid number:", args[1])
		return
	}
	fmt.Println(a + b)
}
