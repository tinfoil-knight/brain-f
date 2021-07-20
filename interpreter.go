package main

import "errors"

type Interpreter struct {
	source   []byte
	cells    []byte
	position int
}

const DEFAULT_CELLS = 30000

var ErrInvalidSize = errors.New("max cells must be greater than 1")

func New(source []byte) *Interpreter {
	return &Interpreter{
		source:   source,
		cells:    make([]byte, DEFAULT_CELLS),
		position: 0,
	}
}

func (in *Interpreter) Run() {
	for _, cell := range in.source {
		in.interpret(cell)
	}
}

func (in *Interpreter) interpret(cell byte) {
	switch cell {
	case '>':
		if in.position == len(in.cells)-1 {
			in.position = 0
		} else {
			in.position++
		}
	case '<':
		if int(in.position) == 0 {
			in.position = len(in.cells) - 1
		} else {
			in.position--
		}
	case '+':
		in.cells[in.position]++
	case '-':
		in.cells[in.position]--
	case ',':
	case '.':
	case '[':
	case ']':

	}

}
