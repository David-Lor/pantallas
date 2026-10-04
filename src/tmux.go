package main

import (
	"os"
	"os/exec"
	"strings"
)

func CheckTmux() bool {
	out, _ := exec.Command(Settings.TmuxPath, "--help").CombinedOutput()
	return strings.Contains(string(out), "usage: tmux")
}

func TmuxNewSession(name string) {
	runInteractiveTmuxCmd("new-session", "-s", name)
}

func TmuxListSessions() []string {
	output, err := exec.Command(Settings.TmuxPath, "list-sessions", "-F", "#{session_name}").Output()
	if err != nil {
		return nil
	}

	return strings.Split(string(output), "\n")
}

func TmuxAttachSession(name string) {
	runInteractiveTmuxCmd("attach-session", "-t", name)
}

func TmuxKillSessions(sessions []string) {
	for _, session := range sessions {
		TmuxKillSession(session)
	}
}

func TmuxKillSession(name string) {
	runInteractiveTmuxCmd("kill-session", "-t", name)
}

func runInteractiveTmuxCmd(arg ...string) {
	cmd := exec.Command(Settings.TmuxPath, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}
