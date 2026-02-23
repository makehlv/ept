package swagger

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (r *SwaggerRepository) GetServers() (map[string]string, error) {
	path := r.swaggersPath()
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read swagger file: %w", err)
	}
	out := make(map[string]string)
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		idx := strings.Index(line, "=")
		if idx < 0 {
			continue
		}
		k := strings.TrimSpace(line[:idx])
		v := strings.TrimSpace(line[idx+1:])
		v = strings.Trim(v, `"`)
		if k != "" {
			out[k] = v
		}
	}
	return out, nil
}

func (r *SwaggerRepository) SaveServerSpec(serverName, specPath string) error {
	path := r.swaggersPath()
	serverName = strings.TrimSpace(serverName)
	if serverName == "" {
		return fmt.Errorf("server name is empty")
	}
	specPath = strings.TrimSpace(specPath)
	if !filepath.IsAbs(specPath) {
		if wd, err := os.Getwd(); err == nil {
			specPath = filepath.Join(wd, specPath)
		}
	}

	entries := make(map[string]string)
	if data, err := os.ReadFile(path); err == nil {
		for _, line := range strings.Split(string(data), "\n") {
			line = strings.TrimSpace(line)
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			idx := strings.Index(line, "=")
			if idx < 0 {
				continue
			}
			k := strings.TrimSpace(line[:idx])
			v := strings.TrimSpace(line[idx+1:])
			v = strings.Trim(v, `"`)
			if k != "" {
				entries[k] = v
			}
		}
	}
	entries[serverName] = specPath

	var b strings.Builder
	for k, v := range entries {
		b.WriteString(k)
		b.WriteString("=")
		b.WriteString(v)
		b.WriteString("\n")
	}
	return os.WriteFile(path, []byte(b.String()), 0644)
}

func (r *SwaggerRepository) LoadSpec(specPath string) ([]byte, error) {
	data, err := os.ReadFile(specPath)
	if err != nil {
		return nil, fmt.Errorf("load spec: %w", err)
	}
	return data, nil
}

func (r *SwaggerRepository) swaggersPath() string {
	p := r.config.SwaggersPath
	if p == "" {
		p = "swaggers.txt"
	}
	if !filepath.IsAbs(p) {
		if wd, err := os.Getwd(); err == nil {
			p = filepath.Join(wd, p)
		}
	}
	return p
}
