package main

import (
	"fmt"
	"os"
)

func main() {
	LoadSettings()
	if !CheckTmux() {
		fmt.Println("tmux not installed")
		os.Exit(1)
	}

	fmt.Println("ok")
}
