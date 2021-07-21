package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	switch len(args) {
	case 0:
		startPrompt()
	case 1:
		runFile(args[0])
	default:
		panic(errors.New("usage: bf <path-to-file>"))
	}
}

func startPrompt() {
	instance := New([]byte{})
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("=> ")
		text, _ := reader.ReadString('\n')
		instance.RunREPL([]byte(text))
	}
}

func runFile(path string) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	instance := New(bytes)
	instance.Run()
}
