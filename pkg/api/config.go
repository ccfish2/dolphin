package api

// Endpoint is an API Endpoint for a parsed API specification.
type Endpoint struct {
	Method      string
	Path        string
	Description string
}
type PathSet map[string]Endpoint
