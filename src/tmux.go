package main

import "os/exec"

func CheckTmux() bool {
	err := exec.Command(Settings.TmuxPath, "ls").Run()
	return err == nil
}
