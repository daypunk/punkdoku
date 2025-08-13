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
		if m.timerEnabled {
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
		if m.timerEnabled {
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
	return m, tea.Tick(130*time.Millisecond, func(time.Time) tea.Msg { return flashDoneMsg{Row: mv.Row, Col: mv.Col} })
}

func (m Model) applyUndo() Model {
	if len(m.undoStack) == 0 { return m }
	last := m.undoStack[len(m.undoStack)-1]
	m.undoStack = m.undoStack[:len(m.undoStack)-1]
	m.board.Values[last.Row][last.Col] = last.Prev
	m.redoStack = append(m.redoStack, last)
	m.cursorRow, m.cursorCol = last.Row, last.Col
	return m
}

func (m Model) applyRedo() Model {
	if len(m.redoStack) == 0 { return m }
	last := m.redoStack[len(m.redoStack)-1]
	m.redoStack = m.redoStack[:len(m.redoStack)-1]
	m.board.Values[last.Row][last.Col] = last.Next
	m.undoStack = append(m.undoStack, last)
	m.cursorRow, m.cursorCol = last.Row, last.Col
	return m
}

func clamp(v, lo, hi int) int {
	if v < lo { return lo }
	if v > hi { return hi }
	return v
}

func (m Model) StatusLine() string {
	auto := "Auto:Off"
	if m.autoCheck { auto = "Auto:On" }
	tm := "Timer:Off"
	if m.timerEnabled { tm = fmt.Sprintf("Timer:%s", m.elapsed.Truncate(time.Second)) }
	return m.styles.Status.Render(fmt.Sprintf("%s | %s | Pos %d,%d", auto, tm, m.cursorRow+1, m.cursorCol+1))
}
