package service

import (
	"errors"
	"strings"
)

func Recover(alias string) (string, error) {
	if strings.TrimSpace(alias) == "" {
		return "", errors.New("invalid alias")
	}
	if url, found := Storage.Get(alias); found {
		return url, nil
	}
	return "", errors.New("alias not found")
}
