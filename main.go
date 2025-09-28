package main

import (
  "fmt"
  "os"
)

func printGlobalHelp() {
    fmt.Println("Usage: mycli <command> [options]")
    fmt.Println("\nAvailable commands:")
    fmt.Println("  greet   Print a greeting with optional flags")
    fmt.Println("  sum     Add numbers together")
    fmt.Println("  time    Show the current time")
    fmt.Println("\nUse \"mycli <command> --help\" for more information about a command.")
}

func main() {
	if len(os.Args) < 2 {
        printGlobalHelp()
        return
	}

    // Global help
    if os.Args[1] == "--help" || os.Args[1] == "-h" {
        printGlobalHelp()
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
        printGlobalHelp()
    }
}
