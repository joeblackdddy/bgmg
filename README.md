# bgmg

**Be** **G**reat, **M**ake **G**reat 🐐

A CLI tool to display cool ASCII art for success/failure in GitHub Actions workflows

## Installation

```bash
go install github.com/joeblackdddy/bgmg/cmd/bgmg@latest
```

After installation, make sure `$HOME/go/bin` is in your PATH:

```bash
# Add to .bashrc or .zshrc
export PATH=$PATH:$HOME/go/bin
```

## Usage

```bash
# Show success art
# No arguments or 0
bgmg 0

# Show failure art
# Any non-zero argument
bgmg 1
```

## GitHub Actions Example

### Method 1: Using `go install` (Recommended)

The simplest way:

```yaml
name: Example Workflow

on: [push]

jobs:
  test:
    runs-on: ubuntu-latest

    steps:
      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.24'

      - name: Install bgmg
        run: go install github.com/joeblackdddy/bgmg/cmd/bgmg@latest

      - name: Run your tests
        run: echo "All tests passed!"

      - name: Show success
        if: success()
        run: bgmg 0

      - name: Show failure
        if: failure()
        run: bgmg 1
```

## License

MIT
