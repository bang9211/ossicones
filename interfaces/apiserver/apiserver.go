package apiserver

type APIServer interface {
	// Serve listens and serves the API Server.
	Serve()
	// Close closes the API Server.
	Close()
}
