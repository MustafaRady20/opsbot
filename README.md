# opsbot

A single Go binary that monitors servers and sites, sends a morning digest, and
captures notes, expenses and reminders through Telegram.

Status: early development.

## Requirements

- Go 1.27 or newer

## Build

```sh
make build   # builds bin/opsbot with the version from git describe
make test    # runs unit tests
make lint    # checks gofmt and runs go vet
```

## Usage

| Command | What it does |
| --- | --- |
| `opsbot check` | Checks every target once, prints one line each, and exits 1 if any is down |
| `opsbot version` | Prints the version |
| `opsbot --help` | Lists commands and flags |

Global flags:

| Flag | Default | Meaning |
| --- | --- | --- |
| `--config` | `config.yaml` | Path to the config file (used from T04) |
