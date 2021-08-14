package httpserver

type HTTPServer interface {
	Serve()
	Close()
}
