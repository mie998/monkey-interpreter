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
	fmt.Printf("Hello! %s✋ This is the Monkye programming language repl system!!\n", user.Username)
	fmt.Printf("Command it!\n")
	repl.Start(os.Stdin, os.Stdout)

}
