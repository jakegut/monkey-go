package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	if len(os.Args) == 2 {
		file := os.Args[1]
		f, err := os.OpenFile(file, os.O_RDONLY, 0700)
		if err != nil {
			panic(err)
		}
		repl.ParseFile(f)
	} else {
		fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	}
}
