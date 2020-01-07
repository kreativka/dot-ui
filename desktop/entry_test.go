package desktop

import (
	"reflect"
	"strings"
	"testing"
)

func TestIsTrue(t *testing.T) {
	cases := []struct {
		in  string
		out bool
	}{
		{"true", true},
		{"True", false},
		{"false", false},
		{"False", false},
		{"Something something", false},
	}

	for _, tt := range cases {
		got := isTrue(tt.in)
		if tt.out != got {
			t.Errorf("isTrue(%q) => wanted %t got %t", tt.in, tt.out, got)
		}
	}
}

func TestParseStr(t *testing.T) {
	cases := []struct {
		oldStr, newStr, out string
	}{
		{
			"", "nvim", "nvim",
		},
		{"nvim", "neovim", "nvim"},
	}

	for _, tt := range cases {
		got := parseStr(tt.oldStr, tt.newStr)
		if got != tt.out {
			t.Errorf("parseStr(%s, %s) =>\nwanted %q\ngot %q", tt.oldStr, tt.newStr, tt.out, got)
		}
	}
}

func TestParseLine(t *testing.T) {
	cases := []struct {
		in, lang      string
		err           error
		inEnt, outEnt Entry
	}{
		{
			in:     "Exec=nvim %F",
			lang:   "Name[pl]",
			err:    nil,
			inEnt:  Entry{},
			outEnt: Entry{Exec: "nvim %F"},
		},
		{
			in:     "Name=Neovim",
			lang:   "Name[pl]",
			err:    nil,
			inEnt:  Entry{},
			outEnt: Entry{Name: "Neovim"},
		},
		{
			in:     "Type=Application",
			lang:   "Name[pl]",
			err:    nil,
			inEnt:  Entry{},
			outEnt: Entry{Type: "Application"},
		},
		{
			in:     "Terminal=true",
			lang:   "Name[pl]",
			err:    nil,
			inEnt:  Entry{},
			outEnt: Entry{Term: true},
		},
		{
			in:     "Hidden=true",
			lang:   "Name[pl]",
			err:    ErrHiddenEntry,
			inEnt:  Entry{},
			outEnt: Entry{},
		},
		{
			in:     "Hidden=false",
			lang:   "Name[pl]",
			err:    nil,
			inEnt:  Entry{},
			outEnt: Entry{},
		},
		{
			in:     "NoDisplay=false",
			lang:   "Name[pl]",
			err:    nil,
			inEnt:  Entry{},
			outEnt: Entry{},
		},
		{
			in:     "NoDisplay=true",
			lang:   "Name[pl]",
			err:    ErrHiddenEntry,
			inEnt:  Entry{},
			outEnt: Entry{},
		},
		{
			in:     "Name[pl]=Neovim",
			lang:   "Name[pl]",
			err:    nil,
			inEnt:  Entry{Name: "Neovim"},
			outEnt: Entry{Name: "Neovim"},
		},
		{
			in:     "Name[pl]=Neovim",
			lang:   "Name[pl]",
			err:    nil,
			inEnt:  Entry{Name: "Something not localized"},
			outEnt: Entry{Name: "Neovim"},
		},
		{
			in:     "OtherField=Something blah",
			lang:   "Name[pl]",
			err:    nil,
			inEnt:  Entry{},
			outEnt: Entry{},
		},
		{
			in:     "OtherFieldSomething blah",
			lang:   "Name[pl]",
			err:    nil,
			inEnt:  Entry{},
			outEnt: Entry{},
		},
		{
			in:     "=O==th=erFieldSomething =-blah=",
			lang:   "Name[pl]",
			err:    nil,
			inEnt:  Entry{},
			outEnt: Entry{},
		},
		{
			in:     "",
			lang:   "Name[pl]",
			err:    nil,
			inEnt:  Entry{},
			outEnt: Entry{},
		},
	}

	for _, tt := range cases {
		err := parseLine(tt.in, tt.lang, &tt.inEnt)
		if err != tt.err || !reflect.DeepEqual(tt.outEnt, tt.inEnt) {
			t.Errorf("given(%s, %s, %+v): expected %+v, actual %+v", tt.in, tt.lang, tt.inEnt, tt.outEnt, tt.inEnt)
		}
	}
}

func TestParseEntry(t *testing.T) {
	var nvimDesktop = `
[Desktop Entry]
Name=Neovim
GenericName=Text Editor
GenericName[de]=Texteditor
Comment=Edit text files
TryExec=nvim
Exec=nvim %F
Terminal=true
Type=Application
Keywords=Text;editor;
Icon=nvim
Categories=Utility;TextEditor;
StartupNotify=false
MimeType=text/english;text/plain;text/x-makefile;text/x-c++hdr;text/x-c++src;text/x-chdr;text/x-csrc;text/x-java;text/x-moc;text/x-pascal;text/x-tcl;text/x-tex;application/x-shellscript;text/x-c;text/x-c++;
X-Desktop-File-Install-Version=0.23
`
	var cases = []struct {
		in      string
		lang    string
		useLang bool
		out     *Entry
		err     error
	}{
		{
			in:      nvimDesktop,
			lang:    "",
			useLang: true,
			out:     &Entry{Name: "Neovim", Exec: "nvim %F", Term: true, Type: "Application"},
			err:     nil,
		}, {
			in:      nvimDesktop,
			lang:    "",
			useLang: false,
			out:     &Entry{Name: "Neovim", Exec: "nvim %F", Term: true, Type: "Application"},
			err:     nil,
		}, {
			in:      nvimDesktop,
			lang:    "pl",
			useLang: false,
			out:     &Entry{Name: "Neovim", Exec: "nvim %F", Term: true, Type: "Application"},
			err:     nil,
		},
		{
			in:      nvimDesktop,
			lang:    "pl",
			useLang: true,
			out:     &Entry{Name: "Neovim", Exec: "nvim %F", Term: true, Type: "Application"},
			err:     nil,
		},
	}

	for _, tt := range cases {
		res, err := parseEntry(strings.NewReader(tt.in), tt.lang)
		if !reflect.DeepEqual(tt.out, res) || err != tt.err {
			t.Errorf("%q", err)
		}
	}
}

var entries = []*Entry{
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

func TestFlatten(t *testing.T) {
	tests := []struct {
		name string
		in   []*Entry
		out  [][]string
	}{
		{
			name: "simple flatten",
			in: []*Entry{
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
			in: []*Entry{
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
			in: []*Entry{
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
			got := Flatten(tt.in)
			if !reflect.DeepEqual(tt.out, got) {
				t.Errorf("flatten(entries) wanted %q got %q", tt.out, got)
			}
		})
	}
}

func BenchmarkFlatten(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Flatten(entries)
	}
}
