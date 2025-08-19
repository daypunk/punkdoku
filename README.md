## punkdoku

<div align="center">
  <img src="https://github.com/user-attachments/assets/8bb7cf23-51b1-4db6-86f8-ef8c96bf6661" width="600" alt="punkdoku">
</div>

<div align="center">
  <a href="https://github.com/daypunk/punkdoku/releases/latest">
    <img src="https://img.shields.io/github/v/release/daypunk/punkdoku?style=flat&logo=github&color=ff6b6b" alt="Latest Release"/>
  </a>
  <a href="https://github.com/daypunk/punkdoku/releases">
    <img src="https://img.shields.io/github/downloads/daypunk/punkdoku/total?style=flat&logo=download&color=4ecdc4" alt="Downloads"/>
  </a>
  <a href="https://golang.org/">
    <img src="https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go&logoColor=white" alt="Go Version"/>
  </a>
  <a href="https://opensource.org/licenses/MIT">
    <img src="https://img.shields.io/badge/License-MIT-yellow?style=flat" alt="License"/>
  </a>
</div>
<p align="center">
  <img src="https://img.shields.io/badge/Platform-macOS%20%7C%20Linux-brightgreen?style=flat&logo=terminal" alt="Platforms"/>
  <img src="https://img.shields.io/badge/TUI-Bubble%20Tea-ff79c6?style=flat" alt="TUI Framework"/>
  <img src="https://img.shields.io/badge/Binary%20Size-4.8MB-blue?style=flat" alt="Binary Size"/>
</p><br>

<div align="center">
  <img width="600" height="146" alt="Image" src="https://github.com/user-attachments/assets/304f1911-3ec1-4311-9892-e7c5fb0d6f89" />
</div>


<p align="center">
  <img src="https://img.shields.io/badge/🧩-Sudoku-purple?style=flat" alt="Sudoku"/>
  <img src="https://img.shields.io/badge/📅-Daily%20Seed%20Puzzles-purple?style=flat" alt="Daily Seed Puzzles"/>
</p>

### 개요
`punkdoku`는 macOS, Linux에서 동일하게 동작하는 터미널 스도쿠 게임입니다. Go로 작성되었고 Bubble Tea와 Lipgloss로 보기 좋은 TUI를 제공합니다. 퍼즐은 난이도별로 생성되며, 4개의 난이도 모드는 나노초 기반 시드를 사용해 퍼즐을 생성하고, Daily 모드는 UTC 날짜 기반 시드를 사용해 모든 사용자가 동일한 퍼즐을 받습니다. 입력 애니메이션, Undo, Auto‑Check, 타이머 등을 지원합니다.

### Overview
`punkdoku` is a cross‑platform terminal Sudoku for macOS and Linux. It's written in Go, built on Bubble Tea and Lipgloss for a clean, responsive TUI. Puzzles are generated per difficulty with a focus on uniqueness and reproducibility; Daily mode uses a UTC date‑based seed so everyone plays the same grid. The game ships with input flashes, undo/redo, auto‑check, and a compact timer.

## Installation

punkdoku is a **terminal application**. After installation, run it from the terminal using the `punkdoku` command.

### 🍺 Homebrew (macOS/Linux - Recommended)
The easiest installation method.

```bash
# Install Homebrew first, if you don't have it!
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

# Install punkdoku
brew install daypunk/tap/punkdoku
```

### 📥 Direct Download (All Platforms)

#### macOS
```bash
# Using curl (recommended)
curl -L https://github.com/daypunk/punkdoku/releases/latest/download/punkdoku-macos -o /usr/local/bin/punkdoku
chmod +x /usr/local/bin/punkdoku

# If curl is not available, download manually from browser
# 1. Download punkdoku-macos from https://github.com/daypunk/punkdoku/releases/latest
# 2. Run these commands in terminal:
# mv ~/Downloads/punkdoku-macos /usr/local/bin/punkdoku
# chmod +x /usr/local/bin/punkdoku
```

#### Linux
```bash
# Using curl (recommended)
curl -L https://github.com/daypunk/punkdoku/releases/latest/download/punkdoku-linux -o /usr/local/bin/punkdoku
chmod +x /usr/local/bin/punkdoku

# If curl is not available, use wget
wget -O /usr/local/bin/punkdoku https://github.com/daypunk/punkdoku/releases/latest/download/punkdoku-linux
chmod +x /usr/local/bin/punkdoku

# If neither is available, download manually from browser
# 1. Download punkdoku-linux from https://github.com/daypunk/punkdoku/releases/latest
# 2. Run these commands in terminal:
# mv ~/Downloads/punkdoku-linux /usr/local/bin/punkdoku
# chmod +x /usr/local/bin/punkdoku
```

### 🚀 Running the Game
After installation, run the game from the **terminal** using:

```bash
punkdoku
```

## How to Play

1. **Launch**: Run `punkdoku` in your terminal
2. **Navigate**: Use arrow keys to move around
3. **Input**: Press `1-9` to place numbers, `0` to clear
4. **Settings**: Press `a` for auto-check, `t` for timer

### Difficulty Levels
- **Easy**: Perfect for beginners
- **Normal**: Balanced challenge
- **Hard**: Requires strategy
- **Lunatic**: Expert level
- **Daily**: Everyone gets the same puzzle based on the date

### Controls
- **Arrow keys**: Navigate the grid
- **1-9**: Place numbers
- **0/Space**: Clear cell
- **u**: Undo last move
- **a**: Toggle auto-check (highlights mistakes)
- **t**: Toggle timer
- **m**: Return to main menu
- **q**: Quit

## Features

- **Clean interface**: Minimalist design that stays out of your way
- **Daily puzzles**: Same puzzle for everyone, changes each day
- **Smart generation**: Every puzzle has exactly one solution
- **Undo/Redo**: Full move history
- **Auto-check**: Optional real-time error highlighting
- **Timer**: Track your solving speed

## Technical Details

The game generates puzzles by starting with a complete solution and carefully removing numbers while ensuring uniqueness. Daily puzzles use the UTC date as a seed, so everyone worldwide gets the same puzzle.

- Generator uses backtracking with randomization
- Solver verifies puzzle uniqueness
- UI built with Bubble Tea and Lipgloss
- Single binary, no dependencies

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

```bash
# Run locally
go run ./cmd/punkdoku

# Build
go build -o punkdoku ./cmd/punkdoku
```

Built for modern terminals with Unicode support. Works best with monospace fonts and true color support.

This project is licensed under the MIT License. Please credit the original source when reusing.
