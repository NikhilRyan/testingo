package config

import "time"

type Config struct {
	Timeout time.Duration
	Verbose bool
}

var defaultConfig = Config{
	Timeout: 30 * time.Second,
	Verbose: false,
}

func SetConfig(config Config) {
	defaultConfig = config
}

func GetConfig() Config {
	return defaultConfig
}
