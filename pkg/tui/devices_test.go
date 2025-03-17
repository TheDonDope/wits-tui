package tui

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDevicesHomeModel(t *testing.T) {
	t.Run("Initialization", func(t *testing.T) {
		model := initialDevicesHomeModel()

		expectedTitle := breadcrumbTitle(homeTitle, devicesTitle)
		assert.Equal(t, expectedTitle, model.hm.title)
	})

	t.Run("Init", func(t *testing.T) {
		model := initialDevicesHomeModel()
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
					model := initialDevicesHomeModel()
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
			model := initialDevicesHomeModel()
			msg := tea.KeyMsg{Type: tea.KeyEscape}

			updatedModel, cmd := model.Update(msg)
			assert.IsType(t, MenuModel{}, updatedModel)
			assert.Nil(t, cmd)
		})

		t.Run("PropagationToHomeModel", func(t *testing.T) {
			model := initialDevicesHomeModel()
			originalHM := model.hm

			// Send a harmless message
			msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'a'}}
			_, cmd := model.Update(msg)

			assert.Same(t, originalHM, model.hm)
			assert.Nil(t, cmd)
		})
	})

	t.Run("View", func(t *testing.T) {
		model := initialDevicesHomeModel()
		view := model.View()

		assert.Contains(t, view, devicesTitle)
		assert.Contains(t, view, homeTitle)
		assert.NotEmpty(t, view)
	})

	t.Run("TitleFormatting", func(t *testing.T) {
		model := initialDevicesHomeModel()
		expected := breadcrumbTitle(homeTitle, devicesTitle)
		assert.Equal(t, expected, model.hm.title)
	})
}
