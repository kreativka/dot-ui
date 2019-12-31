package dotui

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
)

type label struct {
	// Face defines the text style.
	Font text.Font
	// Color is the text color.
	Color color.RGBA
	// Alignment specify the text alignment.
	Alignment text.Alignment
	// MaxLines limits the number of lines. Zero means no limit.
	MaxLines int
	Text     string

	shaper *text.Shaper
}

func (t *theme) Label(size unit.Value, txt string) label {
	return label{
		Text:  txt,
		Color: t.Color.Text,
		Font: text.Font{
			Size: size,
		},
		MaxLines: 1,
		shaper:   t.Shaper,
	}
}

func (t *theme) LabelOdd(size unit.Value, txt string) label {
	return label{
		Text:  txt,
		Color: t.Color.InvText,
		Font: text.Font{
			Size: size,
		},
		MaxLines: 1,
		shaper:   t.Shaper,
	}
}

func (l label) Layout(gtx *layout.Context) {
	paint.ColorOp{Color: l.Color}.Add(gtx.Ops)
	tl := widget.Label{Alignment: l.Alignment, MaxLines: l.MaxLines}
	tl.Layout(gtx, l.shaper, l.Font, l.Text)
}

func (t *theme) Body(txt string) label {
	return t.Label(t.TextSize, txt)
}

func (t *theme) BodyOdd(txt string) label {
	return t.LabelOdd(t.TextSize, txt)
}
