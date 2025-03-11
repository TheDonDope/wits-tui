package tui

import (
	"fmt"

	"github.com/TheDonDope/wits-tui/pkg/service"
	"github.com/TheDonDope/wits-tui/pkg/storage"
	tea "github.com/charmbracelet/bubbletea"
)

var (
	strainStore   storage.StrainStore
	strainService service.StrainService
)

// MainMenu is a list of options for the main menu.
var MainMenu = []string{
	"🌿 Strains",
	"🚀 Devices",
	"🔧 Settings",
	"📊 Stats"}

// MenuModel is the model for the main menu.
type MenuModel struct {
	cursor int
	items  []string
	menu   string
}

// InitialMenuModel returns the initial model for the main menu.
func InitialMenuModel(ssvc service.StrainService, sstr storage.StrainStore) MenuModel {
	strainStore = sstr
	strainService = ssvc
	return MenuModel{
		items: MainMenu,
		menu:  "main",
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
				m.cursor = len(m.menu) - 1 // Wrap to last item
			}
		case "down", "j":
			m.cursor++
			if m.cursor >= len(m.menu) {
				m.cursor = 0 // Wrap to first item
			}
		case "1", "2", "3", "4":
			idx := int(msg.String()[0] - '1') // Convert key to index
			if idx < len(m.menu) {
				m.cursor = idx // Jump to selected menu item
			}
		case "enter":
			return onMenuSelected(m)
		case "esc":
			return InitialMenuModel(strainService, strainStore), nil
		}
	}
	return m, nil
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (m MenuModel) View() string {
	s := "🥦 Welcome to Wits!\n\n"
	if m.menu == "main" {
		for i, item := range m.items {
			cursor := " "
			if m.cursor == i {
				cursor = "➡️ "
			}
			s += fmt.Sprintf("%s(%d): %s\n", cursor, i+1, item)
		}
	} else {
		s += onSubmenuSelected(m)
	}
	s += "\nPress ctrl+c or q to quit."
	if m.menu != "main" {
		s += "\nPress esc to return to main menu."
	}
	return s
}

// onMenuSelected returns a model for the selected menu.
func onMenuSelected(m MenuModel) (tea.Model, tea.Cmd) {
	switch m.menu {
	case "main":
		switch m.cursor {
		case 0:
			return MenuModel{
				items: StrainsSubmenu,
				menu:  "strains"}, nil
		case 1:
			return MenuModel{
				items: DevicesSubmenu,
				menu:  "devices"}, nil
		case 2:
			return MenuModel{
				items: SettingsSubmenu,
				menu:  "settings"}, nil
		case 3:
			return MenuModel{
				items: StatsSubmenu,
				menu:  "stats"}, nil
		}
	case "strains":
		switch m.cursor {
		case 0:
			return AddStrain(strainService), nil
		case 1:
			return ListStrains(strainService), nil
		}
	}
	return m, nil
}

// onSubmenuSelected renders the selected submenu and its items.
func onSubmenuSelected(m MenuModel) string {
	s := fmt.Sprintf("%s Menu:\n", m.menu)
	for i, item := range m.items {
		cursor := " "
		if m.cursor == i {
			cursor = "➡️ "
		}
		s += fmt.Sprintf("%s(%d): %s\n", cursor, i+1, item)
	}
	return s
}
