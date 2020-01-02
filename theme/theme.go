package theme

import (
	"image/color"

	"gioui.org/font"
	"gioui.org/text"
	"gioui.org/unit"
)

// Theme is simple theme
type Theme struct {
	Shaper *text.Shaper
	Color  struct {
		Primary color.RGBA
		Text    color.RGBA
		Hint    color.RGBA
		InvText color.RGBA
		BgEven  color.RGBA
		BgCurr  color.RGBA
	}
	TextSize unit.Value
}

// NewTheme returns simple theme struct
func NewTheme() *Theme {
	rv := &Theme{
		Shaper:   font.Default(),
		TextSize: unit.Sp(16),
	}

	rv.Color.Primary = rgb(0x3f51b5)
	rv.Color.Text = color.RGBA{A: 254, R: 21, G: 17, B: 24}
	rv.Color.Hint = rgb(0xbbbbbb)
	rv.Color.InvText = color.RGBA{A: 254, R: 244, G: 242, B: 245}
	rv.Color.BgCurr = color.RGBA{A: 255, R: 39, G: 37, B: 55}
	rv.Color.BgEven = color.RGBA{A: 255, R: 244, G: 242, B: 245}

	return rv
}

// Caption returns label for profiling
func (t *Theme) Caption(txt string) label {
	return t.Label(t.TextSize.Scale(12.0/16.0), txt)
}

// Editor return editor
func (t *Theme) Editor(hint string) editor {
	return editor{
		Font: text.Font{
			Size: t.TextSize,
		},
		Color:     t.Color.Text,
		shaper:    t.Shaper,
		Hint:      hint,
		ColorHint: t.Color.Hint,
	}
}
