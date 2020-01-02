package dotui

import (
	"reflect"
	"testing"

	"github.com/kreativka/dot-ui/desktop"
)

var entries = []*desktop.Entry{
	{Exec: "/opt/google/chrome/google-chrome --profile-directory=Default --app-id=pnoffddplpippgcfjdhbmhkofpnaalpg",
		Name: "Chrome Dev Editor",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "imv %F",
		Name: "imv",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "/opt/google/chrome/google-chrome --profile-directory=Default --app-id=emefpkhgihlhfddcjfghpndaeliajgjj",
		Name: "TIDAL",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "/opt/google/chrome/google-chrome --profile-directory=Default --app-id=hmjkmjkepdijhoojdojkdfohbdgmmhki",
		Name: "Google Keep – notatki i listy",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "/opt/google/chrome/google-chrome --profile-directory=Default --app-id=ohmmkhmmmpcnpikjeljgnaoabkaalbgc",
		Name: "Chrome Apps & Extensions Developer Tool",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "/usr/bin/flatpak run --branch=stable --arch=x86_64 --command=obs com.obsproject.Studio",
		Name: "OBS Studio",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "/usr/bin/flatpak run --branch=stable --arch=x86_64 --command=pulseeffects com.github.wwmm.pulseeffects",
		Name: "PulseEffects",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "/usr/bin/flatpak run --branch=stable --arch=x86_64 --command=feedreader --file-forwarding org.gnome.FeedReader @@u %U @@",
		Name: "FeedReader",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "/usr/bin/flatpak run --branch=stable --arch=x86_64 --command=geary --file-forwarding org.gnome.Geary @@u %U @@",
		Name: "Geary",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "/usr/bin/flatpak run --branch=stable --arch=x86_64 --command=com.github.johnfactotum.Foliate --file-forwarding com.github.johnfactotum.Foliate @@ %F @@",
		Name: "Foliate",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "/usr/bin/flatpak run --branch=stable --arch=x86_64 --command=com.github.lainsce.timetable com.github.lainsce.timetable",
		Name: "Timetable",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "/usr/bin/flatpak run --branch=stable --arch=x86_64 --command=Postman --file-forwarding com.getpostman.Postman @@u %U @@",
		Name: "Postman",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "/usr/bin/flatpak run --branch=stable --arch=x86_64 --command=thunderbird --file-forwarding org.mozilla.Thunderbird @@u %u @@",
		Name: "Klient poczty Thunderbird",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "/usr/bin/flatpak run --branch=stable --arch=x86_64 --command=gitkraken --file-forwarding com.axosoft.GitKraken @@u %U @@",
		Name: "GitKraken",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "/usr/bin/flatpak run --branch=stable --arch=x86_64 --command=keepassxc --file-forwarding org.keepassxc.KeePassXC @@ %f @@",
		Name: "KeePassXC",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "pavucontrol",
		Name: "Sterowanie głośnością PulseAudio",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "/usr/bin/gnome-characters",
		Name: "Znaki",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "gcm-viewer",
		Name: "Podgląd profilu kolorów",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "ranger",
		Name: "ranger",
		Term: true,
		Type: "Application",
		URL:  ""},
	{Exec: "file-roller %U",
		Name: "Archiwa",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "wineboot",
		Name: "Wine Boot",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "gnome-contacts",
		Name: "Kontakty",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "cheese",
		Name: "Cheese",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "gnome-disks",
		Name: "Dyski",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "com.github.phase1geo.minder %f",
		Name: "Minder",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "swell-foop",
		Name: "Swell Foop",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "winecfg",
		Name: "Wine Configuration",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "simple-scan",
		Name: "Skaner dokumentów",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "gnome-system-monitor",
		Name: "Monitor systemu",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "lutris %U",
		Name: "Lutris",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "gnome-calendar",
		Name: "Kalendarz",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "winefile",
		Name: "Wine File",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "rhythmbox %U",
		Name: "Rhythmbox",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "gapplication launch org.gnome.Weather",
		Name: "Pogoda",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "gnome-screenshot --interactive",
		Name: "Zrzut ekranu",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "wine wordpad.exe",
		Name: "Wine Wordpad",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "winemine",
		Name: "WineMine",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "libreoffice --impress %U",
		Name: "LibreOffice Impress",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "wine winhlp32.exe",
		Name: "Wine Help",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "baobab",
		Name: "Wykorzystanie dysku",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "gnome-system-monitor",
		Name: "Monitor systemu GNOME",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "calibre --detach %F",
		Name: "calibre",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "transmission-gtk %U",
		Name: "Transmission",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "nm-connection-editor",
		Name: "Zaawansowana konfiguracja sieci",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "gnome-calculator",
		Name: "Kalkulator",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "eog %U",
		Name: "Obrazy",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "totem %U",
		Name: "Filmy",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "/usr/bin/google-chrome-stable %U",
		Name: "Google Chrome",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "gitg --no-wd %U",
		Name: "gitg",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "gnome-terminal",
		Name: "Terminal",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "virt-manager",
		Name: "Menedżer maszyn wirtualnych",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "firefox-wayland --name firefox-wayland %u",
		Name: "Firefox on Wayland",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "gapplication launch org.gnome.Maps %U",
		Name: "Mapy",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "dconf-editor",
		Name: "Edytor dconf",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "gnome-logs",
		Name: "Dzienniki",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "evince %U",
		Name: "Przeglądarka dokumentów",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "wine oleview.exe",
		Name: "Wine OLE View",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "/usr/bin/brave-browser-stable %U",
		Name: "Brave Web Browser",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "nautilus --new-window %U",
		Name: "Pliki",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "opera %U",
		Name: "Opera",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "wine uninstaller.exe",
		Name: "Wine Software Uninstaller",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "mpv --player-operation-mode=pseudo-gui -- %U",
		Name: "Odtwarzacz mpv",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "gnome-clocks",
		Name: "Zegar",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "gnome-software %U",
		Name: "Oprogramowanie",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "yelp %u",
		Name: "Pomoc",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "regedit",
		Name: "Regedit",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "libreoffice --calc %U",
		Name: "LibreOffice Calc",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "libreoffice --writer %U",
		Name: "LibreOffice Writer",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "libreoffice --draw %U",
		Name: "LibreOffice Draw",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "notepad",
		Name: "Notepad",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "gnome-photos",
		Name: "Zdjęcia",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "alacritty",
		Name: "Alacritty",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "gnome-abrt",
		Name: "Zgłaszanie problemów",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "gedit %U",
		Name: "Edytor tekstu",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "gnome-tweaks",
		Name: "Dostrajanie",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "/usr/share/code/code --no-sandbox --unity-launch %F",
		Name: "Visual Studio Code",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "rygel-preferences",
		Name: "Preferencje usługi Rygel",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "nvim %F",
		Name: "Neovim",
		Term: true,
		Type: "Application",
		URL:  ""},
	{Exec: "gnome-boxes %U",
		Name: "Boxes",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "mediawriter",
		Name: "Fedora Media Writer",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "/usr/bin/ThermalMonitor",
		Name: "thermald Monitor",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "gnome-control-center",
		Name: "Ustawienia",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "gnome-font-viewer %u",
		Name: "Czcionki",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "firefox %u",
		Name: "Firefox",
		Term: false,
		Type: "Application",
		URL:  ""},
}

