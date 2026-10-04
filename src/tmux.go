package main

import (
	"os"
	"os/exec"
	"strings"
)

func CheckTmux() bool {
	err := exec.Command(Settings.TmuxPath, "ls").Run()
	return err == nil
}

func TmuxNewSession(name string) {
	cmd := exec.Command(Settings.TmuxPath, "new-session", "-s", name)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

func TmuxListSessions() ([]string, error) {
	output, err := exec.Command(Settings.TmuxPath, "list-sessions", "-F", "#{session_name}").Output()
	if err != nil {
		return nil, err
	}

	return strings.Split(string(output), "\n"), nil
}

func TmuxAttachSession(name string) {
	cmd := exec.Command(Settings.TmuxPath, "attach-session", "-t", name)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}
