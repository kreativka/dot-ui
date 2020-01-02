package dotui

import (
	"flag"
	"log"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/kreativka/dot-ui/theme"
)

func newApp(w *app.Window) *App {
	a := &App{
		w:         w,
		profiling: false,
	}

	a.env.editor = &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
	a.env.localized = parseFlag()
	a.env.theme = theme.NewTheme()

	return a
}

func parseFlag() bool {
	flagL := flag.Bool("l", false, "use localized applications' names")
	flag.Parse()

	return *flagL
}

// RunUI runs ui
func RunUI() {
	const (
		height = float32(20*30 + 28)
		width  = float32(1600 / 3)
		name   = "dot-ui"
	)

	gofont.Register()

	go func() {
		w := app.NewWindow(
			app.Size(unit.Dp(width), unit.Dp(height)),
			app.Title(name),
		)
		if err := newApp(w).run(); err != nil {
			log.Fatal(err)
		}
	}()

	app.Main()
}
