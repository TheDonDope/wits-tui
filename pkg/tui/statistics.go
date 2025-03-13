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

// StatisticsHomeModel is the tea.Model for the Statistics appliance.
type StatisticsHomeModel struct {
	hm *HomeModel
}

// initialStatisticsHomeModel returns a new StatisticsHomeModel, with the following contents:
//   - rendered title
func initialStatisticsHomeModel() *StatisticsHomeModel {
	s := &StatisticsHomeModel{
		hm: initialHomeModel(),
	}
	s.hm.Title(breadcrumbTitle(s.hm.title, statisticsTitle))
	return s
}

// StatisticsHomeModel implementation of tea.Model interface -------------------

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (shm *StatisticsHomeModel) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (shm *StatisticsHomeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
	hm, cmd := shm.hm.Update(msg)
	shm.hm = hm.(*HomeModel)
	return shm, cmd
}

// View renders the StatisticsHomeModel UI, which is just a string. The view is
// rendered after every Update.
func (shm *StatisticsHomeModel) View() string {
	return shm.hm.View()
}
