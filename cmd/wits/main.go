// Package main is the entry point for the Wits TUI application.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime/debug"

	"github.com/TheDonDope/wits-tui/cmd/wits/home"
	"github.com/TheDonDope/wits-tui/pkg/version"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var (
	// Version contains the application version number. It's set via ldflags
	// when building.
	Version = ""

	// CommitSHA contains the SHA of the commit that this application was built
	// against. It's set via ldflags when building.
	CommitSHA = ""

	// CommitDate contains the date of the commit that this application was
	// built against. It's set via ldflags when building.
	CommitDate = ""

	rootCmd = &cobra.Command{
		Use:          "wits",
		Short:        "A tui for cannabis patients and users",
		Long:         "Wits is the weed information tracking system, aimed to help cannabis patients and users.",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return home.Command.RunE(cmd, args)
		},
	}
)

func init() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true

	if len(CommitSHA) >= 7 {
		vt := rootCmd.VersionTemplate()
		rootCmd.SetVersionTemplate(vt[:len(vt)-1] + " (" + CommitSHA[0:7] + ")\n")
	}
	if Version == "" {
		if info, ok := debug.ReadBuildInfo(); ok && info.Main.Sum != "" {
			Version = info.Main.Version
		} else {
			Version = "unknown (built from source)"
		}
	}
	rootCmd.Version = Version

	version.Version = Version
	version.CommitSHA = CommitSHA
	version.CommitDate = CommitDate
}

func main() {
	log.Println("🚀 🖥️  (cmd/wits/main.go) main()")
	ctx := context.Background()
	loadEnvironment()
	ensureWitsFolders()
	f, err := tea.LogToFile(fmt.Sprintf("%s/%s/%s", os.Getenv("WITS_DIR"), os.Getenv("LOG_DIR"), os.Getenv("LOG_FILE")), "debug")
	if err != nil {
		log.Fatalf("🚨 🖥️  (cmd/wits/main.go) ❓ 🗒️  Failed setting the debug log file: %v \n", err)
	}
	defer f.Close()
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}

}

func loadEnvironment() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("🚨 🖥️  (cmd/wits/main.go) ❓ 🗒️  Failed to load configuration from environment: %v \n", err)
	}
	log.Println("✅ 🖥️  (cmd/wits/main.go) loadEnvironment()")
}

func ensureWitsFolders() error {
	log.Println("✅ 🖥️  (cmd/wits/main.go) ensureWitsFolders()")
	return os.MkdirAll(fmt.Sprintf("%s/%s", os.Getenv("WITS_DIR"), os.Getenv("LOG_DIR")), os.ModePerm)
}
