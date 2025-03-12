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

var settingsActions = map[settingsAction]string{
	appereance:       markedText("🎨 &Appearance"),
	keybindings:      markedText("⌨️ &Keybindings"),
	localization:     markedText("🌍 &Localization"),
	backupAndRestore: markedText("💾 &Backup & Restore")}

// SettingsAppliance is the tea.Model for the Settings appliance.
type SettingsAppliance struct {
	hv *HomeView
}

// NewSettingsAppliance returns a new SettingsAppliance, with the following contents:
//   - rendered title
func NewSettingsAppliance() *SettingsAppliance {
	s := &SettingsAppliance{
		hv: NewHomeView(),
	}
	s.hv.Title(breadcrumbTitle(s.hv.title, settingsTitle))
	return s
}

// SettingsAppliance implementation of tea.Model interface ---------------------

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (s *SettingsAppliance) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (s *SettingsAppliance) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return s, tea.Quit
		case "esc":
			return NewMenuModel(), nil
		}
	}

	var cmd tea.Cmd
	hv, cmd := s.hv.Update(msg)
	s.hv = hv.(*HomeView)
	return s, cmd
}

// View renders the StatisticsAppliance UI, which is just a string. The view is
// rendered after every Update.
func (s *SettingsAppliance) View() string {
	return s.hv.View()
}
