package ui

import (
	"github.com/charmbracelet/lipgloss"
	"punkdoku/internal/theme"
)

type UIStyles struct {
	App           lipgloss.Style
	Panel         lipgloss.Style
	Banner        lipgloss.Style
	MenuItem      lipgloss.Style
	MenuItemSelected lipgloss.Style
	Hint          lipgloss.Style

	BoolTrue      lipgloss.Style
	BoolFalse     lipgloss.Style

	Board         lipgloss.Style
	RowSep        lipgloss.Style
	ColSep        lipgloss.Style
	Cell          lipgloss.Style
	CellFixed     lipgloss.Style
	CellSelected  lipgloss.Style
	CellDuplicate lipgloss.Style
	CellConflict  lipgloss.Style
	Status        lipgloss.Style
	StatusError   lipgloss.Style

	DiffBox       lipgloss.Style
}

func BuildStyles(t theme.Theme) UIStyles {
	gridColor := lipgloss.Color(t.Palette.GridLine)
	accent := lipgloss.Color(t.Palette.Accent)
	gray := lipgloss.Color("#9ca3af")
	red := lipgloss.Color("#ef4444")
	return UIStyles{
		App:              lipgloss.NewStyle().Foreground(lipgloss.Color(t.Palette.Foreground)),
		Panel:            lipgloss.NewStyle().Padding(0, 4).Margin(1, 4).Border(lipgloss.RoundedBorder()).BorderForeground(gridColor),
		Banner:           lipgloss.NewStyle().Foreground(accent).Bold(true),
		MenuItem:         lipgloss.NewStyle().Foreground(lipgloss.Color(t.Palette.Foreground)),
		MenuItemSelected: lipgloss.NewStyle().Foreground(accent).Bold(true),
		Hint:             lipgloss.NewStyle().Foreground(accent),

		BoolTrue:  lipgloss.NewStyle().Foreground(lipgloss.Color("#22c55e")).Bold(true),
		BoolFalse: lipgloss.NewStyle().Foreground(gray),

		Board:         lipgloss.NewStyle(),
		RowSep:        lipgloss.NewStyle().Foreground(gridColor),
		ColSep:        lipgloss.NewStyle().Foreground(gridColor),
		Cell:          lipgloss.NewStyle().Background(lipgloss.Color(t.Palette.CellBaseBG)).Padding(0, 1),
		CellFixed:     lipgloss.NewStyle().Background(lipgloss.Color(t.Palette.CellFixedBG)).Foreground(lipgloss.Color(t.Palette.CellFixedFG)).Padding(0, 1).Bold(true),
		CellSelected:  lipgloss.NewStyle().Background(lipgloss.Color(t.Palette.CellSelectedBG)).Padding(0, 1).Bold(true),
		CellDuplicate: lipgloss.NewStyle().Background(lipgloss.Color(t.Palette.CellDuplicateBG)).Padding(0, 1),
		CellConflict:  lipgloss.NewStyle().Background(lipgloss.Color(t.Palette.CellConflictBG)).Padding(0, 1).Bold(true),
		Status:        lipgloss.NewStyle().Foreground(gray),
		StatusError:   lipgloss.NewStyle().Foreground(red).Bold(true),

		DiffBox: lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(accent).Padding(1, 4),
	}
}
