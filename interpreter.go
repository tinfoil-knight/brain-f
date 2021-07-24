package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Interpreter struct {
	source         []byte
	cells          []byte
	cellPosition   int
	sourcePosition int
}

const DEFAULT_CELLS = 30000

var ErrInvalidSize = errors.New("max cells must be greater than 1")

func New(source []byte) *Interpreter {
	return &Interpreter{
		source:         source,
		cells:          make([]byte, DEFAULT_CELLS),
		cellPosition:   0,
		sourcePosition: 0,
	}
}

func (in *Interpreter) Run() {
	for _, cell := range in.source {
		in.interpret(cell)
	}
	fmt.Println()
}

func (in *Interpreter) RunREPL(source []byte) {
	printLine := false
	for _, cell := range source {
		in.interpret(cell)
		if !printLine && cell == '.' {
			printLine = true
		}
	}
	if printLine {
		fmt.Println()
	}
}

func (in *Interpreter) interpret(cell byte) {
	switch cell {
	case '>':
		if in.cellPosition == len(in.cells)-1 {
			in.cellPosition = 0
		} else {
			in.cellPosition++
		}
	case '<':
		if int(in.cellPosition) == 0 {
			in.cellPosition = len(in.cells) - 1
		} else {
			in.cellPosition--
		}
	case '+':
		in.cells[in.cellPosition]++
	case '-':
		in.cells[in.cellPosition]--
	case ',':
		fmt.Print("~ ")
		reader := bufio.NewReader(os.Stdin)
		char, _ := reader.ReadByte()
		in.cells[in.cellPosition] = char
	case '.':
		fmt.Printf("%c", in.cells[in.cellPosition])
	case '[':
	case ']':

	}
	in.sourcePosition++
}
