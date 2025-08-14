package generator

import (
	"errors"
	"time"
)

// Difficulty represents puzzle difficulty tiers.
type Difficulty int

const (
	Easy Difficulty = iota
	Normal
	Hard
	Lunatic
)

// DailySeed returns a stable seed based on UTC date (YYYY-MM-DD).
func DailySeed(t time.Time) string {
	utc := t.UTC()
	return utc.Format("2006-01-02")
}

// Params controls generation knobs derived from difficulty.
type Params struct {
	// number of blanks/removed cells; higher -> harder
	RemovedCells int
	// backtracking timeout to avoid worst-cases
	Timeout time.Duration
}

// paramsFor maps Difficulty to generation parameters.
func paramsFor(d Difficulty) Params {
	switch d {
	case Easy:
		return Params{RemovedCells: 38, Timeout: 150 * time.Millisecond}
	case Normal:
		return Params{RemovedCells: 46, Timeout: 150 * time.Millisecond}
	case Hard:
		return Params{RemovedCells: 52, Timeout: 200 * time.Millisecond}
	case Lunatic:
		return Params{RemovedCells: 58, Timeout: 250 * time.Millisecond}
	default:
		return Params{RemovedCells: 46, Timeout: 150 * time.Millisecond}
	}
}

// Grid is a 9x9 Sudoku grid. 0 represents empty.
type Grid [9][9]uint8

// ErrTimeout is returned when generation exceeds the configured timeout.
var ErrTimeout = errors.New("generation timed out")

// Generate creates a Sudoku puzzle with the given difficulty and seed.
// - If seed is empty, uses current time for randomness.
// - For Daily mode, pass seed from DailySeed(date).
// Returns a puzzle grid with 0 as blanks, aimed at single-solution.
func Generate(d Difficulty, seed string) (Grid, error) {
	p := paramsFor(d)
	return generateWithParams(p, seed)
}

// GenerateDaily creates a daily puzzle based on UTC date.
func GenerateDaily(date time.Time) (Grid, error) {
	return Generate(Normal, DailySeed(date))
}

// generateWithParams contains the core generation pipeline.
func generateWithParams(p Params, seed string) (Grid, error) {
	// 1) Create a full valid solution via randomized backtracking
	full, err := randomizedFullSolution(seed, p.Timeout)
	if err != nil {
		return Grid{}, err
	}
	// 2) Remove cells according to difficulty while keeping uniqueness if possible
	puzzle, err := carveCellsUnique(full, p.RemovedCells, seed, p.Timeout)
	if err != nil {
		return Grid{}, err
	}
	return puzzle, nil
}
