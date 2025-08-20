# punkdoku

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
  <img width="600" alt="Image" src="https://github.com/user-attachments/assets/304f1911-3ec1-4311-9892-e7c5fb0d6f89" />
</div>

<p align="center">
  <img src="https://img.shields.io/badge/🧩-Sudoku-purple?style=flat" alt="Sudoku"/>
  <img src="https://img.shields.io/badge/📅-Daily%20Seed%20Puzzles-purple?style=flat" alt="Daily Seed Puzzles"/>
</p>

### 개요
`punkdoku`는 macOS, Linux에서 동일하게 동작하는 터미널 스도쿠 게임입니다. Go로 작성되었고 Bubble Tea와 Lipgloss로 보기 좋은 TUI를 제공합니다. 퍼즐은 난이도별로 생성되며, 4개의 난이도 모드는 나노초 기반 시드를 사용해 퍼즐을 생성하고, Daily 모드는 UTC 날짜 기반 시드를 사용해 모든 사용자가 동일한 퍼즐을 받습니다. 입력 애니메이션, Undo, Auto‑Check, 타이머 등을 지원합니다.

### Overview
`punkdoku` is a cross‑platform terminal Sudoku for macOS and Linux. It's written in Go, built on Bubble Tea and Lipgloss for a clean, responsive TUI. Puzzles are generated per difficulty with a focus on uniqueness and reproducibility; Daily mode uses a UTC date‑based seed so everyone plays the same grid. The game ships with input flashes, undo/redo, auto‑check, and a compact timer.

## Quick Start

### Option 1: Homebrew (Recommended)
```bash
brew install daypunk/tap/punkdoku
```

### Option 2: Manual Download

#### macOS
```bash
# Download the binary (choose based on your Mac's processor)
# For Intel Macs:
curl -L -o punkdoku https://github.com/daypunk/punkdoku/releases/latest/download/punkdoku-macos-amd64

# For Apple Silicon Macs:
curl -L -o punkdoku https://github.com/daypunk/punkdoku/releases/latest/download/punkdoku-macos-arm64

# Make it executable
chmod +x punkdoku

# Move to your PATH (optional)
sudo mv punkdoku /usr/local/bin/
```

#### Linux
```bash
# Download the binary
curl -L -o punkdoku https://github.com/daypunk/punkdoku/releases/latest/download/punkdoku-linux

# Make it executable
chmod +x punkdoku

# Move to your PATH (optional)
sudo mv punkdoku /usr/local/bin/
```

### 🚀 Running the Game
After installation, run the game from the **terminal** using:

```bash
punkdoku
```

## How to Play

Run `punkdoku` in your terminal and use:
- **Arrow keys** to navigate
- **1-9** to place numbers
- **0** or **Space** to clear cells
- **u** to undo
- **a** to toggle auto-check
- **t** to toggle timer
- **m** to return to menu
- **q** to quit

## Game Modes

- **Easy** - Good for beginners
- **Normal** - Balanced challenge
- **Hard** - Requires strategy
- **Lunatic** - Expert level
- **Daily(=Normal)** - Same puzzle for everyone, changes daily

## Features

- Clean, minimalist interface
- Daily puzzles with shared seeds
- Smart puzzle generation (unique solutions only)
- Undo/redo functionality
- Real-time error checking
- Built-in timer
- No external dependencies

## Development

```bash
# Run locally
go run ./cmd/punkdoku

# Build binary
go build -o punkdoku ./cmd/punkdoku
```

Requires Go 1.23+ and works best with terminals that support Unicode and true color.

## License

MIT License - feel free to use and modify as needed.
