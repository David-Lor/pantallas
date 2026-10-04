package main

import (
	"fmt"

	"github.com/pterm/pterm"
)

const (
	mainMenu_newSession   = "✨New session"
	mainMenu_openSession  = "📂Open session"
	mainMenu_killSessions = "💥Kill sessions"
	mainMenu_exit         = "👋Exit"
)

func MainMenu() bool {
	options := []string{
		mainMenu_newSession,
		mainMenu_openSession,
		mainMenu_killSessions,
		mainMenu_exit,
	}

	printer := pterm.DefaultInteractiveSelect.
		WithOptions(options)

	option, err := printer.Show("Select an option")
	if err != nil {
		return true
	}

	switch option {
	case mainMenu_newSession:
		NewSession()
		return true
	case mainMenu_openSession:
		OpenSession()
		return true
	case mainMenu_killSessions:
		KillSessions()
		return true
	case mainMenu_exit:
		return false
	default:
		return true
	}
}

func NewSession() {
	sessionName, err := pterm.DefaultInteractiveTextInput.Show("✨Enter the new session name")
	// TODO Option to cancel
	if err != nil || sessionName == "" {
		return
	}

	TmuxNewSession(sessionName)
}

func OpenSession() {
	ok, allSessions := listSessions()
	if !ok {
		return
	}
	if allSessions == nil {
		pterm.Println("🙅There are no active sessions")
		return
	}

	printer := pterm.DefaultInteractiveSelect.
		WithOptions(allSessions).
		WithFilterInputPlaceholder("🔍 Start typing")

	option, err := printer.Show("Select an option")
	if err != nil {
		return
	}

	TmuxAttachSession(option)
}

func KillSessions() {
	ok, allSessions := listSessions()
	if !ok {
		return
	}
	if allSessions == nil {
		pterm.Println("🙅There are no active sessions")
		return
	}

	printer := pterm.DefaultInteractiveMultiselect.
		WithOptions(allSessions).
		WithFilterInputPlaceholder("🔍 Start typing")

	selectedOptions, err := printer.Show("💥Select the panels to kill")
	if err != nil {
		return
	}

	TmuxKillSessions(selectedOptions)
	// TODO parse errors, show ok/ko for killed panels
}

func listSessions() (ok bool, sessions []string) {
	sessions, err := TmuxListSessions()
	if err != nil {
		pterm.Error.Println(fmt.Sprintf("failed listing sessions: %v", err))
		return false, nil
	}

	return true, sessions
}
