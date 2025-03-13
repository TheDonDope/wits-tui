package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	maxWidth  = 120
	homeTitle = "🥦 Wits"
)

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

// HomeModelBuilder builds home models from a set of given mandatory and optional components:
// title, list (table), list buttons, preview.
type HomeModelBuilder interface {
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

// HomeModel implements both tui.HomeModelBuilder and tea.Model interfaces to
// act as a base for the different appliances.
type HomeModel struct {
	lg     *lipgloss.Renderer
	styles *Styles
	width  int

	title string

	listView   tea.Model
	listBar    tea.Model
	listExtras tea.Model
	preview    tea.Model
}

// initialHomeModel returns a new HomeModel with empty content.
func initialHomeModel() *HomeModel {
	m := &HomeModel{width: maxWidth, title: homeTitle}
	m.lg = lipgloss.DefaultRenderer()
	m.styles = NewStyles(m.lg)
	return m
}

// HomeModel implementation of tui.HomeModelBuilder interface --------------------

// Title sets the title to render
func (hm *HomeModel) Title(t string) {
	hm.title = t
}

// List sets the list to render
func (hm *HomeModel) List(l tea.Model) {
	hm.listView = l
}

// Bar sets the list bar to render
func (hm *HomeModel) Bar(b tea.Model) {
	hm.listBar = b
}

// Extras sets the extras to render
func (hm *HomeModel) Extras(e tea.Model) {
	hm.listExtras = e
}

// Preview sets the preview to render
func (hm *HomeModel) Preview(p tea.Model) {
	hm.preview = p
}

// HomeModel implementation of tea.Model interface ------------------------------

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (hm *HomeModel) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (hm *HomeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q", "ctrl+c":
			return hm, tea.Quit
		}
	}
	var cmds []tea.Cmd

	return hm, tea.Batch(cmds...)
}

// View renders the HomeModel, which is just a string. The view is
// rendered after every Update.
func (hm *HomeModel) View() string {
	s := hm.styles

	header := hm.appBoundaryView(hm.title)

	body := lipgloss.JoinVertical(lipgloss.Left, hm.decoratedList(), hm.decoratedListBarAndExtras(), hm.decoratedPreview())

	return s.Base.Render(header + "\n" + body)
}

// appBoundaryView returns boundary view for the application with the given text.
func (hm *HomeModel) appBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		hm.width,
		lipgloss.Left,
		hm.styles.HeaderText.Render(text),
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceForeground(indigo),
	)
}

// decoratedList returns the rendered list view, if existing, or otherwise empty
// content.
func (hm *HomeModel) decoratedList() string {
	if hm.listView == nil {
		return "\n\n"
	}
	return hm.listView.View() + "\n\n"
}

// decoratedListBarAndExtras returns the rendered list bar and extras, depending
// on their existence.
func (hm *HomeModel) decoratedListBarAndExtras() string {
	var b strings.Builder
	if hm.listBar != nil {
		b.WriteString(hm.listBar.View())
	}
	if hm.listExtras != nil {
		b.WriteString(hm.listExtras.View())
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, b.String()) + "\n\n"
}

// decoratedPreview returns the rendered preview, if existing, or otherwise
// empty content
func (hm *HomeModel) decoratedPreview() string {
	if hm.preview == nil {
		return ""
	}
	return hm.preview.View() + "\n\n"
}

// markedText returns an string with its marked character (denoted by an `&`)
// underlined by using ANSI escape codes
func markedText(s string) string {
	if idx := strings.Index(s, "&"); idx != -1 && idx+1 < len(s) {
		return fmt.Sprintf("%s\033[4m%s\033[0m%s", s[:idx], string(s[idx+1]), s[idx+2:])
	}
	return s
}

// breadcrumTitle returns a breadcrumb navigation representation of all given nodes,
// separated by a ` > `.
func breadcrumbTitle(nodes ...string) string {
	var b strings.Builder
	for i, node := range nodes {
		b.WriteString(node)
		if i < len(nodes)-1 {
			b.WriteString(" > ")
		}
	}
	return b.String()
}
