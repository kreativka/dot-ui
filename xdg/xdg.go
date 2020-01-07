package xdg

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// XDG contains xdg model dirs
type XDG struct {
	dataHome string
	dataDirs []string
}

// DataHome returns xdg home
func (x XDG) DataHome() string {
	return x.dataHome
}

//DataDirs returns xdg data dirs
func (x XDG) DataDirs() []string {
	return x.dataDirs
}

// NewXDG returns simple xdg model
// Returns error when $HOME is not set, don't check if directory exists.
func NewXDG() (*XDG, error) {
	rv := &XDG{}

	dataHome, err := parseDataHome()
	if err != nil {
		return nil, err
	}

	rv.dataHome = dataHome
	rv.dataDirs = parseDataDirs()

	return rv, nil
}

func parseDataDirs() []string {
	var rv []string

	if s, ok := os.LookupEnv("XDG_DATA_DIRS"); s != "" && ok {
		dirs := strings.Split(s, ":")
		for _, el := range dirs {
			dir := filepath.Clean(el)
			if dir != "." {
				rv = append(rv, el)
			}
		}
	}

	if len(rv) == 0 {
		rv = []string{"/usr/local/share/", "/usr/share/"}
	}

	return rv
}

func parseDataHome() (string, error) {
	homeDir, err := parseHome()
	if err != nil {
		return "", err
	}

	if s, ok := os.LookupEnv("XDG_DATA_HOME"); ok {
		if s != "" {
			dir := filepath.Clean(s)
			if dir != "." {
				return dir, nil
			}
		}
	}

	return path.Join(homeDir, ".local/share"), nil
}

func parseHome() (string, error) {
	var rv string

	if s, ok := os.LookupEnv("HOME"); ok {
		s = filepath.Clean(s)
		if s == "." {
			return rv, errors.New("$HOME is not set")
		}
		rv = s
	}

	return rv, nil
}
