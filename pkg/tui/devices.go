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

// DevicesAppliance is the tea.Model for the Devices appliance.
type DevicesAppliance struct {
	hv *HomeView
}

// NewDevicesAppliance returns a new DevicesAppliance, with the following contents:
//   - rendered title
func NewDevicesAppliance() *DevicesAppliance {
	d := &DevicesAppliance{
		hv: NewHomeView(),
	}
	d.hv.Title(breadcrumbTitle(d.hv.title, devicesTitle))
	return d
}

// StatisticsAppliance implementation of tea.Model interface -------------------

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (d *DevicesAppliance) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (d *DevicesAppliance) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return d, tea.Quit
		case "esc":
			return NewMenuModel(), nil
		}
	}

	var cmd tea.Cmd
	// FIXME: these 2 lines seems wonky
	hv, cmd := d.hv.Update(msg)
	d.hv = hv.(*HomeView)
	return d, cmd
}

// View renders the DevicesAppliance UI, which is just a string. The view is
// rendered after every Update.
func (d *DevicesAppliance) View() string {
	return d.hv.View()
}
