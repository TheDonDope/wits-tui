package home

import (
	"log"

	"github.com/TheDonDope/wits-tui/pkg/tui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// Command is the home command.
var Command = &cobra.Command{
	Use:   "home",
	Short: "Launch the main menu",
	Args:  cobra.NoArgs,
	RunE: func(_ *cobra.Command, _ []string) error {
		_, err := tea.NewProgram(tui.InitialMenuModel(), tea.WithAltScreen()).Run()
		if err != nil {
			log.Fatalf("🚨 🖥️  (cmd/wits/main.go) ❓ 🗒️  Error starting program: %v \n", err)
			return err
		}
		return nil
	},
}
