package humanizeurl

import (
	"testing"
)

func TestHumanize(t *testing.T) {

	var tests = map[string]string{
		"http://sekimura.org":      "sekimura.org",
		" http://sekimura.org":     "sekimura.org",
		"https://sekimura.org":     "sekimura.org",
		"http://www.sekimura.org":  "sekimura.org",
		"www.sekimura.org":         "sekimura.org",
		"http://sekimura.org/":     "sekimura.org",
		"sekimura.org/":            "sekimura.org",
		"http://sekimura.org/foo/": "sekimura.org/foo",

		"file:///Users/sindresorhus/dev/humanize-url/": "file:///Users/sindresorhus/dev/humanize-url",
	}

	for orig, expected := range tests {
		actual, err := humanize(orig)
		if err != nil {
			t.Error("humanize failed", err, expected)
		}
		if actual != expected {
			t.Error("humanize did not match expeted value", actual, expected)
		}
	}
}
