package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"punkdoku/internal/config"
	"punkdoku/internal/ui"
)

func main() {
	// Legacy flags kept but ignored when menu is used
	_ = flag.Bool("daily", false, "Generate daily puzzle")
	_ = flag.String("difficulty", "normal", "Difficulty: easy|normal|hard|nightmare")
	flag.Parse()

	cfg, _ := config.Load()
	app := ui.NewApp(cfg)
	if _, err := tea.NewProgram(app, tea.WithAltScreen()).Run(); err != nil {
		fmt.Fprintln(os.Stderr, "ui error:", err)
		os.Exit(1)
	}
}
