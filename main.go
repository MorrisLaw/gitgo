package main

import (
	"fmt"
	"os"
	"strings"

	// Aliasing library name with flag.
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
	if flag.NFlag() == 0 {
		printUsage()
	}

	/* If multiple users are passed separated by commas, store them in a "users" 
	   array. */
	users := strings.Split(user, ",")
	fmt.Printf("Searching user(s): %s\n", users)

	/* Essentially a for each loop where we iterate over each user and pass
	   through to the getUsers(user) function. */
	for _, u := range users {
		result := getUsers(u)
		fmt.Println(`Username: `, result.Login)
		fmt.Println(`Name:     `, result.Name)
		fmt.Println(`Email:    `, result.Email)
		fmt.Println(`Bio:      `, result.Bio)
		fmt.Println("")
	}
}

func init() {
	/* Pass in the user variable that was declared at package level (above).
	   The "&" character means we are passing the variable "by reference"
	   (as opposed to "by value"), meaning: we don't want to pass a copy of
	   the user variable. We want to pass the original variable. */ 
	flag.StringVarP(&user, "user", "u", "", "Search Users")
}

func printUsage() {
	
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}