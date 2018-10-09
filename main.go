package main

import (
	"fmt"
	"os"
)

type Value int

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

func (b *Board) setTest1() {
	b.values[0] = 1
	b.values[1] = 0
	b.values[2] = 3
	b.values[3] = 0
	b.values[4] = 0
	b.values[5] = 6
	b.values[6] = 0
	b.values[7] = 8
	b.values[8] = 0
	b.values[9] = 0
	b.values[10] = 5
	b.values[11] = 0
	b.values[12] = 0
	b.values[13] = 8
	b.values[14] = 0
	b.values[15] = 1
	b.values[16] = 2
	b.values[17] = 0
	b.values[18] = 7
	b.values[19] = 0
	b.values[20] = 9
	b.values[21] = 1
	b.values[22] = 0
	b.values[23] = 3
	b.values[24] = 0
	b.values[25] = 5
	b.values[26] = 6
	b.values[27] = 0
	b.values[28] = 3
	b.values[29] = 0
	b.values[30] = 0
	b.values[31] = 6
	b.values[32] = 7
	b.values[33] = 0
	b.values[34] = 9
	b.values[35] = 0
	b.values[36] = 5
	b.values[37] = 0
	b.values[38] = 7
	b.values[39] = 8
	b.values[40] = 0
	b.values[41] = 0
	b.values[42] = 0
	b.values[43] = 3
	b.values[44] = 0
	b.values[45] = 8
	b.values[46] = 0
	b.values[47] = 1
	b.values[48] = 0
	b.values[49] = 3
	b.values[50] = 0
	b.values[51] = 5
	b.values[52] = 0
	b.values[53] = 7
	b.values[54] = 0
	b.values[55] = 4
	b.values[56] = 0
	b.values[57] = 0
	b.values[58] = 7
	b.values[59] = 8
	b.values[60] = 0
	b.values[61] = 1
	b.values[62] = 0
	b.values[63] = 6
	b.values[64] = 0
	b.values[65] = 8
	b.values[66] = 0
	b.values[67] = 0
	b.values[68] = 2
	b.values[69] = 0
	b.values[70] = 4
	b.values[71] = 0
	b.values[72] = 0
	b.values[73] = 1
	b.values[74] = 2
	b.values[75] = 0
	b.values[76] = 4
	b.values[77] = 5
	b.values[78] = 0
	b.values[79] = 7
	b.values[80] = 8
	return
}

func (b *Board) takeInput() {
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

func (b *Board) solve() {
	solvedCount := 0
	changed := true
	for solvedCount < 81 && changed == true {
		solvedCount = 0
		changed = false
		for loc, val := range b.values {
			if val != 0 {
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

func main() {
	b := Board{}
	// b.takeInput()
	b.setTest1()
	b.solve()
}
