package tui

import tea "github.com/charmbracelet/bubbletea"

const statsTitle = "Statistics"

type statsAction int

const (
	usageHistory statsAction = iota
	trends
	dosageTracker
)

// statsActions is a list of options for the stats appliance.
var statsActions = map[statsAction]string{
	usageHistory:  "📅 Usage History",
	trends:        "📈 Trends",
	dosageTracker: "🔢 Dosage Tracker"}

// StatsAppliance ...
type StatsAppliance struct {
	hv *HomeView
}

// NewStatsAppliance ...
func NewStatsAppliance() *StatsAppliance {
	s := &StatsAppliance{
		hv: NewHomeView(),
	}
	return s
}

// Init ...
func (s *StatsAppliance) Init() tea.Cmd {
	s.hv.title = statsTitle
	return nil
}

// Update ...
func (s *StatsAppliance) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
func (s *StatsAppliance) View() string {
	return s.hv.View()
}
