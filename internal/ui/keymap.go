package ui

import (
	"strings"
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	Up, Down, Left, Right key.Binding
	Undo, Redo            key.Binding
	ToggleAuto            key.Binding
	ToggleTimer           key.Binding
	Help                  key.Binding
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		Up:          key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("↑/k", "Up/위로")),
		Down:        key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("↓/j", "Down/아래로")),
		Left:        key.NewBinding(key.WithKeys("left", "h"), key.WithHelp("←/h", "Left/왼쪽")),
		Right:       key.NewBinding(key.WithKeys("right", "l"), key.WithHelp("→/l", "Right/오른쪽")),
		Undo:        key.NewBinding(key.WithKeys("ctrl+z", "u"), key.WithHelp("Ctrl+Z/u", "Undo/되돌리기")),
		Redo:        key.NewBinding(key.WithKeys("ctrl+y", "ctrl+r"), key.WithHelp("Ctrl+Y/R", "Redo/다시하기")),
		ToggleAuto:  key.NewBinding(key.WithKeys("a"), key.WithHelp("a", "Auto-Check/자동 체크")),
		ToggleTimer: key.NewBinding(key.WithKeys("t"), key.WithHelp("t", "Timer/타이머")),
		Help:        key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "Help/도움말")),
	}
}

func (km *KeyMap) ApplyBindings(bindings map[string][]string) {
	if bindings == nil { return }
	set := func(b *key.Binding, keys []string, help string) {
		if len(keys) == 0 { return }
		*b = key.NewBinding(key.WithKeys(keys...), key.WithHelp(strings.Join(keys, "/"), help))
	}
	if v, ok := bindings["up"]; ok { set(&km.Up, v, "Up/위로") }
	if v, ok := bindings["down"]; ok { set(&km.Down, v, "Down/아래로") }
	if v, ok := bindings["left"]; ok { set(&km.Left, v, "Left/왼쪽") }
	if v, ok := bindings["right"]; ok { set(&km.Right, v, "Right/오른쪽") }
	if v, ok := bindings["undo"]; ok { set(&km.Undo, v, "Undo/되돌리기") }
	if v, ok := bindings["redo"]; ok { set(&km.Redo, v, "Redo/다시하기") }
	if v, ok := bindings["auto"]; ok { set(&km.ToggleAuto, v, "Auto-Check/자동 체크") }
	if v, ok := bindings["timer"]; ok { set(&km.ToggleTimer, v, "Timer/타이머") }
	if v, ok := bindings["help"]; ok { set(&km.Help, v, "Help/도움말") }
}
