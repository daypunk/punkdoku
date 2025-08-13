package generator

import (
	"math/rand"
	"time"

	"punkdoku/internal/solver"
)

// randomizedFullSolution builds a complete valid Sudoku solution using randomized DFS.
func randomizedFullSolution(seed string, timeout time.Duration) (Grid, error) {
	var rng *rand.Rand
	if seed == "" {
		rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	} else {
		rng = rand.New(rand.NewSource(int64(hashStringToUint64(seed))))
	}
	deadline := time.Now().Add(timeout)
	var g Grid
	if fillCellRandom(&g, 0, 0, rng, deadline) {
		return g, nil
	}
	return Grid{}, ErrTimeout
}

func fillCellRandom(g *Grid, row, col int, rng *rand.Rand, deadline time.Time) bool {
	if time.Now().After(deadline) {
		return false
	}
	nextRow, nextCol := row, col+1
	if nextCol == 9 {
		nextRow++
		nextCol = 0
	}
	if row == 9 {
		return true
	}
	vals := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}
	rng.Shuffle(len(vals), func(i, j int) { vals[i], vals[j] = vals[j], vals[i] })
	for _, v := range vals {
		if isSafe(*g, row, col, v) {
			g[row][col] = v
			if fillCellRandom(g, nextRow, nextCol, rng, deadline) {
				return true
			}
			g[row][col] = 0
		}
	}
	return false
}

func isSafe(g Grid, row, col int, v uint8) bool {
	for i := 0; i < 9; i++ {
		if g[row][i] == v || g[i][col] == v {
			return false
		}
	}
	r0 := (row / 3) * 3
	c0 := (col / 3) * 3
	for r := r0; r < r0+3; r++ {
		for c := c0; c < c0+3; c++ {
			if g[r][c] == v {
				return false
			}
		}
	}
	return true
}

// carveCellsUnique removes cells while trying to keep a single solution.
func carveCellsUnique(full Grid, targetRemoved int, seed string, timeout time.Duration) (Grid, error) {
	puzzle := full
	var rng *rand.Rand
	if seed == "" {
		rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	} else {
		rng = rand.New(rand.NewSource(int64(hashStringToUint64(seed) + 0x9e3779b97f4a7c15)))
	}
	deadline := time.Now().Add(timeout)
	cells := make([]int, 81)
	for i := 0; i < 81; i++ { cells[i] = i }
	rng.Shuffle(len(cells), func(i, j int) { cells[i], cells[j] = cells[j], cells[i] })
	removed := 0
	for _, idx := range cells {
		if time.Now().After(deadline) {
			break
		}
		r := idx / 9
		c := idx % 9
		backup := puzzle[r][c]
		puzzle[r][c] = 0
		// Check uniqueness using solver.CountSolutions up to 2
		if solver.CountSolutions(convertToSolverGrid(puzzle), 50*time.Millisecond, 2) != 1 {
			puzzle[r][c] = backup
			continue
		}
		removed++
		if removed >= targetRemoved {
			break
		}
	}
	return puzzle, nil
}

func convertToSolverGrid(g Grid) solver.Grid {
	var s solver.Grid
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			s[r][c] = g[r][c]
		}
	}
	return s
}

// Simple FNV-1a 64-bit hash for seed strings.
func hashStringToUint64(s string) uint64 {
	const (
		offset64 = 1469598103934665603
		prime64  = 1099511628211
	)
	h := uint64(offset64)
	for i := 0; i < len(s); i++ {
		h ^= uint64(s[i])
		h *= prime64
	}
	return h
}
