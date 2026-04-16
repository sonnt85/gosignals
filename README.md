# gosignals

Cross-platform OS signal utilities — convert signal names to `os.Signal` values and send signals to processes with optional child-process propagation.

## Installation

```bash
go get github.com/sonnt85/gosignals
```

## Features

- Convert signal name strings (e.g. `"SIGTERM"`, `"TERM"`) to `os.Signal` values
- Cross-platform: Linux, macOS, and Windows implementations
- `Kill` function that optionally sends the signal to an entire process group (Unix) or uses `taskkill /F /T` (Windows)
- Full signal name map covering all common Unix signals on Linux and macOS

## Usage

```go
sig, err := gosignals.ToSignal("SIGUSR1")
if err != nil {
    log.Fatal(err)
}

// Send signal to process (and optionally its children)
proc, _ := os.FindProcess(pid)
err = gosignals.Kill(proc, sig, true) // true = send to process group
```

## API

- `ToSignal(signalName string) (os.Signal, error)` — convert a signal name to `os.Signal`; the `SIG` prefix is added automatically if missing; falls back to `SIGTERM` for unknown names on Unix
- `Kill(process *os.Process, sig os.Signal, sigChildren bool) error` — send a signal to a process; on Unix with `sigChildren=true`, negates the PID to target the process group; on Windows uses `taskkill /F /T`

## License

MIT License - see [LICENSE](LICENSE) for details.
