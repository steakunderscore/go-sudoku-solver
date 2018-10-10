package gss_test

import (
	"testing"

	. "github.com/steakunderscore/go-sudoku-solver/pkg/gss"
)

// Test the half implemented algorithm. Note
// this should be removed once Solve is
// properly implemented
func TestSolveHalf(t *testing.T) {
	var values = [9 * 9]Value{
		1, 0, 3, 0, 0, 6, 0, 8, 0,
		0, 5, 0, 0, 8, 0, 1, 2, 0,
		7, 0, 9, 1, 0, 3, 0, 5, 6,
		0, 3, 0, 0, 6, 7, 0, 9, 0,
		5, 0, 7, 8, 0, 0, 0, 3, 0,
		8, 0, 1, 0, 3, 0, 5, 0, 7,
		0, 4, 0, 0, 7, 8, 0, 1, 0,
		6, 0, 8, 0, 0, 2, 0, 4, 0,
		0, 1, 2, 0, 4, 5, 0, 7, 8}
	var solution = [9 * 9]Value{
		1, 2, 3, 4, 9, 6, 7, 8, 5,
		4, 5, 6, 9, 8, 0, 1, 2, 3,
		7, 8, 9, 1, 2, 3, 0, 5, 6,
		0, 3, 0, 2, 6, 7, 0, 9, 0,
		5, 0, 7, 8, 0, 4, 9, 3, 0,
		8, 0, 1, 0, 3, 9, 5, 0, 7,
		0, 4, 0, 0, 7, 8, 3, 1, 9,
		6, 0, 8, 0, 0, 2, 0, 4, 0,
		0, 1, 2, 6, 4, 5, 0, 7, 8}

	var b = Board{Values: values}
	b.Solve()
	if b.Values != solution {
		t.Error("Expected to solve, got ", b.Values)
	}
}

func TestSolveFull(t *testing.T) {
	t.Skip("Algorithm not fully implemented")
	var values = [9 * 9]Value{
		1, 0, 3, 0, 0, 6, 0, 8, 0,
		0, 5, 0, 0, 8, 0, 1, 2, 0,
		7, 0, 9, 1, 0, 3, 0, 5, 6,
		0, 3, 0, 0, 6, 7, 0, 9, 0,
		5, 0, 7, 8, 0, 0, 0, 3, 0,
		8, 0, 1, 0, 3, 0, 5, 0, 7,
		0, 4, 0, 0, 7, 8, 0, 1, 0,
		6, 0, 8, 0, 0, 2, 0, 4, 0,
		0, 1, 2, 0, 4, 5, 0, 7, 8}
	var solution = [9 * 9]Value{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
		4, 5, 6, 7, 8, 9, 1, 2, 3,
		7, 8, 9, 1, 2, 3, 4, 5, 6,
		2, 3, 4, 5, 6, 7, 8, 9, 1,
		5, 6, 7, 8, 9, 1, 2, 3, 4,
		8, 9, 1, 2, 3, 4, 5, 6, 7,
		3, 4, 5, 6, 7, 8, 9, 1, 2,
		6, 7, 8, 9, 1, 2, 3, 4, 5,
		9, 1, 2, 3, 4, 5, 6, 7, 8}

	var b = Board{Values: values}
	b.Solve()
	if b.Values != solution {
		t.Error("Expected to solve, got ", b.Values)
	}
}
