package desktop

import (
	"testing"
)

func TestTrimRight(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{"", ""},
		{"@@ %f @@", ""},
		{"@@ %F @@ a", "@@ %F @@ a"},
		{"@@u %u @@", ""},
		{"@@u %U @@", ""},
		{" %u", ""},
		{" %U", ""},
		{" %f", ""},
		{" %F", ""},
		{
			"/usr/bin/flatpak run --branch=stable --arch=x86_64 --command=signal --file-forwarding org.signal.Signal @@u %U @@",
			"/usr/bin/flatpak run --branch=stable --arch=x86_64 --command=signal --file-forwarding org.signal.Signal",
		},
		{"nvim %F", "nvim"},
	}

	for _, tt := range cases {
		res := TrimRight(tt.in)
		if tt.out != res {
			t.Errorf("trimRight(%q)\nwanted %q\ngot %q", tt.in, tt.out, res)
		}
	}
}

func BenchmarkTrimRight(b *testing.B) {
	cases := []struct {
		name string
		f    func(string) string
		in   string
	}{
		{
			"flatpak",
			TrimRight,
			"/usr/bin/flatpak run --branch=stable --arch=x86_64 --command=signal --file-forwarding org.signal.Signal @@u %U @@",
		},
		{
			"non Flatpak",
			TrimRight,
			"nvim %F",
		},
	}

	for _, bench := range cases {
		b.Run(bench.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = bench.f(bench.in)
			}
		})
	}
}
