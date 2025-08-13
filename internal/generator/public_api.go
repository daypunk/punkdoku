package generator

// Puzzle returns a copy of the grid suitable for rendering: 0 means blank.
func (g Grid) Puzzle() [9][9]uint8 {
	return g
}
