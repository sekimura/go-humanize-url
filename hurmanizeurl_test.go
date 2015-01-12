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
		{"file:///Users/sindresorhus/dev/humanize-url/", "file:///Users/sindresorhus/dev/humanize-url"},
	}

	for i := range tests {
		actual, err := humanize(tests[i].value)
		if err != nil {
			t.Error("humanize failed", err)
		}
		if actual != tests[i].expected {
			t.Error("humanize did not match expected value", actual, tests[i].expected)
		}
	}
}
