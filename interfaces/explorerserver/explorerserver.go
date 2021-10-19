package explorerserver

type ExplorerServer interface {
	// Serve listens and serves the Explorer Server.
	Serve() error
	// Close closes the Explorer Server.
	Close() error
}
