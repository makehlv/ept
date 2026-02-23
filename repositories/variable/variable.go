package variable

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (r *VariableRepository) varsPath() string {
	p := r.config.VarsPath
	if p == "" {
		p = "vars.txt"
	}
	if !filepath.IsAbs(p) {
		if wd, err := os.Getwd(); err == nil {
			p = filepath.Join(wd, p)
		}
	}
	return p
}

func (r *VariableRepository) GetVars() (map[string]string, error) {
	path := r.varsPath()
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return map[string]string{}, nil
		}
		return nil, fmt.Errorf("read vars file: %w", err)
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

func (r *VariableRepository) Add(key, value string) error {
	varsPath := r.varsPath()
	key = strings.TrimSpace(key)
	if key == "" {
		return fmt.Errorf("key is empty")
	}

	entries := make(map[string]string)
	if data, err := os.ReadFile(varsPath); err == nil {
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
	entries[key] = value

	var b strings.Builder
	for k, v := range entries {
		b.WriteString(k)
		b.WriteString("=")
		b.WriteString(v)
		b.WriteString("\n")
	}
	return os.WriteFile(varsPath, []byte(b.String()), 0644)
}
