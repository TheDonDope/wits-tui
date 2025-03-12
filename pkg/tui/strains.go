package tui

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"

	can "github.com/TheDonDope/wits-tui/pkg/cannabis"
	"github.com/TheDonDope/wits-tui/pkg/service"
	"github.com/TheDonDope/wits-tui/pkg/storage"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/google/uuid"
)

const strainsTitle = "🌿 Strains"

type strainsAction int

const (
	addStrain strainsAction = iota
	viewStrain
	editStrain
	deleteStrain
)

var strainsActions = map[strainsAction]string{
	addStrain:    markedText("➕ &Add Strain"),
	viewStrain:   markedText("📋 &View Strains"),
	editStrain:   markedText("✏️ &Edit Strain"),
	deleteStrain: markedText("❌ &Delete Strain")}

var (
	strainStore   storage.StrainStore
	strainService service.StrainService
)

// StrainsAppliance is the tea.Model for the Strains appliance
type StrainsAppliance struct {
	hv *HomeView
}

// NewStrainsAppliance returns a new StrainsAppliance, with the following contents:
//   - rendered title
func NewStrainsAppliance() *StrainsAppliance {
	sstr, err := storage.NewStrainStoreYMLFile()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading data from YML File: %v", err)
		os.Exit(1)
	}
	strainStore = sstr
	strainService = service.NewStrainService(strainStore)
	s := &StrainsAppliance{
		hv: NewHomeView(),
	}
	s.hv.Title(breadcrumbTitle(s.hv.title, strainsTitle))
	//s.hv.List(ListStrains())
	return s
}

// StrainsAppliance implementation of tea.Model interface ----------------------

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (s *StrainsAppliance) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (s *StrainsAppliance) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
	hv, cmd := s.hv.Update(msg)
	s.hv = hv.(*HomeView)
	return s, cmd
}

// View renders the StrainsAppliance UI, which is just a string. The view is
// rendered after every Update.
func (s *StrainsAppliance) View() string {
	return s.hv.View()
}

// geneticsOptions returns a list of genetic options for the user to choose from.
func geneticsOptions() []huh.Option[can.GeneticType] {
	var genetics []huh.Option[can.GeneticType]
	for k, v := range can.Genetics {
		genetics = append(genetics, huh.NewOption(v, k))
	}
	sort.Slice(genetics, func(i, j int) bool {
		return genetics[i].Value < genetics[j].Value
	})
	return genetics
}

// radiationOptions returns a list of options for radiation treatment for the user to choose from.
func radiationOptions() []huh.Option[bool] {
	var radiations []huh.Option[bool]
	radiations = append(radiations, huh.NewOption("No", false))
	radiations = append(radiations, huh.NewOption("Yes", true))
	sort.Slice(radiations, func(i, j int) bool {
		// We consider `false`` to be "smaller" than `true``, so we put `false`` before `true`
		return !radiations[i].Value && radiations[j].Value
	})
	return radiations
}

// terpeneOptions returns a list of terpene options for the user to choose from.
func terpeneOptions() []huh.Option[*can.Terpene] {
	var terpenes []huh.Option[*can.Terpene]
	for _, t := range can.Terpenes {
		terpenes = append(terpenes, huh.NewOption(t.Name, t))
	}
	sort.Slice(terpenes, func(i, j int) bool {
		return terpenes[i].Value.Name < terpenes[j].Value.Name
	})
	return terpenes
}

// newStrainForm returns a form for creating a new strain.
func newStrainForm() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Key("strain").
				Title("Strain").
				Description("The product name"),

			huh.NewInput().
				Key("cultivar").
				Title("Cultivar").
				Description("The plant name"),

			huh.NewInput().
				Key("manufacturer").
				Title("Manufacturer").
				Description("The producing company"),

			huh.NewInput().
				Key("country").
				Title("Country").
				Description("The country of origin"),

			huh.NewSelect[can.GeneticType]().
				Key("genetic").
				Options(geneticsOptions()...).
				Title("Genetic").
				Description("The phenotype"),

			huh.NewSelect[bool]().
				Key("radiated").
				Options(radiationOptions()...).
				Title("Radiated").
				Description("If the plant was radiation treated"),

			huh.NewInput().
				Key("thc").
				Title("THC (%)").
				Description("The THC content"),

			huh.NewInput().
				Key("cbd").
				Title("CBD (%)").
				Description("The CBD content"),

			huh.NewMultiSelect[*can.Terpene]().
				Key("terpenes").
				Options(terpeneOptions()...).
				Title("Terpenes").
				Description("The contained terpenes"),

			huh.NewInput().
				Key("amount").
				Title("Amount (g)").
				Description("The weight"),
		),
	)
}

// newStrainFromForm creates a new strain entity from the given form data.
func newStrainFromForm(form *huh.Form) *can.Strain {
	thc, err := strconv.ParseFloat(form.GetString("thc"), 64)
	if err != nil {
		thc = 0
	}
	cbd, err := strconv.ParseFloat(form.GetString("cbd"), 64)
	if err != nil {
		cbd = 0
	}
	amount, err := strconv.ParseFloat(form.GetString("amount"), 64)
	if err != nil {
		amount = 0
	}
	return &can.Strain{
		ID:           uuid.New(),
		Strain:       form.GetString("strain"),
		Cultivar:     form.GetString("cultivar"),
		Manufacturer: form.GetString("manufacturer"),
		Country:      form.GetString("country"),
		Genetic:      form.Get("genetic").(can.GeneticType),
		Radiated:     form.GetBool("radiated"),
		THC:          thc,
		CBD:          cbd,
		Terpenes:     form.Get("terpenes").([]*can.Terpene),
		Amount:       amount,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

// AddStrain opens a form for adding a new strain and returns the created strain object.
func AddStrain() tea.Model {
	form := newStrainForm()

	if err := form.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running strain creation form: %v\n", err)
		os.Exit(1)
	}
	strain := newStrainFromForm(form)
	strainService.AddStrain(strain)
	return ListStrains()
}

// StrainsListItem is a list item for strains.
type StrainsListItem struct {
	value *can.Strain
}

// StrainsListItem implementation of list.Item interface -----------------------

// FilterValue is the value we use when filtering against this item when
// we're filtering the list.
func (sli StrainsListItem) FilterValue() string {
	return sli.value.Cultivar
}

// Title returns the title for the list item.
func (sli StrainsListItem) Title() string {
	return sli.value.Strain
}

// Description returns the description for the list item.
func (sli StrainsListItem) Description() string {
	return fmt.Sprintf("Genetic: %s, THC/CBD: %.1f%% %.1f%%", can.Genetics[sli.value.Genetic], sli.value.THC, sli.value.CBD)
}

// StrainsListModel is a tea.Model for the strains list.
type StrainsListModel struct {
	list list.Model
}

// ListStrains creates a new model for the strains list.
func ListStrains() *StrainsListModel {
	items := []list.Item{}
	for _, strain := range strainService.GetStrains() {
		items = append(items, StrainsListItem{value: strain})
	}

	l := list.New(items, list.NewDefaultDelegate(), 60, 30)
	l.Title = "Entries"

	return &StrainsListModel{list: l}
}

// StrainsListModel implementation of tea.Model interface ----------------------

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (slm StrainsListModel) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (slm StrainsListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return slm, tea.Quit
		}
	}

	var cmd tea.Cmd
	slm.list, cmd = slm.list.Update(msg)
	return slm, cmd
}

// View renders the StrainListModel UI, which is just a string. The view is
// rendered after every Update.
func (slm StrainsListModel) View() string {
	return slm.list.View()
}
