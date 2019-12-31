package dotui

import (
	"testing"

	"github.com/kreativka/dot-ui/desktop"
)

var longerInitial = ents{
	names: []*desktop.Entry{
		{Name: "Alacritty"},
		{Name: "Archiwa"},
		{Name: "Boxes"},
		{Name: "Brave Web Browser"},
		{Name: "Cheese"},
		{Name: "Chrome Apps & Extensions Developer Tool"},
		{Name: "Chrome Dev Editor"},
		{Name: "Czcionki"},
		{Name: "Dyski"},
		{Name: "Dzienniki"},
		{Name: "Edytor tekstu"},
		{Name: "FeedReader"},
		{Name: "Filmy"},
		{Name: "Firefox"},
		{Name: "Geary"},
		{Name: "GoLand"},
		{Name: "Google Chrome"},
		{Name: "Google Keep - notatki i listy"},
		{Name: "Kalendarz"},
		{Name: "Kalkulator"},
		{Name: "Klient poczty Thunderbird"},
		{Name: "Kontakty"},
		{Name: "LibreOffice Calc"},
		{Name: "LibreOffice Draw"},
		{Name: "LibreOffice Impress"},
		{Name: "LibreOffice Writer"},
		{Name: "Lutris"},
		{Name: "Mapy"},
		{Name: "Menedżer maszyn wirtualnych"},
		{Name: "Minder"},
		{Name: "Monitor systemu"},
		{Name: "Monitor systemu GNOME"},
		{Name: "Neovim"},
		{Name: "Notepad"},
	},
}

func TestEntriesCursorDown(t *testing.T) {
	initial := ents{
		names: []*desktop.Entry{
			{Name: "Alacritty"},
			{Name: "Archiwa"},
			{Name: "Boxes"},
			{Name: "Brave Web Browser"},
			{Name: "Cheese"},
			{Name: "Chrome Apps & Extensions Developer Tool"},
			{Name: "Chrome Dev Editor"},
			{Name: "Czcionki"},
			{Name: "Dyski"},
		},
	}
	tests := []struct {
		in    ents
		times int
		out   int
	}{
		{
			in:    initial,
			times: 1,
			out:   1,
		},
		{
			in:    initial,
			times: 7,
			out:   7,
		},
		{
			in:    initial,
			times: 8,
			out:   8,
		},
		{
			in:    initial,
			times: 20,
			out:   8,
		},
	}

	for _, tt := range tests {
		for i := 0; i < tt.times; i++ {
			tt.in.CursorDown()
		}
		if tt.in.curr != tt.out {
			t.Errorf("cursorDown() %d times got %d expected %d", tt.times, tt.in.curr, tt.out)
		}
	}
}

func TestEntriesChangeUp(t *testing.T) {
	initial := ents{
		names: []*desktop.Entry{
			{Name: "Alacritty"},
			{Name: "Archiwa"},
			{Name: "Boxes"},
			{Name: "Brave Web Browser"},
			{Name: "Cheese"},
			{Name: "Chrome Apps & Extensions Developer Tool"},
			{Name: "Chrome Dev Editor"},
			{Name: "Czcionki"},
			{Name: "Dyski"},
		},
		curr:  8,
		end:   8,
		start: 0,
		limit: 0,
	}
	tests := []struct {
		in    ents
		times int
		out   int
	}{
		{
			in:    initial,
			times: 1,
			out:   7,
		},
		{
			in:    initial,
			times: 7,
			out:   1,
		},
		{
			in:    initial,
			times: 8,
			out:   0,
		},
	}

	for _, tt := range tests {
		for i := 0; i < tt.times; i++ {
			tt.in.CursorUp()
		}
		if tt.in.curr != tt.out {
			t.Errorf("cursorUp() times %d got %d expected %d", tt.times, tt.in.curr, tt.out)
		}
	}
}

// {
// 	name: "keep highlighted on bottom",
// },
// {
// 	name: "keep highlighted on top",
// },
// {
// 	name: "hande cursor down",
// },
// {
// 	name: "handle cursor up",
// },
// {
// 	name: "handle change of height/width",
// },