var entriesFilteredByA = []*desktop.Entry{
	{Exec: "alacritty",
		Name: "Alacritty",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "calibre --detach %F",
		Name: "calibre",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "gnome-calendar",
		Name: "Kalendarz",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "gnome-calculator",
		Name: "Kalkulator",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "/opt/google/chrome/google-chrome --profile-directory=Default --app-id=emefpkhgihlhfddcjfghpndaeliajgjj",
		Name: "TIDAL",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "gnome-terminal",
		Name: "Terminal",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "/usr/bin/flatpak run --branch=stable --arch=x86_64 --command=com.github.lainsce.timetable com.github.lainsce.timetable",
		Name: "Timetable",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "/usr/bin/ThermalMonitor",
		Name: "thermald Monitor",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "libreoffice --calc %U",
		Name: "LibreOffice Calc",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "/opt/google/chrome/google-chrome --profile-directory=Default --app-id=hmjkmjkepdijhoojdojkdfohbdgmmhki",
		Name: "Google Keep – notatki i listy",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "/usr/share/code/code --no-sandbox --unity-launch %F",
		Name: "Visual Studio Code",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "firefox-wayland --name firefox-wayland %u",
		Name: "Firefox on Wayland",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "/opt/google/chrome/google-chrome --profile-directory=Default --app-id=ohmmkhmmmpcnpikjeljgnaoabkaalbgc",
		Name: "Chrome Apps & Extensions Developer Tool",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "gnome-abrt",
		Name: "Zgłaszanie problemów",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "wine uninstaller.exe",
		Name: "Wine Software Uninstaller",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "virt-manager",
		Name: "Menedżer maszyn wirtualnych",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "pavucontrol",
		Name: "Sterowanie głośnością PulseAudio",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "nautilus --new-window %U",
		Name: "Pliki",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "gapplication launch org.gnome.Maps %U",
		Name: "Mapy",
		Term: false,
		Type: "Application",
		URL:  ""},
	{Exec: "gapplication launch org.gnome.Weather",
		Name: "Pogoda",
		Term: false,
		Type: "Application",
		URL:  ""},
}

