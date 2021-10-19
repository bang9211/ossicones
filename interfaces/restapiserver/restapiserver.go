package restapiserver

type RESTAPIServer interface {
	// Serve listens and serves the REST API Server.
	Serve() error
	// Close closes the REST API Server.
	Close() error
}
