package tui

import (
	"fmt"
	"regexp"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStatisticsHomeModel(t *testing.T) {
	t.Run("Initialization", func(t *testing.T) {
		model := initialStatisticsHomeModel()
		expectedTitle := breadcrumbTitle(homeTitle, statisticsTitle)
		assert.Equal(t, expectedTitle, model.hm.title)
	})

	t.Run("Init", func(t *testing.T) {
		model := initialStatisticsHomeModel()
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
					model := initialStatisticsHomeModel()
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
			model := initialStatisticsHomeModel()
			msg := tea.KeyMsg{Type: tea.KeyEscape}

			updatedModel, cmd := model.Update(msg)
			assert.IsType(t, MenuModel{}, updatedModel)
			assert.Nil(t, cmd)
		})

		t.Run("HomeModelPropagation", func(t *testing.T) {
			model := initialStatisticsHomeModel()
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
		model := initialStatisticsHomeModel()
		view := model.View()

		assert.Contains(t, view, statisticsTitle)
		assert.Contains(t, view, homeTitle)
		assert.NotEmpty(t, view)
	})

	t.Run("TitleFormatting", func(t *testing.T) {
		model := initialStatisticsHomeModel()
		expected := breadcrumbTitle(homeTitle, statisticsTitle)
		assert.Equal(t, expected, model.hm.title)
	})

	t.Run("StatisticsActions", func(t *testing.T) {
		ansiRegex := regexp.MustCompile(`\x1b\[[0-9;]*m`)
		stripANSI := func(s string) string { return ansiRegex.ReplaceAllString(s, "") }

		tests := []struct {
			action     statisticsAction
			fullText   string
			markedChar string
		}{
			{usageHistory, "📅 Usage History", "U"},
			{trends, "📈 Trends", "T"},
			{dosageTracker, "🔢 Dosage Tracker", "D"},
		}

		for _, tt := range tests {
			t.Run(tt.fullText, func(t *testing.T) {
				// Verify full text without ANSI codes
				cleanText := stripANSI(statisticsActions[tt.action])
				assert.Equal(t, tt.fullText, cleanText)

				// Verify underlined character format
				ansiPattern := fmt.Sprintf(`\x1b\[4m%s\x1b\[0m`, tt.markedChar)
				assert.Regexp(t, regexp.MustCompile(ansiPattern), statisticsActions[tt.action])
			})
		}
	})
}
