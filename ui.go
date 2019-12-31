package dotui

import (
	"log"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/unit"
)

func newApp(w *app.Window) *App {
	a := &App{
		w:         w,
		profiling: false,
	}

	a.env.theme = newTheme()

	return a
}

// RunUI runs ui
func RunUI() {
	gofont.Register()

	height := float32(20*30 + 28)
	width := float32(1600 / 3)

	go func() {
		w := app.NewWindow(
			app.Size(unit.Dp(width), unit.Dp(height)),
			app.Title("dot-ui"),
		)
		if err := newApp(w).run(); err != nil {
			log.Fatal(err)
		}
	}()

	app.Main()
}
