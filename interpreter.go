package main

import "errors"

type Interpreter struct {
	source []byte
	cells  []byte
}

var ErrInvalidSize = errors.New("max cells must be greater than 1")

func New(source []byte) *Interpreter {
	return &Interpreter{
		source: source,
		cells:  make([]byte, 30),
	}
}

func (in *Interpreter) Run() {
	for _, cell := range in.source {
		in.interpret(cell)
	}
}

func (in *Interpreter) interpret(cell byte) {

}
