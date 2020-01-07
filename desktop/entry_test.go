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
