package desktop

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strings"
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
func NewEntry(fn, lang string, useLang bool) (*Entry, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	rv, err := parseEntry(f, lang, useLang)
	if err == ErrHiddenEntry {
		return nil, err

	}
	if err != nil || rv.Type == "" || rv.Exec == "" || rv.Name == "" {
		return nil, ErrInvalidEntry
	}

	return rv, nil
}

// Sort returns slice of entries sorted by name in order
func Sort(slice []*Entry, descending bool) {
	if descending {
		sort.Slice(slice, func(i, j int) bool {
			return slice[i].Name > slice[j].Name
		})
		return
	}
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].Name < slice[j].Name
	})
}

// parseEntry parses perfectly formed .desktop files. Section starts at col 0
// like # for comments. Every field Name, Type, Exec starts at col 0 and
// has = right after its name.
//
// Skip empty lines and lines shorter than THE keys. On the first pass of
// line: check line's first char and skip C like comments, X like X-Something,
// I like Icon and so on, if character isn't in exclude list, then bravely
// parse this line.
// If [Desktop Action is seen return what is parsed to this moment as
// it should be the whole [Desktop Entry]
func parseEntry(r io.Reader, lang string, useLang bool) (*Entry, error) {
	var err error
	rv := &Entry{}
	b := bufio.NewScanner(r)
	for b.Scan() {
		l := b.Text()
		if l == "" && len(l) < 5 {
			continue
		}
		if strings.HasPrefix(l, "[Desktop Action") {
			break
		}
		switch l[0] {
		case ' ', '#', '[', 'C', 'G', 'I', 'K', 'M', 'S', 'V', 'X':
			continue
		default:
			if err := parseLine(l, lang, useLang, rv); err == ErrHiddenEntry {
				return nil, err
			}
		}
	}
	if err := b.Err(); err != nil {
		return nil, err
	}
	return rv, err
}

// parseLine returns error when Hidden or NoDisplay fields are set to true,
// otherwise returns nil for good fields and empty lines or even for very
// strange fields.
// parseLine updates directly fields in entry struct.
func parseLine(str, lang string, useLang bool, entry *Entry) error {
	s := strings.SplitN(str, "=", 2)
	if len(s) == 0 || len(s) > 2 {
		return nil
	}

	switch s[0] {
	case "Exec":
		entry.Exec = parseStr(entry.Exec, s[1])
	case "Hidden", "NoDisplay":
		if isTrue(s[1]) {
			return ErrHiddenEntry
		}
	case lang: // Name[pl]
		if useLang {
			entry.Name = s[1]
		}
	case "Name":
		entry.Name = parseStr(entry.Name, s[1])
	case "Terminal":
		entry.Term = isTrue(s[1])
	case "Type":
		entry.Type = parseStr(entry.Type, s[1])
	default:
		return nil
	}
	return nil
}

// parseStr returns newStr only when oldStr is empty.
func parseStr(oldStr, newStr string) string {
	if oldStr == "" {
		return newStr
	}
	return oldStr
}

// isTrue optimistically takes that freedesktop's spec is followed.
// Values of type boolean must either be the string true or false.
func isTrue(str string) bool {
	return str == "true"
}
