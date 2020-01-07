package theme

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

// Label returns themed label
func (t *Theme) Label(size unit.Value, txt string) label {
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

// Label returns themed lable for odd rows in list
func (t *Theme) LabelOdd(size unit.Value, txt string) label {
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

// Layout this label
func (l label) Layout(gtx *layout.Context) {
	paint.ColorOp{Color: l.Color}.Add(gtx.Ops)
	tl := widget.Label{Alignment: l.Alignment, MaxLines: l.MaxLines}
	tl.Layout(gtx, l.shaper, l.Font, l.Text)
}

// Body is a wrapper around themed label
func (t *Theme) Body(txt string) label {
	return t.Label(t.TextSize, txt)
}

// BodyOdd is a wrapper around themed label for odd rows
func (t *Theme) BodyOdd(txt string) label {
	return t.LabelOdd(t.TextSize, txt)
}