var names = []*desktop.Entry{
	{Name: "Alacritty"},
	{Name: "Archives"},
	{Name: "Boxes"},
	{Name: "Brave Web Browser"},
	{Name: "Cheese"},
	{Name: "Chrome Apps & Extensions Developer Tool"},
	{Name: "Chrome Dev Editor"},
	{Name: "Fonts"},
	{Name: "Disks"},
	{Name: "Logs"},
	{Name: "Text editor"},
	{Name: "FeedReader"},
	{Name: "Movies"},
	{Name: "Firefox"},
	{Name: "Geary"},
	{Name: "GoLand"},
	{Name: "Google Chrome"},
	{Name: "Google Keep - notes and lists"},
	{Name: "Calendar"},
	{Name: "Kalkulator"},
	{Name: "Thunderbird email client"},
	{Name: "Contacts"},
	{Name: "LibreOffice Calc"},
	{Name: "LibreOffice Draw"},
	{Name: "LibreOffice Impress"},
	{Name: "LibreOffice Writer"},
	{Name: "Lutris"},
	{Name: "Maps"},
	{Name: "Virtal machines manager"},
	{Name: "Minder"},
	{Name: "System monitor"},
	{Name: "System monitor GNOME"},
	{Name: "Neovim"},
	{Name: "Notepad"},
}

