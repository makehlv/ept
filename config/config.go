package config

type Config struct {
	Constants    *map[string]string
	SwaggersPath string // путь к файлу server=path (по умолчанию swaggers.txt в текущей директории)
	VarsPath     string // путь к файлу key=value (по умолчанию vars.txt в текущей директории)
}

func NewConfig() *Config {
	return &Config{}
}
