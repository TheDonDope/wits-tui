package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

const devicesTitle = "🚀 Devices"

type devicesAction int

const (
	createDevice devicesAction = iota
	viewDevice
	editDevice
	deleteDevice
)

var devicesActions = map[devicesAction]string{
	createDevice: markedText("➕ &Add Device"),
	viewDevice:   markedText("📋 &View Devices"),
	editDevice:   markedText("✏️ &Edit Device"),
	deleteDevice: markedText("❌ &Delete Device")}

// DevicesHomeModel is the tea.Model for the Devices appliance.
type DevicesHomeModel struct {
	hm *HomeModel
}

// initialDevicesHomeModel returns a new DevicesHomeModel, with the following contents:
//   - rendered title
func initialDevicesHomeModel() *DevicesHomeModel {
	d := &DevicesHomeModel{
		hm: initialHomeModel(),
	}
	d.hm.Title(breadcrumbTitle(d.hm.title, devicesTitle))
	return d
}

// DevicesHomeModel implementation of tea.Model interface -------------------

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (dhm *DevicesHomeModel) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (dhm *DevicesHomeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return dhm, tea.Quit
		case "esc":
			return InitialMenuModel(), nil
		}
	}

	var cmd tea.Cmd
	hm, cmd := dhm.hm.Update(msg)
	dhm.hm = hm.(*HomeModel)
	return dhm, cmd
}

// View renders the DevicesHomeModel UI, which is just a string. The view is
// rendered after every Update.
func (dhm *DevicesHomeModel) View() string {
	return dhm.hm.View()
}
