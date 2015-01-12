package humanizeurl

import (
	"net/url"
	"strings"

	"github.com/sekimura/go-normalize-url"
)

func Humanize(s string) (string, error) {
	s = strings.TrimSpace(s)
	s, err := normalizeurl.Normalize(s)
	if err != nil {
		return s, err
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	h := u.String()
	if strings.HasPrefix(u.Scheme, "http") {
		h = strings.TrimPrefix(h, u.Scheme+"://")
	}

	return h, nil
}
