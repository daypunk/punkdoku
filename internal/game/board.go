package game

import (
	"time"
)

type Grid [9][9]uint8

type Move struct {
	Row int
	Col int
	Prev uint8
	Next uint8
	At   time.Time
}

type Board struct {
	Given [9][9]bool
	Values Grid
}

func NewBoardFromPuzzle(p Grid) Board {
	var b Board
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			v := p[r][c]
			if v != 0 {
				b.Given[r][c] = true
			}
			b.Values[r][c] = v
		}
	}
	return b
}

func (b *Board) IsGiven(row, col int) bool { return b.Given[row][col] }

func (b *Board) SetValue(row, col int, v uint8) (prev uint8, ok bool) {
	if b.Given[row][col] {
		return b.Values[row][col], false
	}
	prev = b.Values[row][col]
	b.Values[row][col] = v
	return prev, true
}

func InBounds(row, col int) bool { return row >= 0 && row < 9 && col >= 0 && col < 9 }

// DuplicateMap marks cells that duplicate the selected cell's value across row/col/box.
func DuplicateMap(g Grid, selRow, selCol int) [9][9]bool {
	var dup [9][9]bool
	v := g[selRow][selCol]
	if v == 0 { return dup }
	for i := 0; i < 9; i++ {
		if g[selRow][i] == v && i != selCol { dup[selRow][i] = true }
		if g[i][selCol] == v && i != selRow { dup[i][selCol] = true }
	}
	r0 := (selRow/3)*3
	c0 := (selCol/3)*3
	for r := r0; r < r0+3; r++ {
		for c := c0; c < c0+3; c++ {
			if (r != selRow || c != selCol) && g[r][c] == v {
				dup[r][c] = true
			}
		}
	}
	return dup
}

// DuplicateMapAll marks any duplicates in rows, columns, or 3x3 blocks across the entire grid.
func DuplicateMapAll(g Grid) [9][9]bool {
	var dup [9][9]bool
	// rows
	for r := 0; r < 9; r++ {
		count := map[uint8]int{}
		for c := 0; c < 9; c++ {
			v := g[r][c]
			if v != 0 { count[v]++ }
		}
		for c := 0; c < 9; c++ {
			v := g[r][c]
			if v != 0 && count[v] > 1 { dup[r][c] = true }
		}
	}
	// cols
	for c := 0; c < 9; c++ {
		count := map[uint8]int{}
		for r := 0; r < 9; r++ {
			v := g[r][c]
			if v != 0 { count[v]++ }
		}
		for r := 0; r < 9; r++ {
			v := g[r][c]
			if v != 0 && count[v] > 1 { dup[r][c] = true }
		}
	}
	// blocks
	for br := 0; br < 3; br++ {
		for bc := 0; bc < 3; bc++ {
			count := map[uint8]int{}
			for r := br*3; r < br*3+3; r++ {
				for c := bc*3; c < bc*3+3; c++ {
					v := g[r][c]
					if v != 0 { count[v]++ }
				}
			}
			for r := br*3; r < br*3+3; r++ {
				for c := bc*3; c < bc*3+3; c++ {
					v := g[r][c]
					if v != 0 && count[v] > 1 { dup[r][c] = true }
				}
			}
		}
	}
	return dup
}

// ConflictMap marks cells that violate Sudoku constraints (duplicates), excluding givens.
func ConflictMap(values Grid, given [9][9]bool) [9][9]bool {
	all := DuplicateMapAll(values)
	var bad [9][9]bool
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if given[r][c] { continue }
			bad[r][c] = all[r][c]
		}
	}
	return bad
}
