package main

import "fmt"

// it prints "Hello <name>!"

func cmdGreet(args []string) {
	name := "World"
	if len(args) > 0 {
		name = args[0]
	}
	fmt.Printf("Hello %s!\n", name)
}
