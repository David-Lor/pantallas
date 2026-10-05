package main

import (
	"github.com/pterm/pterm"
)

const (
	mmNewSession   = "✨New session"
	mmOpenSession  = "📂Open session"
	mmKillSessions = "💥Kill sessions"
	mmExit         = "👋Exit"

	msgEnterNewSession  = "✨Enter the new session name"
	msgNoActiveSessions = "🙅There are no active sessions"
	msgSearch           = "🔍Start typing"
	msgSelectOption     = "Select an option"
	msgSelectKill       = "💥Select the panels to kill"
)

func MainMenu() bool {
	options := []string{
		mmOpenSession,
		mmNewSession,
		mmKillSessions,
		mmExit,
	}

	// Hide menu options when no sessions available
	if len(TmuxListSessions()) == 0 {
		options[0] = "" // open session
		options[2] = "" // kill sessions
	}

	printer := pterm.DefaultInteractiveSelect.
		WithOptions(ClearArray(options))

	option, err := printer.Show("Select an option")
	if err != nil {
		return true
	}

	switch option {
	case mmNewSession:
		NewSession()
		return true
	case mmOpenSession:
		OpenSession()
		return true
	case mmKillSessions:
		KillSessions()
		return true
	case mmExit:
		return false
	default:
		return true
	}
}

func NewSession() {
	sessionName, err := pterm.DefaultInteractiveTextInput.Show(msgEnterNewSession)
	// TODO Option to cancel
	if err != nil || sessionName == "" {
		return
	}

	TmuxNewSession(sessionName)
}

func OpenSession() {
	allSessions := TmuxListSessions()
	if allSessions == nil {
		pterm.Println(msgNoActiveSessions)
		return
	}

	printer := pterm.DefaultInteractiveSelect.
		WithOptions(allSessions).
		WithFilterInputPlaceholder(msgSearch)

	option, err := printer.Show(msgSelectOption)
	if err != nil {
		return
	}

	TmuxAttachSession(option)
}

func KillSessions() {
	allSessions := TmuxListSessions()
	if allSessions == nil {
		pterm.Println(msgNoActiveSessions)
		return
	}

	printer := pterm.DefaultInteractiveMultiselect.
		WithOptions(allSessions).
		WithFilterInputPlaceholder(msgSearch)

	selectedOptions, err := printer.Show(msgSelectKill)
	if err != nil {
		return
	}

	TmuxKillSessions(selectedOptions)
	// TODO parse errors, show ok/ko for killed panels
}
