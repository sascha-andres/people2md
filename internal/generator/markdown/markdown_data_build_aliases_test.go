package markdown

import "testing"

// TestMarkdownData_BuildAlias tests the buildAlias function to ensure it correctly replaces special characters in strings.
func TestMarkdownData_BuildAlias(t *testing.T) {
	var tests = map[string]struct {
		name string
		want string
	}{
		"one umlaut": {
			name: "one umläut",
			want: "one umlaut",
		},
		"two umlauts": {
			name: "two umläüts",
			want: "two umlauts",
		},
	}
	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			if got := buildAlias(test.name); got != test.want {
				t.Errorf("buildAlias() = [%v], want [%v]", got, test.want)
			}
		})
	}
}
