package solver

import "time"

// Grid matches generator's grid representation
 type Grid [9][9]uint8

// Solve attempts to fill the grid in-place using backtracking.
// Returns whether a solution was found before timeout.
func Solve(g *Grid, timeout time.Duration) bool {
	deadline := time.Now().Add(timeout)
	return solveBacktrack(g, deadline)
}

// CountSolutions counts up to maxCount solutions for uniqueness check.
func CountSolutions(g Grid, timeout time.Duration, maxCount int) int {
	deadline := time.Now().Add(timeout)
	count := 0
	var dfs func(*Grid) bool
	dfs = func(cur *Grid) bool {
		if time.Now().After(deadline) {
			return true
		}
		row, col, ok := findEmpty(*cur)
		if !ok {
			count++
			return count >= maxCount
		}
		cands := candidates(*cur, row, col)
		for _, v := range cands {
			cur[row][col] = v
			if dfs(cur) {
				return true
			}
			cur[row][col] = 0
		}
		return false
	}
	copyGrid := g
	dfs(&copyGrid)
	return count
}

func solveBacktrack(g *Grid, deadline time.Time) bool {
	if time.Now().After(deadline) {
		return false
	}
	row, col, ok := findEmpty(*g)
	if !ok {
		return true
	}
	cands := candidates(*g, row, col)
	for _, v := range cands {
		g[row][col] = v
		if solveBacktrack(g, deadline) {
			return true
		}
		g[row][col] = 0
	}
	return false
}

func findEmpty(g Grid) (int, int, bool) {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if g[r][c] == 0 {
				return r, c, true
			}
		}
	}
	return 0, 0, false
}

func candidates(g Grid, row, col int) []uint8 {
	used := [10]bool{}
	for i := 0; i < 9; i++ {
		used[g[row][i]] = true
		used[g[i][col]] = true
	}
	r0 := (row / 3) * 3
	c0 := (col / 3) * 3
	for r := r0; r < r0+3; r++ {
		for c := c0; c < c0+3; c++ {
			used[g[r][c]] = true
		}
	}
	var out []uint8
	for v := 1; v <= 9; v++ {
		if !used[v] {
			out = append(out, uint8(v))
		}
	}
	return out
}
