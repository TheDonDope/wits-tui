package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var appliances = []string{
	"[=== Strains ===]",
	"[=== Devices ===]",
	"[=== Settings ===]",
	"[=== Statistics ===]",
}

// MenuModel is the tea.Model for the main menu.
type MenuModel struct {
	cursor int
	items  []string
}

// InitialMenuModel returns the initial model for the main menu.
func InitialMenuModel() MenuModel {
	return MenuModel{
		items: appliances,
	}
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (m MenuModel) Init() tea.Cmd {
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
		case "enter":
			return onMenuSelected(m)
		case "esc":
			return InitialMenuModel(), nil
		}
	}
	return m, nil
}

// View renders the program's UI using Lipgloss for styling.
func (m MenuModel) View() string {
	// Create a fancy header style using Lipgloss.
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("230")).     // Light text color.
		Background(lipgloss.Color("#4CAF50")). // Calm emerald green background.
		Padding(1, 4).
		Border(lipgloss.RoundedBorder(), true).
		Align(lipgloss.Center).
		Width(40)
	header := headerStyle.Render(" Wits")

	// Set a fixed width for items.
	itemWidth := 30

	// Define a style for unselected menu items.
	itemStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), true).
		Padding(0, 1).
		Width(itemWidth).
		Align(lipgloss.Center).
		Foreground(lipgloss.Color("240")).
		Background(lipgloss.Color("0"))

	// Define a style for the selected menu item: bold with highlighted background.
	selectedStyle := lipgloss.NewStyle().
		Bold(true).
		Border(lipgloss.RoundedBorder(), true).
		Padding(0, 1).
		Width(itemWidth).
		Align(lipgloss.Center).
		Foreground(lipgloss.Color("230")). // Light text color for contrast.
		Background(lipgloss.Color("236"))  // Highlight background.

	// Render each menu item.
	s := header + "\n\n"
	for i, item := range m.items {
		var rendered string
		if m.cursor == i {
			rendered = selectedStyle.Render(item)
		} else {
			rendered = itemStyle.Render(item)
		}
		s += rendered + "\n\n"
	}
	s += "\nPress ctrl+c or q to quit."

	// Wrap the entire view in a container style that centers the block.
	containerStyle := lipgloss.NewStyle().
		Width(80). // Set the container width to 80 (adjust as needed)
		Align(lipgloss.Center)
	return containerStyle.Render(s)
}

// onMenuSelected returns a model for the selected menu.
func onMenuSelected(m MenuModel) (tea.Model, tea.Cmd) {
	switch m.cursor {
	case 0:
		// Open the strains view.
		shm := initialStrainsHomeModel()
		return shm, shm.onStrainsListed()
	case 1:
		return initialDevicesHomeModel(), nil
	case 2:
		return initialSettingsModel(), nil
	case 3:
		return initialStatisticsHomeModel(), nil
	}
	return m, nil
}
