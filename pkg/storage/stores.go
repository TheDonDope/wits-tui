package storage

import (
	"os"
)

var witsFolder = os.Getenv("WITS_DIR")

const (
	// StoreInMemory defines that data is only persistant at runtime and will die with the programs exit.
	StoreInMemory = "in-memory"
	// StoreYMLFile defines that data will be loaded and persisted to disk to a specified YML file within the .wits folder.
	StoreYMLFile = "yml-file"
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
