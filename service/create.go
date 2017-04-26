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
	return alias, nil
}
