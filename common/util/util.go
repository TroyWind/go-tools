package util

import (
	"path/filepath"
	"strings"
	"os"
	"runtime"
	"os/user"
)

// Gopath returns all GOPATHs
func Gopath() []string {
	defaultGoPath := runtime.GOROOT()
	if u, err := user.Current(); err == nil {
		defaultGoPath = filepath.Join(u.HomeDir, "go")
	}
	path := Getenv("GOPATH", defaultGoPath)
	paths := strings.Split(path, string(os.PathListSeparator))
	for idx, path := range paths {
		if abs, err := filepath.Abs(path); err == nil {
			paths[idx] = abs
		}
	}
	return paths
}

// Getenv returns the values of the environment variable, otherwise
//returns the default if variable is not set
func Getenv(key, userDefault string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return userDefault
}