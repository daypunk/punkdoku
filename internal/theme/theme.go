package theme

import "github.com/charmbracelet/lipgloss"

type Palette struct {
	Background string
	Foreground string
	GridLine   string
	CellBaseBG string
	CellFixedFG string
	CellFixedBG string
	CellSelectedBG string
	CellDuplicateBG string
	CellConflictBG string
	Accent string
}

type Theme struct {
	Name    string
	Palette Palette
}

func Dark() Theme {
	return Theme{
		Name: "dark",
		Palette: Palette{
			Background:      "#0f1117",
			Foreground:      "#e6edf3",
			GridLine:        "#2b303b",
			CellBaseBG:      "#151a21",
			CellFixedFG:     "#a0b3c5",
			CellFixedBG:     "#10141a",
			CellSelectedBG:  "#2d4f7c",
			CellDuplicateBG: "#4c3a15",
			CellConflictBG:  "#7c2d2d",
			Accent:          "#7aa2f7",
		},
	}
}

func Light() Theme {
	return Theme{
		Name: "light",
		Palette: Palette{
			Background:      "#ffffff",
			Foreground:      "#111111",
			GridLine:        "#cfcfcf",
			CellBaseBG:      "#f7f7f7",
			CellFixedFG:     "#6b7280",
			CellFixedBG:     "#efefef",
			CellSelectedBG:  "#cfe3ff",
			CellDuplicateBG: "#fff2cc",
			CellConflictBG:  "#ffd6d6",
			Accent:          "#2563eb",
		},
	}
}

func Neon() Theme {
	return Theme{
		Name: "neon",
		Palette: Palette{
			Background:      "#140019",
			Foreground:      "#f2f2f2",
			GridLine:        "#5d0073",
			CellBaseBG:      "#1a0021",
			CellFixedFG:     "#d9a6ff",
			CellFixedBG:     "#22002b",
			CellSelectedBG:  "#1f4b99",
			CellDuplicateBG: "#664400",
			CellConflictBG:  "#7a0044",
			Accent:          "#39ff14",
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
			CellFixedFG:     "#9fb7c6",
			CellFixedBG:     "#0e141a",
			CellSelectedBG:  "#063a3e",
			CellDuplicateBG: "#3d2d0a",
			CellConflictBG:  "#5b1515",
			Accent:          "#00e5ff",
		},
	}
}

func BaseStyle(t Theme) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(lipgloss.Color(t.Palette.Foreground)).Background(lipgloss.Color(t.Palette.Background))
}
