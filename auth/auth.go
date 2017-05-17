package auth

// Provider interface
type Provider interface {
	Name() string
	GenerateAuthToken() (string, error) // called per request
}
