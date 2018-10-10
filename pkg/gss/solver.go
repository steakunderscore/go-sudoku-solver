package gss

func (b *Board) Solve() {
	solvedCount := 0
	changed := true
	for solvedCount < 81 && changed == true {
		solvedCount = 0
		changed = false
		for loc, val := range b.Values {
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
	}
}
