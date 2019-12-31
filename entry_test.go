package dotui

import (
	"testing"

	"github.com/kreativka/dot-ui/desktop"
)

func TestEntriesHandleResize(t *testing.T) {
	names := []*desktop.Entry{
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
	}
	tests := []struct {
		in           ents
		out          ents
		shouldResize bool
		winH, entryH int
	}{
		{
			in: ents{
				names: names,
				curr:  0,
				end:   0,
				start: 0,
				limit: 0,
			},
			out: ents{
				names: names,
				start: 0,
				end:   0,
				limit: 26,
				curr:  0,
			},
			shouldResize: true,
			winH:         580,
			entryH:       21,
		},
		{
			in: ents{
				names: names,
				curr:  29,
				start: 2,
				limit: 28,
			},
			out: ents{
				names: names,
				start: 2,
				limit: 26,
				curr:  27,
			},
			shouldResize: true,
			winH:         580,
			entryH:       21,
		},
		{
			in: ents{
				names: names,
				curr:  0,
				start: 0,
				limit: 26,
			},
			out: ents{
				names: names,
				start: 0,
				limit: 26,
				curr:  0,
			},
			shouldResize: false,
			winH:         580,
			entryH:       21,
		},
		{
			in: ents{
				names: names,
				curr:  30,
				start: 5,
				limit: 26,
			},
			out: ents{
				names: names,
				start: 0,
				limit: 36,
				curr:  30,
			},
			shouldResize: true,
			winH:         790,
			entryH:       21,
		},
		{
			in: ents{
				names: names,
				curr:  30,
				start: 5,
				limit: 26,
			},
			out: ents{
				names: names,
				start: 2,
				limit: 32,
				curr:  30,
			},
			shouldResize: true,
			winH:         700,
			entryH:       21,
		}}

	for _, tt := range tests {
		res := tt.in.handleResize(tt.winH, tt.entryH)
		if res != tt.shouldResize || isNotEqual(tt.in, tt.out, false) {
			t.Errorf("handleResize() got %t expected %t", res, tt.shouldResize)
		}
	}
}

func isNotEqual(in, out ents, withNames bool) bool {
	if withNames {
		if len(in.names) != len(out.names) {
			return false
		}
	}
	return in.curr != out.curr || in.start != out.start || in.end != out.end ||
		in.iter != out.iter || in.limit != out.limit || in.filter != out.filter
}
