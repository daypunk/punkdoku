// git tag v1.0.1
// git push origin v1.0.1

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

	currentDiff   string
	game          Model
}

func NewApp(cfg config.Config) App {
	th := theme.DetectTheme()
	return App{
		state:        stateMenu,
		cfg:          cfg,
		th:           th,
		styles:       BuildStyles(th),
		menuItems:    []string{"Easy", "Normal", "Hard", "Lunatic", "Daily"},
		selectedIdx:  1,
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
			case "left", "h":
				a.selectedIdx = clamp(a.selectedIdx-1, 0, len(a.menuItems)-1)
			case "right", "l":
				a.selectedIdx = clamp(a.selectedIdx+1, 0, len(a.menuItems)-1)
			case "a":
				a.autoCheck = !a.autoCheck
			case "t":
				a.timerEnabled = !a.timerEnabled
			case "enter":
				gm, cmd := a.startGame()
				a.game = gm
				a.state = stateGame
				return a, cmd
			case "q", "esc", "ctrl+c":
				return a, tea.Quit
			}
		case tea.WindowSizeMsg:
			a.width, a.height = m.Width, m.Height
		}
		return a, nil
	case stateGame:
		// intercept main menu key
		if kmsg, isKey := msg.(tea.KeyMsg); isKey {
			if kmsg.String() == "m" {
				a.state = stateMenu
				return a, nil
			}
		}
		gm, cmd := a.game.Update(msg)
		if v, ok := gm.(Model); ok { a.game = v }
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
	case "Lunatic":
		g, err = generator.Generate(generator.Lunatic, "")
	}
	if err != nil { return a.game, nil }
	cfg := a.cfg
	cfg.AutoCheck = a.autoCheck
	cfg.TimerEnabled = a.timerEnabled
	a.currentDiff = sel
	m := New(g, a.th, cfg)
	// 적응형 색상 사용
	adaptiveColors := theme.NewAdaptiveColors(a.th)
	diffColors := adaptiveColors.GetDifficultyColors()
	hex := diffColors[sel]
	if hex == "" {
		hex = a.th.Palette.Accent
	}
	style := lipgloss.NewStyle().Foreground(lipgloss.Color(hex))
	m.styles.RowSep = style
	m.styles.ColSep = style
	// Fixed 숫자도 구분선과 동일한 색상 사용
	m.styles.CellFixed = m.styles.CellFixed.Foreground(lipgloss.Color(hex))
	return m, m.Init()
}

func (a App) viewMenu() string {
	banner := `                       __       __      __        
    ____  __  ______  / /______/ /___  / /____  __
   / __ \/ / / / __ \/ // / __  / __ \/ // / / / /
  / /_/ / /_/ / / / /   </ /_/ / /_/ /   </ /_/ / 
 /  ___/\____/_/ /_/_/|_|\____/\____/_/|_|\____/ 
/_/                            sudoku for punks
`

	// Options
	optAC := fmt.Sprintf("Auto-Check (a): %s", boolText(a.styles, a.autoCheck))
	optTM := fmt.Sprintf("Timer (t): %s", boolText(a.styles, a.timerEnabled))

	// Adaptive colors
	adaptiveColors := theme.NewAdaptiveColors(a.th)
	accentColors := adaptiveColors.GetAccentColors()
	
	// Difficulty list (horizontal), Daily last
	diffs := []string{"Easy", "Normal", "Hard", "Lunatic", "Daily"}
	var items []string
	selectedStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(accentColors["selected"])).Bold(true)
	for i, name := range diffs {
		prefix := "  "
		if i == a.selectedIdx { prefix = "✭ " }
		label := prefix + name
		if i == a.selectedIdx {
			items = append(items, selectedStyle.Render(label))
		} else {
			items = append(items, a.styles.MenuItem.Render(label))
		}
	}
	gap := strings.Repeat(" ", 4)
	diffRow := strings.Join(items, gap)

	// Adaptive gradient colors
	gradientColors := adaptiveColors.GetGradientColors()
	bannerGrad := gradientColors["banner"]
	leftHex := bannerGrad[0]
	rightHex := bannerGrad[1]
	
	title := gradientText("Select difficulty", leftHex, rightHex)
	box := renderGradientBox(diffRow, 2, leftHex, rightHex)  // 4에서 2로 줄임

	// Gradient banner (line by line)
	var gb strings.Builder
	for i, l := range strings.Split(strings.TrimRight(banner, "\n"), "\n") {
		gb.WriteString(gradientText(l, leftHex, rightHex))
		if i <  len(strings.Split(strings.TrimRight(banner, "\n"), "\n"))-1 { gb.WriteString("\n") }
	}
	gradientBanner := gb.String()

	// Compose content with explicit 2-line top/bottom padding
	content := "\n\n" + gradientBanner + "\n\n\n" + optAC + "\n" + optTM + "\n\n\n" + title + "\n" + box + "\n\n"
	panel := a.styles.Panel.Render(content)
	if a.width > 0 && a.height > 0 {
		return a.styles.App.Render(lipgloss.Place(a.width, a.height, lipgloss.Center, lipgloss.Center, panel))
	}
	return a.styles.App.Render(panel)
}

