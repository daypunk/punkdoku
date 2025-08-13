package ui

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"punkdoku/internal/config"
	"punkdoku/internal/generator"
	"punkdoku/internal/theme"
)

type appState int

const (
	stateMenu appState = iota
	stateGame
)

type App struct {
	state         appState
	cfg           config.Config
	th            theme.Theme
	styles        UIStyles

	menuItems     []string
	selectedIdx   int
	autoCheck     bool
	timerEnabled  bool

	width         int
	height        int

	game          Model
}

func NewApp(cfg config.Config) App {
	th := theme.Punk()
	return App{
		state:        stateMenu,
		cfg:          cfg,
		th:           th,
		styles:       BuildStyles(th),
		menuItems:    []string{"Daily", "Easy", "Normal", "Hard", "Nightmare"},
		selectedIdx:  2,
		autoCheck:    cfg.AutoCheck,
		timerEnabled: cfg.TimerEnabled,
	}
}

func (a App) Init() tea.Cmd { return nil }

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch a.state {
	case stateMenu:
		switch m := msg.(type) {
		case tea.KeyMsg:
			s := m.String()
			switch s {
			case "up", "k":
				a.selectedIdx = clamp(a.selectedIdx-1, 0, len(a.menuItems)-1)
			case "down", "j":
				a.selectedIdx = clamp(a.selectedIdx+1, 0, len(a.menuItems)-1)
			case "a":
				a.autoCheck = !a.autoCheck
			case "t":
				a.timerEnabled = !a.timerEnabled
			case "enter":
				gm, cmd := a.startGame()
				if cmd != nil { return a, cmd }
				a.game = gm
				a.state = stateGame
				return a, nil
			case "q", "esc", "ctrl+c":
				return a, tea.Quit
			}
		case tea.WindowSizeMsg:
			a.width, a.height = m.Width, m.Height
		}
		return a, nil
	case stateGame:
		gm, cmd := a.game.Update(msg)
		if v, ok := gm.(Model); ok {
			a.game = v
		}
		return a, cmd
	}
	return a, nil
}

func (a App) View() string {
	switch a.state {
	case stateMenu:
		return a.viewMenu()
	case stateGame:
		return a.viewGame()
	}
	return ""
}

func (a *App) startGame() (Model, tea.Cmd) {
	var g generator.Grid
	var err error
	sel := a.menuItems[a.selectedIdx]
	switch sel {
	case "Daily":
		g, err = generator.GenerateDaily(time.Now())
	case "Easy":
		g, err = generator.Generate(generator.Easy, "")
	case "Normal":
		g, err = generator.Generate(generator.Normal, "")
	case "Hard":
		g, err = generator.Generate(generator.Hard, "")
	case "Nightmare":
		g, err = generator.Generate(generator.Nightmare, "")
	}
	if err != nil { return a.game, nil }
	cfg := a.cfg
	cfg.AutoCheck = a.autoCheck
	cfg.TimerEnabled = a.timerEnabled
	m := New(g, a.th, cfg)
	return m, nil
}

func (a App) viewMenu() string {
	banner := `                       __       __      __        
    ____  __  ______  / /______/ /___  / /____  __
   / __ \/ / / / __ \/ //_/ __  / __ \/ //_/ / / /
  / /_/ / /_/ / / / / ,< / /_/ / /_/ / ,< / /_/ / 
 / .___/\__,_/_/ /_/_/|_|\__,_/\____/_/|_|\__,_/  
/_/                          sudoku for punkers                  
`

	left := &strings.Builder{}
	left.WriteString(a.styles.Banner.Render(banner))
	left.WriteString("\n")
	left.WriteString(fmt.Sprintf("Auto-Check: %v  (a)\n", a.autoCheck))
	left.WriteString(fmt.Sprintf("Timer: %v  (t)\n\n", a.timerEnabled))
	left.WriteString("Select difficulty (↑/↓, Enter):\n")
	for i, it := range a.menuItems {
		item := a.styles.MenuItem.Render(it)
		if i == a.selectedIdx {
			item = a.styles.MenuItemSelected.Render(it)
			left.WriteString("> " + item + "\n")
		} else {
			left.WriteString("  " + item + "\n")
		}
	}
	left.WriteString("\n")
	left.WriteString(a.styles.Hint.Render("a Auto-Check  t Timer  Enter Start  q Quit"))

	panel := a.styles.Panel.Render(left.String())
	if a.width > 0 && a.height > 0 {
		return a.styles.App.Render(lipgloss.Place(a.width, a.height, lipgloss.Center, lipgloss.Center, panel))
	}
	return a.styles.App.Render(panel)
}

func (a App) viewGame() string {
	content := a.game.View()
	innerWidth := lipgloss.Width(content) + 4
	centered := lipgloss.PlaceHorizontal(innerWidth, lipgloss.Center, content)
	panel := a.styles.Panel.Render(centered)
	if a.width > 0 && a.height > 0 {
		return a.styles.App.Render(lipgloss.Place(a.width, a.height, lipgloss.Center, lipgloss.Center, panel))
	}
	return a.styles.App.Render(panel)
}
