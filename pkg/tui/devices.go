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

// devicesActions is a list of options for the devices appliance.
var devicesActions = map[devicesAction]string{
	createDevice: "➕ Add Device",
	viewDevice:   "📋 View Devices",
	editDevice:   "✏️ Edit Device",
	deleteDevice: "❌ Delete Device"}

// DevicesAppliance ...
type DevicesAppliance struct {
	hv *HomeView
}

// NewDevicesAppliance ...
func NewDevicesAppliance() *DevicesAppliance {
	d := &DevicesAppliance{
		hv: NewHomeView(),
	}
	d.hv.Title(breadcrumbTitle(d.hv.title, devicesTitle))
	return d
}

// Init ...
func (d *DevicesAppliance) Init() tea.Cmd {
	return nil
}

// Update ...
func (d *DevicesAppliance) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return d, tea.Quit
		case "esc":
			return InitialMenuModel(), nil
		}
	}

	var cmd tea.Cmd
	// FIXME: these 2 lines seems wonky
	hv, cmd := d.hv.Update(msg)
	d.hv = hv.(*HomeView)
	return d, cmd
}

// View ...
func (d *DevicesAppliance) View() string {
	return d.hv.View()
}
