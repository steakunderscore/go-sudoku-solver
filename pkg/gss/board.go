package gss

import (
	"fmt"
	"io"
	"os"
)

type Board struct {
	Values [9 * 9]Value
}

func (b *Board) finalizeValue(value Value, loc int) {
	b.Values[loc] = value
}

func (b *Board) Print(out io.Writer) {
	fmt.Fprintln(out, "Result:")
	for i := 0; i < 9; i++ {
		row := b.row(i)
		for j := 0; j < (9 - 1); j++ {
			fmt.Fprintf(out, "%d ", row[j])
		}
		fmt.Fprintf(out, "%d\n", row[8])
	}
	fmt.Println()
}

func (b *Board) TakeInput(in io.Reader, out io.Writer) {
	fmt.Fprintln(out, "Enter starting Sudoku:")
	for i := 0; i < 9*9; i += 9 {
		_, err := fmt.Fscanf(in, "%d %d %d %d %d %d %d %d %d\n", &b.Values[i+0], &b.Values[i+1], &b.Values[i+2], &b.Values[i+3], &b.Values[i+4], &b.Values[i+5], &b.Values[i+6], &b.Values[i+7], &b.Values[i+8])
		if err != nil {
			fmt.Fprintf(out, "Input not formatted\n")
			os.Exit(2)
		}
	}
}

func (b *Board) row(i int) []Value {
	return b.Values[i*9 : i*9+9]
}

func (b *Board) col(column int) []Value {
	r := make([]Value, 9)
	for i := 0; i < 9; i++ {
		r[i] = b.Values[column+i*9]
	}
	return r
}

func (b *Board) findSquare(location int) (sRow int, sCol int) {
	sRow = (location % 9) / 3
	sCol = location / (3 * 9)
	return sRow, sCol
}

func (b *Board) findRow(location int) (row int) {
	return location / 9
}

func (b *Board) findCol(location int) (col int) {
	return location % 9
}

func (b *Board) square(sRow int, sCol int) []Value {
	r := make([]Value, 9)
	for i := 0; i < 3; i++ {
		toLoc := i * 3
		fromLoc := i*9 + (sRow * 9 * 3)
		for j, v := range b.Values[fromLoc : fromLoc+3] {
			r[toLoc+j] = v
		}
	}
	return r
}

func (b *Board) taken(location int) map[Value]bool {
	taken := make(map[Value]bool)
	rv := b.row(b.findRow(location))
	cv := b.col(b.findCol(location))
	sq := b.square(b.findSquare(location))
	for _, v := range rv {
		if v == 0 {
			continue
		}
		taken[v] = true
	}
	for _, v := range cv {
		if v == 0 {
			continue
		}
		taken[v] = true
	}
	for _, v := range sq {
		if v == 0 {
			continue
		}
		taken[v] = true
	}
	return taken
}
