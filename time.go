package main

import (
  "fmt"
  "time"
)

// this is souppose to show the current time!

func cmdTime(args []string) {
	fmt.Println(time.Now().Format(time.RFC3339))
}
