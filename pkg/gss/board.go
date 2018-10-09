package gss

import (
	"fmt"
	"os"
)

type Board struct {
	values [9 * 9]Value
}

func (b *Board) finalizeValue(value Value, loc int) {
	b.values[loc] = value
}

func (b *Board) print() {
	for i := 0; i < 9; i++ {
		row := b.row(i)
		fmt.Printf("%d %d %d %d %d %d %d %d %d\n", row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8])
	}
	fmt.Println()
}

func (b *Board) TakeInput() {
	fmt.Println("Please enter starting Sudoku:")
	for i := 0; i < 9*9; i += 9 {
		_, err := fmt.Scanf("%d %d %d %d %d %d %d %d %d\n", &b.values[i+0], &b.values[i+1], &b.values[i+2], &b.values[i+3], &b.values[i+4], &b.values[i+5], &b.values[i+6], &b.values[i+7], &b.values[i+8])
		if err != nil {
			fmt.Printf("Input not formatted\n")
			os.Exit(2)
		}
	}
}

func (b *Board) row(i int) []Value {
	return b.values[i*9 : i*9+9]
}

func (b *Board) col(column int) []Value {
	r := make([]Value, 9)
	for i := 0; i < 9; i++ {
		r[i] = b.values[column+i*9]
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
		for j, v := range b.values[fromLoc : fromLoc+3] {
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

func (b *Board) Solve() {
	solvedCount := 0
	changed := true
	for solvedCount < 81 && changed == true {
		solvedCount = 0
		changed = false
		for loc, val := range b.values {
			if val.Solved() {
				solvedCount++
				continue
			}
			taken := b.taken(loc)
			count := 0
			for j := Value(1); j <= Value(9); j++ {
				if taken[j] {
					count++
				}
			}
			if count == 8 {
				for missing := Value(1); missing <= Value(9); missing++ {
					if !taken[missing] {
						b.finalizeValue(missing, loc)
						solvedCount++
						changed = true
						break
					}
				}
			}
		}
		b.print()
	}
}
