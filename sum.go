package main

import (
    "flag"
    "fmt"
)

// cmdSum handles the "sum" subcommand with flags
func cmdSum(args []string) {
    sumCmd := flag.NewFlagSet("sum", flag.ExitOnError)

    a := sumCmd.Int("a", 0, "First integer to add")
    b := sumCmd.Int("b", 0, "Second integer to add")

    // Custom usage
    sumCmd.Usage = func() {
        fmt.Println("Usage: mycli sum [options]")
        fmt.Println("\nOptions:")
        sumCmd.PrintDefaults()
    }

    err := sumCmd.Parse(args)
    if err != nil {
        return
    }

    result := *a + *b
    fmt.Printf("%d + %d = %d\n", *a, *b, result)
}

