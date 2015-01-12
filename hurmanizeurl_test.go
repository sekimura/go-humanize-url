package humanizeurl

import (
	"testing"
)

func TestHumanize(t *testing.T) {

	tests := []struct {
		value    string
		expected string
	}{
		{"http://sekimura.org", "sekimura.org"},
		{" http://sekimura.org", "sekimura.org"},
		{"https://sekimura.org", "sekimura.org"},
		{"http://www.sekimura.org", "sekimura.org"},
		{"www.sekimura.org", "sekimura.org"},
		{"http://sekimura.org/", "sekimura.org"},
		{"sekimura.org/", "sekimura.org"},
		{"http://sekimura.org/foo/", "sekimura.org/foo"},
		{"file:///Users/sekimura/dev/humanize-url/", "file:///Users/sekimura/dev/humanize-url"},
	}

	for i := range tests {
		actual, err := Humanize(tests[i].value)
		if err != nil {
			t.Error("Humanize failed", err)
		}
		if actual != tests[i].expected {
			t.Error("Humanize did not match expected value", actual, tests[i].expected)
		}
	}
}
