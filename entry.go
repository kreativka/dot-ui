package dotui

import (
	"strings"

	"github.com/kreativka/dot-ui/desktop"
)

type ents struct {
	names                    []*desktop.Entry
	list                     [][]string
	filter                   string
	curr, iter, limit, start int
}

func (e *ents) Next() bool {
	e.iter++
	if e.iter < e.limit && e.iter < len(e.names) {
		return true
	}

	return false
}

func (e *ents) Value() string {
	return e.names[e.currIndex()].Name
}

func (e *ents) Reset() {
	e.iter = -1
}

func (e ents) currIndex() int {
	return e.iter + e.start
}

func (e *ents) IsCurrentHighlighted() bool {
	return e.curr == e.currIndex()
}

func (e ents) IsCurrentEven() bool {
	return e.iter%2 == 0
}

func (e ents) CurrentSelection() *desktop.Entry {
	return e.names[e.curr]
}

func (e ents) Len() int {
	return len(e.names)
}

func (e *ents) CursorDown() bool {
	if e.curr < len(e.names)-1 {
		e.curr++
		if e.curr >= e.limit && e.curr < len(e.names) &&
			e.start+e.limit+1 <= len(e.names) {
			e.start++
		}

		return true
	}

	return false
}

func (e *ents) CursorUp() bool {
	if e.curr > 0 {
		if e.curr == e.start && e.start > 0 {
			e.start--
		}
		e.curr--

		return true
	}

	return false
}

func (e *ents) handleResize(height, entryHeight int) {
	limit := height / entryHeight

	// Check if we should hide last non-fully visible entry.
	// 0.8 on scale 0 - 1 how much of entry is visible
	// 0.8 makes it almost fully visible
	offset := float32(height) / float32(entryHeight)
	if offset-float32(limit) < 0.8 {
		limit--
	}

	switch {
	case limit == e.limit:
		return
	case limit < e.limit:
		// When shrinking hits cursor position, keep it on the last visible
		// entry.
		if e.curr >= e.start+limit {
			if e.start+limit < len(e.names)-1 && e.start+limit > 0 {
				e.curr = limit + e.start - 1
			}
		}
	case limit > e.limit:
		// When growing continue showing next entries, when at the end: start
		// showing previous entries.
		if e.start > 0 && e.start+limit >= len(e.names) {
			for ; e.start > 0; e.start-- {
				if e.start+limit <= len(e.names) {
					break
				}
			}
		}
	}

	e.limit = limit
}

func flatten(entries []*desktop.Entry) [][]string {
	names := make([]string, 0, len(entries))
	execs := make([]string, 0, len(entries))

	for _, v := range entries {
		names = append(names, v.Name)

		switch s := strings.Split(v.Exec, " "); {
		case !strings.HasPrefix(s[0], "/") && len(s) == 1 && len(v.Name) > 1:
			execs = append(execs, trimRight(v.Exec))
		case !strings.HasPrefix(s[0], "/") && len(s[0]) > 2:
			execs = append(execs, trimRight(s[0]))
		default:
			execs = append(execs, "")
		}
	}

	return [][]string{names, execs}
}
