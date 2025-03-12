package storage

import (
	"os"
)

const witsFolder = ".wits"

// PersistanceType is the enum for the different types of data storage.
type PersistanceType int

const (
	// InMemory defines that data is only persistant at runtime and will die with the programs exit.
	InMemory PersistanceType = iota
	// YMLFile defines that data will be loaded and persisted to disk to a specified YML file within the .wits folder.
	YMLFile
)

// EnsureWitsFolder creates the `.wits` directory if missing.
func EnsureWitsFolder() error {
	return os.MkdirAll(witsFolder, os.ModePerm)
}

// ReadFile reads the contents of the file with the given name (path).
func ReadFile(name string) ([]byte, error) {
	if err := EnsureWitsFolder(); err != nil {
		return nil, err
	}
	return os.ReadFile(name)
}

// WriteFile writes the given data to the file by the given name (path).
func WriteFile(name string, data []byte) error {
	return os.WriteFile(name, data, 0644)
}
