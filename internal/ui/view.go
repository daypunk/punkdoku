package ui

import (
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"punkdoku/internal/game"
)

func boardString(m Model) string {
	var b strings.Builder
	var dup [9][9]bool
	if m.autoCheck {
		dup = game.DuplicateMap(m.board.Values, m.cursorRow, m.cursorCol)
	}
	var conf [9][9]bool
	if m.autoCheck {
		conf = game.ConflictMap(m.board.Values, m.board.Given)
	}
	// 셀의 시각적 폭 계산(패딩 포함)
	cellWidth := lipgloss.Width(m.styles.Cell.Render("0"))

	buildLine := func(left, mid, right string) string {
		seg := strings.Repeat("─", cellWidth)
		var sb strings.Builder
		sb.WriteString(left)
		for c := 0; c < 9; c++ {
			sb.WriteString(seg)
			if c == 8 {
				sb.WriteString(right)
			} else if c == 2 || c == 5 {
				sb.WriteString(mid)
			}
		}
		return sb.String()
	}

	// 상단/하단 프레임과 블록 가로 구분선 (코너 포함)
	topBorder := m.styles.RowSep.Render(buildLine("╭", "┬", "╮"))
	midBorder := m.styles.RowSep.Render(buildLine("├", "┼", "┤"))
	botBorder := m.styles.RowSep.Render(buildLine("╰", "┴", "╯"))

	// 상단 프레임
	b.WriteString(topBorder)
	b.WriteString("\n")

	for r := 0; r < 9; r++ {
		// 좌측 외곽선
		b.WriteString(m.styles.ColSep.Render("│"))
		for c := 0; c < 9; c++ {
			// 블록 경계에서만 세로 구분선 출력
			if c > 0 && c%3 == 0 {
				b.WriteString(m.styles.ColSep.Render("│"))
			}
			cell := m.cellView(r, c, dup[r][c], conf[r][c])
			b.WriteString(cell)
		}
		// 우측 외곽선
		b.WriteString(m.styles.ColSep.Render("│"))
		b.WriteString("\n")
		if r == 2 || r == 5 {
			b.WriteString(midBorder)
			b.WriteString("\n")
		}
	}
	// 하단 프레임
	b.WriteString(botBorder)
	return b.String()
}

func Render(m Model) string {
	board := boardString(m)
	// 고정 폭으로 상태줄 중앙 정렬 (46은 스도쿠 보드의 실제 폭)
	status := lipgloss.PlaceHorizontal(46, lipgloss.Center, m.StatusLine())
	// 보드와 상태줄 사이 2줄 공백
	return board + "\n\n\n" + status
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
