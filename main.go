package main

import (
	"fmt"
	"os"
	"strings"

	flag "github.com/ogier/pflag"
)

// Flags
var (
	user string
)

func main() {

	// Parse flags.
	flag.Parse()

	// If user does not supply flags, print usage.
	// We can clean this up later by putting this into its own
	//function
	if flag.NFlag() == 0 {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	users = strings.Split(user, ",")
	fmt.Printf("Searching user(s): %s\n", users)
}

func init() {
	flag.StringVarP(&user, "user", "u", "", "Search Users")
}
