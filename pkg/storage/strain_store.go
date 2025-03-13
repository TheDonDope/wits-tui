package storage

import (
	"errors"
	"fmt"
	"os"
	"sync"

	can "github.com/TheDonDope/wits-tui/pkg/cannabis"
	"gopkg.in/yaml.v3"
)

var strainsFile = fmt.Sprintf("%s/strains.yml", os.Getenv("WITS_DIR"))

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

// StrainStoreInMemory is the in memory implementation of the StrainStore interface.
type StrainStoreInMemory struct {
	mu      sync.Mutex
	strains map[string]*can.Strain
}

// AddStrain adds a strain to the store, using its product name as the key.
func (ssim *StrainStoreInMemory) AddStrain(s *can.Strain) error {
	ssim.mu.Lock()
	defer ssim.mu.Unlock()

	if _, exists := ssim.strains[s.Strain]; exists {
		return ErrStrainAlreadyExists
	}
	ssim.strains[s.Strain] = s
	return nil
}

// GetStrains returns all strains in the store as a slice.
func (ssim *StrainStoreInMemory) GetStrains() []*can.Strain {
	ssim.mu.Lock()
	defer ssim.mu.Unlock()

	var strains []*can.Strain
	for _, s := range ssim.strains {
		strains = append(strains, s)
	}
	return strains
}

// FindStrainByProduct finds a strain in the store by product name.
func (ssim *StrainStoreInMemory) FindStrainByProduct(p string) (*can.Strain, error) {
	ssim.mu.Lock()
	defer ssim.mu.Unlock()

	strain, exists := ssim.strains[p]
	if !exists {
		return nil, ErrStrainNotFound
	}
	return strain, nil
}

// StrainStoreYMLFile is the yaml file storage implementation of the StrainStore
// interface.
type StrainStoreYMLFile struct {
	mu      sync.Mutex
	strains map[string]*can.Strain
}

// AddStrain adds a strain to the store, using its product name as the key.
func (ssyf *StrainStoreYMLFile) AddStrain(s *can.Strain) error {
	ssyf.mu.Lock()
	defer ssyf.mu.Unlock()

	if _, exists := ssyf.strains[s.Strain]; exists {
		return ErrStrainAlreadyExists
	}
	ssyf.strains[s.Strain] = s

	data, err := yaml.Marshal(ssyf.strains)
	if err != nil {
		return err
	}
	return os.WriteFile(strainsFile, data, 0644)
}

// GetStrains returns all strains in the store as a slice.
func (ssyf *StrainStoreYMLFile) GetStrains() []*can.Strain {
	ssyf.mu.Lock()
	defer ssyf.mu.Unlock()

	var strains []*can.Strain
	for _, s := range ssyf.strains {
		strains = append(strains, s)
	}
	return strains
}

// FindStrainByProduct finds a strain in the store by product name.
func (ssyf *StrainStoreYMLFile) FindStrainByProduct(p string) (*can.Strain, error) {
	ssyf.mu.Lock()
	defer ssyf.mu.Unlock()

	strain, exists := ssyf.strains[p]
	if !exists {
		return nil, ErrStrainNotFound
	}
	return strain, nil
}

// NewStrainStore returns a new StrainStore implementation depending on the
// configured storage mode in the environment variable.
func NewStrainStore() StrainStore {
	storageMode := os.Getenv("STORAGE_MODE")
	switch storageMode {
	case StoreInMemory:
		return &StrainStoreInMemory{}
	case StoreYMLFile:
		ssyf := &StrainStoreYMLFile{}
		data, err := os.ReadFile(strainsFile)
		if err != nil {
			if os.IsNotExist(err) {
				return ssyf
			}
		}
		err = yaml.Unmarshal(data, ssyf.strains)
		if err != nil {
			return nil
		}
		return ssyf
	}
	return nil
}
