package service

import (
	"errors"
	"strings"
)

var (
	saved = map[string]string{
		"f1": "http://valid.com",
		"f2": "http://valid.com/2",
		"f3": "http://valid.com/3",
	}
)

func Recover(alias string) (string, error) {
	if strings.TrimSpace(alias) == "" {
		return "", errors.New("invalid alias")
	}
	if url, found := saved[alias]; found {
		return url, nil
	}
	return "", errors.New("alias not found")
}
