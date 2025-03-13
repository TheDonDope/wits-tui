// Package main is the entry point for the Wits TUI application.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/TheDonDope/wits-tui/pkg/tui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
)

func main() {
	loadEnvironment()
	ensureWitsFolders()
	configureLogging()

	_, err := tea.NewProgram(tui.InitialMenuModel(), tea.WithAltScreen()).Run()
	if err != nil {
		log.Fatalf("Error starting program: %v", err)
	}
}

func loadEnvironment() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load configuration from environment: %v", err)
	}
}

func ensureWitsFolders() error {
	return os.MkdirAll(fmt.Sprintf("%s/%s", os.Getenv("WITS_DIR"), os.Getenv("LOG_DIR")), os.ModePerm)
}

func configureLogging() {
	if len(os.Getenv("LOG_LEVEL")) > 0 {
		f, err := tea.LogToFile(fmt.Sprintf("%s/%s/%s", os.Getenv("WITS_DIR"), os.Getenv("LOG_DIR"), os.Getenv("LOG_FILE")), "debug")
		if err != nil {
			log.Fatalf("Failed setting the debug log file: %v", err)
		}
		defer f.Close()
	}
}
