package dotui

import (
	"image/color"
	"log"

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

func (m *mainPage) filterEntries(filter string) {
	entries := m.env.currList

	// No filter
	if entries.filter == filter {
		return
	}

	// Filter was cleared, so reset to initial state
	if filter == "" && entries.Len() != len(m.env.entries) {
		entries.names = m.env.entries
		entries.filter = filter
		entries.list = flatten(entries.names)

		return
	}

	var list []*desktop.Entry

	// Get the smallest possible list for searching
	if len(filter) > len(entries.filter) {
		list = entries.names
	} else {
		list = m.env.entries
		entries.list = flatten(list)
	}

	entries.filter = filter

	var nl []*desktop.Entry
	seen := make(map[int]bool)

	matches := fuzzy.Find(filter, entries.list[0])
	for _, match := range matches {
		if !seen[match.Index] {
			nl = append(nl, list[match.Index])
			seen[match.Index] = true
		}
	}

	matches = fuzzy.Find(filter, entries.list[1])
	for _, match := range matches {
		if !seen[match.Index] {
			nl = append(nl, list[match.Index])
			seen[match.Index] = true
		}
	}

	if len(nl) > len(m.env.currList.names) {
		entries.start = 0
		entries.curr = 0
	}

	if len(nl) < len(m.env.currList.names) {
		if entries.start > len(nl) {
			entries.start = 0
		}

		if entries.curr >= len(nl) {
			entries.curr = len(nl) - 1
		}
	}

	entries.names = nl
	entries.list = flatten(nl)
}

func (m *mainPage) Layout(gtx *layout.Context) {
	th := m.env.theme
	dims := f32.Point{
		X: float32(gtx.Constraints.Width.Max),
		Y: th.TextSize.V + 5,
	}
	entries := m.env.currList

	head := layout.Rigid(func() {
		layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Rigid(func() {
				th.Body("Run: ").Layout(gtx)
			}),
			layout.Rigid(func() {
				th.Editor("start typing...").Layout(gtx, m.env.editor)
				if entries.filter != m.env.editor.Text() {
					m.filterEntries(m.env.editor.Text())
				}
			}),
		)
	})

	if entries.handleResize(gtx.Constraints.Height.Max, int(dims.Y)) {
		log.Println("new limits applied")
	}

	body := layout.Rigid(func() {
		entries.Reset()
		for entries.Next() {
			layout.Stack{}.Layout(gtx, layout.Stacked(func() {
				name, err := entries.Value()
				if err != nil {
					name = err.Error()
				}
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
		head, body,
	)
}

func fillBg(gtx *layout.Context, maxDims f32.Point, color color.RGBA) {
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{Rect: f32.Rectangle{Max: maxDims}}.Add(gtx.Ops)
}
