package dotui

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"

	"gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/io/profile"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/kreativka/dot-ui/desktop"
)

// Env struct contains bs
type Env struct {
	theme   *theme
	entries []*desktop.Entry
	editor  *widget.Editor

	currList *ents
}

// App struct contains almost everything
type App struct {
	w    *app.Window
	env  Env
	page *mainPage

	// profiling.
	profiling   bool
	profile     profile.Event
	lastMallocs uint64
}

func (a App) executeCurrent() {
	entry := a.env.currList.CurrentSelection()

	app := strings.Split(trimRight(entry.Exec), " ")
	cmd := exec.Command(app[0], app[1:]...)

	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	if err := cmd.Start(); err != nil {
		log.Fatalln(err)
	}

	os.Exit(0)
}

func (a *App) moveCursorUp() {
	if a.env.currList.CursorUp() {
		a.w.Invalidate()
	}
}

func (a *App) moveCursorDown() {
	if a.env.currList.CursorDown() {
		a.w.Invalidate()
	}
}

func (a *App) layoutTimings(gtx *layout.Context) {
	for _, e := range gtx.Events(a) {
		if e, ok := e.(profile.Event); ok {
			a.profile = e
		}
	}

	profile.Op{Key: a}.Add(gtx.Ops)

	var mstats runtime.MemStats
	runtime.ReadMemStats(&mstats)
	mallocs := mstats.Mallocs - a.lastMallocs
	a.lastMallocs = mstats.Mallocs

	layout.Align(layout.NE).Layout(gtx, func() {
		in := &layout.Inset{}
		in.Top = unit.Max(gtx, unit.Dp(1), in.Top)
		in.Layout(gtx, func() {
			txt := fmt.Sprintf("m: %d %s", mallocs, a.profile.Timings)
			lbl := a.env.theme.Caption(txt)
			lbl.Font.Variant = "Mono"
			lbl.Layout(gtx)
		})
	})
}

func (a *App) run() error {
	gtx := layout.NewContext(a.w.Queue())
	a.page = &mainPage{env: &a.env}

	for e := range a.w.Events() {
		switch e := e.(type) {
		case key.Event:
			switch e.Name {
			case key.NameReturn, key.NameEnter:
				a.executeCurrent()
			case key.NameEscape:
				os.Exit(0)
			case "N":
				if e.Modifiers&key.ModShortcut != 0 {
					a.moveCursorDown()
				}
			case key.NameDownArrow:
				a.moveCursorDown()
			case "P":
				if e.Modifiers&key.ModShortcut != 0 {
					a.moveCursorUp()
				}
			case key.NameUpArrow:
				a.moveCursorUp()
			}
		case system.DestroyEvent:
			return e.Err
		case system.StageEvent:
			if e.Stage >= system.StageRunning {
				if a.env.editor == nil {
					a.env.editor = &widget.Editor{
						SingleLine: true,
						Submit:     true,
					}
				}

				if a.env.entries == nil {
					xdgDirs, err := appendDirs()
					if err != nil {
						log.Fatalln(err)
					}

					entries, err := walk(xdgDirs, true)
					if err != nil {
						log.Fatalln(err)
					}

					a.env.entries = entries
					if a.env.currList == nil {
						a.env.currList = &ents{
							names: entries,
							end:   0,
							curr:  0,
							start: 0,
						}
						a.env.currList.list = flatten(entries)
					}
				}
			}
		case system.FrameEvent:
			gtx.Reset(e.Config, e.Size)

			in := &layout.Inset{
				Top:    e.Insets.Top,
				Left:   e.Insets.Left,
				Right:  e.Insets.Right,
				Bottom: e.Insets.Bottom,
			}
			in.Layout(gtx, func() {
				a.page.Layout(gtx)
			})

			if a.profiling {
				a.layoutTimings(gtx)
			}

			e.Frame(gtx.Ops)
		}
	}

	return nil
}
