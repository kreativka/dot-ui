package dotui

import (
	"reflect"
	"testing"

	"github.com/kreativka/dot-ui/desktop"
)

func BenchmarkFilterMatches(b *testing.B) {
	elems := desktop.Flatten(entries)

	for i := 0; i < b.N; i++ {
		_ = filterMatches(entries, elems, "A")
	}
}

func TestFilterMatches(t *testing.T) {
	type testIn struct {
		entries []*desktop.Entry
		lists   [][]string
		filter  string
	}

	tests := []struct {
		name string
		in   testIn
		out  []*desktop.Entry
	}{
		{
			in: testIn{
				entries: []*desktop.Entry{
					{Name: "Alacritty"},
					{Name: "Archives"},
				},
				lists: [][]string{
					{"Alacritty", "Archives"},
					{"", ""},
				},
				filter: "A",
			},
			out: []*desktop.Entry{
				{Name: "Archives"},
				{Name: "Alacritty"},
			},
		},
		{
			in: testIn{
				entries: []*desktop.Entry{
					{Name: "Alacritty"},
					{Name: "Archives"},
					{Name: "Menedżer maszyn wirtualnych", Exec: "virt-manager"},
				},
				lists: [][]string{
					{"Alacritty", "Archives", "Menedżer maszyn wirtualnych"},
					{"", "", "virt-manager"},
				},
				filter: "virt",
			},
			out: []*desktop.Entry{
				{Name: "Menedżer maszyn wirtualnych", Exec: "virt-manager"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := filterMatches(tt.in.entries, tt.in.lists, tt.in.filter)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("filterMatches(entries, lists, filter) wanted %p got %p",
					tt.out, got)
			}
		})
	}
}

func TestFilterEntries(t *testing.T) {
	flattenNames := [][]string{
		{"Alacritty", "Archives", "Boxes", "Brave Web Browser", "Cheese", "Firefox"},
		{"", "", "", "", "", ""},
	}
	names := []*desktop.Entry{
		{Name: "Alacritty"},
		{Name: "Archives"},
		{Name: "Boxes"},
		{Name: "Brave Web Browser"},
		{Name: "Cheese"},
		{Name: "Firefox"},
	}
	tests := []struct {
		name, filter string
		in           mainPage
		out          mainPage
	}{
		{
			name:   "empty filter",
			filter: "",
			in: mainPage{env: &Env{
				entries: names,
				currList: &ents{
					names: names,
				},
			}},
			out: mainPage{env: &Env{
				entries: names,
				currList: &ents{
					names: names,
				},
			}},
		},
		{
			name:   "filter is A and m.filter is empty",
			filter: "A",
			in: mainPage{env: &Env{
				entries: names,
				currList: &ents{
					names: names,
					list:  flattenNames,
				},
			}},
			out: mainPage{env: &Env{
				entries: names,
				currList: &ents{
					list: [][]string{
						{"Archives",
							"Alacritty",
							"Brave Web Browser"},
						{"", "", ""},
					},
					names: []*desktop.Entry{
						{Name: "Archives"},
						{Name: "Alacritty"},
						{Name: "Brave Web Browser"},
					},
					filter: "A",
				},
			}},
		},
		{
			name:   "filter is Al and m.filter is A",
			filter: "Al",
			in: mainPage{env: &Env{
				entries: names,
				currList: &ents{
					names:  names,
					list:   flattenNames,
					filter: "A",
				},
			}},
			out: mainPage{env: &Env{
				entries: names,
				currList: &ents{
					list: [][]string{
						{"Alacritty"},
						{""},
					},
					names: []*desktop.Entry{
						{Name: "Alacritty"},
					},
					filter: "Al",
				},
			}},
		},
		{
			name:   "filter is Alacritty and m.filter is Alacrittyl",
			filter: "Alacritty",
			in: mainPage{env: &Env{
				entries: names,
				currList: &ents{
					names:  nil,
					list:   nil,
					filter: "Alacrittyl",
					curr:   0,
				},
			}},
			out: mainPage{env: &Env{
				entries: names,
				currList: &ents{
					list: [][]string{
						{"Alacritty"},
						{""},
					},
					names: []*desktop.Entry{
						{Name: "Alacritty"},
					},
					filter: "Alacritty",
					curr:   0,
				},
			}},
		},
		{
			name:   "filter is Al and m.filter is A but start is 2",
			filter: "Al",
			in: mainPage{env: &Env{
				entries: names,
				currList: &ents{
					names:  names,
					list:   flattenNames,
					filter: "A",
					start:  2,
				},
			}},
			out: mainPage{env: &Env{
				entries: names,
				currList: &ents{
					list: [][]string{
						{"Alacritty"},
						{""},
					},
					names: []*desktop.Entry{
						{Name: "Alacritty"},
					},
					filter: "Al",
				},
			}},
		},
		{
			name:   "filter is Al and m.filter is A but curr is 2",
			filter: "Al",
			in: mainPage{env: &Env{
				entries: names,
				currList: &ents{
					names:  names,
					list:   flattenNames,
					filter: "A",
					curr:   2,
				},
			}},
			out: mainPage{env: &Env{
				entries: names,
				currList: &ents{
					list: [][]string{
						{"Alacritty"},
						{""},
					},
					names: []*desktop.Entry{
						{Name: "Alacritty"},
					},
					filter: "Al",
				},
			}},
		},
		{
			name:   "filter is A and m.filter is Al",
			filter: "A",
			in: mainPage{env: &Env{
				entries: names,
				currList: &ents{
					list: [][]string{
						{"Alacritty"},
						{""},
					},
					names: []*desktop.Entry{
						{Name: "Alacritty"},
					},
					filter: "Al",
				},
			}},
			out: mainPage{env: &Env{
				entries: names,
				currList: &ents{
					list: [][]string{
						{"Archives", "Alacritty", "Brave Web Browser"},
						{"", "", ""},
					},
					names: []*desktop.Entry{
						{Name: "Archives"}, {Name: "Alacritty"},
						{Name: "Brave Web Browser"},
					},
					filter: "A",
				},
			}},
		},
		{
			name:   "filter is cleared and m.filter is A",
			filter: "",
			in: mainPage{env: &Env{
				entries: names,
				currList: &ents{
					list: [][]string{
						{"Archives", "Alacritty", "Brave Web Browser"},
						{"", "", ""},
					},
					names: []*desktop.Entry{
						{Name: "Archives"}, {Name: "Alacritty"},
						{Name: "Brave Web Browser"},
					},
					filter: "A",
				},
			}},
			out: mainPage{env: &Env{
				entries: names,
				currList: &ents{
					names: names,
					list:  flattenNames,
				},
			}},
		},
		{
			name:   "edge case curr is high, start > 0, resized and filter is Al",
			filter: "Al",
			in: mainPage{env: &Env{
				entries: entries,
				currList: &ents{
					names:  entries,
					list:   desktop.Flatten(entries),
					curr:   21,
					limit:  10,
					start:  10,
					filter: "A",
				},
			}},
			out: mainPage{env: &Env{
				entries: entries,
				currList: &ents{
					names:  entriesFilteredByA,
					list:   desktop.Flatten(entriesFilteredByA),
					curr:   9,
					limit:  10,
					start:  0,
					filter: "Al",
				},
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in.filterEntries(tt.filter)
			if !reflect.DeepEqual(tt.in, tt.out) {
				t.Errorf("%s", tt.name)
			}
		})
	}
}
