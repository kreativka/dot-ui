package dotui

import (
	"image/color"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"github.com/kreativka/dot-ui/desktop"
	"github.com/sahilm/fuzzy"
)

type mainPage struct {
	env *Env
}

func (m *mainPage) Layout(gtx *layout.Context) {
	th := m.env.theme
	entries := m.env.currList
	dims := f32.Point{
		X: float32(gtx.Constraints.Width.Max),
		Y: th.TextSize.V + 5,
	}

	m.filterEntries(m.env.editor.Text())

	entries.handleResize(gtx.Constraints.Height.Max, int(dims.Y))

	head := layout.Rigid(func() {
		layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Rigid(func() {
				th.Body("Run: ").Layout(gtx)
			}),
			layout.Rigid(func() {
				th.Editor("start typing...").Layout(gtx, m.env.editor)
			}),
		)
	})

	list := layout.Rigid(func() {
		entries.Reset()
		for entries.Next() {
			layout.Stack{}.Layout(gtx, layout.Stacked(func() {
				name := entries.Value()
				switch {
				case entries.IsCurrentHighlighted():
					fillBg(gtx, dims, th.Color.BgCurr)
					th.BodyOdd(name).Layout(gtx)
				case entries.IsCurrentEven():
					fillBg(gtx, dims, th.Color.BgEven)
					th.Body(name).Layout(gtx)
				default:
					th.Body(name).Layout(gtx)
				}
			}))
			op.TransformOp{}.Offset(f32.Point{Y: dims.Y}).Add(gtx.Ops)
		}
	})

	layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		head, list,
	)
}

func (m *mainPage) filterEntries(pattern string) {
	entries := m.env.currList

	// No filter
	if entries.filter == pattern {
		return
	}

	// Filter was cleared, reset to initial state
	if pattern == "" && entries.Len() != len(m.env.entries) {
		entries.filter = pattern
		entries.names = m.env.entries
		entries.list = desktop.Flatten(entries.names)

		return
	}

	var list []*desktop.Entry

	if len(pattern) > len(entries.filter) {
		list = entries.names
	} else {
		list = m.env.entries
		entries.list = desktop.Flatten(list)
	}

	res := filterMatches(list, entries.list, pattern)
	if len(res) >= len(m.env.currList.names) ||
		(len(res) < len(m.env.currList.names) && entries.start > len(res)) ||
		entries.start+entries.limit >= len(res) {
		entries.start = 0
	}

	if entries.curr >= len(res) && len(res) > 0 {
		entries.curr = len(res) - 1
	}

	if entries.curr > entries.start+entries.limit {
		entries.curr = entries.start + entries.limit - 1
	}

	entries.filter = pattern
	entries.list = desktop.Flatten(res)
	entries.names = res
}

func filterMatches(entries []*desktop.Entry, lists [][]string, pattern string) []*desktop.Entry {
	rv := make([]*desktop.Entry, 0, len(entries))
	seen := make(map[int]bool, len(entries))

	for _, elems := range lists {
		matches := fuzzy.Find(pattern, elems)
		for _, match := range matches {
			if !seen[match.Index] {
				rv = append(rv, entries[match.Index])
				seen[match.Index] = true
			}
		}
	}

	return rv
}

func fillBg(gtx *layout.Context, maxDims f32.Point, color color.RGBA) {
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{Rect: f32.Rectangle{Max: maxDims}}.Add(gtx.Ops)
}
