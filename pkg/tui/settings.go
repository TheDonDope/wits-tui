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

// SettingsHomeModel is the tea.Model for the Settings appliance.
type SettingsHomeModel struct {
	hm *HomeModel
}

// initialSettingsModel returns a new SettingsHomeModel, with the following contents:
//   - rendered title
func initialSettingsModel() *SettingsHomeModel {
	s := &SettingsHomeModel{
		hm: initialHomeModel(),
	}
	s.hm.Title(breadcrumbTitle(s.hm.title, settingsTitle))
	return s
}

// SettingsHomeModel implementation of tea.Model interface ---------------------

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (shm *SettingsHomeModel) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (shm *SettingsHomeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return shm, tea.Quit
		case "esc":
			return InitialMenuModel(), nil
		}
	}

	var cmd tea.Cmd
	hv, cmd := shm.hm.Update(msg)
	shm.hm = hv.(*HomeModel)
	return shm, cmd
}

// View renders the SettingsHomeModel UI, which is just a string. The view is
// rendered after every Update.
func (shm *SettingsHomeModel) View() string {
	return shm.hm.View()
}
