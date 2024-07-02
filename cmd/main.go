package main

import (
	"fmt"
	"github.com/symelbak/kamelang/internal/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello Kamel: %s\n Welcome to Kamelang!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)

}
