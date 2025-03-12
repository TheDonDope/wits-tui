package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

const statisticsTitle = "📊 Statistics"

type statisticsAction int

const (
	usageHistory statisticsAction = iota
	trends
	dosageTracker
)

// statisticsActions is a list of options for the stats appliance.
var statisticsActions = map[statisticsAction]string{
	usageHistory:  markedText("📅 &Usage History"),
	trends:        markedText("📈 &Trends"),
	dosageTracker: markedText("🔢 &Dosage Tracker")}

// StatisticsAppliance ...
type StatisticsAppliance struct {
	hv *HomeView
}

// NewStatisticsAppliance ...
func NewStatisticsAppliance() *StatisticsAppliance {
	s := &StatisticsAppliance{
		hv: NewHomeView(),
	}
	s.hv.Title(breadcrumbTitle(s.hv.title, statisticsTitle))
	return s
}

// Init ...
func (s *StatisticsAppliance) Init() tea.Cmd {
	return nil
}

// Update ...
func (s *StatisticsAppliance) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
func (s *StatisticsAppliance) View() string {
	return s.hv.View()
}
