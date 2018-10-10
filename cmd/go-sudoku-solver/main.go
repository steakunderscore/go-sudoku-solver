package main

import (
	"os"

	gss "github.com/steakunderscore/go-sudoku-solver/pkg/gss"
)

func main() {
	b := gss.Board{}
	b.TakeInput(os.Stdin, os.Stdout)
	b.Solve()
	b.Print(os.Stdout)
}
