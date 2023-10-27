package main

import (
	"fmt"
	"interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is my programming language \n", user.Username)
	fmt.Printf("Fell free to type in the program / command")
	repl.Start(os.Stdin, os.Stdout)
}
