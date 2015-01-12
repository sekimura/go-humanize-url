package humanizeurl

import (
	"net/url"
	"strings"
)

func humanize(s string) (string, error) {
	s = strings.TrimSpace(s)

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	h := u.String()
	if strings.HasPrefix(u.Scheme, "http") {
		h = strings.TrimPrefix(h, u.Scheme+"://")
	}

	h = strings.TrimPrefix(h, "www.")

	h = strings.TrimSuffix(h, "/")

	return h, nil
}
