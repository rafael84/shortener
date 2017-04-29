package service

import "github.com/rafael84/shortener/persistence"

var (
	Alphabet = "abc"

	Storage   = persistence.NewMemory()
	generated = 0
)
