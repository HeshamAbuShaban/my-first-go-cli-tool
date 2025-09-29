package main

import (
    "flag"
    "fmt"
    "time"
)

// cmdTime handles the "time" subcommand with flags
func cmdTime(args []string) {
    timeCmd := flag.NewFlagSet("time", flag.ExitOnError)

    utc := timeCmd.Bool("utc", false, "Show time in UTC instead of local")
    format := timeCmd.String("format", "15:04:05", "Time format (Go layout)")

    // Custom usage
    timeCmd.Usage = func() {
        fmt.Println("Usage: mycli time [options]")
        fmt.Println("\nOptions:")
        timeCmd.PrintDefaults()
        fmt.Println("\nFormat uses Go's time layout, e.g.:")
        fmt.Println("  2001-09-11 08:44:05   -> full date/time")
        fmt.Println("  15:04                 -> hours:minutes")
    }

    err := timeCmd.Parse(args)
    if err != nil {
        return
    }

    now := time.Now()
    if *utc {
        now = now.UTC()
    }

    fmt.Println(now.Format(*format))
}

