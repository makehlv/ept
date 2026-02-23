package httprequests

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func (r *HttpRequestRepository) baseDir() string {
	baseDir := r.config.HttpRequestsPath
	if !filepath.IsAbs(baseDir) {
		if wd, err := os.Getwd(); err == nil {
			baseDir = filepath.Join(wd, baseDir)
		}
	}
	return baseDir
}

func (r *HttpRequestRepository) ServerRequestsDir(serverName string) string {
	return filepath.Join(r.baseDir(), strings.TrimSpace(serverName))
}

func (r *HttpRequestRepository) SaveCurlIfNotExists(serverName, operationId, curlCmd string) (string, error) {
	serverName = strings.TrimSpace(serverName)
	operationId = strings.TrimSpace(operationId)
	if serverName == "" {
		return "", fmt.Errorf("server name is empty")
	}
	if operationId == "" {
		return "", fmt.Errorf("operationId is empty")
	}

	dir := r.ServerRequestsDir(serverName)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", fmt.Errorf("create dir %q: %w", dir, err)
	}

	path := filepath.Join(dir, operationId+".http")

	if existing, err := os.ReadFile(path); err == nil {
		ts := time.Now().Format("2006.01.02T15:04")

		var b strings.Builder
		b.WriteString("# generated at: ")
		b.WriteString(ts)
		b.WriteString("\n")
		b.WriteString(curlCmd)
		b.WriteString("\n\n")
		b.Write(existing)

		if err := os.WriteFile(path, []byte(b.String()), 0o644); err != nil {
			return "", fmt.Errorf("write file %q: %w", path, err)
		}
		return path, nil
	} else if !os.IsNotExist(err) {
		return "", fmt.Errorf("read file %q: %w", path, err)
	}

	if err := os.WriteFile(path, []byte(curlCmd+"\n"), 0o644); err != nil {
		return "", fmt.Errorf("write file %q: %w", path, err)
	}

	return path, nil
}
