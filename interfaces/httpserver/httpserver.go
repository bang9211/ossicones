package httpserver

type HTTPServer interface {
	// Serve listens and serves the HTTP Server.
	Serve()
	// Close closes the HTTP Server.
	Close()
}
