package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

const settingsTitle = "🔧 Settings"

type settingsAction int

const (
	appereance settingsAction = iota
	keybindings
	localization
	backupAndRestore
)

// settingsActions is a list of options for the settings appliance.
var settingsActions = map[settingsAction]string{
	appereance:       markedText("🎨 &Appearance"),
	keybindings:      markedText("⌨️ &Keybindings"),
	localization:     markedText("🌍 &Localization"),
	backupAndRestore: markedText("💾 &Backup & Restore")}

// SettingsAppliance ...
type SettingsAppliance struct {
	hv *HomeView
}

// NewSettingsAppliance ...
func NewSettingsAppliance() *SettingsAppliance {
	s := &SettingsAppliance{
		hv: NewHomeView(),
	}
	s.hv.Title(breadcrumbTitle(s.hv.title, settingsTitle))
	return s
}

// Init ...
func (s *SettingsAppliance) Init() tea.Cmd {
	return nil
}

// Update ...
func (s *SettingsAppliance) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return s, tea.Quit
		case "esc":
			return InitialMenuModel(), nil
		}
	}

	var cmd tea.Cmd
	hv, cmd := s.hv.Update(msg)
	s.hv = hv.(*HomeView)
	return s, cmd
}

// View ...
func (s *SettingsAppliance) View() string {
	return s.hv.View()
}
