package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// MaxWidth defines the maximum boundaries the application has to adhere to
const MaxWidth = 120

var (
	red    = lipgloss.AdaptiveColor{Light: "#FE5F86", Dark: "#FE5F86"}
	indigo = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
	green  = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
)

// Styles defines the used lipgloss styles
type Styles struct {
	Base,
	HeaderText,
	Status,
	StatusHeader,
	Highlight,
	ErrorHeaderText,
	Help lipgloss.Style
}

// NewStyles returns the new configured styles using the given renderer.
func NewStyles(lg *lipgloss.Renderer) *Styles {
	s := Styles{}
	s.Base = lg.NewStyle().
		Padding(1, 4, 0, 1)
	s.HeaderText = lg.NewStyle().
		Foreground(indigo).
		Bold(true).
		Padding(0, 1, 0, 2)
	s.Status = lg.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(indigo).
		PaddingLeft(1).
		MarginTop(1)
	s.StatusHeader = lg.NewStyle().
		Foreground(green).
		Bold(true)
	s.Highlight = lg.NewStyle().
		Foreground(lipgloss.Color("212"))
	s.ErrorHeaderText = s.HeaderText.
		Foreground(red)
	s.Help = lg.NewStyle().
		Foreground(lipgloss.Color("240"))
	return &s
}

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
}

// HomeView holds the default implementation for the HomeViewBuilder
type HomeView struct {
	lg     *lipgloss.Renderer
	styles *Styles
	width  int

	title string

	listView   tea.Model
	listBar    tea.Model
	listExtras tea.Model
	preview    tea.Model
}

// NewHomeView returns a new HomeView with empty content.
func NewHomeView() *HomeView {
	m := &HomeView{width: MaxWidth}
	m.lg = lipgloss.DefaultRenderer()
	m.styles = NewStyles(m.lg)
	return m
}

// HomeViewBuilder interface ---------------------------------------------------

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

// tea.Model interface ---------------------------------------------------------

// Init ...
func (hv *HomeView) Init() tea.Cmd {
	return nil
}

// Update ...
func (hv *HomeView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q", "ctrl+c":
			return hv, tea.Quit
		}
	}
	var cmds []tea.Cmd

	return hv, tea.Batch(cmds...)
}

// View ...
func (hv *HomeView) View() string {
	s := hv.styles

	header := hv.appBoundaryView(hv.title)

	body := lipgloss.JoinVertical(lipgloss.Left, hv.decoratedList(), hv.decoratedListBarAndExtras(), hv.decoratedPreview())

	return s.Base.Render(header + "\n" + body)
}

func (hv *HomeView) appBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		hv.width,
		lipgloss.Left,
		hv.styles.HeaderText.Render(text),
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceForeground(indigo),
	)
}

func (hv *HomeView) decoratedList() string {
	if hv.listView == nil {
		return "\n\n"
	}
	return hv.listView.View() + "\n\n"
}

func (hv *HomeView) decoratedListBarAndExtras() string {
	var builder []string
	if hv.listBar != nil {
		builder = append(builder, hv.listBar.View())
	}
	if hv.listExtras != nil {
		builder = append(builder, hv.listExtras.View())
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, builder...) + "\n\n"
}

func (hv *HomeView) decoratedPreview() string {
	if hv.preview == nil {
		return ""
	}
	return hv.preview.View() + "\n\n"
}

// markedText returns an string with its marked character (denoted by an `&`)
// underlined by using ANSI escape codes
func markedText(s string) string {
	if idx := strings.Index(s, "&"); idx != -1 && idx+1 < len(s) {
		return fmt.Sprintf("%s\033[4m%s\033[0m%s", s[:idx], string(s[idx+1]), s[idx+2:])
	}
	return s
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
