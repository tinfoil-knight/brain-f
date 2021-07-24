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
	bracketStack   []int
}

const DEFAULT_CELLS = 30000

var ErrInvalidSize = errors.New("max cells must be greater than 1")
var ErrInvalidSyntax = errors.New("invalid syntax")

func New(source []byte) *Interpreter {
	return &Interpreter{
		source:         source,
		cells:          make([]byte, DEFAULT_CELLS),
		cellPosition:   0,
		sourcePosition: 0,
		bracketStack:   make([]int, 0),
	}
}

func (in *Interpreter) Run() {
	in.interpret()
	fmt.Println()
}

func (in *Interpreter) RunREPL(source []byte) {
	in.source = source
	in.interpret()
	for _, a := range source {
		if a == '.' {
			fmt.Println()
			return
		}
	}
}

func (in *Interpreter) interpret() {
	end := false
	for !end {
		switch in.source[in.sourcePosition] {
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
			if in.cells[in.cellPosition] != 0 {
				in.bracketStack = append(in.bracketStack, in.sourcePosition)
			} else {
				c := 0
			Loop:
				for {
					in.sourcePosition++
					if in.sourcePosition == len(in.source) {
						// TODO: Perform this check beforehand
						panic(ErrInvalidSyntax)
					}
					switch in.source[in.sourcePosition] {
					case '[':
						c++
					case ']':
						if c > 0 {
							c--
						} else {
							break Loop
						}
					}
				}
			}
		case ']':
			a := in.bracketStack
			if len(a) == 0 {
				panic(ErrInvalidSyntax)
			}
			x, a := a[len(a)-1], a[:len(a)-1]
			in.bracketStack = a
			in.sourcePosition = x - 1
		}
		in.sourcePosition++
		if in.sourcePosition == len(in.source) {
			end = true
		}
	}
}
