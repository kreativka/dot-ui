package desktop

import (
	"os"
	"sort"
)

// Entry struct for fields from file.desktop
type Entry struct {
	Exec string
	Name string
	Term bool
	Type string
	URL  string
}

// NewEntry returns new entry with parsed fields from given file.
// returns errInvalid
func NewEntry(fn, lang string) (*Entry, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	rv, err := parseEntry(f, lang)
	if err == ErrHiddenEntry {
		return nil, err
	}

	if err != nil || rv.Type == "" || rv.Exec == "" || rv.Name == "" {
		return nil, ErrInvalidEntry
	}

	return rv, nil
}

// Sort returns slice of entries sorted by name in order
func Sort(entries []*Entry, descending bool) {
	if descending {
		sort.Slice(entries, func(i, j int) bool {
			return entries[i].Name > entries[j].Name
		})

		return
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name < entries[j].Name
	})
}
