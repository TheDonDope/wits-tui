package tui

import (
	"fmt"
	"regexp"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSettingsHomeModel(t *testing.T) {
	t.Run("Initialization", func(t *testing.T) {
		model := initialSettingsModel()
		expectedTitle := breadcrumbTitle(homeTitle, settingsTitle)
		assert.Equal(t, expectedTitle, model.hm.title)
	})

	t.Run("Init", func(t *testing.T) {
		model := initialSettingsModel()
		cmd := model.Init()
		assert.Nil(t, cmd)
	})

	t.Run("Update", func(t *testing.T) {
		t.Run("QuitKeys", func(t *testing.T) {
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
					model := initialSettingsModel()
					msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(tt.key)}

					_, cmd := model.Update(msg)
					if tt.wantCmd != nil {
						require.NotNil(t, cmd)
						assert.Equal(t, tt.wantCmd(), cmd())
					}
				})
			}
		})

		t.Run("EscapeKey", func(t *testing.T) {
			model := initialSettingsModel()
			msg := tea.KeyMsg{Type: tea.KeyEscape}

			updatedModel, cmd := model.Update(msg)
			assert.IsType(t, MenuModel{}, updatedModel)
			assert.Nil(t, cmd)
		})

		t.Run("HomeModelPropagation", func(t *testing.T) {
			model := initialSettingsModel()
			originalTitle := model.hm.title

			// Simulate HomeModel update
			model.hm.Title("Modified Title")
			msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'a'}}
			_, cmd := model.Update(msg)

			assert.NotEqual(t, originalTitle, model.hm.title)
			assert.Nil(t, cmd)
		})
	})

	t.Run("View", func(t *testing.T) {
		model := initialSettingsModel()
		view := model.View()

		assert.Contains(t, view, settingsTitle)
		assert.Contains(t, view, homeTitle)
		assert.NotEmpty(t, view)
	})

	t.Run("TitleFormatting", func(t *testing.T) {
		model := initialSettingsModel()
		expected := breadcrumbTitle(homeTitle, settingsTitle)
		assert.Equal(t, expected, model.hm.title)
	})
	t.Run("SettingsActions", func(t *testing.T) {
		ansiRegex := regexp.MustCompile(`\x1b\[[0-9;]*m`)
		stripANSI := func(s string) string { return ansiRegex.ReplaceAllString(s, "") }

		tests := []struct {
			action     settingsAction
			fullText   string
			markedChar string
		}{
			{appereance, "🎨 Appearance", "A"},
			{keybindings, "⌨️ Keybindings", "K"},
			{localization, "🌍 Localization", "L"},
			{backupAndRestore, "💾 Backup & Restore", "B"},
		}

		for _, tt := range tests {
			t.Run(tt.fullText, func(t *testing.T) {
				// Verify full text without ANSI codes
				cleanText := stripANSI(settingsActions[tt.action])
				assert.Equal(t, tt.fullText, cleanText)

				// Verify underlined character format
				ansiPattern := fmt.Sprintf(`\x1b\[4m%s\x1b\[0m`, tt.markedChar)
				assert.Regexp(t, regexp.MustCompile(ansiPattern), settingsActions[tt.action])
			})
		}
	})
}
