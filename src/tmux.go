package main

import (
	"os"
	"os/exec"
)

func CheckTmux() bool {
	err := exec.Command(Settings.TmuxPath, "ls").Run()
	return err == nil
}

func TmuxNewSession(name string) {
	cmd := exec.Command(Settings.TmuxPath, "new", "-s", name)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}
