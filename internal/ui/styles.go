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
	
	// Adaptive colors
	adaptiveColors := theme.NewAdaptiveColors(t)
	accentColors := adaptiveColors.GetAccentColors()
	
	gray := lipgloss.Color("#9ca3af")
	if t.Name == "light" {
		gray = lipgloss.Color("#6b7280") // darker gray for light theme
	}
	
	// 화이트모드에서 메뉴 아이템 색상 조정
	menuItemColor := lipgloss.Color(t.Palette.Foreground)
	statusColor := gray // 다크모드에서 상태줄 회색
	if t.Name == "light" {
		menuItemColor = lipgloss.Color("#000000") // 화이트모드에서 미선택 난이도 검은색
		statusColor = menuItemColor // 화이트모드에서 상태줄 검은색
	}
	
	return UIStyles{
		App:              lipgloss.NewStyle().Foreground(lipgloss.Color(t.Palette.Foreground)),
		Panel:            lipgloss.NewStyle().Padding(0, 4).Margin(1, 4).Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color(accentColors["panel"])),
		Banner:           lipgloss.NewStyle().Foreground(accent).Bold(true),
		MenuItem:         lipgloss.NewStyle().Foreground(menuItemColor),
		MenuItemSelected: lipgloss.NewStyle().Foreground(accent).Bold(true),
		Hint:             lipgloss.NewStyle().Foreground(accent),

		BoolTrue:  lipgloss.NewStyle().Foreground(lipgloss.Color("#16a34a")).Bold(true), // 다크모드에서 어두운 초록색
		BoolFalse: lipgloss.NewStyle().Foreground(gray),

		Board:         lipgloss.NewStyle(),
		RowSep:        lipgloss.NewStyle().Foreground(gridColor),
		ColSep:        lipgloss.NewStyle().Foreground(gridColor),
		Cell:          lipgloss.NewStyle().Background(lipgloss.Color(t.Palette.CellBaseBG)).Foreground(lipgloss.Color(t.Palette.CellBaseFG)).Padding(0, 1),
		CellFixed:     lipgloss.NewStyle().Background(lipgloss.Color(t.Palette.CellFixedBG)).Foreground(lipgloss.Color(t.Palette.CellFixedFG)).Padding(0, 1).Bold(true),
		CellSelected:  lipgloss.NewStyle().Background(lipgloss.Color(t.Palette.CellSelectedBG)).Foreground(lipgloss.Color(t.Palette.CellSelectedFG)).Padding(0, 1).Bold(true),
		CellDuplicate: lipgloss.NewStyle().Background(lipgloss.Color(t.Palette.CellDuplicateBG)).Padding(0, 1),
		CellConflict:  lipgloss.NewStyle().Background(lipgloss.Color(t.Palette.CellConflictBG)).Padding(0, 1).Bold(true),
		Status:        lipgloss.NewStyle().Foreground(statusColor), // 다크모드에서 회색, 화이트모드에서 검은색
		StatusError:   lipgloss.NewStyle().Foreground(lipgloss.Color(accentColors["error"])).Bold(true),

		DiffBox: lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(accent).Padding(1, 4),
	}
}
