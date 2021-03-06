# brain-f

This is an interpreter written in Go for the esoteric language [BrainF\*\*k](https://esolangs.org/wiki/Brainfuck). See some of the examples [here](https://en.wikipedia.org/wiki/Brainfuck#Examples).

## Getting Started

### Pre-requisites

- [Go >= 1.14](https://golang.org/)
- [GNU Make](https://www.gnu.org/software/make/)

### Creating a build

```bash
make build
```

This will create an executable binary with the name `bf`.

### Usage

- Starting a REPL

```bash
./bf
```

- Running a file

```bash
./bf <file-path>
```

### Caveat(s)

The number of cells for the interpreter are fixed and moving beyond the last cells causes it to wrap (and vice versa in the case of moving left of the first cell).

## License

[MIT](./LICENSE)

## Author

- **_Kunal Kundu_** [@tinfoil-knight](https://github.com/tinfoil-knight)

## Attribution

- [Brainfuck Visualizer](https://fatiherikli.github.io/brainfuck-visualizer/#) by [fatiherikli](https://github.com/fatiherikli)
