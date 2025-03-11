package storage

import (
	"errors"
	"os"
	"sync"

	can "github.com/TheDonDope/wits-tui/pkg/cannabis"
	"gopkg.in/yaml.v3"
)

const strainsFile = ".wits/strains.yml"

var (
	// ErrStrainNotFound is returned when a strain is not found in the store.
	ErrStrainNotFound = errors.New("Strain with that product name not found")
	// ErrStrainAlreadyExists is returned when a strain with the same product name already exists in the store.
	ErrStrainAlreadyExists = errors.New("Strain with that product name already exists")
)

// StrainStore is an interface for storing strains.
type StrainStore interface {
	AddStrain(s *can.Strain) error
	GetStrains() []*can.Strain
	FindStrainByProduct(p string) (*can.Strain, error)
}

// StrainStoreType is a holder for a collection of strains.
type StrainStoreType struct {
	Persistance PersistanceType
	mu          sync.Mutex
	Strains     map[string]*can.Strain
}

// NewStrainStoreInMemory creates a new in-memory Strain store.
func NewStrainStoreInMemory() *StrainStoreType {
	return &StrainStoreType{
		Persistance: InMemory,
		Strains:     make(map[string]*can.Strain),
	}
}

// NewStrainStoreYMLFile creates a new Strain store which persists data to a yml file in the .wits directory.
func NewStrainStoreYMLFile() (*StrainStoreType, error) {
	store := &StrainStoreType{
		Persistance: YMLFile,
		Strains:     make(map[string]*can.Strain),
	}
	data, err := ReadFile(strainsFile)
	if err != nil {
		if os.IsNotExist(err) {
			return store, nil
		}
		return nil, err
	}
	err = yaml.Unmarshal(data, store.Strains)
	if err != nil {
		return nil, err
	}
	return store, nil
}

// AddStrain adds a strain to the store, using its product name as the key.
func (sstr *StrainStoreType) AddStrain(s *can.Strain) error {
	sstr.mu.Lock()
	defer sstr.mu.Unlock()

	if _, exists := sstr.Strains[s.Strain]; exists {
		return ErrStrainAlreadyExists
	}
	sstr.Strains[s.Strain] = s

	if sstr.Persistance == YMLFile {
		data, err := yaml.Marshal(sstr.Strains)
		if err != nil {
			return err
		}
		return WriteFile(strainsFile, data)
	}

	return nil
}

// GetStrains returns all strains in the store as a slice.
func (sstr *StrainStoreType) GetStrains() []*can.Strain {
	sstr.mu.Lock()
	defer sstr.mu.Unlock()

	var strains []*can.Strain
	for _, s := range sstr.Strains {
		strains = append(strains, s)
	}
	return strains
}

// FindStrainByProduct finds a strain in the store by product name.
func (sstr *StrainStoreType) FindStrainByProduct(p string) (*can.Strain, error) {
	sstr.mu.Lock()
	defer sstr.mu.Unlock()

	strain, exists := sstr.Strains[p]
	if !exists {
		return nil, ErrStrainNotFound
	}
	return strain, nil
}
