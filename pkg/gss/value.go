package gss

type Value int

const UnsolvedValue Value = 0

func (v Value) Solved() bool {
	return v != UnsolvedValue
}
