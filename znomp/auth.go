package znomp

// Auth type of API authentication to use.
type Auth uint8

const (
	// HTTPBasicAuth constants for HTTP Basic Authentication.
	HTTPBasicAuth = Auth(iota)
)
