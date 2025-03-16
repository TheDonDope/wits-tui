package tui

import (
	"fmt"
	"log"
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
	sortedGenetics  = sortedGeneticsList()
	sortedRadiation = sortedRadiationList()
	sortedTerpenes  = sortedTerpenesList()
)

type strainsListedMsg struct {
	items []list.Item
}

type strainSubmittedMsg struct {
	strain *can.Strain
}

// StrainsHomeModel is the tea.Model for the Strains appliance
type StrainsHomeModel struct {
	hm      *HomeModel
	service service.StrainService
}

// initialStrainsHomeModel returns a new StrainsHomeModel, with the following contents:
//   - rendered title
func initialStrainsHomeModel() *StrainsHomeModel {
	log.Println("💬 💾  (pkg/tui/strains.go) initialStrainsHomeModel()")
	s := &StrainsHomeModel{
		hm:      initialHomeModel(),
		service: service.NewStrainService(storage.NewStrainStore()),
	}
	s.hm.Title(breadcrumbTitle(s.hm.title, strainsTitle))
	s.hm.List(initialStrainListModel())
	return s
}

// StrainsHomeModel implementation of tea.Model interface ----------------------

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (shm *StrainsHomeModel) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (shm *StrainsHomeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return shm, tea.Quit
		case "esc":
			return InitialMenuModel(), nil
		case "alt+n", "ctrl+n":
			return shm, onStrainAdded()
		}
	case strainSubmittedMsg:
		shm.service.AddStrain(msg.strain)
		// TODO: redirect to home view?
		return shm, shm.onStrainsListed()
	case strainsListedMsg:
		return shm.hm.Update(msg)
	}

	var cmd tea.Cmd
	hm, cmd := shm.hm.Update(msg)
	shm.hm = hm.(*HomeModel)
	return shm, cmd
}

// View renders the StrainsHomeModel UI, which is just a string. The view is
// rendered after every Update.
func (shm *StrainsHomeModel) View() string {
	return shm.hm.View()
}

// onStrainsListed retrieves all strains from the service and returns a message
// containing the results as an slice of list items.
func (shm *StrainsHomeModel) onStrainsListed() tea.Cmd {
	return func() tea.Msg {
		items := []list.Item{}
		strains := shm.service.GetStrains()

		if len(strains) == 0 {
			items = append(items, StrainListItem{value: &can.Strain{
				Strain:   "No strains available, press alt+n to create a new one.",
				Cultivar: "",
				THC:      0,
				CBD:      0,
			}})
		} else {
			for _, strain := range strains {
				items = append(items, StrainListItem{value: strain})
			}
		}
		return strainsListedMsg{items}
	}
}

// onStrainAdded runs the form to add a strain and on submission sends a message
// with the parsed strain data from the form.
func onStrainAdded() tea.Cmd {
	form := initialStrainForm()

	if err := form.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running strain creation form: %v\n", err)
		return nil // Return nil to prevent further processing
	}

	strain := parseStrain(form)
	if strain == nil {
		return nil // Do nothing if the user canceled or input was invalid
	}
	return func() tea.Msg { return strainSubmittedMsg{strain} }

}

// sortedGeneticsList returns a list of genetic options for the user to choose from.
func sortedGeneticsList() []huh.Option[can.GeneticType] {
	var genetics []huh.Option[can.GeneticType]
	for k, v := range can.Genetics {
		genetics = append(genetics, huh.NewOption(v, k))
	}
	sort.Slice(genetics, func(i, j int) bool {
		return genetics[i].Value < genetics[j].Value
	})
	return genetics
}

// sortedRadiationList returns a list of options for radiation treatment for the user to choose from.
func sortedRadiationList() []huh.Option[bool] {
	var radiations []huh.Option[bool]
	radiations = append(radiations, huh.NewOption("No", false))
	radiations = append(radiations, huh.NewOption("Yes", true))
	sort.Slice(radiations, func(i, j int) bool {
		// We consider `false`` to be "smaller" than `true``, so we put `false`` before `true`
		return !radiations[i].Value && radiations[j].Value
	})
	return radiations
}

