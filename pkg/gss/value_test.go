package gss_test

import (
	"testing"

	gss "github.com/steakunderscore/go-sudoku-solver/pkg/gss"
)

func TestSolved(t *testing.T) {
	var v gss.Value
	v = 0
	r := v.Solved()
	if r == true {
		t.Error("Expected false, got ", r)
	}

	v = 8
	r = v.Solved()
	if r == false {
		t.Error("Expected true, got ", r)
	}
}
