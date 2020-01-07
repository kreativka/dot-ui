package dirs

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"syscall"

	"github.com/kreativka/dot-ui/desktop"
	"golang.org/x/text/language"
)

// Walk through xdg directories and return slice of desktop entries from .dot
// files found in that directories.
func Walk(dirs []string, useLang bool) ([]*desktop.Entry, error) {
	if len(dirs) == 0 {
		return nil, ErrEmptyList
	}

	var nameLC string
	if useLang {
		nameLC = addLocale()
	}

	var rv []*desktop.Entry
	seen := make(map[string]bool)

	for i, dir := range dirs {
		f, err := os.Open(dir)
		if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOENT {
			continue
		}
		if err != nil {
			return nil, err
		}
		defer f.Close()

		files, err := f.Readdirnames(-1)
		if err != nil {
			return nil, err
		}

		for _, fn := range files {
			if filepath.Ext(fn) != ".desktop" {
				continue
			}

			el, err := desktop.NewEntry(path.Join(dir, fn), nameLC)
			if err == desktop.ErrHiddenEntry ||
				err == desktop.ErrInvalidEntry || el == nil {
				continue
			}
			if err != nil {
				return nil, err
			}

			if seen[el.Name] {
				el.Name = fmt.Sprintf("%s %s%d", el.Name, "EXTRA-", i)
			}
			seen[el.Name] = true

			rv = append(rv, el)
		}
	}

	return rv, nil
}

func addLocale() string {
	if str := os.Getenv("LANG"); str != "" {
		ll := language.Make(str)
		base, _ := ll.Base()

		return fmt.Sprintf("Name[%s]", base.String())
	}

	return ""
}
