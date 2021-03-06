package desktop

import "strings"

// TrimRight returns string without placeholders for addittional arguments
func TrimRight(cmd string) string {
	if len(cmd) < 2 || cmd == "" {
		return cmd
	}

	patterns := []string{
		"@@ %f @@", "@@ %F @@", "@@u %u @@", "@@u %U @@",
		" %f", " %F", " %u", " %U",
	}

	for _, suf := range patterns {
		if strings.Contains(cmd, suf) {
			cmd = strings.TrimRight(cmd, suf)
		}
	}

	return cmd
}
