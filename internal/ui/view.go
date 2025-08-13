package ui

import (
	"strings"
	"time"

	"punkdoku/internal/game"
)

func Render(m Model) string {
	var b strings.Builder
	dup := game.DuplicateMap(m.board.Values, m.cursorRow, m.cursorCol)
	var conf [9][9]bool
	if m.autoCheck {
		conf = game.ConflictMap(m.board.Values, m.solution, m.board.Given)
	}
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if c > 0 {
				sep := " "
				if c%3 == 0 { sep = m.styles.ColSep.Render("│") } else { sep = " " }
				b.WriteString(sep)
			}
			cell := m.cellView(r, c, dup[r][c], conf[r][c])
			b.WriteString(cell)
		}
		b.WriteString("\n")
		if r == 2 || r == 5 {
			b.WriteString(m.styles.RowSep.Render(strings.Repeat("─", 9*3+2)))
			b.WriteString("\n")
		}
	}
	b.WriteString("\n")
	b.WriteString(m.StatusLine())
	return b.String()
}

func (m Model) cellView(r, c int, isDup, isConf bool) string {
	v := m.board.Values[r][c]
	str := "·"
	if v != 0 { str = string('0'+v) }
	style := m.styles.Cell
	if m.board.Given[r][c] {
		style = m.styles.CellFixed
	}
	if isDup {
		style = m.styles.CellDuplicate
	}
	if isConf {
		style = m.styles.CellConflict
	}
	if r == m.cursorRow && c == m.cursorCol {
		style = m.styles.CellSelected
	}
	if deadline, ok := m.flashes[[2]int{r, c}]; ok {
		if time.Now().Before(deadline) {
			style = style.Bold(true)
		}
	}
	return style.Render(str)
}
