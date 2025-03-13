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

var statisticsActions = map[statisticsAction]string{
	usageHistory:  markedText("📅 &Usage History"),
	trends:        markedText("📈 &Trends"),
	dosageTracker: markedText("🔢 &Dosage Tracker")}

// StatisticsAppliance is the tea.Model for the Statistics appliance.
type StatisticsAppliance struct {
	hm *HomeModel
}

// NewStatisticsAppliance returns a new StatisticsAppliance, with the following contents:
//   - rendered title
func NewStatisticsAppliance() *StatisticsAppliance {
	s := &StatisticsAppliance{
		hm: initialHomeModel(),
	}
	s.hm.Title(breadcrumbTitle(s.hm.title, statisticsTitle))
	return s
}

// StatisticsAppliance implementation of tea.Model interface -------------------

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (s *StatisticsAppliance) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (s *StatisticsAppliance) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
	hv, cmd := s.hm.Update(msg)
	s.hm = hv.(*HomeModel)
	return s, cmd
}

// View renders the StatisticsAppliance UI, which is just a string. The view is
// rendered after every Update.
func (s *StatisticsAppliance) View() string {
	return s.hm.View()
}
