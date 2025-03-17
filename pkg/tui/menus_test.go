package tui

import (
	"io"
	"log"
	"os"
	"regexp"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestMain handles global test setup
func TestMain(m *testing.M) {
	// Disable log output during tests
	log.SetOutput(io.Discard)
	os.Exit(m.Run())
}

func TestMenuModel(t *testing.T) {
	ansiRegex := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	stripANSI := func(s string) string { return ansiRegex.ReplaceAllString(s, "") }

	t.Run("Initialization", func(t *testing.T) {
		model := InitialMenuModel()
		assert.Equal(t, 0, model.cursor)
		assert.Len(t, model.items, 4)
	})

	t.Run("View", func(t *testing.T) {
		model := InitialMenuModel()
		view := model.View()
		cleanView := stripANSI(view)

		// Verify header and footer
		assert.Contains(t, cleanView, "Welcome to Wits")
		assert.Contains(t, cleanView, "Press ctrl+c or q to quit")

		// Verify menu items
		assert.Contains(t, cleanView, "🌿 Strains")
		assert.Contains(t, cleanView, "🚀 Devices")
		assert.Contains(t, cleanView, "🔧 Settings")
		assert.Contains(t, cleanView, "📊 Statistics")
	})

	t.Run("Navigation", func(t *testing.T) {
		tests := []struct {
			name        string
			keys        []string
			expectIndex int
		}{
			{"DownArrow", []string{"down"}, 1},
			{"UpArrow", []string{"up"}, 3},
			{"WrapDown", []string{"down", "down", "down", "down"}, 0},
			{"WrapUp", []string{"up", "up"}, 2},
			{"KKey", []string{"k"}, 3},
			{"JKey", []string{"j"}, 1},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// Start with concrete MenuModel type
				model := InitialMenuModel()
				for _, key := range tt.keys {
					// Update and type assert
					updated, _ := model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(key)})
					var ok bool
					model, ok = updated.(MenuModel)
					require.True(t, ok, "Should maintain MenuModel type")
				}
				assert.Equal(t, tt.expectIndex, model.cursor)
			})
		}
	})

	t.Run("Shortcuts", func(t *testing.T) {
		tests := []struct {
			name      string
			shortcut  string
			checkType any
		}{
			{"AltS", "alt+s", &StrainsHomeModel{}},
			{"AltD", "alt+d", &DevicesHomeModel{}},
			{"AltE", "alt+e", &SettingsHomeModel{}},
			{"AltT", "alt+t", &StatisticsHomeModel{}},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				model := InitialMenuModel()
				updated, _ := model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(tt.shortcut)})
				assert.IsType(t, tt.checkType, updated)
			})
		}
	})

	t.Run("Selection", func(t *testing.T) {
		tests := []struct {
			name      string
			cursorPos int
			checkType any
		}{
			{"Strains", 0, &StrainsHomeModel{}},
			{"Devices", 1, &DevicesHomeModel{}},
			{"Settings", 2, &SettingsHomeModel{}},
			{"Statistics", 3, &StatisticsHomeModel{}},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				model := InitialMenuModel()
				model.cursor = tt.cursorPos
				updated, _ := model.Update(tea.KeyMsg{Type: tea.KeyEnter})
				assert.IsType(t, tt.checkType, updated)
			})
		}
	})

	t.Run("QuitHandling", func(t *testing.T) {
		tests := []struct {
			name    string
			key     string
			wantCmd tea.Cmd
		}{
			{"QKey", "q", tea.Quit},
			{"CtrlC", "ctrl+c", tea.Quit},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				model := InitialMenuModel()
				_, cmd := model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(tt.key)})
				if tt.wantCmd != nil {
					require.NotNil(t, cmd)
					assert.Equal(t, tt.wantCmd(), cmd())
				}
			})
		}
	})

	t.Run("MarkedTextFormatting", func(t *testing.T) {
		expectedItems := []string{
			"🌿 Strains",
			"🚀 Devices",
			"🔧 Settings",
			"📊 Statistics",
		}

		for i, item := range InitialMenuModel().items {
			t.Run(expectedItems[i], func(t *testing.T) {
				// Verify clean text
				clean := stripANSI(item)
				assert.Equal(t, expectedItems[i], clean)

				// Verify underline formatting
				assert.Regexp(t, `\x1b\[4m.?\x1b\[0m`, item)
			})
		}
	})
}
