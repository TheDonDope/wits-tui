package storage

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sync"

	can "github.com/TheDonDope/wits-tui/pkg/cannabis"
	"gopkg.in/yaml.v3"
)

const strainsFile = "strains.yml"

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
	log.Printf("💬 💾  (pkg/storage/strain_store.go) AddStrain(s *can.Strain: %v) \n", s.ID)
	ssim.mu.Lock()
	defer ssim.mu.Unlock()

	if _, exists := ssim.strains[s.Strain]; exists {
		log.Printf("🚨 💾  (pkg/storage/strain_store.go) 🗒️  Failed to add already existing strain: %v \n", s.ID)
		return ErrStrainAlreadyExists
	}
	ssim.strains[s.Strain] = s
	log.Printf("✅ 💾  (pkg/storage/strain_store.go) AddStrain() -> len(ssim.strains): %v \n", len(ssim.strains))
	return nil
}

// GetStrains returns all strains in the store as a slice.
func (ssim *StrainStoreInMemory) GetStrains() []*can.Strain {
	log.Println("💬 💾  (pkg/storage/strain_store.go) GetStrains()")
	ssim.mu.Lock()
	defer ssim.mu.Unlock()

	var strains []*can.Strain
	for _, s := range ssim.strains {
		strains = append(strains, s)
	}
	log.Printf("✅ 💾  (pkg/storage/strain_store.go) GetStrains() -> len(strains): %v \n", len(strains))
	return strains
}

// FindStrainByProduct finds a strain in the store by product name.
func (ssim *StrainStoreInMemory) FindStrainByProduct(p string) (*can.Strain, error) {
	log.Printf("💬 💾  (pkg/storage/strain_store.go) FindStrainByProduct(p string: %v) \n", p)
	ssim.mu.Lock()
	defer ssim.mu.Unlock()

	strain, exists := ssim.strains[p]
	if !exists {
		log.Printf("🚨 💾  (pkg/storage/strain_store.go) 🗒️  Strain with product name %v does not exist. \n", p)
		return nil, ErrStrainNotFound
	}
	log.Printf("✅ 💾  (pkg/storage/strain_store.go) FindStrainByProduct() -> strain: %v (%v)\n", strain.Strain, strain.ID)
	return strain, nil
}

// String returns a formatted string representation of StrainStoreInMemory.
func (ssim *StrainStoreInMemory) String() string {
	ssim.mu.Lock()
	defer ssim.mu.Unlock()

	var strains []string
	for _, strain := range ssim.strains {
		strains = append(strains, strain.String())
	}

	return fmt.Sprintf("StrainStoreInMemory: len(strains): %v", len(strains))
}

// StrainStoreYMLFile is the yaml file storage implementation of the StrainStore
// interface.
type StrainStoreYMLFile struct {
	mu      sync.Mutex
	strains map[string]*can.Strain
}

// AddStrain adds a strain to the store, using its product name as the key.
func (ssyf *StrainStoreYMLFile) AddStrain(s *can.Strain) error {
	log.Printf("💬 💾  (pkg/storage/strain_store.go) AddStrain(s *can.Strain: %v) \n", s.ID)
	ssyf.mu.Lock()
	defer ssyf.mu.Unlock()

	if _, exists := ssyf.strains[s.Strain]; exists {
		log.Printf("🚨 💾  (pkg/storage/strain_store.go) 🗒️  Failed to add already existing strain: %v \n", s.ID)
		return ErrStrainAlreadyExists
	}
	ssyf.strains[s.Strain] = s

	data, err := yaml.Marshal(ssyf.strains)
	if err != nil {
		log.Printf("🚨 💾  (pkg/storage/strain_store.go) 🗒️  Failed to marshal strain with error: %v \n", err)
		return err
	}
	log.Println("✅ 💾  (pkg/storage/strain_store.go) AddStrain()")
	return os.WriteFile(fmt.Sprintf("%s/%s", os.Getenv("WITS_DIR"), strainsFile), data, 0644)
}

// GetStrains returns all strains in the store as a slice.
func (ssyf *StrainStoreYMLFile) GetStrains() []*can.Strain {
	log.Println("💬 💾  (pkg/storage/strain_store.go) GetStrains()")
	ssyf.mu.Lock()
	defer ssyf.mu.Unlock()

	var strains []*can.Strain
	for _, s := range ssyf.strains {
		strains = append(strains, s)
	}
	log.Printf("✅ 💾  (pkg/storage/strain_store.go) GetStrains() -> len(strains): %v \n", len(strains))
	return strains
}

// FindStrainByProduct finds a strain in the store by product name.
func (ssyf *StrainStoreYMLFile) FindStrainByProduct(p string) (*can.Strain, error) {
	log.Printf("💬 💾  (pkg/storage/strain_store.go) FindStrainByProduct(p string: %v) \n", p)
	ssyf.mu.Lock()
	defer ssyf.mu.Unlock()

	strain, exists := ssyf.strains[p]
	if !exists {
		log.Printf("🚨 💾  (pkg/storage/strain_store.go) 🗒️  Strain with product name %v does not exist. \n", p)
		return nil, ErrStrainNotFound
	}
	log.Printf("✅ 💾  (pkg/storage/strain_store.go) FindStrainByProduct() -> strain: %v (%v) \n", strain.Strain, strain.ID)
	return strain, nil
}

// String returns a formatted string representation of StrainStoreYMLFile.
func (ssyf *StrainStoreYMLFile) String() string {
	ssyf.mu.Lock()
	defer ssyf.mu.Unlock()

	var strains []string
	for _, strain := range ssyf.strains {
		strains = append(strains, strain.String())
	}

	return fmt.Sprintf("StrainStoreYMLFile: len(strains): %v", len(strains))
}

// NewStrainStore returns a new StrainStore implementation depending on the
// configured storage mode in the environment variable.
func NewStrainStore() StrainStore {
	storageMode := os.Getenv("STORAGE_MODE")
	log.Printf("💬 💾  (pkg/storage/strain_store.go) NewStrainStore() -> storageMode: %v \n", storageMode)
	switch storageMode {
	case StoreInMemory:
		return &StrainStoreInMemory{
			strains: make(map[string]*can.Strain),
		}
	case StoreYMLFile:
		ssyf := &StrainStoreYMLFile{
			strains: make(map[string]*can.Strain),
		}
		data, err := os.ReadFile(fmt.Sprintf("%s/%s", os.Getenv("WITS_DIR"), strainsFile))
		if err != nil {
			if os.IsNotExist(err) {
				log.Println("ℹ️  💾  (pkg/storage/strain_store.go) 🗒️  Strain file not existing. Returning new empty store.")
				return ssyf
			}
		}
		err = yaml.Unmarshal(data, ssyf.strains)
		if err != nil {
			log.Printf("🚨 💾  (pkg/storage/strain_store.go) 🗒️  Failed unmarshal strain data with error: %v. Returning new empty store. \n", err)
			return ssyf
		}
		log.Printf("✅ 💾  (pkg/storage/strain_store.go) NewStrainStore() -> store: %v \n", ssyf)
		return ssyf
	}
	return nil
}
