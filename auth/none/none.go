package none

// Config struct
type Config struct{}

// Name is the unique name of this provider
func (config Config) Name() string {
	return "none"
}

// GetProvider reads a config file and unmarshalls its contents.
func GetProvider(authConfig *string) Config {
	return Config{}
}

// GenerateAuthToken not implemented
func (config Config) GenerateAuthToken() (string, error) {
	return "", nil
}
