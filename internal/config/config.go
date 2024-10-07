package config

// Config holds application configuration
type Config struct {
	DefaultName string
}

// LoadConfig loads and returns the application configuration
func LoadConfig() Config {
	// In a real scenario, this could load from a file or environment variables
	return Config{
		DefaultName: "World", // Default value if no name is provided
	}
}
