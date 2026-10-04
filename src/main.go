package main

import (
	"os"

	"github.com/pterm/pterm"
)

func main() {
	LoadSettings()
	if !CheckTmux() {
		pterm.Error.Println("tmux is not installed")
		os.Exit(1)
	}

	for {
		if !MainMenu() {
			break
		}
		// TODO clear?
	}
}
