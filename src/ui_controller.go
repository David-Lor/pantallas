package main

import (
	"strings"

	"github.com/pterm/pterm"
)

const (
	mainMenu_newSession  = "New session"
	mainMenu_openSession = "Open session"
	mainMenu_killSession = "Kill session"
	mainMenu_exit        = "Exit"
)

func MainMenu() bool {
	options := []string{
		mainMenu_newSession,
		mainMenu_openSession,
		mainMenu_killSession,
		mainMenu_exit,
	}

	printer := pterm.DefaultInteractiveSelect.
		WithOptions(options).
		WithFilterInputPlaceholder("🔍 Start typing")

	option, err := printer.Show("select option")
	if err != nil {
		return true
	}

	switch option {
	case mainMenu_newSession:
		NewSession()
		return true
	case mainMenu_exit:
		return false
	default:
		return true
	}
}

func NewSession() {
	sessionName, err := pterm.DefaultInteractiveTextInput.Show("Enter new session name")
	sessionName = strings.TrimSpace(sessionName)
	if err != nil || sessionName == "" {
		return
	}

	TmuxNewSession(sessionName)
}
