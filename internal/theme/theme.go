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
			Foreground:      "#111111",
			GridLine:        "#cfcfcf",
			CellBaseBG:      "#f7f7f7",
			CellBaseFG:      "#1f2937",  // 어두운 회색 (사용자 입력)
			CellFixedFG:     "#6b7280",
			CellFixedBG:     "#efefef",
			CellSelectedBG:  "#cfe3ff",
			CellSelectedFG:  "#1e40af",  // 어두운 파랑 (선택된 셀)
			CellDuplicateBG: "#fff2cc",
			CellConflictBG:  "#ffd6d6",
			Accent:          "#2563eb",
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
			CellBaseBG:      "#0b0f14",
			CellBaseFG:      "#e7f9ff",  // 밝은 청백색 (사용자 입력)
			CellFixedFG:     "#9fb7c6",
			CellFixedBG:     "#0e141a",
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
	// Check environment variables that might indicate light theme
	if termenv.HasDarkBackground() {
		return false
	}
	
	// Check for environment variables indicating light theme
	term := strings.ToLower(os.Getenv("TERM"))
	colorterm := strings.ToLower(os.Getenv("COLORTERM"))
	
	// Some terminal emulators set specific variables for light themes
	if strings.Contains(term, "light") || strings.Contains(colorterm, "light") {
		return true
	}
	
	// Check for common light theme terminal configurations
	if iterm2Profile := os.Getenv("ITERM_PROFILE"); iterm2Profile != "" {
		profileLower := strings.ToLower(iterm2Profile)
		if strings.Contains(profileLower, "light") || strings.Contains(profileLower, "bright") {
			return true
		}
	}
	
	// macOS Terminal.app profile detection
	if termProgram := os.Getenv("TERM_PROGRAM"); termProgram == "Apple_Terminal" {
		// This is a basic heuristic - in practice you might need more sophisticated detection
		// Check if we can query the terminal for background color
		output := termenv.NewOutput(termenv.DefaultOutput())
		if output.Profile == termenv.TrueColor {
			// TrueColor terminals might support background color queries
			// but this is complex to implement reliably
		}
	}
	
	// Default to dark background if we can't determine
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
			"selected":  "#eab308", // darker yellow for light bg
			"panel":     "#6b7280", // darker gray for light bg border
			"success":   "#16a34a", // darker green for light bg
			"error":     "#dc2626", // red for light bg
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
