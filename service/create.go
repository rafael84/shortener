package service

import (
	"errors"
	parser "net/url"
)

func Create(url, alias string) (string, error) {
	if url == "" {
		return "", errors.New("url is required")
	}
	u, err := parser.Parse(url)
	if err != nil || u.Scheme == "" {
		return "", errors.New("url is invalid")
	}
	if alias == "" {
		alias = Encode(Storage.Count())
	}
	if err := Storage.Set(alias, url); err != nil {
		return "", err
	}
	return alias, nil
}

func Encode(n int) string {
	if n == 0 {
		return string(Alphabet[0])
	}
	t := make([]byte, 0)
	lenA := len(Alphabet)
	for n > 0 {
		t = append(t, Alphabet[n%lenA])
		n = n / lenA
	}
	for l, r := 0, len(t)-1; l < r; l, r = l+1, r-1 {
		t[l], t[r] = t[r], t[l]
	}
	return string(t)
}
