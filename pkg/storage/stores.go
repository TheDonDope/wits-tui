package storage

const (
	// StoreInMemory defines that data is only persistant at runtime and will die with the programs exit.
	StoreInMemory = "in-memory"
	// StoreYMLFile defines that data will be loaded and persisted to disk to a specified YML file within the .wits folder.
	StoreYMLFile = "yml-file"
)
