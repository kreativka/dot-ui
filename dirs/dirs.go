package dirs

import (
	"os"
	"path"

	"github.com/kreativka/dot-ui/xdg"
)

// AppendXDGDirs returns slice of xdg standard dirs to be later checked for .dot
// files
func AppendXDGDirs() ([]string, error) {
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
