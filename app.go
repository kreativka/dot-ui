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
	"github.com/kreativka/dot-ui/theme"
)

// Env struct contains bs
type Env struct {
	theme   *theme.Theme
	entries []*desktop.Entry
	editor  *widget.Editor

	currList  *ents
	localized bool
}

// App struct contains almost everything
type App struct {
	w    *app.Window
	env  Env
	page *mainPage

	// profiling
	profiling   bool
	profile     profile.Event
	lastMallocs uint64
}

func (a App) executeCurrent() {
	term := []string{"alacritty", "--command"}
	entry := a.env.currList.CurrentSelection()
	app := strings.Split(trimRight(entry.Exec), " ")

	var cmd *exec.Cmd
	if entry.Term {
		cmd = exec.Command(term[0], append(term[1:], app[0:]...)...)
	} else {
		cmd = exec.Command(app[0], app[1:]...)
	}

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

func (a *App) populateEntries() error {
	if a.env.entries != nil {
		return nil
	}

	xdgDirs, err := appendDirs()
	if err != nil {
		return err
	}

	entries, err := walk(xdgDirs, a.env.localized)
	if err != nil {
		return err
	}

	a.env.entries = entries

	a.env.currList = &ents{
		names: entries,
		curr:  0,
		start: 0,
		list:  flatten(entries),
	}

	return nil
}

func (a *App) run() error {
	gtx := layout.NewContext(a.w.Queue())
	a.page = &mainPage{env: &a.env}

	if err := a.populateEntries(); err != nil {
		return err
	}

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