func boolText(s UIStyles, v bool) string {
	if v { return s.BoolTrue.Render("ON") }
	return s.BoolFalse.Render("OFF")
}

func (a App) viewGame() string {
	// 고정 폭 사용 (스도쿠 보드는 항상 동일한 크기)
	innerWidth := 58  // 메인화면과 너비 맞춤
	
	boardAndStatus := Render(a.game)

	label := a.currentDiff
	if a.currentDiff == "Daily" { label = "Daily Seed" }
	headerText := label + " Mode"
	// Adaptive colors for headers
	adaptiveColors := theme.NewAdaptiveColors(a.th)
	gradientColors := adaptiveColors.GetGradientColors()
	
	var header string
	switch a.currentDiff {
	case "Easy":
		easyGrad := gradientColors["easy"]
		header = gradientText(headerText, easyGrad[0], easyGrad[1])
	case "Normal":
		normalGrad := gradientColors["normal"]
		header = gradientText(headerText, normalGrad[0], normalGrad[1])
	case "Hard":
		hardGrad := gradientColors["hard"]
		header = gradientText(headerText, hardGrad[0], hardGrad[1])
	case "Lunatic":
		lunaticGrad := gradientColors["lunatic"]
		header = gradientText(headerText, lunaticGrad[0], lunaticGrad[1])
	case "Daily":
		dailyGrad := gradientColors["daily"]
		header = gradientText(headerText, dailyGrad[0], dailyGrad[1])
	default:
		header = lipgloss.NewStyle().Foreground(lipgloss.Color(a.th.Palette.Accent)).Bold(true).Render(headerText)
	}

	headerCentered := lipgloss.PlaceHorizontal(innerWidth, lipgloss.Center, header)
	centered := lipgloss.PlaceHorizontal(innerWidth, lipgloss.Center, boardAndStatus)
	// 간격: 상단 1줄 + 헤더 + 1줄(빈 줄 보이도록 개행 2개) + 보드(내부 보드-상태 2줄) + 하단 1줄
	body := "\n" + headerCentered + "\n\n" + centered + "\n"
	panel := a.styles.Panel.Render(body)
	if a.width > 0 && a.height > 0 {
		return a.styles.App.Render(lipgloss.Place(a.width, a.height, lipgloss.Center, lipgloss.Center, panel))
	}
	return a.styles.App.Render(panel)
}

// Helpers: gradient text and gradient bordered box
func renderGradientBox(content string, padX int, leftHex, rightHex string) string {
	w := lipgloss.Width(content) + padX*2
	top := lipgloss.NewStyle().Foreground(lipgloss.Color(leftHex)).Render("╭") + gradientLine("─", w, leftHex, rightHex) + lipgloss.NewStyle().Foreground(lipgloss.Color(rightHex)).Render("╮")
	bottom := lipgloss.NewStyle().Foreground(lipgloss.Color(leftHex)).Render("╰") + gradientLine("─", w, leftHex, rightHex) + lipgloss.NewStyle().Foreground(lipgloss.Color(rightHex)).Render("╯")
	left := lipgloss.NewStyle().Foreground(lipgloss.Color(leftHex)).Render("│")
	right := lipgloss.NewStyle().Foreground(lipgloss.Color(rightHex)).Render("│")
	middle := left + strings.Repeat(" ", padX) + content + strings.Repeat(" ", padX) + right
	return strings.Join([]string{top, middle, bottom}, "\n")
}

func gradientLine(ch string, width int, fromHex, toHex string) string {
	colors := gradientColors(fromHex, toHex, width)
	var b strings.Builder
	for i := 0; i < width; i++ {
		b.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color(colors[i])).Render(ch))
	}
	return b.String()
}

func gradientText(text, leftHex, rightHex string) string {
	colors := gradientColors(leftHex, rightHex, len(text))
	var b strings.Builder
	idx := 0
	for _, ch := range text { // rune-safe
		b.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color(colors[idx])).Bold(true).Render(string(ch)))
		idx++
	}
	return b.String()
}

func gradientColors(fromHex, toHex string, steps int) []string {
	r1, g1, b1 := hexToRGB(fromHex)
	r2, g2, b2 := hexToRGB(toHex)
	out := make([]string, steps)
	for i := 0; i < steps; i++ {
		if steps == 1 { out[i] = fromHex; continue }
		t := float64(i) / float64(steps-1)
		r := int(float64(r1) + (float64(r2)-float64(r1))*t)
		g := int(float64(g1) + (float64(g2)-float64(g1))*t)
		b := int(float64(b1) + (float64(b2)-float64(b1))*t)
		out[i] = fmt.Sprintf("#%02x%02x%02x", r, g, b)
	}
	return out
}

func hexToRGB(hex string) (int, int, int) {
	h := strings.TrimPrefix(hex, "#")
	if len(h) != 6 { return 255, 255, 255 }
	var r, g, b int
	fmt.Sscanf(h, "%02x%02x%02x", &r, &g, &b)
	return r, g, b
}

