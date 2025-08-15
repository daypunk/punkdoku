package ui

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/key"
	"punkdoku/internal/config"
	"punkdoku/internal/game"
	"punkdoku/internal/generator"
	"punkdoku/internal/solver"
	"punkdoku/internal/theme"
)

type timerTickMsg struct{}

type flashDoneMsg struct{ Row, Col int }

type Model struct {
	keymap       KeyMap
	styles       UIStyles
	theme        theme.Theme

	board        game.Board
	solution     game.Grid
	cursorRow    int
	cursorCol    int
	autoCheck    bool
	timerEnabled bool
	startTime    time.Time
	elapsed      time.Duration
	completed    bool

	undoStack    []game.Move
	redoStack    []game.Move
	flashes      map[[2]int]time.Time
	showHelp     bool
}

func New(p generator.Grid, th theme.Theme, cfg config.Config) Model {
	b := game.NewBoardFromPuzzle(game.Grid(p))
	// Solve once for auto-check
	sg := b.Values
	if s := solveCopy(b.Values); s != nil {
		sg = *s
	}
	km := DefaultKeyMap()
	km.ApplyBindings(cfg.Bindings)
	m := Model{
		keymap:       km,
		styles:       BuildStyles(th),
		theme:        th,
		board:        b,
		solution:     sg,
		cursorRow:    0,
		cursorCol:    0,
		autoCheck:    cfg.AutoCheck,
		timerEnabled: cfg.TimerEnabled,
		startTime:    time.Now(),
		flashes:      map[[2]int]time.Time{},
	}
	return m
}

func solveCopy(g game.Grid) *game.Grid {
	var sg solver.Grid
	for r := 0; r < 9; r++ { for c := 0; c < 9; c++ { sg[r][c] = g[r][c] } }
	if solver.Solve(&sg, 2*time.Second) {
		var out game.Grid
		for r := 0; r < 9; r++ { for c := 0; c < 9; c++ { out[r][c] = sg[r][c] } }
		return &out
	}
	return nil
}

func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd
	if m.timerEnabled {
		cmds = append(cmds, tea.Tick(time.Second, func(time.Time) tea.Msg { return timerTickMsg{} }))
	}
	return tea.Batch(cmds...)
}

func (m Model) View() string { return Render(m) }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKey(msg)
	case timerTickMsg:
		if m.timerEnabled && !m.completed {
			m.elapsed = time.Since(m.startTime)
			return m, tea.Tick(time.Second, func(time.Time) tea.Msg { return timerTickMsg{} })
		}
		return m, nil
	case flashDoneMsg:
		delete(m.flashes, [2]int{msg.Row, msg.Col})
		return m, nil
	}
	return m, nil
}

func (m Model) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	k := msg
	if key.Matches(k, m.keymap.Help) {
		m.showHelp = !m.showHelp
		return m, nil
	}
	if key.Matches(k, m.keymap.ToggleAuto) {
		m.autoCheck = !m.autoCheck
		return m, nil
	}
	if key.Matches(k, m.keymap.ToggleTimer) {
		m.timerEnabled = !m.timerEnabled
		if m.timerEnabled && !m.completed {
			m.startTime = time.Now().Add(-m.elapsed)
			return m, tea.Tick(time.Second, func(time.Time) tea.Msg { return timerTickMsg{} })
		}
		return m, nil
	}
	if key.Matches(k, m.keymap.Undo) {
		m = m.applyUndo()
		return m, nil
	}
	if key.Matches(k, m.keymap.Redo) {
		m = m.applyRedo()
		return m, nil
	}
	s := k.String()
	switch s {
	case "up", "k":
		m.cursorRow = clamp(m.cursorRow-1, 0, 8)
	case "down", "j":
		m.cursorRow = clamp(m.cursorRow+1, 0, 8)
	case "left", "h":
		m.cursorCol = clamp(m.cursorCol-1, 0, 8)
	case "right", "l":
		m.cursorCol = clamp(m.cursorCol+1, 0, 8)
	case " ", "0":
		return m.applyInput(0)
	case "1", "2", "3", "4", "5", "6", "7", "8", "9":
		v := uint8(s[0]-'0')
		return m.applyInput(v)
	case "q", "esc", "ctrl+c":
		return m, tea.Quit
	}
	return m, nil
}

