package config

import "time"

// Config holds runtime configuration for the service.
type Config struct {
	Host          string        `env:"HOST"`
	Port          int           `env:"PORT"`
	ReadTimeout   time.Duration `env:"READ_TIMEOUT"`
	WriteTimeout  time.Duration `env:"WRITE_TIMEOUT"`
	ShutdownGrace time.Duration `env:"SHUTDOWN_GRACE"`
}

// Load loads configuration. Replace with your preferred loader (env, file, etc.).
func Load() *Config {
	// TODO: load from env / file / flags
	return &Config{
		Host:          "0.0.0.0",
		Port:          8080,
		ReadTimeout:   5 * time.Second,
		WriteTimeout:  10 * time.Second,
		ShutdownGrace: 5 * time.Second,
	}
}
