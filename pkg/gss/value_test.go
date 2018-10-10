package gss_test

import (
	"testing"

	. "github.com/steakunderscore/go-sudoku-solver/pkg/gss"
)

func TestSolved(t *testing.T) {
	var v Value
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
