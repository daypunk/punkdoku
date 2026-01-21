<h1 align="center">Punkdoku</h1>

<div align="center">
  <img src="https://github.com/user-attachments/assets/8bb7cf23-51b1-4db6-86f8-ef8c96bf6661" width="600" alt="punkdoku">
</div>

<div align="center">
  <!-- Repo Stats -->
  <a href="https://github.com/daypunk/punkdoku/stargazers"><img src="https://img.shields.io/github/stars/daypunk/punkdoku?style=plastic&logo=apachespark&label=Stars&logoColor=FFD700&labelColor=222&color=FFD700" alt="GitHub Stars"/></a><a href="https://github.com/daypunk/punkdoku/releases/latest"><img src="https://img.shields.io/github/v/release/daypunk/punkdoku?style=plastic&logo=github&logoColor=white&label=Release&labelColor=222&color=ccc" alt="Latest Release"/></a><a href="https://github.com/daypunk/punkdoku/releases"><img src="https://img.shields.io/github/downloads/daypunk/punkdoku/total?style=plastic&logo=github&logoColor=white&label=Downloads&labelColor=222&color=ccc" alt="Downloads"/></a>
</div>
<br>

<div align="center">
  <!-- Tech Stack -->
  <img src="https://img.shields.io/badge/Go-1.23+-66E3FF?style=plastic&logo=go&labelColor=FFF" alt="Go Version"/>
  <img src="https://img.shields.io/badge/TUI-Bubble%20Tea-FFB3C7?style=plastic&logo=ntfy&labelColor=FFF&logoColor=FF79C6" alt="TUI Framework"/>
</div>


<div align="center">
  <img src="https://img.shields.io/badge/macOS-ccc?style=plastic&logo=apple&labelColor=fff&logoColor=black" alt="macOS"/>
  <img src="https://img.shields.io/badge/Linux-ccc?style=plastic&logo=linux&labelColor=fff&logoColor=black" alt="Linux"/>
  <img src="https://img.shields.io/badge/Binary-4.8MB-ccc?style=plastic&labelColor=fff" alt="Binary Size"/>
  <a href="https://opensource.org/licenses/MIT">
    <img src="https://img.shields.io/badge/License-MIT-ccc?style=plastic&&labelColor=fff&logoColor=white" alt="License"/>
  </a>
</div>
<br>

<div align="center">
  <img width="600" alt="Image" src="https://github.com/user-attachments/assets/304f1911-3ec1-4311-9892-e7c5fb0d6f89" />
</div>

<p align="center">
  <img src="https://img.shields.io/badge/🧩-Sudoku-ccc?style=plastic&labelColor=fff" alt="Sudoku"/>
  <img src="https://img.shields.io/badge/📅-Daily%20Seed%20Puzzles-ccc?style=plastic&labelColor=fff" alt="Daily Seed Puzzles"/>
</p>

### 개요
**`punkdoku`** 는 macOS, Linux에서 즐기는 **터미널 스도쿠 게임**입니다. Go로 작성되었고 Bubble Tea와 Lipgloss로 보기 좋은 TUI를 제공합니다. 퍼즐은 난이도별로 생성되며, 4개의 난이도 모드는 밀리초 기반 시드를 사용해 퍼즐을 생성하고, Daily 모드는 UTC 날짜 기반 시드를 사용해 모든 사용자가 동일한 퍼즐을 받습니다. 입력 애니메이션, Undo, Auto‑Check, 타이머 등을 지원합니다.

### Overview
**`punkdoku`** is a **terminal Sudoku game** that runs identically on macOS and Linux. It is written in Go and provides a visually appealing TUI built with Bubble Tea and Lipgloss. Puzzles are generated per difficulty level: the four difficulty modes use a millisecond-based seed to create unique puzzles, while Daily mode uses a UTC date-based seed so that all players receive the same puzzle. The game supports input animations, undo, auto-check, and a timer.

## Quick Start

### Option 1: 🍺 Homebrew (Recommended)
```bash
# Download
brew install daypunk/tap/punkdoku

# Run 🧩
punkdoku
```

### Option 2: Manual Download

#### macOS
```bash
# Download
curl -L -o punkdoku https://github.com/daypunk/punkdoku/releases/latest/download/punkdoku-macos-$(uname -m | sed 's/x86_64/amd64/')

# Make it executable
chmod +x punkdoku

# Move to your PATH
sudo mv punkdoku /usr/local/bin/

# Run 🧩
punkdoku
```

#### Linux
```bash
# Download
curl -L -o punkdoku https://github.com/daypunk/punkdoku/releases/latest/download/punkdoku-linux

# Make it executable
chmod +x punkdoku

# Move to your PATH
sudo mv punkdoku /usr/local/bin/

# Run 🧩
punkdoku
```

## How to Play

Run **`punkdoku`** in your terminal and use:
- **Arrow keys** to navigate
- **1-9** to place numbers
- **0** or **Space** to clear cells
- **u** to undo
- **a** to toggle auto-check
- **t** to toggle timer
- **m** to return to menu
- **q** to quit

## Game Modes

- **🍼 Easy** - Good for beginners
- **🌞 Normal** - Balanced challenge
- **🌚 Hard** - Requires strategy
- **🥀 Lunatic** - Expert level
- **🌞 Daily(=Normal)** - Same puzzle for everyone, changes daily

## Features

- Cute! minimalist interface
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
