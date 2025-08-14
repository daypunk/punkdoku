## punkdoku

### 개요
`punkdoku`는 macOS, Linux, Windows에서 동일하게 동작하는 터미널 스도쿠 게임입니다. Go로 작성되었고 Bubble Tea와 Lipgloss로 보기 좋은 TUI를 제공합니다. 퍼즐은 난이도별로 생성되며, Daily 모드는 UTC 날짜 기반 시드를 사용해 모든 사용자가 동일한 퍼즐을 받습니다. 입력 애니메이션, Undo, Auto‑Check, 타이머 등을 지원합니다.

### Overview
`punkdoku` is a cross‑platform terminal Sudoku for macOS, Linux, and Windows. It's written in Go, built on Bubble Tea and Lipgloss for a clean, responsive TUI. Puzzles are generated per difficulty with a focus on uniqueness and reproducibility; Daily mode uses a UTC date‑based seed so everyone plays the same grid. The game ships with input flashes, undo/redo, auto‑check, and a compact timer.

## Features
- **Cross‑platform TUI**: Runs the same on macOS, Linux, and Windows terminals.
- **Difficulty presets**: Easy, Normal, Hard, Nightmare, and Daily.
- **Reproducible Daily**: UTC `YYYY‑MM‑DD` seeded generation.
- **Single‑solution focus**: Generator removes cells with a uniqueness check.
- **Auto‑Check toggle** and **Undo/Redo**.
- **Compact timer** with fixed‑width `MM:SS` to avoid UI jitter.
- **Modern look**: Box‑drawing grid, rounded frame, cyan accent; block separators match difficulty color (gradients for Hard/Nightmare).
- **Single binary**: distribute as a single executable per OS.

## Installation

### One‑line install via Homebrew (Recommended)
```bash
brew install daypunk/tap/punkdoku
```

### One‑line install via curl (Direct download)
```bash
# macOS (arm64/amd64)
curl -L https://github.com/daypunk/punkdoku/releases/latest/download/punkdoku-macos -o /usr/local/bin/punkdoku && chmod +x /usr/local/bin/punkdoku

# Linux (x86_64)
curl -L https://github.com/daypunk/punkdoku/releases/latest/download/punkdoku-linux -o /usr/local/bin/punkdoku && chmod +x /usr/local/bin/punkdoku

# Windows (PowerShell)
Invoke-WebRequest -Uri "https://github.com/daypunk/punkdoku/releases/latest/download/punkdoku-windows.exe" -OutFile "punkdoku.exe"
```

### Run
```bash
punkdoku
```

## Usage

### Getting Started
1. **Launch the game**: Run `punkdoku` from your terminal
2. **Select difficulty**: Use arrow keys or `h j k l` to navigate, `Enter` to confirm
3. **Start playing**: Navigate with arrow keys, input numbers 1-9, clear with 0 or Space

### Main Menu
- **Easy**: ~38 cells removed, suitable for beginners
- **Normal**: ~46 cells removed, balanced challenge
- **Hard**: ~52 cells removed, requires strategy
- **Nightmare**: ~58 cells removed, expert level
- **Daily**: Same puzzle for everyone based on UTC date

### Game Controls
- **Navigation**: Arrow keys or `h j k l` (Vim-style)
- **Input**: `1-9` to set a number, `0` or `Space` to clear
- **Undo**: `u` or `Ctrl+Z`
- **Redo**: `Ctrl+Y` or `Ctrl+R`
- **Toggle Auto-Check**: `a` (highlights conflicts immediately)
- **Toggle Timer**: `t` (shows completion time)
- **Main Menu**: `m` (return to difficulty selection)
- **Quit**: `q`, `Esc`, or `Ctrl+C`

### Game Features
- **Auto-Check**: When enabled, shows duplicate numbers and conflicts in real-time
- **Timer**: Tracks completion time (displayed as MM:SS)
- **Undo/Redo**: Full move history with keyboard shortcuts
- **Visual Feedback**: Selected cell highlighting, input animations, color-coded difficulty themes

### Game States
- **Playing**: Normal gameplay with status line showing Auto-Check, Timer, and controls
- **Completed**: Success message with completion time (if timer enabled)
- **Try Again**: Appears when grid is filled but incorrect

### Tips
- Use Auto-Check to catch mistakes early
- Daily puzzles are the same for everyone on the same date
- The timer stops automatically when you complete the puzzle
- You can return to the main menu anytime with `m`

## Configuration

User settings are stored at `~/.punkdoku/config.yaml` (auto‑check and timer). Defaults are sensible; you can run without any config file.

## Technical Notes

### Generator
- Pipeline:
  1) Build a full valid solution with randomized DFS/backtracking.
  2) Carve cells according to difficulty while trying to keep a single solution.
- Difficulty → parameters:
  - Easy: remove ~38 cells, ~150ms timeout
  - Normal: remove ~46 cells, ~150ms timeout
  - Hard: remove ~52 cells, ~200ms timeout
  - Nightmare: remove ~58 cells, ~250ms timeout
- Uniqueness: after each removal, check with the solver (`CountSolutions`) and skip removals that break uniqueness (stop after 2 solutions for speed).
- Seeding: Daily mode uses `UTC YYYY‑MM‑DD`, hashed via 64‑bit FNV‑1a to seed `math/rand` for reproducibility.
- Performance: bounded by timeouts to avoid pathological search; puzzles aim to generate within a few hundred milliseconds.

### Solver
- Classic backtracking with early exits.
- `CountSolutions` is used during carving to test uniqueness (stop after 2 solutions).

### UI Architecture
- Bubble Tea model: `Model` manages board state, timer, auto‑check, input flashes, and history.
- Lipgloss styles: centralized in `internal/ui/styles.go` and `internal/theme/theme.go` for consistent colors.
- The grid uses Unicode box‑drawing characters with rounded outer corners. 3×3 separators share the difficulty color; Hard and Nightmare use a gradient.
- The timer and status fields are formatted with fixed widths to avoid horizontal jitter; the game panel is centered.

### Layout Stability
- Main menu and in‑game panels are vertically spaced for readability.
- The board and status line are rendered with explicit, stable newlines to maintain a consistent 1‑1‑2‑1 layout in game.

### Terminals
- A modern, true‑color, monospace terminal is recommended (e.g., Windows Terminal, iTerm2, GNOME Terminal).
- If box‑drawing characters appear misaligned, verify the terminal font is truly monospace and ligatures are disabled.

## Project Structure
```text
cmd/
  punkdoku/
    main.go              # entrypoint
internal/
  config/               # YAML load/save under ~/.punkdoku
  game/                 # board state, moves, duplicates/conflicts
  generator/            # difficulty params, daily seeding, carving
  solver/               # backtracking solver + uniqueness counting
  theme/                # color palette (Punk theme)
  ui/                   # Bubble Tea app, model, view, styles, keymap
```

## Development
- Run: `go run ./cmd/punkdoku`
- Build: `CGO_ENABLED=0 go build -ldflags "-s -w" -o punkdoku ./cmd/punkdoku`
- Lint/build with your usual Go toolchain. The code favors clarity and explicit naming.


