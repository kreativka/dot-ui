package dotui

import (
	"os"
	"path"

	"github.com/kreativka/dot-ui/xdg"
)

func appendDirs() ([]string, error) {
	const appDir = "/applications"

	XDG, err := xdg.NewXDG()
	if err != nil {
		return nil, err
	}

	var rv []string
	rv = append(rv, path.Join(XDG.DataHome(), appDir))

	for _, dir := range XDG.DataDirs() {
		if _, err := os.Stat(dir); !os.IsNotExist(err) {
			rv = append(rv, path.Join(dir, appDir))
		}
	}

	return rv, nil
}
