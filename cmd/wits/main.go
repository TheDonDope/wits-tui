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
	log.Println("🚀 🖥️  (cmd/wits/main.go) main()")
	loadEnvironment()
	ensureWitsFolders()
	configureLogging()

	_, err := tea.NewProgram(tui.InitialMenuModel(), tea.WithAltScreen()).Run()
	if err != nil {
		log.Fatalf("🚨 🖥️  (cmd/wits/main.go) ❓❓❓ ❓ 🗒️  Error starting program: %v \n", err)
	}
}

func loadEnvironment() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("🚨 🖥️  (cmd/wits/main.go) ❓❓❓ ❓ 🗒️  Failed to load configuration from environment: %v \n", err)
	}
	log.Println("✅ 🖥️  (cmd/wits/main.go) loadEnvironment()")
}

func configureLogging() {
	if len(os.Getenv("LOG_LEVEL")) > 0 {
		log.Println("💬 🖥️  (cmd/wits/main.go) configureLogging()")
		f, err := tea.LogToFile(fmt.Sprintf("%s/%s/%s", os.Getenv("WITS_DIR"), os.Getenv("LOG_DIR"), os.Getenv("LOG_FILE")), "debug")
		if err != nil {
			log.Fatalf("🚨 🖥️  (cmd/wits/main.go) ❓❓❓ ❓ 🗒️  Failed setting the debug log file: %v \n", err)
		}
		defer f.Close()
		log.Println("✅ 🖥️  (cmd/wits/main.go) configureLogging()")
	}
}

func ensureWitsFolders() error {
	log.Println("✅ 🖥️  (cmd/wits/main.go) ensureWitsFolders()")
	return os.MkdirAll(fmt.Sprintf("%s/%s", os.Getenv("WITS_DIR"), os.Getenv("LOG_DIR")), os.ModePerm)
}
