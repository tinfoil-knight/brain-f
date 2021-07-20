package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	switch len(args) {
	case 0:
		startPrompt()
	case 1:
		runFile(args[1])
	default:
		panic(errors.New("usage: bf <path-to-file>"))
	}
}

func startPrompt() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		raw, _ := reader.ReadString('\n')
		text := strings.TrimSpace(raw)
		// TODO
		fmt.Println(text)
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
