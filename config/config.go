package config

type Config struct {
	Constants    *map[string]string
	SwaggersPath string
	VarsPath     string
}

func NewConfig() *Config {
	return &Config{}
}
