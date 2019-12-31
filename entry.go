package dotui

import (
	"errors"
	"fmt"
	"strings"

	"github.com/kreativka/dot-ui/desktop"
)

type ents struct {
	list        [][]string
	names       []*desktop.Entry
	curr        int
	start, end  int
	iter, limit int
	filter      string
}

func flatten(entries []*desktop.Entry) [][]string {
	rv := make([][]string, 2)
	for _, v := range entries {
		rv[0] = append(rv[0], v.Name)

		switch s := strings.Split(v.Exec, " "); {
		case !strings.HasPrefix(s[0], "/") && len(s) == 1 && len(v.Name) > 1:
			rv[1] = append(rv[1], trimRight(v.Exec))
		case !strings.HasPrefix(s[0], "/") && len(s[0]) > 2:
			rv[1] = append(rv[1], trimRight(s[0]))
		default:
			rv[1] = append(rv[1], "")
		}
	}

	return rv
}

func (e ents) String() string {
	return fmt.Sprintf(
		"len(names) = %d, curr = %d, start = %d, end = %d, iter = %d, "+
			"limit = %d, filter = %s",
		len(e.names), e.curr, e.start, e.end, e.iter, e.limit, e.filter,
	)
}

func (e ents) currIndex() int {
	return e.iter + e.start
}

func (e *ents) Next() bool {
	e.iter++
	if e.iter < e.limit && e.iter < len(e.names) {
		return true
	}

	return false
}

func (e *ents) Reset() {
	e.iter = -1
}

func (e *ents) IsCurrentHighlighted() bool {
	return e.curr == e.currIndex()
}

func (e ents) IsCurrentEven() bool {
	return e.iter%2 == 0
}

func (e *ents) Value() (string, error) {
	if e.currIndex() >= len(e.names) || e.currIndex() < 0 {
		return "", errors.New("index out of bounds")
	}

	return e.names[e.currIndex()].Name, nil
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

func (e *ents) handleResize(height, entryHeight int) bool {
	limit := height / entryHeight

	// Check if we should hide last non-fully visible entry.
	// 0.8 on scale 0 - 1 how much of entry is visible
	// 0.8 makes it almost fully visible
	offset := float32(height) / float32(entryHeight)
	if offset-float32(limit) < 0.8 {
		limit--
	}

	if e.limit == limit {
		return false
	}

	if limit < e.limit {
		// When shrinking hits cursor position, keep it on the last visible
		// entry.
		if e.curr >= e.start+limit {
			if e.start+limit < len(e.names)-1 && e.start+limit > 0 {
				e.curr = limit + e.start - 1
			}
		}
	}

	if limit > e.limit {
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

	return true
}