func TestFlatten(t *testing.T) {
	tests := []struct {
		name string
		in   []*desktop.Entry
		out  [][]string
	}{
		{
			name: "simple flatten",
			in: []*desktop.Entry{
				{Name: "Alacritty"},
				{Name: "Archives"},
			},
			out: [][]string{
				{"Alacritty", "Archives"},
				{"", ""},
			},
		},
		{
			name: "flatten some execs",
			in: []*desktop.Entry{
				{Name: "Alacritty", Exec: "/usr/bin/alacritty"},
				{Name: "Archives", Exec: "file-roller"},
			},
			out: [][]string{
				{"Alacritty", "Archives"},
				{"", "file-roller"},
			},
		},
		{
			name: "nicely flat app names and execs",
			in: []*desktop.Entry{
				{Name: "Alacritty"},
				{Name: "Archives"},
				{Name: "Virtual Machines Manager", Exec: "virt-manager"},
				{Name: "Boxes", Exec: "gnome-boxes %U"},
			},
			out: [][]string{
				{"Alacritty", "Archives", "Virtual Machines Manager", "Boxes"},
				{"", "", "virt-manager", "gnome-boxes"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := flatten(tt.in)
			if !reflect.DeepEqual(tt.out, got) {
				t.Errorf("flatten(entries) wanted %q got %q", tt.out, got)
			}
		})
	}
}

func TestEntriesIteratorNext(t *testing.T) {
	tests := []struct {
		in  ents
		out int
	}{
		{
			in: ents{
				names: names,
				curr:  0,
				start: 0,
				limit: 1,
			},
			out: 1,
		},
		{
			in: ents{
				names: names,
				curr:  0,
				start: 0,
				limit: 9,
			},
			out: 9,
		},
		{
			in: ents{
				names: names,
				curr:  0,
				start: 2,
				limit: 42,
			},
			out: 34,
		},
	}

	for _, tt := range tests {
		tt.in.Reset()

		i := 0
		for tt.in.Next() {
			i++
		}

		if i != tt.out {
			t.Errorf("iterator got %d expected %d", i, tt.out)
		}
	}
}

func TestEntriesIteratorValue(t *testing.T) {
	tests := []struct {
		in  ents
		out []string
	}{
		{
			in: ents{
				names: names,
				curr:  0,
				start: 0,
				limit: 1,
			},
			out: []string{"Alacritty"},
		},
		{
			in: ents{
				names: names,
				curr:  1,
				start: 1,
				limit: 4,
			},
			out: []string{"Archives", "Boxes", "Brave Web Browser", "Cheese"},
		},
		{
			in: ents{
				names: names,
				curr:  32,
				start: 32,
				limit: 2,
			},
			out: []string{"Neovim", "Notepad"},
		},
	}

	for _, tt := range tests {
		tt.in.Reset()

		var got []string
		for tt.in.Next() {
			got = append(got, tt.in.Value())
		}

		if !reflect.DeepEqual(got, tt.out) {
			t.Errorf("iterator got %q expected %q", got, tt.out)
		}
	}
}

func TestEntriesHandleResize(t *testing.T) {
	tests := []struct {
		name         string
		in           ents
		out          ents
		shouldResize bool
		winH, entryH int
	}{
		{
			name: "proper initialize",
			in: ents{
				names: names,
				curr:  0,
				start: 0,
				limit: 0,
			},
			out: ents{
				names: names,
				start: 0,
				limit: 26,
				curr:  0,
			},
			shouldResize: true,
			winH:         580,
			entryH:       21,
		},
		{
			name: "change size",
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
			name: "same size",
			in: ents{
				names: names,
				curr:  0,
				start: 0,
				limit: 26,
			},
			out: ents{
				names: names,
				curr:  0,
				start: 0,
				limit: 26,
			},
			shouldResize: false,
			winH:         580,
			entryH:       21,
		},
		{
			name: "cursor moved, limit exceeds len",
			in: ents{
				names: names,
				start: 5,
				limit: 26,
				curr:  30,
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
			name: "make it smaller, when cursor moved near the end",
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
		t.Run(tt.name, func(t *testing.T) {
			tt.in.handleResize(tt.winH, tt.entryH)
			if isNotEqual(tt.in, tt.out) {
				t.Errorf("handleResize(%d, %d)", tt.winH, tt.entryH)
			}
		})
	}
}

func TestEntriesCursorDown(t *testing.T) {
	tests := []struct {
		in    ents
		times int
		out   int
	}{
		{
			in:    ents{names: names, curr: 0, start: 0},
			times: 1,
			out:   1,
		},
		{
			in:    ents{names: names, curr: 0, start: 0},
			times: 7,
			out:   7,
		},
		{
			in:    ents{names: names, curr: 0, start: 0},
			times: 8,
			out:   8,
		},
		{
			in:    ents{names: names, curr: 0, start: 0},
			times: 40,
			out:   33,
		},
	}

	for _, tt := range tests {
		for i := 0; i < tt.times; i++ {
			tt.in.CursorDown()
		}
		if tt.in.curr != tt.out {
			t.Errorf("CursorDown() %d times got %d expected %d", tt.times, tt.in.curr, tt.out)
		}
	}
}

func TestEntriesCursorUp(t *testing.T) {
	tests := []struct {
		in    ents
		times int
		out   int
	}{
		{
			in:    ents{curr: 8},
			times: 1,
			out:   7,
		},
		{
			in:    ents{curr: 8},
			times: 7,
			out:   1,
		},
		{
			in:    ents{curr: 8},
			times: 8,
			out:   0,
		},
		{
			in:    ents{curr: 7},
			times: 8,
			out:   0,
		},
		{
			in:    ents{curr: 0},
			times: 8,
			out:   0,
		},
		{
			in:    ents{curr: 2, start: 2},
			times: 8,
			out:   0,
		},
	}

	for _, tt := range tests {
		for i := 0; i < tt.times; i++ {
			tt.in.CursorUp()
		}
		if tt.in.curr != tt.out {
			t.Errorf("CursorUp() times %d got %d expected %d", tt.times, tt.in.curr, tt.out)
		}
	}
}

func BenchmarkFlatten(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = flatten(entries)
	}
}

func isNotEqual(in, out ents) bool {
	return in.curr != out.curr || in.start != out.start ||
		in.iter != out.iter || in.limit != out.limit || in.filter != out.filter
}
