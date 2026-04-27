package main

import (
	"fmt"
	"os"
	"os/user"

	"maelho.github.io/monkey/repl"
)

func main() {
	_, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Println("This is the Monkey programming language!")
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
