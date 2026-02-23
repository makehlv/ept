package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	RootDirPath      string
	SwaggersPath     string
	VarsPath         string
	HttpRequestsPath string
}

func NewConfig() *Config {
	return &Config{
		RootDirPath:      "eptfiles",
		VarsPath:         "eptfiles/vars.txt",
		SwaggersPath:     "eptfiles/swaggers.txt",
		HttpRequestsPath: "eptfiles/generated/httprequests",
	}
}

func (c *Config) BaseDir() string {
	p := c.RootDirPath
	if !filepath.IsAbs(p) {
		if wd, err := os.Getwd(); err == nil {
			p = filepath.Join(wd, p)
		}
	}
	return p
}
