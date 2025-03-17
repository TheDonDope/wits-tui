package service

import (
	"io"
	"log"
	"os"
	"testing"
	"time"

	can "github.com/TheDonDope/wits-tui/pkg/cannabis"
	"github.com/TheDonDope/wits-tui/pkg/storage"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestMain handles global test setup
func TestMain(m *testing.M) {
	// Disable log output during tests
	log.SetOutput(io.Discard)
	os.Exit(m.Run())
}

// testStrain generates a consistent test strain with fixed values
func testStrain() *can.Strain {
	testUUID := uuid.MustParse("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	testTime := time.Date(2023, time.October, 5, 12, 0, 0, 0, time.UTC)

	return &can.Strain{
		ID:           testUUID,
		Strain:       "Test Strain",
		Cultivar:     "Test Cultivar",
		Manufacturer: "Test Manufacturer",
		Country:      "Test Country",
		Genetic:      can.Sativa,
		Radiated:     false,
		THC:          20.0,
		CBD:          0.5,
		Terpenes:     []*can.Terpene{},
		Amount:       3.5,
		CreatedAt:    testTime,
		UpdatedAt:    testTime,
	}
}

// mockStrainStore implements storage.StrainStore for testing
type mockStrainStore struct {
	addStrainCalls []*can.Strain
	addStrainErr   error

	getStrainsCalls  int
	getStrainsResult []*can.Strain

	findStrainByProductCalls  []string
	findStrainByProductResult *can.Strain
	findStrainByProductErr    error
}

func (m *mockStrainStore) AddStrain(s *can.Strain) error {
	m.addStrainCalls = append(m.addStrainCalls, s)
	return m.addStrainErr
}

func (m *mockStrainStore) GetStrains() []*can.Strain {
	m.getStrainsCalls++
	return m.getStrainsResult
}

func (m *mockStrainStore) FindStrainByProduct(p string) (*can.Strain, error) {
	m.findStrainByProductCalls = append(m.findStrainByProductCalls, p)
	return m.findStrainByProductResult, m.findStrainByProductErr
}

func TestStrainService(t *testing.T) {
	t.Run("AddStrain", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			store := &mockStrainStore{}
			svc := NewStrainService(store)
			strain := testStrain()

			err := svc.AddStrain(strain)

			require.NoError(t, err)
			assert.Len(t, store.addStrainCalls, 1)
			assert.Equal(t, strain, store.addStrainCalls[0])
		})

		t.Run("Error", func(t *testing.T) {
			store := &mockStrainStore{
				addStrainErr: storage.ErrStrainAlreadyExists,
			}
			svc := NewStrainService(store)
			strain := testStrain()

			err := svc.AddStrain(strain)

			assert.ErrorIs(t, err, storage.ErrStrainAlreadyExists)
			assert.Len(t, store.addStrainCalls, 1)
		})
	})

	t.Run("GetStrains", func(t *testing.T) {
		t.Run("Empty", func(t *testing.T) {
			store := &mockStrainStore{}
			svc := NewStrainService(store)

			result := svc.GetStrains()

			assert.Empty(t, result)
			assert.Equal(t, 1, store.getStrainsCalls)
		})

		t.Run("WithResults", func(t *testing.T) {
			expected := []*can.Strain{testStrain()}
			store := &mockStrainStore{
				getStrainsResult: expected,
			}
			svc := NewStrainService(store)

			result := svc.GetStrains()

			assert.Equal(t, expected, result)
			assert.Equal(t, 1, store.getStrainsCalls)
		})
	})

	t.Run("FindStrainByProduct", func(t *testing.T) {
		t.Run("Found", func(t *testing.T) {
			expected := testStrain()
			store := &mockStrainStore{
				findStrainByProductResult: expected,
			}
			svc := NewStrainService(store)
			productName := "Test Strain"

			result, err := svc.FindStrainByProduct(productName)

			require.NoError(t, err)
			assert.Equal(t, expected, result)
			assert.Len(t, store.findStrainByProductCalls, 1)
			assert.Equal(t, productName, store.findStrainByProductCalls[0])
		})

		t.Run("NotFound", func(t *testing.T) {
			store := &mockStrainStore{
				findStrainByProductErr: storage.ErrStrainNotFound,
			}
			svc := NewStrainService(store)
			productName := "Non-existent Strain"

			_, err := svc.FindStrainByProduct(productName)

			assert.ErrorIs(t, err, storage.ErrStrainNotFound)
			assert.Len(t, store.findStrainByProductCalls, 1)
		})
	})
}
