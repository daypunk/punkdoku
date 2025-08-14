package theme

import (
	"os"
	"strings"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

type Palette struct {
	Background string
	Foreground string
	GridLine   string
	CellBaseBG string
	CellBaseFG string  // 사용자 입력 문자 색상
	CellFixedFG string
	CellFixedBG string
	CellSelectedBG string
	CellSelectedFG string // 선택된 셀 문자 색상
	CellDuplicateBG string
	CellConflictBG string
	Accent string
}

type Theme struct {
	Name    string
	Palette Palette
}



func Light() Theme {
	return Theme{
		Name: "light",
		Palette: Palette{
			Background:      "#ffffff",
			Foreground:      "#000000",  // 메인 텍스트만 검은색
			GridLine:        "#cfcfcf",  // 그리드 라인 원복
			CellBaseBG:      "",         // 투명 배경
			CellBaseFG:      "#000000",  // 입력 텍스트 검은색
			CellFixedFG:     "#6b7280",  // 고정 숫자 원복
			CellFixedBG:     "",         // 투명 배경
			CellSelectedBG:  "#cfe3ff",  // 선택된 셀 배경 원복
			CellSelectedFG:  "#000000",  // 선택된 셀 텍스트 검은색
			CellDuplicateBG: "#fff2cc",  // 중복은 배경 유지
			CellConflictBG:  "#ffd6d6",  // 충돌은 배경 유지
			Accent:          "#ff6600",  // 주황색 액센트
		},
	}
}



// Punk returns a single vivid cyan-accent theme used across the app.
func Punk() Theme {
	return Theme{
		Name: "punk",
		Palette: Palette{
			Background:      "#000000",
			Foreground:      "#e7f9ff",
			GridLine:        "#1f2937",
			CellBaseBG:      "",         // 미선택 셀 배경 투명
			CellBaseFG:      "#e7f9ff",  // 밝은 청백색 (사용자 입력)
			CellFixedFG:     "#9fb7c6",
			CellFixedBG:     "",         // 고정 셀 배경도 투명
			CellSelectedBG:  "#063a3e",
			CellSelectedFG:  "#00e5ff",  // 사이언 (선택된 셀)
			CellDuplicateBG: "#3d2d0a",
			CellConflictBG:  "#5b1515",
			Accent:          "#00e5ff",
		},
	}
}

// DetectTheme automatically detects the terminal background and returns appropriate theme
func DetectTheme() Theme {
	// Try to detect if terminal has light background
	if hasLightBackground() {
		return Light()
	}
	
	// Default to punk theme for dark backgrounds
	return Punk()
}

// hasLightBackground attempts to detect if the terminal has a light background
func hasLightBackground() bool {
	// Primary detection using termenv
	if !termenv.HasDarkBackground() {
		return true
	}
	
	// Additional environment variable checks
	term := strings.ToLower(os.Getenv("TERM"))
	colorterm := strings.ToLower(os.Getenv("COLORTERM"))
	termProgram := strings.ToLower(os.Getenv("TERM_PROGRAM"))
	
	// Check for light theme indicators in environment variables
	lightIndicators := []string{"light", "bright", "white"}
	for _, indicator := range lightIndicators {
		if strings.Contains(term, indicator) || 
		   strings.Contains(colorterm, indicator) || 
		   strings.Contains(termProgram, indicator) {
			return true
		}
	}
	
	// Check iTerm2 profile
	if iterm2Profile := strings.ToLower(os.Getenv("ITERM_PROFILE")); iterm2Profile != "" {
		for _, indicator := range lightIndicators {
			if strings.Contains(iterm2Profile, indicator) {
				return true
			}
		}
	}
	
	// Force detection for testing - uncomment to test light mode
	// return true
	
	// Default to dark background
	return false
}

// AdaptiveColors provides theme-aware color mappings
type AdaptiveColors struct {
	theme Theme
}

func NewAdaptiveColors(t Theme) AdaptiveColors {
	return AdaptiveColors{theme: t}
}

// GetDifficultyColors returns colors for each difficulty level adapted to the theme
func (ac AdaptiveColors) GetDifficultyColors() map[string]string {
	if ac.theme.Name == "light" {
		return map[string]string{
			"Easy":      "#0891b2", // darker cyan for light bg
			"Normal":    "#16a34a", // darker green for light bg
			"Hard":      "#dc2626", // red for light bg
			"Lunatic": "#7c2d92", // purple for light bg
			"Daily":     "#16a34a", // darker green for light bg
		}
	}
	// Dark theme colors (original)
	return map[string]string{
		"Easy":      "#06b6d4", // cyan
		"Normal":    "#22c55e", // green
		"Hard":      "#f59e0b", // orange
		"Lunatic": "#7c3aed", // violet
		"Daily":     "#22c55e", // green
	}
}

// GetGradientColors returns gradient color pairs adapted to the theme
func (ac AdaptiveColors) GetGradientColors() map[string][2]string {
	if ac.theme.Name == "light" {
		return map[string][2]string{
			"banner":    {"#7c2d92", "#be185d"}, // darker purple to darker pink
			"easy":      {"#0891b2", "#1e40af"}, // cyan to blue for light bg
			"normal":    {"#16a34a", "#eab308"}, // green to yellow for light bg
			"daily":     {"#16a34a", "#eab308"}, // green to yellow for light bg
			"hard":      {"#dc2626", "#991b1b"}, // red gradient for light bg
			"lunatic": {"#7c2d92", "#be185d"}, // purple to pink for light bg
			"complete":  {"#7c2d92", "#be185d"}, // success gradient for light bg
		}
	}
	// Dark theme gradients (original)
	return map[string][2]string{
		"banner":    {"#7c3aed", "#ec4899"}, // violet to pink
		"easy":      {"#06b6d4", "#3b82f6"}, // cyan to blue for dark bg
		"normal":    {"#22c55e", "#facc15"}, // green to yellow for dark bg
		"daily":     {"#22c55e", "#facc15"}, // green to yellow for dark bg
		"hard":      {"#f59e0b", "#ef4444"}, // orange to red
		"lunatic": {"#7c3aed", "#ec4899"}, // violet to pink
		"complete":  {"#7c3aed", "#ec4899"}, // violet to pink
	}
}

// GetAccentColors returns various accent colors adapted to the theme
func (ac AdaptiveColors) GetAccentColors() map[string]string {
	if ac.theme.Name == "light" {
		return map[string]string{
			"selected":  "#ff6600", // 주황색 선택
			"panel":     "#6b7280", // 패널 테두리 원복
			"success":   "#16a34a", // 성공 메시지 원복
			"error":     "#dc2626", // 빨간색 에러
		}
	}
	// Dark theme accents (original)
	return map[string]string{
		"selected":  "#facc15", // bright yellow
		"panel":     "#374151", // gray
		"success":   "#22c55e", // green
		"error":     "#ef4444", // red
	}
}

func BaseStyle(t Theme) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(lipgloss.Color(t.Palette.Foreground)).Background(lipgloss.Color(t.Palette.Background))
}
