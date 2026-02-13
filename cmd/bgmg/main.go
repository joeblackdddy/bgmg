package main

import (
	"os"
	"github.com/joeblackdddy/bgmg/internal/cli"
)

func main() {
	if err := cli.Run(); err != nil {
		os.Exit(1)
	}
}
