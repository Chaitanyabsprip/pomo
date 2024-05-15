# Pomo

`pomo` is a simple command line pomodoro timer.

## Features

Pomo features very basic set of features including start, stop and show running
time. Additionally:

- It also allows to pause timers.
- If you're using it with watch or as a statusline element in tmux or neovim,
  the symbol will flash
- You can have multiple binaries, named differently running different timers

## Installation

Make sure you have go installed.  
You can install pomo by cloning the directory and running make install

```bash
  git clone https://github.com/Chaitanyabsprip/pomo.git
  sudo make install
```

Or using go CLI

```sh
go install github.com/Chaitanyabsprip/pomo/cmd/pomo
```

## Usage/Examples

- Start a timer

```sh
pomo start
```

You can optionally pass the duration in `2h54m31s` format or `hr` or `hour` to
start a timer until the start of the next hour.

```sh
pomo start 52m30s
pomo start hour
```

- Stop a timer

```sh
pomo Stop
```

- Pause a timer

```sh
pomo pause
```

- Print the pending duration on stdout

```sh
pomo
```

## Acknowledgements

- [rwxrob/pomo](https://github.com/rwxrob/pomo)
