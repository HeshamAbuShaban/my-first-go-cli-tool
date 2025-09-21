package main

import (
  "fmt"
  "os"
)

func main() {
  if len(os.Args) < 2 {
	fmt.Println("Usage: mycli <command> [args...]")
	return
  }

  cmd := os.Args[1]
  args := os.Args[2:]

  switch cmd {
  case "greet":
	  cmdGreet(args)
  case "sum":
	  cmdSum(args)
  case "time":
	  cmdTime(args)
  default:
	  fmt.Println("Unknown command:", cmd)
  }
}
