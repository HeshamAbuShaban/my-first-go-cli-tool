package main

import (
    "flag"
    "fmt"
    "strings"
)

// it prints "Hello <name>!"
// cmdGreet handles the "greet" subcommand with flags
func cmdGreet(args []string) {
    // Create a new FlagSet just for "greet"
    greetCmd := flag.NewFlagSet("greet", flag.ExitOnError)

    name := greetCmd.String("name", "World", "The name to greet")
    shout := greetCmd.Bool("shout", false, "Whether to shout the greeting")

	// Override the default usage message
    greetCmd.Usage = func() {
        fmt.Println("Usage: mycli greet [options]")
        fmt.Println("\nOptions:")
        greetCmd.PrintDefaults()
    }

    // Parse only the args passed after "greet"
    err := greetCmd.Parse(args)
    if err != nil {
        return
    }
	
	greeting := fmt.Sprintf("Hello %s!", *name)
    
	if *shout {
        greeting = strings.ToUpper(greeting)
    }
    fmt.Println(greeting)
}

