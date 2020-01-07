package desktop

import "strings"

// Flatten given slice of entries, return two slices of strings containing
// applications names, and executables. It's needed for fuzzy searching, this
// will become "index".
func Flatten(entries []*Entry) [][]string {
	names := make([]string, 0, len(entries))
	execs := make([]string, 0, len(entries))

	for _, v := range entries {
		names = append(names, v.Name)

		switch s := strings.Split(v.Exec, " "); {
		case !strings.HasPrefix(s[0], "/") && len(s) == 1 && len(v.Name) > 1:
			execs = append(execs, TrimRight(v.Exec))
		case !strings.HasPrefix(s[0], "/") && len(s[0]) > 2:
			execs = append(execs, TrimRight(s[0]))
		default:
			execs = append(execs, "")
		}
	}

	return [][]string{names, execs}
}