func (m Model) applyInput(v uint8) (tea.Model, tea.Cmd) {
	if m.board.IsGiven(m.cursorRow, m.cursorCol) {
		return m, nil
	}
	prev, ok := m.board.SetValue(m.cursorRow, m.cursorCol, v)
	if !ok { return m, nil }
	mv := game.Move{Row: m.cursorRow, Col: m.cursorCol, Prev: prev, Next: v, At: time.Now()}
	m.undoStack = append(m.undoStack, mv)
	m.redoStack = nil
	m.flashes[[2]int{m.cursorRow, m.cursorCol}] = time.Now().Add(120 * time.Millisecond)
	if isSolved(m.board.Values, m.solution) {
		m.completed = true
	}
	return m, tea.Tick(130*time.Millisecond, func(time.Time) tea.Msg { return flashDoneMsg{Row: mv.Row, Col: mv.Col} })
}

func (m Model) applyUndo() Model {
	if len(m.undoStack) == 0 { return m }
	last := m.undoStack[len(m.undoStack)-1]
	m.undoStack = m.undoStack[:len(m.undoStack)-1]
	m.board.Values[last.Row][last.Col] = last.Prev
	m.redoStack = append(m.redoStack, last)
	m.cursorRow, m.cursorCol = last.Row, last.Col
	m.completed = isSolved(m.board.Values, m.solution)
	return m
}

func (m Model) applyRedo() Model {
	if len(m.redoStack) == 0 { return m }
	last := m.redoStack[len(m.redoStack)-1]
	m.redoStack = m.redoStack[:len(m.redoStack)-1]
	m.board.Values[last.Row][last.Col] = last.Next
	m.undoStack = append(m.undoStack, last)
	m.cursorRow, m.cursorCol = last.Row, last.Col
	m.completed = isSolved(m.board.Values, m.solution)
	return m
}

func clamp(v, lo, hi int) int {
	if v < lo { return lo }
	if v > hi { return hi }
	return v
}

func (m Model) StatusLine() string {
	// Completed UI
	if m.completed {
		adaptiveColors := theme.NewAdaptiveColors(m.theme)
		gradientColors := adaptiveColors.GetGradientColors()
		completeGrad := gradientColors["complete"]
		var completeText string
		if m.timerEnabled {
			secs := int(m.elapsed.Truncate(time.Second).Seconds())
			mins := (secs / 60) % 100
			s := secs % 60
			timeStr := fmt.Sprintf("%02d:%02d", mins, s)
			completeText = fmt.Sprintf("✭ Clear %s ! Tap 'm' to quit ✭", timeStr)
		} else {
			completeText = "✭ Clear! Tap 'm' to quit ✭"
		}
		return gradientText(completeText, completeGrad[0], completeGrad[1])
	}
	// All filled but not solved → Try again
	if allFilled(m.board.Values) && !isSolved(m.board.Values, m.solution) {
		return m.styles.StatusError.Render("✭ Try again... ✭")
	}
	// Normal status (fixed width segments)
	var auto string
	if m.autoCheck {
		auto = m.styles.Status.Render("Auto: ") + m.styles.BoolTrue.Render("ON ") // Auto: 부분은 회색, ON 부분만 초록색
	} else {
		auto = m.styles.Status.Render("Auto: OFF")
	}
	
	var timerStr string
	if m.timerEnabled {
		secs := int(m.elapsed.Truncate(time.Second).Seconds())
		mins := (secs / 60) % 100
		s := secs % 60
		timeValue := fmt.Sprintf("%02d:%02d", mins, s)
		timerStr = m.styles.Status.Render("Timer: ") + m.styles.BoolTrue.Render(timeValue) // Timer: 부분은 회색, 시간만 초록색
	} else {
		timerStr = m.styles.Status.Render("Timer: --:--")
	}
	
	separator := m.styles.Status.Render(" | ")
	undoHint := m.styles.Status.Render("Undo: u")
	mainHint := m.styles.Status.Render("Main: m")
	
	return auto + separator + timerStr + separator + undoHint + separator + mainHint
}

func allFilled(g game.Grid) bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if g[r][c] == 0 { return false }
		}
	}
	return true
}

func isSolved(cur game.Grid, sol game.Grid) bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if cur[r][c] != sol[r][c] {
				return false
			}
		}
	}
	return true
}
