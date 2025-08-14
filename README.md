## punkdoku

### 개요
`punkdoku`는 macOS, Linux, Windows에서 동일하게 동작하는 터미널 스도쿠 게임입니다. Go로 작성되었고 Bubble Tea와 Lipgloss로 보기 좋은 TUI를 제공합니다. 퍼즐은 난이도별로 생성되며, Daily 모드는 UTC 날짜 기반 시드를 사용해 모든 사용자가 동일한 퍼즐을 받습니다. 입력 애니메이션, Undo, Auto‑Check, 타이머 등을 지원합니다.

### Overview
`punkdoku` is a cross‑platform terminal Sudoku for macOS, Linux, and Windows. It’s written in Go, built on Bubble Tea and Lipgloss for a clean, responsive TUI. Puzzles are generated per difficulty with a focus on uniqueness and reproducibility; Daily mode uses a UTC date‑based seed so everyone plays the same grid. The game ships with input flashes, undo/redo, auto‑check, and a compact timer.

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

### One‑line install (after publishing releases)
Replace `YOUR_ORG` with your GitHub org or username.

- macOS (arm64):
```bash
sudo curl -L https://github.com/YOUR_ORG/punkdoku/releases/latest/download/punkdoku-macos -o /usr/local/bin/punkdoku && sudo chmod +x /usr/local/bin/punkdoku && punkdoku
```

- Linux (x86_64):
```bash
sudo curl -L https://github.com/YOUR_ORG/punkdoku/releases/latest/download/punkdoku-linux -o /usr/local/bin/punkdoku && sudo chmod +x /usr/local/bin/punkdoku && punkdoku
```

- Windows (PowerShell):
```powershell
iwr -useb https://github.com/YOUR_ORG/punkdoku/releases/latest/download/punkdoku-windows.exe -OutFile $env:USERPROFILE\punkdoku.exe; & $env:USERPROFILE\punkdoku.exe
```

### Build from source
Prerequisites:
- Go 1.23+

Commands:
```bash
go mod tidy
CGO_ENABLED=0 go build -ldflags "-s -w" -o punkdoku ./cmd/punkdoku
./punkdoku
```

Run without building:
```bash
go run ./cmd/punkdoku
```

### Cross‑compiling examples
```bash
# macOS (arm64)
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o dist/punkdoku-macos ./cmd/punkdoku

# Linux (x86_64)
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o dist/punkdoku-linux ./cmd/punkdoku

# Windows (x86_64)
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o dist/punkdoku-windows.exe ./cmd/punkdoku
```

> Tip: Publish the three files in `dist/` as GitHub Release assets to enable the one‑line installs above.

## Usage

Launch `punkdoku` and pick a difficulty from the menu. The game view shows `{Difficulty} Mode`, the board, and a status line. With Auto‑Check enabled, conflicts are highlighted immediately. When solved, the status line changes to a clear message; if the timer is off, no time is shown.

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


