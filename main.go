package main

import (
	"fmt"
	"os"
	"os/user"
	"panda/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic("err")
	}
	fmt.Printf("Hello %s, welcome to panda repl!\n", user.Username)
	fmt.Println("Feel free to type any command")
	repl.Start(os.Stdin, os.Stdout)
}
