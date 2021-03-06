package theme

// Most code in this package comes directly from gioui.org/widget/material.
// Copyright Elias Naur
//
// Due to frequent upstream changes, and really simple theming's needs it's
// just Ctrl+C, Ctrl+V. Plan was to deviate from upstream as little as possible
// keep it simple but to allow a lot of experimenting...

import (
	"image/color"

	"gioui.org/font"
	"gioui.org/text"
	"gioui.org/unit"
)

// Theme is simple copy of gioui.org/widget/material theme
type Theme struct {
	Shaper *text.Shaper
	Color  struct {
		Text    color.RGBA
		Hint    color.RGBA
		InvText color.RGBA
		BgEven  color.RGBA
		BgCurr  color.RGBA
	}
	TextSize unit.Value
}

// NewTheme returns simple theme
func NewTheme() *Theme {
	rv := &Theme{
		Shaper:   font.Default(),
		TextSize: unit.Sp(16),
	}

	rv.Color.Text = color.RGBA{A: 255, R: 21, G: 17, B: 24}
	rv.Color.Hint = color.RGBA{A: 255, R: 187, G: 187, B: 187}
	rv.Color.InvText = color.RGBA{A: 255, R: 244, G: 242, B: 245}
	rv.Color.BgCurr = color.RGBA{A: 255, R: 39, G: 37, B: 55}
	rv.Color.BgEven = color.RGBA{A: 255, R: 244, G: 242, B: 245}

	return rv
}

// Caption returns a label for profiling
func (t *Theme) Caption(txt string) label {
	return t.Label(t.TextSize.Scale(12.0/16.0), txt)
}

// Editor returns an editor
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
