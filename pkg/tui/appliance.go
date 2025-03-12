package tui

import tea "github.com/charmbracelet/bubbletea"

// HomeViewBuilder builds home views from a set of given mandatory and optional components:
// title, list (table), list buttons, preview.
type HomeViewBuilder interface {
	// Title sets the title to render
	Title(t string)
	// List sets the list to render
	List(l tea.Model)
	// Bar sets the list bar to render
	Bar(b tea.Model)
	// Extras sets the extras to render
	Extras(e tea.Model)
	// Preview sets the preview to render
	Preview(p tea.Model)
	// Panel returns the built panel
	Panel() tea.Model
}

// HomeView holds the default implementation for the HomeViewBuilder
type HomeView struct {
	title      string
	titleLabel tea.Model

	listView   tea.Model
	listBar    tea.Model
	listExtras tea.Model
	preview    tea.Model

	panel tea.Model
}

// NewHomeView returns a new HomeView with empty content.
func NewHomeView() *HomeView {
	return &HomeView{}
}

// Title sets the title to render
func (hv *HomeView) Title(t string) {
	hv.title = t
}

// List sets the list to render
func (hv *HomeView) List(l tea.Model) {
	hv.listView = l
}

// Bar sets the list bar to render
func (hv *HomeView) Bar(b tea.Model) {
	hv.listBar = b
}

// Extras sets the extras to render
func (hv *HomeView) Extras(e tea.Model) {
	hv.listExtras = e
}

// Preview sets the preview to render
func (hv *HomeView) Preview(p tea.Model) {
	hv.preview = p
}

// Panel returns the built panel
func (hv *HomeView) Panel() tea.Model {
	if hv.panel == nil {
		hv.panel = hv.buildPanel()
	}
	return hv.panel
}

func (hv *HomeView) buildPanel() tea.Model {
	return nil
}
