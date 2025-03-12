package tui

import tea "github.com/charmbracelet/bubbletea"

const settingsTitle = "Settings"

type settingsAction int

const (
	appereance settingsAction = iota
	keybindings
	localization
	backupAndRestore
)

// settingsActions is a list of options for the settings appliance.
var settingsActions = map[settingsAction]string{
	appereance:       "🎨 Appearance",
	keybindings:      "⌨️ Keybindings",
	localization:     "🌍 Localization",
	backupAndRestore: "💾 Backup & Restore"}

// SettingsAppliance ...
type SettingsAppliance struct {
	hv *HomeView
}

// NewSettingsAppliance ...
func NewSettingsAppliance() *SettingsAppliance {
	s := &SettingsAppliance{
		hv: NewHomeView(),
	}
	return s
}

// Init ...
func (s *SettingsAppliance) Init() tea.Cmd {
	s.hv.title = settingsTitle
	return nil
}

// Update ...
func (s *SettingsAppliance) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
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
