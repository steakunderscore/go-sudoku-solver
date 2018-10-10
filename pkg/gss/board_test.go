package gss_test

import (
	"io"
	"io/ioutil"
	"testing"

	gss "github.com/steakunderscore/go-sudoku-solver/pkg/gss"
)

func TestPrint(t *testing.T) {
	var b = gss.Board{Values: [9 * 9]gss.Value{
		1, 0, 3, 0, 0, 6, 0, 8, 0, 0, 5, 0, 0, 8, 0, 1, 2, 0, 7, 0, 9, 1, 0, 3, 0, 5, 6, 0, 3, 0, 0, 6, 7, 0, 9, 0, 5, 0, 7, 8, 0, 0, 0, 3, 0, 8, 0, 1, 0, 3, 0, 5, 0, 7, 0, 4, 0, 0, 7, 8, 0, 1, 0, 6, 0, 8, 0, 0, 2, 0, 4, 0, 0, 1, 2, 0, 4, 5, 0, 7, 8,
	}}
  expected := "Result:\n1 0 3 0 0 6 0 8 0\n0 5 0 0 8 0 1 2 0\n7 0 9 1 0 3 0 5 6\n0 3 0 0 6 7 0 9 0\n5 0 7 8 0 0 0 3 0\n8 0 1 0 3 0 5 0 7\n0 4 0 0 7 8 0 1 0\n6 0 8 0 0 2 0 4 0\n0 1 2 0 4 5 0 7 8\n"
	r, w := io.Pipe()
	go func() {
		b.Print(w)
		w.Close()
	}()
	stdout, _ := ioutil.ReadAll(r)
	result := string(stdout)
	if result != expected {
		t.Error("Expected\n", expected, "but got\n", result)
	}
}
