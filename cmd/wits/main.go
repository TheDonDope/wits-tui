// Package main is the entry point for the Wits TUI application.
package main

import (
	"fmt"
	"os"

	"github.com/TheDonDope/wits-tui/pkg/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	_, err := tea.NewProgram(tui.InitialMenuModel(), tea.WithAltScreen()).Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error starting program: %v", err)
		os.Exit(1)
	}
}
