# Official Terminal Client for Codossium

## Currently done:
- Basic boilerplate for Bubbletea (TUI) and Cobra (replacement for `flags` library, it will power the CLI and flag handling)
- Basic file upload TUI, no uploading logic yet since the backend is not finalized. Also, some file extensions are now automatically recognized (the likes of Haskell, Go, Rust, Python, JS & TS), which will be useful when we begin implementing the backend. Test feature: `[go run main.go | ./codosseum] upload example-file.go`
- Basic TOML configuration. Creates a file in your configuration directory (e.g. `.config/codosseum/config.toml` on Linux). Test feature: `[go run main.go | ./codosseum] config`