// sortedTerpenesList returns a list of terpene options for the user to choose from.
func sortedTerpenesList() []huh.Option[*can.Terpene] {
	var terpenes []huh.Option[*can.Terpene]
	for _, t := range can.Terpenes {
		terpenes = append(terpenes, huh.NewOption(t.Name, t))
	}
	sort.Slice(terpenes, func(i, j int) bool {
		return terpenes[i].Value.Name < terpenes[j].Value.Name
	})
	return terpenes
}

// initialStrainForm returns a form for creating a new strain.
func initialStrainForm() *huh.Form {
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
				Options(sortedGenetics...).
				Title("Genetic").
				Description("The phenotype"),

			huh.NewSelect[bool]().
				Key("radiated").
				Options(sortedRadiation...).
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
				Options(sortedTerpenes...).
				Title("Terpenes").
				Description("The contained terpenes"),

			huh.NewInput().
				Key("amount").
				Title("Amount (g)").
				Description("The weight"),
		),
	)
}

// parseStrain creates a new strain entity from the given form data.
func parseStrain(form *huh.Form) *can.Strain {
	thc := parseFloatWithDefault(form.GetString("thc"), 0)
	cbd := parseFloatWithDefault(form.GetString("cbd"), 0)
	amount := parseFloatWithDefault(form.GetString("amount"), 0)

	// Handle potential nil values
	var genetic can.GeneticType
	if val, ok := form.Get("genetic").(can.GeneticType); ok {
		genetic = val
	}

	var terpenes []*can.Terpene
	if val, ok := form.Get("terpenes").([]*can.Terpene); ok {
		terpenes = val
	}

	return &can.Strain{
		ID:           uuid.New(),
		Strain:       form.GetString("strain"),
		Cultivar:     form.GetString("cultivar"),
		Manufacturer: form.GetString("manufacturer"),
		Country:      form.GetString("country"),
		Genetic:      genetic,
		Radiated:     form.GetBool("radiated"),
		THC:          thc,
		CBD:          cbd,
		Terpenes:     terpenes,
		Amount:       amount,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

// parseFloatWithDefault parses the given input to a float64. If an error occurs
// the given defaultValue is returned.
func parseFloatWithDefault(input string, defaultValue float64) float64 {
	if val, err := strconv.ParseFloat(input, 64); err == nil {
		return val
	}
	return defaultValue
}

// StrainListItem is a list item for strains.
type StrainListItem struct {
	value *can.Strain
}

// StrainListItem implementation of list.Item interface ------------------------

// FilterValue is the value we use when filtering against this item when
// we're filtering the list.
func (sli StrainListItem) FilterValue() string {
	return sli.value.Cultivar
}

// Title returns the title for the list item.
func (sli StrainListItem) Title() string {
	return sli.value.Strain
}

// Description returns the description for the list item.
func (sli StrainListItem) Description() string {
	return fmt.Sprintf("Amount: %.1f g, THC/CBD: %.1f%% / %.1f%%, Genetic: %s", sli.value.Amount, sli.value.THC, sli.value.CBD, can.Genetics[sli.value.Genetic])
}

// StrainListModel is a tea.Model for the strains list.
type StrainListModel struct {
	list list.Model
}

// initialStrainListModel creates a new model for the strains list, without any
// items.
func initialStrainListModel() *StrainListModel {
	log.Println("💬 💾  (pkg/tui/strains.go) initialStrainListModel()")
	l := list.New([]list.Item{}, list.NewDefaultDelegate(), 60, 30)
	l.Title = "Entries"
	log.Printf("✅ 💾  (pkg/tui/strains.go) initialStrainListModel() -> len(l.Items()): %v \n", len(l.Items()))
	return &StrainListModel{list: l}
}

// StrainListModel implementation of tea.Model interface -----------------------

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (slm *StrainListModel) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (slm *StrainListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return InitialMenuModel(), nil
		case "alt+n", "ctrl+n":
			return slm, onStrainAdded()
		}

	case strainsListedMsg:
		return slm, slm.list.SetItems(msg.items)
	}

	var cmd tea.Cmd
	slm.list, cmd = slm.list.Update(msg)
	return slm, cmd
}

// View renders the StrainListModel UI, which is just a string. The view is
// rendered after every Update.
func (slm *StrainListModel) View() string {
	return slm.list.View()
}
