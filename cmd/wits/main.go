// Package main is the entry point for the Wits TUI application.
package main

import (
	"fmt"
	"os"

	"github.com/TheDonDope/wits-tui/pkg/service"
	"github.com/TheDonDope/wits-tui/pkg/storage"
	"github.com/TheDonDope/wits-tui/pkg/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	strainStore, err := storage.NewStrainStoreYMLFile()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading data from YML File: %v", err)
		os.Exit(1)
	}
	strainService := service.NewStrainService(strainStore)
	_, err = tea.NewProgram(tui.InitialMenuModel(strainService, strainStore), tea.WithAltScreen()).Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error starting program: %v", err)
		os.Exit(1)
	}
}
