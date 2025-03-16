package service

import (
	"fmt"
	"log"

	can "github.com/TheDonDope/wits-tui/pkg/cannabis"
	"github.com/TheDonDope/wits-tui/pkg/storage"
)

// StrainService provides operations on strains.
type StrainService interface {
	AddStrain(s *can.Strain) error
	GetStrains() []*can.Strain
	FindStrainByProduct(p string) (*can.Strain, error)
}

// StrainServiceType provides operations on strains, accessing a store.
type StrainServiceType struct {
	store storage.StrainStore
}

// NewStrainService creates a new service layer for strains.
func NewStrainService(s storage.StrainStore) *StrainServiceType {
	log.Println("💬 🤝  (pkg/svc/strain.go) NewStrainService(s storage.StrainStore)")
	log.Println(fmt.Sprintf("✅ 🤝  (s): %v", s))
	return &StrainServiceType{store: s}
}

// AddStrain adds a strain to the store.
func (svc *StrainServiceType) AddStrain(s *can.Strain) error {
	log.Println("💬 🤝  (pkg/svc/strain.go) AddStrain(s *can.Strain)")
	log.Println(fmt.Sprintf("✅ 🤝  (s): %v", s))
	return svc.store.AddStrain(s)
}

// GetStrains retrieves all strains from the store.
func (svc *StrainServiceType) GetStrains() []*can.Strain {
	log.Println("💬 🤝  (pkg/svc/strain.go) GetStrains()")
	return svc.store.GetStrains()
}

// FindStrainByProduct looks up a strain by its prodcut name.
func (svc *StrainServiceType) FindStrainByProduct(p string) (*can.Strain, error) {
	log.Println("💬 🤝  (pkg/svc/strain.go) FindStrainByProduct(p string)")
	log.Println(fmt.Sprintf("✅ 🤝  (p): %v", p))
	return svc.store.FindStrainByProduct(p)
}
