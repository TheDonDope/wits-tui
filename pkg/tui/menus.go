package tui

import (
	"fmt"

	"github.com/TheDonDope/wits-tui/pkg/storage"
	tea "github.com/charmbracelet/bubbletea"
)

// Appliances is a list of options for the main menu.
var Appliances = []string{
	markedText("🌿 &Strains"),
	markedText("🚀 &Devices"),
	markedText("🔧 S&ettings"),
	markedText("📊 S&tatistics")}

// MenuModel is the model for the main menu.
type MenuModel struct {
	cursor int
	items  []string
}

// InitialMenuModel returns the initial model for the main menu.
func InitialMenuModel() MenuModel {
	return MenuModel{
		items: Appliances,
	}
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (m MenuModel) Init() tea.Cmd {
	storage.EnsureWitsFolder()
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.items) - 1 // Wrap to last item
			}
		case "down", "j":
			m.cursor++
			if m.cursor >= len(m.items) {
				m.cursor = 0 // Wrap to first item
			}
		case "alt+s":
			m.cursor = 0
		case "alt+d":
			m.cursor = 1
		case "alt+e":
			m.cursor = 2
		case "alt+t":
			m.cursor = 3
		case "enter":
			return onMenuSelected(m)
		case "esc":
			return InitialMenuModel(), nil
		}
	}
	return m, nil
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (m MenuModel) View() string {
	s := "🥦 Welcome to Wits!\n\n"
	for i, item := range m.items {
		cursor := " "
		if m.cursor == i {
			cursor = "> "
		}
		s += fmt.Sprintf("%s%s\n", cursor, item)
	}
	s += "\nPress ctrl+c or q to quit."
	return s
}

// onMenuSelected returns a model for the selected menu.
func onMenuSelected(m MenuModel) (tea.Model, tea.Cmd) {
	switch m.cursor {
	case 0:
		return NewStrainsAppliance(), nil
	case 1:
		return NewDevicesAppliance(), nil
	case 2:
		return NewSettingsAppliance(), nil
	case 3:
		return NewStatisticsAppliance(), nil
	}
	return m, nil
}
