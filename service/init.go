package service

import "github.com/rafael84/shortener/persistence"

var (
	Alphabet = "bcdfghjkmnpqrstvwxyz23456789BCDFGHJKMNPQRSTVWXYZ"
	Storage  = persistence.NewMemory()
)
